package lessons

import "fmt"

func pt(value ...any) (n int, err error) {
	return fmt.Println(value...)
}

func ForExample() {
	index := 0
	for ; index <= 5; index++ {
		pt("first", index)
	}
	pt(index)
}
