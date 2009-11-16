package textile

import "regexp"
import "strings"

func TextileToHtml(input string) (output string, ok bool, errtok string) {
	lines := strings.Split(input, "\n", 0);
	for i, line := range lines {
		if blankLine.MatchString(strings.TrimSpace(line)) {
			continue
		}
		line = "<p>" + line;
		for _, pm := range phraseModifiers {
			line = pm.translate(line)
		}
		line += "</p>";
		if i > 0 {
			output += "\n"
		}
		output += line;
	}
	return output, true, "";
}

// Helpers?
// --------
var blankLine = regexp.MustCompile(`^$`)

// Phrase Modifiers
// ----------------

type phraseModifier struct {
	sep	string;
	el	string;
	re	*regexp.Regexp;
}

var phraseModifiers = []phraseModifier{
	phraseModifier{"*", "strong", regexp.MustCompile(`(\*[^ ][^*]*\*)`)},
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
		if pm.re.MatchString(after) {
			//slog("Recurs", after);
			output += pm.translate(after)
		} else {
			output += after
		}
		return output;
	}
	return input;
}


// slog: Stupid Logger. I need to learn more about Go's IO
func slog(lbl string, msg string) {
	print(lbl);
	print(": ^");
	print(msg);
	print("$");
	println();
}
