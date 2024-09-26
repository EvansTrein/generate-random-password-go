package tests

import (
	"testing"
	"GoLearn/internal/logic"
)


func TestIsValidPasswordInCorrect(t *testing.T) {
	fail := []string{"123456", "QWERTY", "qwerty", "qweW123"}

	for _, el := range fail {
		if logic.IsValidPassword(el) {
			t.Errorf("Generated password is %s", el)
		}
	}
}

func TestIsValidPasswordСorrect(t *testing.T) {
	fail := []string{"P!$W5$n_nzOFyK", "0hJuzKH$nUc2eL", "$3nnd92*bT", "H7!vRWtCRI"}

	for _, el := range fail {
		if !logic.IsValidPassword(el) {
			t.Errorf("Generated password is %s", el)
		}
	}
}

func TestIsValidPasswordNoSymbols(t *testing.T) {
	fail := []string{"7WAqIlNszD", "T0BffjQk0P57YnUc2eL", "C5ccRCR", "y0EsF"}

	for _, el := range fail {
		if logic.IsValidPassword(el) {
			t.Errorf("Generated password is %s", el)
		}
	}
}