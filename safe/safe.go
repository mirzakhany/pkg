package safe

import (
	"runtime/debug"
	"time"

	"github.com/sirupsen/logrus"
)

// Try try to run a function with timeout
func Try(a func() error, max time.Duration, extra ...interface{}) {
	x, y := 0, 1
	for {
		err := actual(a, extra...)
		if err == nil {
			return
		}
		t := time.Duration(x) * time.Second
		if t < max {
			x, y = y, x+y
		}
		time.Sleep(t)
	}
}

func actual(a func() error, extra ...interface{}) (e error) {
	defer func() {
		err := recover()
		if err != nil {
			stack := debug.Stack()
			logrus.Warn(string(stack))
			call(extra...)
		}
	}()
	e = a()
	return
}

// Routine run a function and return error stack if need
func Routine(f func(), extra ...interface{}) {
	defer func() {
		e := recover()
		if e != nil {
			stack := debug.Stack()
			logrus.Warn(string(stack))
			call(extra...)
			// TODO : add log to redmine or slack
		}
	}()
	f()
}

func call(extra ...interface{}) {
	for i := range extra {
		if fn, ok := extra[i].(func()); ok {
			fn()
		}
	}
}
