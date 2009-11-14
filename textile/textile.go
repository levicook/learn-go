package textile

import "regexp"
import "strings"

type phraseModifier struct {
	sep	string;
	el	string;
	re	*regexp.Regexp;
}

var phraseModifiers = []phraseModifier{
	phraseModifier{"*", "strong", regexp.MustCompile("\\*[a-z]*\\*")},
	phraseModifier{"+", "ins", regexp.MustCompile("\\+[^+]*\\+")},
	phraseModifier{"-", "del", regexp.MustCompile("-[^\\-]*-")},
	phraseModifier{"_", "em", regexp.MustCompile("_[a-z]*_")},
}

func TextileToHtml(input string) (output string, ok bool, errtok string) {
	output += "<p>";
	for _, pm := range phraseModifiers {
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
	output += "</p>";
	return output, true, "";
}
