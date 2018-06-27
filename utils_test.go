package sooslib

import (
	"errors"
	"testing"
)

func TestCheck(t *testing.T) {

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
