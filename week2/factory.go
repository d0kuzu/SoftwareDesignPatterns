package week2

// NotificationFactory interface
type NotificationFactory interface {
	CreateNotification() Notification
}

// Concrete factories (for Email SMS, Push)

type EmailFactory struct{}

func (f *EmailFactory) CreateNotification() Notification {
	return &EmailNotification{}
}

type SMSFactory struct{}

func (f *SMSFactory) CreateNotification() Notification {
	return &SMSNotification{}
}

type PushFactory struct{}

func (f *PushFactory) CreateNotification() Notification {
	return &PushNotification{}
}
