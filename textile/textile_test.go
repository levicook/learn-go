package textile

import "io"
import "json"
import "os"
import "strings"
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
		check(t, ok, errtok, input, expected, observed, "simplePhraseModifiers");
	}
}

func TestRedClothFixtures(t *testing.T) {
	dirName := "textile_test";
	dirEntries, err := io.ReadDir(dirName);
	doTheAngryErrorThing(t, err);
	for _, de := range dirEntries {
		if strings.HasSuffix(de.Name, "basic.json") {
			bytes, err := io.ReadFile(dirName + "/" + de.Name);
			doTheAngryErrorThing(t, err);
			s := string(bytes);
			json, ok, errtok := json.StringToJson(s);
			if !ok {
				t.Errorf("StringToJson(%#q) => error near %v", s, errtok)
			}
			for i := 0; i < json.Len(); i++ {
				testCase := json.Elem(i);
				name := testCase.Get("name").String() + " (" + de.Name + ")";
				expected := testCase.Get("html").String();
				input := testCase.Get("in").String();
				observed, ok, errtok := TextileToHtml(input);
				check(t, ok, errtok, input, expected, observed, name);
			}
		}
	}
}

func doTheAngryErrorThing(t *testing.T, err os.Error) {
	if err != nil {
		t.Errorf(err.String())
	}
}

//  internal utilities....
func check(t *testing.T, ok bool, errtok string, input string, expected string, observed string, name string) {
	if !ok {
		t.Fatalf("TextileToHtml failed near %s", errtok)
	}
	if expected != observed {
		//t.Errorf("\nTest: %s\nInput: %s\nExpected: %s\nObserved: %s", name, input, expected, observed)
		t.Fatalf("\nTest: %s\nInput: %s\nExpected: %s\nObserved: %s", name, input, expected, observed)
	}
}
