package main

import "fmt"

//Considerations:
//  We can implements method using interface or tis way func (s <Struct>) <NameFunc>() {}
// An struct can implements a interface but an interface cant implements a struct(maybe it's obviously)
//
// To use an interface we need to implements all methods

type Account struct {
	Id    string
	Name  string
	Email string
}

// 1 Way todo
// type CollaboratorRepository struct {
// 	ICollaboratorRepository
// }

// type ICollaboratorRepository interface {
// 	findByCpf(cpf int) bool
// 	findByEmail(email string) bool
// 	findById(id string) Account
// 	list() []Account
// 	delete(id string)
// 	create(account Account) Account
// }

// func findByCpf(cpf int) bool {
// 	result := false
// 	if cpf == 51921744812 {
// 		result = true
// 		fmt.Println("Cpf exists")
// 	}

// 	return result
// }

// func findByEmail(email string) bool {
// 	result := false
// 	if email == "luizsouza@mail.com" {
// 		result = true
// 		fmt.Println("Email exists")
// 	}

// 	return result
// }

// func findById(id string) Account {
// 	result := Account{}
// 	if id == "oi" {
// 		result = Account{
// 			Id:    "oi",
// 			Name:  "luiz",
// 			Email: "luizsouza@mail.com",
// 		}
// 		fmt.Println("account find")
// 	}

// 	return result
// }

// func delete(id string) {
// 	fmt.Println("account Deleted")

// 	return
// }

// func create(account Account) Account {
// 	fmt.Println("account created")

// 	return account
// }

// func list() []Account {
// 	allUsers := []Account{}
// 	fmt.Println("List all users")

// 	return allUsers
// }

//  2 Way todo

type CollaboratorRepository struct{}

func (collaboratorRepository *CollaboratorRepository) findByCpf(cpf int) bool {
	result := false
	if cpf == 51921744812 {
		result = true
		fmt.Println("Cpf exists")
	}

	return result
}

func (collaboratorRepository *CollaboratorRepository) findByEmail(email string) bool {
	result := false
	if email == "luizsouza@mail.com" {
		result = true
		fmt.Println("Email exists")
	}

	return result
}

func (collaboratorRepository *CollaboratorRepository) findById(id string) Account {
	result := Account{}
	if id == "oi" {
		result = Account{
			Id:    "oi",
			Name:  "luiz",
			Email: "luizsouza@mail.com",
		}
		fmt.Println("account find")
	}

	return result
}

func (collaboratorRepository *CollaboratorRepository) delete(id string) {
	fmt.Println("account Deleted")

	return
}

func (collaboratorRepository *CollaboratorRepository) create(account Account) Account {
	fmt.Println("account created")

	return account
}

func (collaboratorRepository *CollaboratorRepository) list() []Account {
	allUsers := []Account{}
	fmt.Println("List all users")

	return allUsers
}

var collaboratorRepository = CollaboratorRepository{}

func main() {

	tes := collaboratorRepository.findById("oi")

	fmt.Println(tes)
}
