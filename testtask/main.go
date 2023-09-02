package main

import (
	"fmt"
	"testtask2/controller"

	"os"
    "os/signal"
    "syscall"
)

func main(){
	controller.Controller()

    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)

    fmt.Println("Сервер запущен. Нажмите Ctrl+C для выхода.")
    <-sigChan

    fmt.Println("Завершение сервера...")
}