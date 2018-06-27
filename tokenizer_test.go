package sooslib

import "testing"

func TestTokenizer(t *testing.T) {
	want := "eac1fbb9986eacd640fdb306d599cb29e2b2db4a"
	hashfiles := []string{"resources/package.json"}
	got := Tokenizer(hashfiles)

	if got != want {
		t.Errorf("Tokenizer(%q) == %q, want %q", hashfiles, got, want)
	}
}
