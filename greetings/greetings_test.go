package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName chama saudações. Alô com um nome, verificando
// para um valor de retorno válido.
func TestHelloName(t *testing.T) {
	name := "Andre"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Andre")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Andre") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty chama greetings.Hello com uma string vazia,
// verificando se há um erro.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
