package main

import (
	"fmt"
	"net"
)

func main() {
	// Создание TCP соединения с сервером
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка при подключении:", err)
		return
	}
	defer conn.Close()

	helmanService := NewHelmanService()
	fmt.Println(helmanService)
}
