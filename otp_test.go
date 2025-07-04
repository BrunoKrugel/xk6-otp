package otp

import (
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

func TestEmailClient(t *testing.T) {
	email := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_APP_PASSWORD")
	sender := os.Getenv("SENDER_EMAIL")

	t.Run("Test EmailClient", func(t *testing.T) {
		otp := Otp{}

		messages, _ := otp.LastOtpCodeBySender(email, password, sender)
		assert.NotEmpty(t, messages, "Expected to receive messages from the email client")
	})

	t.Run("Test EmailClient", func(t *testing.T) {
		otp := Otp{}

		messages, _ := otp.LastOtpCode(email, password, sender, "código")
		assert.NotEmpty(t, messages, "Expected to receive messages from the email client")
	})
}

func TestExtractOTP(t *testing.T) {
	t.Run("Should extract otp method", func(t *testing.T) {
		code := extractOTP("Your OTP is 123456")
		assert.Equal(t, code, "123456")
	})

	t.Run("Should return empty string if no OTP found", func(t *testing.T) {
		code := extractOTP("Your OTP is")
		assert.Equal(t, code, "")
	})
}
