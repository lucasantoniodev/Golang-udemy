package main

import "fmt"

type RepositoryModel interface {
	createUser(name string, password string) uint
	deleteUser(name string)
}

type RepositoryImp struct {
	name     string
	password string
}

func (r RepositoryImp) createUser(name string, password string) uint {
	user := name
	pass := password
	fmt.Printf("Usuário %s com senha %s criado com sucesso!", user, pass)
	return 1
}

func (r RepositoryImp) deleteUser(name string) {
	fmt.Printf("Usuário %s deletado com sucesso!", name)
}

func main() {
	repository := RepositoryImp{"Teco", "123"}
	var methods RepositoryModel = repository
	create := methods.createUser(repository.name, repository.password)
	fmt.Println(create)
}
