package go_test_symbol_lib

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func step(name string) {
	fmt.Printf("\033[33m=== %s\033[39m\n", name)
}

func DebugShow() {
	DebugShow_a()
}

func DebugShow_a() {
	DebugShow_b()
}

func DebugShow_b() {
	fmt.Println("\033[32m###DebugShow###\033[39m\n")

	step("runtime.Caller(i)")

	for i := 0; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if ok == false {
			break
		}
		fn := runtime.FuncForPC(pc)
		file2, line2 := fn.FileLine(pc)

		fmt.Printf("- [%2d] 0x%x %60s:%-5d %60s:%-5d [%s]\n", i, pc, file, line, file2, line2, fn.Name())
	}

	step("runtime.Callers(...) + FuncForPc (discouraged)")

	var (
		pcs = make([]uintptr, 30)
	)
	pcs = pcs[:runtime.Callers(0, pcs)]

	for i, pc := range pcs {
		fn := runtime.FuncForPC(pc)
		file2, line2 := fn.FileLine(pc)
		fmt.Printf("- [%2d] 0x%x %60s:%-5d [%s]\n", i, pc, file2, line2, fn.Name())
	}

	step("runtime.Callers(...) + runtime.CallersFrame()")

	frames := runtime.CallersFrames(pcs)

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		fmt.Printf("- 0x%x %60s:%-5d %s %s\n", frame.PC, frame.File, frame.Line, frame.Function, frame.Func.Name())
	}

	step("debug.Stack() which uses runtime.Stack()")

	fmt.Printf("%s\n", debug.Stack())

	step("panic")
	
	panic("PANIC")
}
