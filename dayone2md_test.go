package dayone2md

import (
	"strings"
	"testing"
)

func TestTrimText(tt *testing.T) {
	escapeNewlines := func(s string) string {
		return strings.ReplaceAll(s, "\n", `\n`)
	}
	testCases := []struct {
		Alias         string
		Text          string
		ExpectedTitle string
		ExpectedText  string
	}{
		{
			Alias:         "simple",
			Text:          "title\nline1\nline2\n\n",
			ExpectedTitle: "title",
			ExpectedText:  "line1\nline2\n",
		},
		{
			Alias:         "simple2",
			Text:          "title\n\nline1\nline2\n\n",
			ExpectedTitle: "title",
			ExpectedText:  "line1\nline2\n",
		},
		{
			Alias:         "no-title",
			Text:          "line1\nline2\n\n",
			ExpectedTitle: "line1",
			ExpectedText:  "line2\n",
		},
		{
			Alias:         "extra-newlines",
			Text:          "\ntitle\n\nline1\nline2\n\n",
			ExpectedTitle: "title",
			ExpectedText:  "line1\nline2\n",
		},
		{
			Alias:         "extra-newlines2",
			Text:          "\n\ntitle\n\nline1\nline2\n\n",
			ExpectedTitle: "title",
			ExpectedText:  "line1\nline2\n",
		},
		{
			Alias:         "oneline",
			Text:          "line1\n",
			ExpectedTitle: "",
			ExpectedText:  "line1\n",
		},
		{
			Alias:         "oneline2",
			Text:          "line1",
			ExpectedTitle: "",
			ExpectedText:  "line1\n",
		},
	}
	for _, tCase := range testCases {
		tt.Run(tCase.Alias, func(t *testing.T) {
			body, title := trimText(tCase.Text)
			if got, want := title, tCase.ExpectedTitle; got != want {
				t.Errorf("bad title: %s, expected: %s", escapeNewlines(got), escapeNewlines(want))
			}
			if got, want := body, tCase.ExpectedText; got != want {
				t.Errorf("bad text: %s, expected: %s", escapeNewlines(got), escapeNewlines(want))
			}
		})
	}
}

func TestWikilinks(t *testing.T) {
	testCases := []struct {
		Alias        string
		Text         string
		Filenames    map[string]string
		ExpectedText string
	}{
		{
			Alias:        "simple",
			Text:         "test [weekend](dayone://view?entryId=8DB538436C2844E89E2D63A71E1EE884) test",
			Filenames:    map[string]string{"8DB538436C2844E89E2D63A71E1EE884": "xyz"},
			ExpectedText: "test [[xyz|weekend]] test",
		},
		{
			Alias:        "simple",
			Text:         "test [weekend](dayone://view?entryId=8DB538436C2844E89E2D63A71E1EE884) [weekend](dayone://view?entryId=8DB538436C2844E89E2D63A71E1EE884) test",
			Filenames:    map[string]string{"8DB538436C2844E89E2D63A71E1EE884": "xyz"},
			ExpectedText: "test [[xyz|weekend]] [[xyz|weekend]] test",
		},
	}
	for _, tCase := range testCases {
		t.Run(tCase.Alias, func(t *testing.T) {
			a := wikilinks(tCase.Text, tCase.Filenames)
			if got, want := a, tCase.ExpectedText; got != want {
				t.Errorf("bad text: %s, expected: %s", got, want)
			}
		})
	}
}
