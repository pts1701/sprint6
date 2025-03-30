package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	/* fmt.Println("Запускаем сервер")

	http.HandleFunc("/", handlers.MainHandle)
	http.HandleFunc("/upload", handlers.UploadHandle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")*/
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)

	srv := server.MyServer(logger)

	err := srv.HttpServer.ListenAndServe()
	if err != nil {
		logger.Fatal("Error starting server: ", err)
	}

}
