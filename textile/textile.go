package textile

import "regexp"
import "strings"

type phraseModifier struct {
	sep	string;
	re	*regexp.Regexp;
	el	string;
}

var phraseModifiers = []phraseModifier{
	phraseModifier{"*", regexp.MustCompile("\\*[a-z]*\\*"), "strong"},
	phraseModifier{"+", regexp.MustCompile("\\+[^+]*\\+"), "ins"},
	phraseModifier{"-", regexp.MustCompile("-[^\\-]*-"), "del"},
	phraseModifier{"_", regexp.MustCompile("_[a-z]*_"), "em"},
}

func TextileToHtml(input string) (output string, ok bool, errtok string) {

	for i := 0; i < len(phraseModifiers); i++ {
		pm := phraseModifiers[i];
		if pm.re.MatchString(input) {
			for _, s := range strings.Split(input, pm.sep, 0) {
				if len(s) > 0 {
					output += "<" + pm.el + ">";
					output += s;
					output += "</" + pm.el + ">";
				}
			}
		}
	}

	return output, true, "";
}
