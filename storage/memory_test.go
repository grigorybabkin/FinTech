package storage

import (
	"testing"
)

func TestURLMemory_GetURLMem(t *testing.T) {

	var urlMem = NewURLMemory()

	for _, test := range urlsTest {
		urlMem.AddURLsMem(test.shortURL, test.fullURL, test.shortURL)
		var fullURLResponse, found = urlMem.GetURLMem(test.shortURL)
		if !found {
			t.Errorf("%v expected to be found", test.fullURL)
		} else if fullURLResponse != test.fullURL {
			t.Errorf("got url:  %v, but wanted %v", fullURLResponse, test.fullURL)
		}
	}
}
