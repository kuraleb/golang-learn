package main

import "fmt"

const (
	exit = "exit"
	auth = "auth"
	reg = "reg"
)

func main() {
	var command string

	userList := []string{"user1_password1", "user2_password2", "user3_password3"}
	productList := make([]string, 0, 10)

	_= productList

	commandList := []string{
		"^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^",
		"Вам доступны следующие команды:",
		"auth - авторизация в магазине",
		"reg - регистрация ",
		"exit - выход из магазина",
	}
	for command != exit{
		for _, v := range commandList {
			fmt.Println(v)
		}
		fmt.Println("Введите команду:")
		fmt.Scan(&command)

		switch command {
		case exit:	
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде \"login_password\":")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			for _, v := range userList {
				i := 0
				if v == command {
					i++
					fmt.Println("Пользователь уже существует")	
					return
				}
				else {
					userList = append(userList, command)
					message:= fmt.Sprintf("Пользователь %v успешно добавлен", command)
					fmt.Println(message)
				}
			}
		case auth: 
		fmt.Println("Введите логин и пароль в таком виде \"login_password\":")
		fmt.Scan(&command) 
		// for _, v := range userList {
		if v == command {
			fmt.Println("Добро пожаловать в магазин!")
		} 
		// }
		fmt.Println("Вы не зарегистрированы") // Сделать так, что бы выводил сообщение, если пользователь не существует
		}
	}
}

// Реализовать следующий API
// add_product - добавляет продукт в список продуктов, который вводится с консоли в ProductList
// order - выводит сообщение, что купили и очищает корзину
// show_cart - выводит список продуктов в корзине