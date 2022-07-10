package storage

import "testing"

func TestAddURLsDB(t *testing.T) {

	for _, test := range urlsTest {
		//if !RowExistsDB("SELECT * FROM urls WHERE full_url=$1", test.fullURL) {
		AddURLsDB(test.fullURL, test.shortURL)
		//}
	}

	for _, test := range urlsTest {
		if !RowExistsDB("SELECT * FROM urls WHERE full_url=$1", test.fullURL) {
			t.Errorf("full url: %v and short url: %v weren't saved in db", test.fullURL, test.shortURL)
		}
	}
}
