package mtproto

import (
	"fmt"
	"github.com/ansel1/merry"
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
	case TL_msg_container:
		names := make([]string, len(x.Items))
		for i, item := range x.Items {
			names[i] = TLName(item)
		}
		text = TLName(x) + " -> [" + strings.Join(names, ", ") + "]"
	case TL_rpc_result:
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

type Logger struct{}

func (l Logger) Error(err error, msg string) {
	logrus.Errorf(msg + ":\n" + merry.Details(err))
}

func (l Logger) Warn(msg string, args ...interface{}) {
	logrus.Warnf(msg, args...)
}

func (l Logger) Info(msg string, args ...interface{}) {
	logrus.Infof(msg, args...)
}

func (l Logger) Debug(msg string, args ...interface{}) {
	logrus.Debugf("\033[90m" + msg + "\033[0m", args...)
}

func (l Logger) Message(isIncoming bool, message TL, id int64) {
	logrus.Debugf("\033[90m" + StringifyMessage(isIncoming, message, id) + "\033[0m")
}
