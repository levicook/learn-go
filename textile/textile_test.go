package textile

import "testing"

var simplePhraseModifiers = map[string]string{
	"*never*": "<p><strong>never</strong></p>",
	"*never* * ever* *sometimes*": "<p><strong>never</strong> * ever* <strong>sometimes</strong></p>",
	"_believe_": "<p><em>believe</em></p>",
	"-Al Gore-": "<p><del>Al Gore</del></p>",
	"+George W. Bush+": "<p><ins>George W. Bush</ins></p>",
}

func TestBasicPhraseModifiers(t *testing.T) {
	for input, expected := range simplePhraseModifiers {
		observed, ok, errtok := TextileToHtml(input);
		check(t, ok, errtok, expected, observed);
	}
}


//  internal utilities....
func check(t *testing.T, ok bool, errtok string, expected string, observed string) {
	if !ok {
		t.Fatalf("TextileToHtml failed near %s", errtok)
	}
	if expected != observed {
		t.Errorf("Expected: %s\nObserved: %s", expected, observed)
	}
}
