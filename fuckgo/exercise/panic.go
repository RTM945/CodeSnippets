package main

import (
	"fmt"
	"runtime"
	"strings"
)

func identifyPanic(panic interface{}) string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	switch {
	case name != "":
		return fmt.Sprintf("%v:%v:%s", name, line, panic)
	case file != "":
		return fmt.Sprintf("%v:%v:%s", file, line, panic)
	}

	return fmt.Sprintf("pc:%x:%s", pc, panic)
}

func recoverPanic() {
	r := recover()
	if r == nil {
		return
	}
	fmt.Println(identifyPanic(r))
}

func createPanic() {
	var s *string
	fmt.Println(*s)
}

func main() {
	defer recoverPanic()
	createPanic()
}
