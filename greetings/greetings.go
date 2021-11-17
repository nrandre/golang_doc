package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello retorna uma saudação para o nome.
func Hello(name string) (string, error) {
	// Se não tiver o nome, retorne um erro com a mensagem.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Se o nome for recebido, retorna o valor que incorpora
	// o nome na mensagem de saudações.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos retorna um mapa que associa cada uma das pessoas
// nomeadas com uma mensagem de saudação.
func Hellos(names []string) (map[string]string, error) {
	// Um mapa para associar nomes a mensagens.
	messages := make(map[string]string)
	// Loop através da fatia de nomes recebida, chamando
	// a função Hello para obter uma mensagem para cada nome.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// No mapa, associe a mensagem recuperada com
		// o nome.
		messages[name] = message
	}
	return messages, nil
}

//init define valores iniciais para variáveis usadas na função.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat retorna um de um conjunto de mensagens de saudação.
// O devolvido a mensagem é selecionada aleatoriamente.
func randomFormat() string {
	//Uma fatia de formatos de mensagem
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met",
	}

	// Retorna um formato de mensagem selecionado aleatoriamente,
	// especificando um índice aleatório para a fatia de formatos.
	return formats[rand.Intn(len(formats))]
}
