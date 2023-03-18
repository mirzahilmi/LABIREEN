package services

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("here.env")
	if err != nil {
		log.Fatalln("Error loading env file")
	}

	c := coreapi.Client{}
	c.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeGopay,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "12345",
			GrossAmt: 200000,
		},
	}

	// 3. Request to Midtrans
	coreApiRes, _ := c.ChargeTransaction(chargeReq)
	fmt.Println("Response :", coreApiRes)
}
