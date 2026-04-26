package logger

import "fmt"

// Pt - Используется для вывода в терминал;
func Pt(value ...any) (n int, err error) {
	return fmt.Println(value...)
}
