package sooslib

import (
	"errors"
	"testing"
)

func Test_check(t *testing.T) {

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("TestCheck should have panicked!")
			}
		}()

		check(errors.New("fail"))
	}()

	func() {
		check(nil)
	}()
}

func Test_cwd(t *testing.T) {

	osGetwd = osGetwdMock

	tests := []struct {
		name string
		want string
	}{
		{
			name: "root",
			want: "/root",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cwd(); got != tt.want {
				t.Errorf("cwd() = %v, want %v", got, tt.want)
			}
		})
	}
}
