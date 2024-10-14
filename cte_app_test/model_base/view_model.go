package modelbase

import (
	"sync"
)

// PropertyChangedEventHandler is a function type for property change event handlers
type PropertyChangedEventHandler func(sender interface{}, propertyName string)

// ViewModel is a base struct for view models
type ViewModel struct {
	propertyChanged []PropertyChangedEventHandler
	mu              sync.RWMutex
}

// AddPropertyChangedHandler adds a new property changed event handler
func (vm *ViewModel) AddPropertyChangedHandler(handler PropertyChangedEventHandler) {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	vm.propertyChanged = append(vm.propertyChanged, handler)
}

// RemovePropertyChangedHandler removes a property changed event handler
func (vm *ViewModel) RemovePropertyChangedHandler(handler PropertyChangedEventHandler) {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	for i, h := range vm.propertyChanged {
		if &h == &handler {
			vm.propertyChanged = append(vm.propertyChanged[:i], vm.propertyChanged[i+1:]...)
			break
		}
	}
}

// OnPropertyChanged notifies listeners that a property value has changed
func (vm *ViewModel) OnPropertyChanged(propertyName string) {
	vm.mu.RLock()
	defer vm.mu.RUnlock()
	for _, handler := range vm.propertyChanged {
		handler(vm, propertyName)
	}
}
