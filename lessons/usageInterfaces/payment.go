package usageInterfaces

import (
	"fmt"
	"math/rand"
	"researching-go/pkg/logger"
)

type PaymentMethods interface {
	Pay(amount int) int
	Cancel() int
	GetId() int
}

type Payment struct {
	name       string
	id         int
	balance    int
	isCanceled bool
}

type PaymentInfo struct {
	payments map[int]PaymentMethods
}

func NewPayment(name string, balance int) *Payment {
	if name != "" {
		logger.Ptc("new payment created: ", name)
		return &Payment{
			name:       name,
			id:         rand.Intn(15),
			balance:    balance,
			isCanceled: false,
		}
	}
	logger.Ptc("name is empty, exit with occurred error")
	return nil
}

func NewPaymentInfo() *PaymentInfo {
	return &PaymentInfo{payments: make(map[int]PaymentMethods)}
}

func (payment *Payment) Pay(amount int) int {
	if payment.balance <= -1000 {
		payment.Cancel()
		logger.Ptc("balance is so small, pay is canceled: ", payment.name)
		return payment.id
	}
	payment.balance -= amount
	logger.Ptc("pay is completed, balance updated")
	return payment.id
}

func (payment *Payment) Cancel() int {

	if payment.isCanceled {
		logger.Ptc("operation already is canceled")
		return payment.id
	}
	payment.isCanceled = true
	logger.Ptc("operation is canceled")
	return payment.id
}

func (payment *Payment) GetId() int {
	return payment.id
}

func (info *PaymentInfo) AddInfo(payment PaymentMethods) {
	if p, ok := payment.(*Payment); ok {
		logger.Ptc("add payment info: ", p.name)
		info.payments[p.id] = payment
		return
	}
	logger.Ptc("add new info: error has occurred")
	return
}

func (info *PaymentInfo) Info(id int) PaymentMethods {
	if payment, ok := info.payments[id]; ok {
		logger.Ptc("data of payment by id: ", payment)
		return payment
	}
	logger.Ptc("payment not found")
	return nil
}

func (info *PaymentInfo) AllInfo() map[int]PaymentMethods {
	if len(info.payments) == 0 {
		logger.Ptc("no payments")
		return nil
	}
	logger.Ptc("all payments data: ", info.payments)
	return info.payments
}

var paymentInfo *PaymentInfo = NewPaymentInfo()

func PaymentTransaction(payment PaymentMethods) {
	if p, ok := payment.(*Payment); ok {
		logger.Ptc("current payment: ", p)
		logger.Ptc("new payment info is created: ", paymentInfo)

		p.Pay(11000)
		paymentInfo.AddInfo(p)
		paymentInfo.Info(p.id)

		pId := p.Pay(100)
		fmt.Printf("payment %d was denied, maybe trouble with balance", pId)

		return
	}
	logger.Ptc("error of payment, try again later")
	return
}

func PaymentExample() {
	PaymentTransaction(NewPayment("Bank", 10000))
	PaymentTransaction(NewPayment("Crypto", 200))
	PaymentTransaction(NewPayment("PayPal", 5))
	logger.Ptc("payment info, result :", paymentInfo)
	ps := ConvertMapToSlicePaymentInfo(paymentInfo)

	r1 := searchPaymentViaSlice(1, ps)
	r2 := searchPaymentViaSlice(2, ps)
	r3 := searchPaymentViaSlice(3, ps)
	r4 := searchPaymentViaSlice(4, ps)
	r5 := searchPaymentViaSlice(5, ps)

	logger.Ptc("founded slice-payment: ", r1)
	logger.Ptc("founded slice-payment: ", r2)
	logger.Ptc("founded slice-payment: ", r3)
	logger.Ptc("founded slice-payment: ", r4)
	logger.Ptc("founded slice-payment: ", r5)
}

func ConvertMapToSlicePaymentInfo(paymentInfo *PaymentInfo) []PaymentMethods {
	var paymentInfoSlice []PaymentMethods
	for _, p := range paymentInfo.payments {
		paymentInfoSlice = append(paymentInfoSlice, p)
	}
	logger.Ptc("payment info: ", paymentInfoSlice)
	return paymentInfoSlice
}

func searchPaymentViaSlice(id int, ps []PaymentMethods) PaymentMethods {
	for _, p := range ps {
		if p.GetId() == id {
			return p
		}
	}
	logger.Ptc("slice: payment not found")
	return nil
}
