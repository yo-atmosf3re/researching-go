package practice

import "researching-go/pkg/logger"

type user struct {
	Name string
}

func changeName(u user) {
	u.Name = "Alex"
}

func ex1() {
	user := user{Name: "Dave"}
	changeName(user)
	logger.Ptc("changed name", user.Name)
}

func changeAge(data map[string]int) {
	data["Ivan"] = 30
}

func ex2() {
	ages := map[string]int{"Ivan": 20}
	changeAge(ages)
	logger.Ptc("changed age", ages)
}

func appendNumber(numbers []int) []int {
	numbers[0] = 999
	return append(numbers, 4)
}

func ex3() {
	mySlice := []int{1, 2, 3}
	newSlice := appendNumber(mySlice)
	logger.Ptc("old slice", mySlice)
	logger.Ptc("new slice", newSlice)
}

func TransferValues() {
	ex1()
	ex2()
	ex3()
}
