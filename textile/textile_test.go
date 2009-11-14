package textile

import "testing"

func TestBasicPhraseModifiers(t *testing.T) {
	tests := map[string]string{
		"*never*": "<strong>never</strong>",
		"_believe_": "<em>believe</em>",
		"-Al Gore-": "<del>Al Gore</del>",
		"+George W. Bush+": "<ins>George W. Bush</ins>",
	};
	for input, expected := range tests {
		observed, ok, errtok := TextileToHtml(input);
		if !ok {
			t.Fatalf("TextileToHtml failed near %s", errtok)
		}
		if observed != expected {
			t.Errorf("[%s] <> [%s]", observed, expected)
		}
	}
}
