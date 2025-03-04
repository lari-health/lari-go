package sms

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func DummyMessage(to string, body string) {
	fmt.Println("Sending message\n", body, "\nto\t", to)
}

func SendMessage(to string, body string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	if len(cwd) == 0 {
		cwd = "root/lari-go"
	}

	// Construct the full path to .env
	envPath := cwd + "/.env"

	fmt.Println("Loading environment variables...")
	err = godotenv.Load(envPath)

	if err != nil {
		log.Fatal("Error loading .env")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TW_ACC_SID"),
		Password: os.Getenv("TW_AUTH"),
	})

	fmt.Println("Sending SMS Message...")

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(os.Getenv("TW_NUM"))
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
