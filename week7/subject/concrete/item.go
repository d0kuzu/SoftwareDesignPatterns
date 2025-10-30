package concrete

import . "SoftwareDesignPatterns/week7/observer"

// Item is the concrete Subject.
type Item struct {
	name         string
	inStock      bool
	observerList []Observer
}

func NewItem(name string) *Item {
	return &Item{
		name:         name,
		inStock:      false,
		observerList: make([]Observer, 0),
	}
}

// Implement the Subject interface methods:

func (i *Item) RegisterObserver(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) RemoveObserver(o Observer) {
	// Simple observer removal
	for idx, obs := range i.observerList {
		if obs == o {
			i.observerList = append(i.observerList[:idx], i.observerList[idx+1:]...)
			break
		}
	}
}

func (i *Item) NotifyObservers(message string) {
	for _, observer := range i.observerList {
		observer.Update(i.name + ": " + message)
	}
}

// UpdateAvailability changes the item's state and notifies observers if the state changed.
func (i *Item) UpdateAvailability(inStock bool) {
	if i.inStock != inStock {
		i.inStock = inStock
		if inStock {
			i.NotifyObservers("Now in stock!")
		} else {
			i.NotifyObservers("Out of stock.")
		}
	}
}
