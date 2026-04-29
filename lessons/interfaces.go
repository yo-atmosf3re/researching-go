package lessons

import "researching-go/pkg/logger"

type Auto interface {
	StepOnGas()
}

type Lada struct{}
type BMW struct{}

func (BMW) StepOnGas() {
	logger.Ptc("step on gas BWM")
}
func (Lada) StepOnGas() {
	logger.Ptc("step on gas Lada")
}

func ride(auto Auto) {
	auto.StepOnGas()
}

func ExampleInterfaces() {
	ride(Lada{})
	ride(BMW{})
}

// interface can implement same methods for types, if methods is defined for these types as in example above
