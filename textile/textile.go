package textile

import "regexp"
import "strings"

func TextileToHtml(input string) (output string, ok bool, errtok string) {
	output += "<p>";
	for _, pm := range phraseModifiers {
		output += pm.translate(input)
	}
	output += "</p>";
	return output, true, "";
}

func split(s string, i int, j int) (before string, match string, after string) {
	return before, match, after
}

// Phrase Modifiers
// ----------------

type phraseModifier struct {
	sep	string;
	el	string;
	re	*regexp.Regexp;
}

var phraseModifiers = []phraseModifier{
	phraseModifier{"*", "strong", regexp.MustCompile(`(\*[^ *]*\*)`)},
	phraseModifier{"+", "ins", regexp.MustCompile(`\+[^+]*\+`)},
	phraseModifier{"-", "del", regexp.MustCompile(`-[^\-]*-`)},
	phraseModifier{"_", "em", regexp.MustCompile(`_[a-z]*_`)},
}

func (pm *phraseModifier) translate(input string) (output string) {
	//slog("\nInput ", input);
	a := pm.re.ExecuteString(input);	//print("a");print(a);println();
	if len(a) > 0 {
		i, j := a[0], a[1];
		before, match, after := input[0:i], input[i:j], input[j:len(input)];
		//slog("Before", before); slog("Match ", match); slog("After ", after);
		output += before;
		for _, s := range strings.Split(match, pm.sep, 0) {
			if len(s) > 0 {
				output += "<" + pm.el + ">";
				output += s;
				output += "</" + pm.el + ">";
			}
		}
		if len(after) > 0 {
			//slog("Recurs", after);
			output += pm.translate(after)
		}
	}
	return output;
}


// slog: Stupid Logger. I need to learn more about Go's IO
func slog(lbl string, msg string) {
	print(lbl);
	print(": ^");
	print(msg);
	print("$");
	println();
}
