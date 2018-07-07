package sooslib

import (
	"testing"
)

func TestBuildRunCmd(t *testing.T) {

	userCurrent = userCurrentMock

	osGetwd = osGetwdMock

	want := `docker run -it -u 0:0 -v /test:/test -v /root:/build/app -p 5000:5000 soos-build:eac1fbb9986eacd640fdb306d599cb29e2b2db4a`

	buildConfig := BuildConfig{
		Build:     ".",
		Image:     "soos-build",
		Ports:     []string{"5000:5000"},
		Volumes:   []string{"/test:/test"},
		Hashfiles: []string{"resources/package.json"},
	}

	got := BuildRunCmd(buildConfig)

	if got != want {
		t.Errorf("BuildRunCmd(%q) == %q, want %q", buildConfig, got, want)
	}

}

func Test_execCmd(t *testing.T) {
	type args struct {
		cmdString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple",
			args: args{cmdString: "echo test"},
			want: true,
		},
		{
			name: "empty",
			args: args{cmdString: "echo -n"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := execCmd(tt.args.cmdString); got != tt.want {
				t.Errorf("execCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
