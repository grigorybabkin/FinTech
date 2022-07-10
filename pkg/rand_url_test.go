package pkg

import (
	"regexp"
	"testing"
)

func TestRandShortURL(t *testing.T) {

	got := RandShortURL(10)

	gotProps := TestRandURL{
		Length:        len(got),
		LetterUpCase:  regexp.MustCompile(`[A-Z]`).MatchString(got),
		LetterLowCase: regexp.MustCompile(`[a-z]`).MatchString(got),
		Digit:         regexp.MustCompile(`[0-9]`).MatchString(got),
		Special:       regexp.MustCompile(`[_]`).MatchString(got),
	}

	if gotProps.Length != want.Length {
		t.Errorf("\ngot String: %s with Length: %d, but wanted: Length: %d",
			got, gotProps.Length, want.Length)
	}

	if gotProps.LetterUpCase != want.LetterUpCase {
		t.Errorf("\ngot String: %s with Property: LetterUpcase: %t, but wanted: "+
			"LetterUpcase: %t", got, gotProps.LetterUpCase, want.LetterUpCase)
	}

	if gotProps.LetterLowCase != want.LetterLowCase {
		t.Errorf("\ngot String: %s with Property: LetterLowCase: %t, but wanted: "+
			"LetterLowCase: %t", got, gotProps.LetterLowCase, want.LetterLowCase)
	}

	if gotProps.Digit != want.Digit {
		t.Errorf("\ngot String: %s with Property: Digit: %t, but wanted: Digit: %t",
			got, gotProps.Digit, want.Digit)
	}

	if gotProps.Special != want.Special {
		t.Errorf("\ngot String: %s with Properties: Special: %t, but wanted: Special: %t",
			got, gotProps.Special, want.Special)
	}
}
