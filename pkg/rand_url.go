package pkg

import (
	"math/rand"
	"time"
)

func RandShortURL(n int) string {

	rand.Seed(time.Now().UnixNano())

	url := make([]byte, n)
	url[0] = lettersUpCase[rand.Intn(len(lettersUpCase))]
	url[1] = lettersLowCase[rand.Intn(len(lettersLowCase))]
	url[2] = digits[rand.Intn(len(digits))]
	url[3] = special[0]

	for i := 4; i < n; i++ {
		url[i] = allSymbs[rand.Intn(len(allSymbs))]
	}
	rand.Shuffle(len(url), func(i, j int) {
		url[i], url[j] = url[j], url[i]
	})

	return string(url)
}
