package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	// Quando usa Marshal, o json é retornado como um slice de bytes
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	// Quando usa Encoder, o json é retornado como um stream
	// que é direcionado para o os.Stdout
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	// jsonPuro é um slice de bytes
	jsonPuro := []byte(`{"n":2,"s":200}`)
	
	var contaX Conta
	// Passamos o endereço da variável contaX para que o Unmarshal
	// faça a alteração na variável
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
}
