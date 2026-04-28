package lessons

import "researching-go/pkg/logger"

type User struct {
	Name        string  // ""
	Age         int     // 0
	PhoneNumber string  // ""
	IsClose     bool    // false
	Rating      float64 // 0
}

func (user User) data() {
	logger.Pt("user data: ", user)
}

func newUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64) *User {
	if name == "" {
		logger.Pt("new: name is empty. Create new user is stopped!")
		return &User{}
	}
	if age > 150 || age < 0 {
		logger.Pt("new: age is not correctly because it's less than 0 or more 150. Create new user is stopped!")
		return &User{}
	}
	if phoneNumber == "" {
		logger.Pt("new: phone number i empty. Create new user is stopped!")
		return &User{}
	}
	if rating > 10.0 || rating < 0.0 {
		logger.Pt("new: rating is not correctly because it's less than 0 or more 10.0. Create new user is stopped!")
		return &User{}
	}
	return &User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		IsClose:     isClose,
		Rating:      rating,
	}
}

func (user *User) setName(name string) {
	if name != "" {
		user.Name = name
		logger.Pt("set name: ", user.Name)
	} else {
		logger.Pt("set name: name is invalid")
	}
}

func (user *User) setAge(age int) {
	if age <= 0 || age > 150 {
		logger.Pt("set age: age is not valid")
		return
	}
	logger.Pt("set age: ", user.Age)
	user.Age = age
}

func (user *User) setPhoneNumber(phoneNumber string) {
	if phoneNumber == "" || len(phoneNumber) != 12 {
		logger.Pt("set phone: phone number is invalid")
	} else {
		logger.Pt("set phone: ", phoneNumber)
		user.PhoneNumber = phoneNumber
	}
}

func (user *User) setIsClose(isClose bool) {
	if user.IsClose == isClose {
		logger.Pt("set isClose: operation is already completed")
		return
	}
	logger.Pt("set isClose: ", user.IsClose)
	user.IsClose = isClose
}

func (user *User) setRating(rating float64) {
	if rating <= 10.0 && rating >= 0.0 {
		user.Rating = rating
		logger.Pt("user set rating: ", user.Rating)
	} else {
		logger.Pt("user set rating: sum rating and user.Rating equal more than 10 or less than 0")
	}
}

func StructuresExample() {
	user := newUser("John", 151, "+78002564411", false, 7.7)

	user.data()

	user.setRating(10)
	user.setRating(11)

	user.data()

	user.setName("Ann")
	user.setAge(26)
	user.setPhoneNumber("+39997778844")
	user.setIsClose(true)
	user.setRating(1.2)

	user.data()
}
