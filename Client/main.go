package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	var Nik string // Запрос Ника у пользователя
	fmt.Print("Введите Ваш ник: ")
	fmt.Scanln(&Nik)
	defer conn.Close()
	go func() {
		io.Copy(os.Stdout, conn)
	}()
	io.Copy(conn, os.Stdin)
	fmt.Printf("%s: exit", Nik)
}

