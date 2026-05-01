package usageInterfaces

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

type SimpleCrypto struct{}
type SimplePayPal struct{}
type SimpleBank struct{}

type SimplePayment interface {
	makePaymentAmount()
	cancelPayment()
	givePaymentInfo()
}

func (SimpleCrypto) makePaymentAmount() {
	logger.Ptc("making payment amount for crypto")
}

func (SimplePayPal) makePaymentAmount() {
	logger.Ptc("making payment amount for payPal")
}

func (SimpleBank) makePaymentAmount() {
	logger.Ptc("making payment amount for bank")
}

func (SimpleCrypto) cancelPayment() {
	logger.Ptc("canceling payment for crypto")
}

func (SimplePayPal) cancelPayment() {
	logger.Ptc("canceling payment for payPal")
}

func (SimpleBank) cancelPayment() {
	logger.Ptc("canceling payment for bank")
}

func (SimpleCrypto) givePaymentInfo() {
	logger.Ptc("give information about payment for crypto")
}

func (SimplePayPal) givePaymentInfo() {
	logger.Ptc("give information about payment for payPal")
}

func (SimpleBank) givePaymentInfo() {
	logger.Ptc("give information about payment for bank")
}

func SimplePaymentOperation(payment SimplePayment) {
	logger.Ptc("payment operation is started")
	payment.givePaymentInfo()
	payment.makePaymentAmount()
	logger.Ptc("1")
	logger.Ptc("2")
	logger.Ptc("3")
	logger.Ptc("occurred some error")
	payment.cancelPayment()
}
