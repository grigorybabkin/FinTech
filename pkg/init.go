package pkg

var (
	lettersUpCase  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lettersLowCase = "abcdefghijklmnopqrstuvwxyz"
	digits         = "1234567890"
	special        = "_"
	allSymbs       = lettersUpCase + lettersLowCase + digits + special
)

type TestRandURL struct {
	Length        int
	LetterUpCase  bool
	LetterLowCase bool
	Digit         bool
	Special       bool
}

var want = TestRandURL{
	Length:        10,
	LetterUpCase:  true,
	LetterLowCase: true,
	Digit:         true,
	Special:       true,
}
