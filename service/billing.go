package service

import (
	"awesomeProject/skillbox/StatusPage/helpers"
	"math"
	"strings"

	log "github.com/sirupsen/logrus"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func NewBillingData(b []bool) BillingData {
	bd := BillingData{}
	bd.CreateCustomer = b[0]
	bd.Purchase = b[1]
	bd.Payout = b[2]
	bd.Recurring = b[3]
	bd.FraudControl = b[4]
	bd.CheckoutPage = b[5]
	return bd
}

func Billing(c chan DataBilling) chan DataBilling {
	l := DataBilling{}
	billingDataCSV := "./simulator/billing.data"
	s, err := helpers.CsvInString(billingDataCSV)
	if err != nil {
		l.err = err
		c <- l
		return c
	}
	rns := strings.Split(s, "")
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	var k float64
	var b []bool
	for i := 0; i < 6; i++ {
		if rns[i] == "1" {
			x := math.Pow(float64(2), float64(i))
			k += x
			b = append(b, true)
		} else {
			b = append(b, false)
		}
	}
	l.bil = NewBillingData(b)
	log.Info("Получены данные billing")
	c <- l
	return c
}
