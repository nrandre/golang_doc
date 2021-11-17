package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Definir propriedades do Logger predefinido, incluindo
	// o prefixo de entrada de registro e um sinalizador para
	// desativar a impressão a hora, o arquivo de origem e o número da linha
	log.SetPrefix("saudações: ")
	log.SetFlags(0)

	// Uma fatia de nomes
	names := []string{"Marcos", "Lilian", "André"}

	// Solicita a mensagem de saudação.
	//message, err := greetings.Hello("André")

	// Solicita a mensage de saudações para os nomes
	message, err := greetings.Hellos(names)

	// Se um erro for retornado, imprima-o no
	//  console e sai do programa.
	if err != nil {
		log.Fatal(err)
	}

	// Se não for retornado erro, imprime o retorno da mensagem
	// no console
	fmt.Println(message)
}
