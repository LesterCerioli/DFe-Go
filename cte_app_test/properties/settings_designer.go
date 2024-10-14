package properties

import (
	"sync"
)

type Settings struct {
	// Add any necessary fields here
}

var (
	defaultInstance *Settings
	once            sync.Once
)

func GetDefault() *Settings {
	once.Do(func() {
		defaultInstance = &Settings{}
	})
	return defaultInstance
}
