# Trabalhando com Json

- Marshal e Unmarshal e também

```go
	// Quando usa Encoder, o json é retornado como um stream
	// que é direcionado para o os.Stdout
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}
```