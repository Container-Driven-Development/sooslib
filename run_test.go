package sooslib

import (
	"os/user"
	"testing"
)

func userCurrentMock() (*user.User, error) {

	usr := user.User{
		HomeDir:  "/root",
		Username: "root",
		Name:     "root",
		Uid:      "0",
		Gid:      "0",
	}

	return &usr, nil
}

func TestBuildRunCmd(t *testing.T) {

	userCurrent = userCurrentMock

	want := `docker run -it -u 0:0 -v .:/code -p 5000:5000 soos-build:eac1fbb9986eacd640fdb306d599cb29e2b2db4a`

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

func Test_execCmd(t *testing.T) {

	cmdString := "echo test"

	got := execCmd(cmdString)

	want := true

	if got != want {
		t.Errorf("execCmd(%q) == %t, want %t", cmdString, got, want)
	}

}
