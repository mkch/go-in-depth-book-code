package main

func main() {
	ForVar()
	ForSharedVar()
}

func ForVar() {
	var prints []func()
	for i := 0; i < 5; i++ {
		prints = append(prints, func() { println(i) })
		i++
	}
	for _, p := range prints {
		p()
	}
}

func ForSharedVar() {
	var prints []func()
	var i int
	for ; i < 5; i++ {
		prints = append(prints, func() { println(i) })
		i++
	}
	for _, p := range prints {
		p()
	}
}
