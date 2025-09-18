package week2

import "fmt"

// Notification interface
type Notification interface {
	Send(message string)
}

// Different types of notifications (Emal, SMS, Push)

type EmailNotification struct{}

func (e *EmailNotification) Send(message string) {
	fmt.Println("Sending EMAIL with message:", message)
}

type SMSNotification struct{}

func (s *SMSNotification) Send(message string) {
	fmt.Println("Sending SMS with message:", message)
}

type PushNotification struct{}

func (p *PushNotification) Send(message string) {
	fmt.Println("Sending PUSH with message:", message)
}
