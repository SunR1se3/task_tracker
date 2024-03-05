package helper

import "html/template"

func GetFunMap() template.FuncMap {
	return map[string]interface{}{
		"N":   Iterator,
		"mod": Mod,
	}
}

func Iterator(start, end int) (stream chan int) {
	stream = make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			stream <- i
		}
		close(stream)
	}()
	return
}

func Mod(a, b int) int {
	return a % b
}
