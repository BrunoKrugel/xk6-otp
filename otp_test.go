package otp_test

import (
	"fmt"
	"os"
	"testing"

	otp "github.com/BrunoKrugel/xk6-otp"
	_ "github.com/joho/godotenv/autoload"
)

func TestEmailClient(t *testing.T) {
	email := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_APP_PASSWORD")
	sender := os.Getenv("SENDER_EMAIL")

	t.Run("Test EmailClient", func(t *testing.T) {
		otp := otp.Otp{}

		messages, _ := otp.LastOtpCodeBySender(email, password, sender)
		fmt.Println(messages)
	})
}
