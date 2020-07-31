package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://github.com/telegramdesktop/tdesktop/tree/dev/Telegram/Resources/tl
// https://github.com/danog/MadelineProto/tree/master/src/danog/MadelineProto
// https://github.com/tdlib/td/tree/master/td/generate/scheme

// https://core.telegram.org/mtproto/TL
// https://core.telegram.org/mtproto/TL-combinator
// identifier#name attr:type attr:type = resultType;

// TODO:
// Pq Ids
// Silent  bool //flag (add flag number)
// Channel TL   // InputChannel (add possible types)
// Use decodeResponse of inner obj when decoding types like:
//  invokeWithLayer#da9b0d0d {X:Type} layer:int query:!X = X;

type Field struct {
	name     string
	typeName string
	flagBit  int
}

func (f Field) isFlag() bool {
	return f.flagBit >= 0
}

type Combinator struct {
	id         string
	name       uint32
	fields     []Field
	typeName   string
	isFunction bool
}

func (c Combinator) hasFlags() bool {
	for _, f := range c.fields {
		if f.typeName == "#" {
			return true
		}
	}
	return false
}

func normalize(s string) string {
	x := []byte(s)
	for i, r := range x {
		if r == '.' {
			x[i] = '_'
		}
	}
	y := string(x)
	if y == "type" {
		return "_type"
	}
	return y
}

func normalizeName(nameStr string) uint32 {
	if nameStr == "" {
		return 0
	}
	if nameStr[0] == '#' {
		nameStr = nameStr[1:]
	}
	nameInt, err := strconv.ParseInt(nameStr, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return uint32(nameInt)
}

func normalizeAttr(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	if strings.HasSuffix(s, "Id") {
		s = s[:len(s)-2] + "ID"
	}
	return s
}

func maybeFlagged(_type string, isFlag bool, flagBit int, args ...string) string {
	argsStr := strings.Join(args, ",")
	if isFlag {
		return fmt.Sprintf("m.Flagged%s(flags, %d, %s),\n", _type, flagBit, argsStr)
	} else {
		return fmt.Sprintf("m.%s(%s),\n", _type, argsStr)
	}
}

func maybeFlaggedExt(_type string, isFlag bool, flagBit int) string {
	if isFlag {
		return fmt.Sprintf("m.Flagged%s(flags, %d)", _type, flagBit)
	} else {
		return fmt.Sprintf("m.%s()", _type)
	}
}

func makeField(name, typeName string) Field {
	flagBit := -1
	if strings.HasPrefix(typeName, "flags.") { //flags.2?string
		var err error
		qPos := strings.Index(typeName, "?")
		dPos := strings.Index(typeName, ".")
		flagBit, err = strconv.Atoi(typeName[dPos+1 : qPos])
		typeName = typeName[qPos+1:]
		if err != nil {
			log.Fatalf("parsing %s: %s", typeName, err)
		}
	}
	return Field{normalize(name), normalize(typeName), flagBit}
}

var fieldsFixForCrcRegexp = regexp.MustCompile(`([Vv])ector<(.*?)>`)

func makeCombinatorDescription(id, fieldsStr, typeName string) string {
	if fieldsStr != "" {
		var fiteredFields []string
		for _, f := range strings.Split(fieldsStr, " ") {
			// for some reason if flagged field has type "true", it is NOT used in crc32 completely
			if strings.HasSuffix(f, "?true") {
				continue
			}
			fiteredFields = append(fiteredFields, f)
		}
		fieldsStr = strings.Join(fiteredFields, " ")
		// for some reason if field is "name:bytes" crc32 will be calculated from "name:string"
		fieldsStr = strings.Replace(fieldsStr, ":bytes", ":string", -1)
		fieldsStr = strings.Replace(fieldsStr, "?bytes", "?string", -1) //same for flags
		// for some reason... again
		fieldsStr = strings.Replace(fieldsStr, "{X:Type}", "X:Type", -1)
	}

	descr := id
	if fieldsStr != "" {
		descr += " " + fieldsStr
	}
	descr += " = " + typeName

	// for come reason if type is "Vector<subtype>" crc32 will be calculated from "Vector subtype"
	// and SOME TIMES it is named "vector<subtype>" (with lower "v")
	descr = fieldsFixForCrcRegexp.ReplaceAllString(descr, "${1}ector $2")
	return descr
}

func parseTLSchema(data []byte) []*Combinator {
	// processing constructors
	var combinators []*Combinator
	isFunction := false

	lineRegexp := regexp.MustCompile(`^(.*?)(#[a-f0-9]*)? (.*)= (.*);$`)
	fieldRegexp := regexp.MustCompile(`^(.*?):(.*)$`)
	total := 0
	for lineNum, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if line == "---functions---" {
			isFunction = true
			continue
		}
		if line == "---types---" {
			isFunction = false
			continue
		}

		match := lineRegexp.FindStringSubmatch(line)
		if len(match) == 0 {
			log.Printf("line %d: wrong combinator: %s", lineNum+1, line)
			continue
		}

		id := strings.TrimSpace(match[1])
		if id == "vector" {
			continue
		}
		name := normalizeName(match[2])
		fieldsStr := strings.TrimSpace(match[3])
		typeName := strings.TrimSpace(match[4])
		total += 1

		// making combinator description string (without id) and checking it's crc32
		descr := makeCombinatorDescription(id, fieldsStr, typeName)
		crc32sum := normalizeName(fmt.Sprintf("%x", crc32.ChecksumIEEE([]byte(descr))))
		if name == 0 {
			log.Printf("WARN: line %d: missing crc32 sum: %s", lineNum+1, line)
			name = crc32sum
		} else if name != crc32sum {
			log.Printf("WARN: line %d: wrong crc32 sum, expected %08x: %s", lineNum+1, crc32sum, line)
		}

		id = normalize(id)
		typeName = normalize(typeName)

		fields := make([]Field, 0, 16)
		for _, fieldStr := range strings.Split(fieldsStr, " ") {
			fieldStr = strings.TrimSpace(fieldStr)
			if fieldStr == "" {
				continue
			}
			if strings.HasPrefix(fieldStr, "{") { //if it is "{X:Type}", just skipping, "!X" type will be written as "TL" later
				continue
			}
			match := fieldRegexp.FindStringSubmatch(fieldStr)
			if len(match) == 0 {
				log.Fatalf("line %d: wrong field: %s", lineNum+1, fieldStr)
			}
			name, typeName := match[1], match[2]
			fields = append(fields, makeField(name, typeName))
		}

		combinators = append(combinators, &Combinator{ normalizeAttr(id), name, fields, typeName, isFunction})
	}
	print("Parsed: ", total)
	return combinators
}

