# Todos os arquivos que estão na mesma pasta devem conter o mesmo nome de Pacote

Exemplo:

Se criarmos um arquivo chamado `a.go` e colocarmos, por exemplo, um nome de pacote diferente
de `package main`, ao tentarmos rodar o projeto com o comando `go run *` dentro do diretório principal.
Ele não rodará, apenas avisará que tem arquivos de pacotes diferentes e não saberá por onde começar.


Outra coisa que é tentado mostrar que tudo que declarar em `a.go` sendo que ele também está no mesmo
pacote `main`. Então todas as contantes, variáveis e funções de `a.go` poderão ser utilizado em `main.go`.

