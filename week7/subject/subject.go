package subject

import "SoftwareDesignPatterns/week7/observer"

// Subject interface defines methods for managing observers.
type Subject interface {
	RegisterObserver(o observer.Observer) // Add an observer
	RemoveObserver(o observer.Observer)   // Remove an observer
	NotifyObservers(message string)       // Notify all observers
}
