package storage

import (
	"sync"
)

type URLMemory struct {
	sync.RWMutex
	urlStorageItems map[string]URL
}

type URL struct {
	FullURL  string
	ShortURL string
}

func NewURLMemory() *URLMemory {

	URLMemoryItems := make(map[string]URL)

	urlMemory := URLMemory{
		urlStorageItems: URLMemoryItems,
	}

	return &urlMemory
}

func (um *URLMemory) AddURLsMem(key string, fullURL, shortURL string) {

	um.Lock()

	defer um.Unlock()

	um.urlStorageItems[key] = URL{
		FullURL:  fullURL,
		ShortURL: shortURL,
	}
}

func (um *URLMemory) GetURLMem(key string) (string, bool) {

	um.RLock()

	defer um.RUnlock()

	item, found := um.urlStorageItems[key]

	// url not found
	if !found {
		return "", false
	}

	return item.FullURL, true
}
