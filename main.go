package main

import . "SoftwareDesignPatterns/week2"

func main() {
	// init notification factory
	var factory NotificationFactory

	// testing Email notification
	factory = &EmailFactory{}
	notification := factory.CreateNotification()
	notification.Send("Hello via Email!")

	// Testing SMS notification
	factory = &SMSFactory{}
	notification = factory.CreateNotification()
	notification.Send("Hello via SMS!")

	// Testing Push notification
	factory = &PushFactory{}
	notification = factory.CreateNotification()
	notification.Send("Hello via Push!")
}
