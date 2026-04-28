package logger

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
)

// Pt - Используется для вывода в терминал;
func Pt(value ...any) (n int, err error) {
	return fmt.Println(value...)
}

func Ptc(value ...any) (n int, err error) {
	scheme := pp.ColorScheme{
		Bool:    pp.Cyan,
		Integer: pp.Magenta,
		Float:   pp.White,
		String:  pp.Green,
	}
	pp.Default.SetColorScheme(scheme)
	return pp.Println(value...)
}
