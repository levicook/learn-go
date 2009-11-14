package textile

import "testing"

func TestBasicPhraseModifiers(t *testing.T) {
	tests := map[string]string{
		"*never*": "<p><strong>never</strong></p>",
		"_believe_": "<p><em>believe</em></p>",
		"-Al Gore-": "<p><del>Al Gore</del></p>",
		"+George W. Bush+": "<p><ins>George W. Bush</ins></p>",
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
