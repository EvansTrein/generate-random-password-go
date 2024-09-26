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

func TestIsValidPasswordNoNumbers(t *testing.T) {
	fail := []string{"C*OdapJXP", "fAdcoUob!XhfC", "XMsqkoylBg#LP", "pLanA$"}

	for _, el := range fail {
		if logic.IsValidPassword(el) {
			t.Errorf("Generated password is %s", el)
		}
	}
}

func TestIsValidPasswordNoLowercase(t *testing.T) {
	fail := []string{"$DB5TLVB", "_7KC$EO9O0", "*Q0F1MEQH_JF", "Y2U3Q!"}

	for _, el := range fail {
		if logic.IsValidPassword(el) {
			t.Errorf("Generated password is %s", el)
		}
	}
}

func TestIsValidPasswordNoUppercase(t *testing.T) {
	fail := []string{"3h_jyg7iyo", "QNKQ7X9#KL5G", "TKLR!ND1SMVIR", "U&ZTLY_0"}

	for _, el := range fail {
		if logic.IsValidPassword(el) {
			t.Errorf("Generated password is %s", el)
		}
	}
}