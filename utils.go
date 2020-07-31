package mtproto

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/ansel1/merry"
)

// https://core.telegram.org/api/errors
func IsError(obj TL, message string) bool {
	err, ok := obj.(RpcError)
	return ok && err.ErrorMessage == message
}

func IsErrorType(obj TL, code int32) bool {
	err, ok := obj.(RpcError)
	return ok && err.ErrorCode == code
}

const FloodWaitErrPrefix = "FLOOD_WAIT_"

func IsFloodError(obj TL) (time.Duration, bool) {
	if err, ok := obj.(RpcError); ok && strings.HasPrefix(err.ErrorMessage, FloodWaitErrPrefix) {
		secs, _ := strconv.ParseInt(err.ErrorMessage[len(FloodWaitErrPrefix):], 10, 64)
		if secs <= 0 {
			secs = 1
		}
		return time.Duration(secs) * time.Second, true
	}
	return 0, false
}

func IsClosedConnErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "use of closed network connection")
}

func Sprint(obj TL) string {
	return fmt.Sprintf("%#v", obj)
}

func UnexpectedTL(name string, obj TL) string {
	return fmt.Sprint("unexpected " + name + ": " + Sprint(obj))
}

func WrongRespError(obj TL) error {
	_type := "response"
	if _, ok := obj.(RpcError); ok {
		_type = "error"
	}
	return merry.Errorf(UnexpectedTL(_type, obj)).WithStackSkipping(1)
}

func SliceConvert(slice interface{}, newSliceType reflect.Type) interface{} {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Slice called with non-slice value of type %T", slice))
		return reflect.New(newSliceType).Elem().Interface()
	}
	if newSliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Slice called with non-slice type of type %T", newSliceType))
	}
	newSlice := reflect.New(newSliceType)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hdr.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
}

func SliceToTL(slice interface{}) []TL {
	return SliceConvert(slice, reflect.TypeOf([]TL{})).([]TL)
}

func SliceToTLStable(slice interface{}) []TL {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Slice called with non-slice value of type %T", slice))
	}
	ret := make([]TL, s.Len())
	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface().(TL)
	}
	return ret
}