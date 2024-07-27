package main

func main() {
	// 任务1
	go func() {
		Prepare1()
		Cook1()
		ServeDish1()
	}()
	// 任务2
	go func() {
		Prepare2()
		Cook2()
		ServeDish2()
	}()
	// 任务3
	go func() {
		Prepare3()
		Cook3()
		ServeDish3()
	}()

	// 必要的等待
}

func Prepare1()   {}
func Cook1()      {}
func ServeDish1() {}
func Prepare2()   {}
func Cook2()      {}
func ServeDish2() {}
func Prepare3()   {}
func Cook3()      {}
func ServeDish3() {}
