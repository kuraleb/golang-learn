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

	for command != exit{
		fmt.Println("Введите команду:") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Scan(&command)

		switch command {
		case exit:	
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде \"login_password\":")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			userList = append(userList, command)
		case auth: 
		fmt.Println("Введите логин и пароль в таком виде \"login_password\":")
		fmt.Scan(&command) 
			for _, v := range userList {
				if v == command {
					fmt.Println("Добро пожаловать в магазин!")
					break
				}
				fmt.Println("Вы не зарегистрированы")
			}
		}
	}
}