func main() {
	if len(os.Args) != 2 {
		println("Usage: " + os.Args[0] + " layer")
		os.Exit(2)
	}
	layer, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// read json file from stdin
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	// parsing
	combinators := parseTLSchema(data)

	// constants
	fmt.Printf(`package mtproto
import (
	"github.com/ansel1/merry"
)

`)
	fmt.Printf("const (\n")
	fmt.Printf("Layer = %d\n", layer)
	for _, c := range combinators {
		fmt.Printf("crc%s = 0x%08x\n", c.id, c.name)
	}
	fmt.Printf(")\n\n")

	// type structs
	for _, c := range combinators {
		//print(fmt.Sprintf("%+v", c), "\n\n\n")
		fmt.Printf("type %s struct {\n", c.id)
		for _, t := range c.fields {
			line := fmt.Sprintf("%s\t", normalizeAttr(t.name))
			switch t.typeName {
			case "true": //flags only
				line += fmt.Sprintf("bool")
			case "int", "#":
				line += fmt.Sprintf("int32")
			case "long":
				line += fmt.Sprintf("int64")
			case "int128":
				line += fmt.Sprintf("[]byte")
			case "int256":
				line += fmt.Sprintf("[]byte")
			case "string":
				line += fmt.Sprintf("string")
			case "double":
				line += fmt.Sprintf("float64")
			case "bytes":
				line += fmt.Sprintf("[]byte")
			case "Vector<int>":
				line += fmt.Sprintf("[]int32")
			case "Vector<long>":
				line += fmt.Sprintf("[]int64")
			case "Vector<string>":
				line += fmt.Sprintf("[]string")
			case "Vector<double>":
				line += fmt.Sprintf("[]float64")
			case "!X":
				line = fmt.Sprintf("%s TL",  normalizeAttr(t.name))
			default:
				var inner string
				n, _ := fmt.Sscanf(strings.Title(t.typeName), "Vector<%s", &inner)
				if n == 1 {
					line = fmt.Sprintf("%s []TL // %s",  normalizeAttr(t.name), inner[:len(inner)-1])
				} else {
					line = fmt.Sprintf("%s TL // %s",  normalizeAttr(t.name), t.typeName)
				}
			}
			if t.isFlag() {
				line += fmt.Sprintf(" // flag")
			}
			fmt.Printf("%s\n", line)
		}
		fmt.Printf("}\n\n")
	}

	// encode funcs
	for _, c := range combinators {
		fmt.Printf("func (e %s) encode() []byte {\n",  c.id)
		fmt.Printf("x := NewEncodeBuf(512)\n")
		fmt.Printf("x.UInt(crc%s)\n", c.id)
		for _, t := range c.fields {
			attrName := "e." + normalizeAttr(t.name)
			if t.isFlag() && t.typeName != "true" {
				fmt.Printf("if e.Flags & %d != 0 {\n", 1<<uint(t.flagBit))
				//attrName = "*" + attrName
			}
			switch t.typeName {
			case "true": //flags only
				fmt.Printf("//flag %s\n", attrName)
			case "int", "#":
				fmt.Printf("x.Int(%s)\n", attrName)
			case "long":
				fmt.Printf("x.Long(%s)\n", attrName)
			case "int128":
				fmt.Printf("x.Bytes(%s)\n", attrName)
			case "int256":
				fmt.Printf("x.Bytes(%s)\n", attrName)
			case "string":
				fmt.Printf("x.String(%s)\n", attrName)
			case "double":
				fmt.Printf("x.Double(%s)\n", attrName)
			case "bytes":
				fmt.Printf("x.StringBytes(%s)\n", attrName)
			case "Vector<int>":
				fmt.Printf("x.VectorInt(%s)\n", attrName)
			case "Vector<long>":
				fmt.Printf("x.VectorLong(%s)\n", attrName)
			case "Vector<string>":
				fmt.Printf("x.VectorString(%s)\n", attrName)
			case "Vector<double>":
				fmt.Printf("x.VectorDouble(%s)\n", attrName)
			case "!X":
				if t.isFlag() {
					fmt.Printf("x.Bytes(%s.encode())\n", attrName[1:])
				} else {
					fmt.Printf("x.Bytes(%s.encode())\n", attrName)
				}
			default:
				var inner string
				n, _ := fmt.Sscanf(strings.Title(t.typeName), "Vector<%s", &inner)
				if n == 1 {
					fmt.Printf("x.Vector(%s) \n", attrName)
				} else {
					fmt.Printf("x.Bytes(%s.encode())\n", attrName)
				}
			}
			if t.isFlag() && t.typeName != "true" {
				fmt.Printf("}\n")
			}
		}
		fmt.Printf("return x.buf\n")
		fmt.Printf("}\n\n")
	}

	// request decode funcs (for funtions)
	for _, c := range combinators {
		if c.isFunction {
			fmt.Printf("func (e %s) decodeResponse(dbuf *DecodeBuf) TL {\n", c.id)
			if c.typeName == "Vector<int>" {
				fmt.Printf("return VectorInt(dbuf.VectorInt())\n")
			} else if c.typeName == "Vector<long>" {
				fmt.Printf("return VectorLong(dbuf.VectorLong())\n")
			} else if strings.HasPrefix(c.typeName, "Vector<") {
				fmt.Printf("return VectorObject(dbuf.Vector())\n")
			} else {
				fmt.Printf("return dbuf.Object()\n")
			}
			fmt.Printf("}\n\n")
		}
	}

	// decode funcs
	fmt.Printf(`
func readFlags(m *DecodeBuf, flagsPtr *int32) int32 {
	flags := m.Int()
	*flagsPtr = flags
	return flags
}

func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {`)

	for _, c := range combinators {
		fmt.Printf("case crc%s:\n", c.id)
		if c.hasFlags() {
			fmt.Printf("var flags int32\n")
		}
		fmt.Printf("r = %s{\n", c.id)
		for _, t := range c.fields {
			isFlag := t.isFlag()
			switch t.typeName {
			case "true": //flags only
				fmt.Printf("flags & %d != 0, //flag #%d\n", 1<<uint(t.flagBit), t.flagBit)
			case "#":
				fmt.Printf("readFlags(m, &flags),\n")
			case "int":
				fmt.Printf(maybeFlagged("Int", isFlag, t.flagBit))
			case "long":
				fmt.Printf(maybeFlagged("Long", isFlag, t.flagBit))
			case "int128":
				fmt.Printf(maybeFlagged("Bytes", isFlag, t.flagBit, "16"))
			case "int256":
				fmt.Printf(maybeFlagged("Bytes", isFlag, t.flagBit, "32"))
			case "string":
				fmt.Printf(maybeFlagged("String", isFlag, t.flagBit))
			case "double":
				fmt.Printf(maybeFlagged("Double", isFlag, t.flagBit))
			case "bytes":
				fmt.Printf(maybeFlagged("StringBytes", isFlag, t.flagBit))
			case "Vector<int>":
				fmt.Printf(maybeFlagged("VectorInt", isFlag, t.flagBit))
			case "Vector<long>":
				fmt.Printf(maybeFlagged("VectorLong", isFlag, t.flagBit))
			case "Vector<string>":
				fmt.Printf(maybeFlagged("VectorString", isFlag, t.flagBit))
			case "Vector<double>":
				fmt.Printf(maybeFlagged("VectorDouble", isFlag, t.flagBit))
			case "!X":
				fmt.Printf(maybeFlagged("Object", isFlag, t.flagBit))
			default:
				var inner string
				n, _ := fmt.Sscanf(strings.Title(t.typeName), "Vector<%s>", &inner)
				if n == 1 {
					fmt.Printf(maybeFlaggedExt("Vector", isFlag, t.flagBit) + ",\n")
				} else {
					fmt.Printf(maybeFlaggedExt("Object", isFlag, t.flagBit) + ",\n")
				}
			}
		}
		fmt.Printf("}\n\n")
	}

	fmt.Printf(`
	default:
		m.err = merry.Errorf("Unknown constructor: \u002508x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}`)
}
