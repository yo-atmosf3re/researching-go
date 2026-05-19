package practice

import "researching-go/pkg/logger"

func l1() {
	i := 20
	f := 2.222

	i = int(f)
	logger.Ptc(i)
	logger.Ptc(f)
}

func l2() {
	const value = 10
	i := 20
	f := 2.222
	i = value
	f = float64(value)
	logger.Ptc(i, f)
}

func l3() {
	var b byte
	var smallI int32
	var bigI uint64

	b = 255
	smallI = 2147483647
	bigI = 18446744073709551615

	b += 1      // 255 + 1 = 0 - integer overflow
	smallI += 1 // -21474836489, same
	bigI += 1   // 0, same

	logger.Ptc(b, smallI, bigI)
}

func IntegerExample() {
	l1()
	l2()
	l3()
}
