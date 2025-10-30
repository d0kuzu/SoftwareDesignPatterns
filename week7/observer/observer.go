package observer

// Observer interface defines the method that will be called upon notification.
type Observer interface {
	Update(string) // Method to receive state change notification
}
