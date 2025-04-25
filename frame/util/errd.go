package util

import (
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

func Recover() {
	if e := recover(); e != nil {
		stack := debug.Stack()
		logrus.WithFields(logrus.Fields{
			"err":   e,
			"stack": string(stack),
		}).Error("Recover")
	}
}

func SafeGo(f func()) {
	if f != nil {
		go func() {
			defer Recover()
			f()
		}()
	}
}
