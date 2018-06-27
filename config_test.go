package sooslib

import (
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {

	configFile := "resources/build.yml"

	want := BuildConfig{
		Build:     ".",
		Image:     "soos-build",
		Ports:     []string{"5000:5000"},
		Volumes:   []string{".:/code"},
		Hashfiles: []string{"resources/package.json"},
	}

	got := Config(configFile)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Config(%q) == %q, want %q", configFile, got, want)
	}
}
