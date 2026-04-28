package lessons

type Worker struct {
	fullName string
	place    string
	age      int
	isLive   bool
	money    float64
}

func NewWorker(fullName string, place string, age int, isLive bool, money float64) *Worker {
	if fullName == "" || place == "" || age < 0 || age > 100 || money < 0.0 {
		return &Worker{}
	}
	return &Worker{
		fullName: fullName,
		place:    place,
		age:      age,
		money:    money,
		isLive:   isLive,
	}
}
