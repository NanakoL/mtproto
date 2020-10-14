package mtproto

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

func TLName(obj interface{}) string {
	return reflect.TypeOf(obj).Name()
}

func StringifyMessage(isIncoming bool, msg TL, id int64) string {
	var text string
	switch x := msg.(type) {
	case MsgContainer:
		names := make([]string, len(x.Items))
		for i, item := range x.Items {
			names[i] = TLName(item)
		}
		text = TLName(x) + " -> [" + strings.Join(names, ", ") + "]"
	case RpcResult:
		text = TLName(x) + " -> " + TLName(x.obj)
	default:
		text = TLName(x)
	}
	if isIncoming {
		text = ">>> " + text
	} else {
		text = "<<< " + text + fmt.Sprintf(" (#%d)", id)
	}
	return text
}

type Logger struct {
	*logrus.Logger
}

func (l *Logger) Error(err error, msg string) {
	l.Errorf(msg + ":\n" + err.Error())
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.Warnf(msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.Infof(msg, args...)
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Debugf("\033[90m"+msg+"\033[0m", args...)
}

func (l *Logger) Message(isIncoming bool, message TL, id int64) {
	l.Debugf("\033[90m" + StringifyMessage(isIncoming, message, id) + "\033[0m")
}
