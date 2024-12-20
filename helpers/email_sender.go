package helpers

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendRegistrationEmail(email, name string) error {
	// Email configuration
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "noreply@rehan.id")
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Welcome to Our Rental Car Service")

	// Set HTML body
	htmlBody := fmt.Sprintf(WelcomeEmailTemplate, name, email)
	mailer.SetBody("text/html", htmlBody)

	// Add plain text alternative
	plainBody := fmt.Sprintf(WelcomeEmailPlainTemplate, name)
	mailer.AddAlternative("text/plain", plainBody)

	// SMTP configuration - use environment variables for sensitive info
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	dialer := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
	)

	// Send the email
	if err := dialer.DialAndSend(mailer); err != nil {
		return fmt.Errorf("failed to send registration email: %v", err)
	}

	return nil
}
