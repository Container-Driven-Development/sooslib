package sooslib

import "testing"

func TestBuildRunCmd(t *testing.T) {

	want := `docker run -it -v .:/code -p 5000:5000 soos-build:eac1fbb9986eacd640fdb306d599cb29e2b2db4a`

	buildConfig := BuildConfig{
		Build:     ".",
		Image:     "soos-build",
		Ports:     []string{"5000:5000"},
		Volumes:   []string{".:/code"},
		Hashfiles: []string{"resources/package.json"},
	}

	got := BuildRunCmd(buildConfig)

	if got != want {
		t.Errorf("BuildRunCmd(%q) == %q, want %q", buildConfig, got, want)
	}

}
