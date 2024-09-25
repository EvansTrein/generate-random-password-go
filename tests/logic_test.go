package tests

import (
	"testing"
	"GoLearn/internal/logic"
)


func TestIsValidPassword(t *testing.T) {
	fail := []string{"123456", "QWERTY", "qwerty"}

	for _, el := range fail {
		if logic.IsValidPassword(el) {
			t.Errorf("Generated password is %s", el)
		}
	}
}