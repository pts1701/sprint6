package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method != http.MethodGet {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	data, err := os.ReadFile("..\\index.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}

func UploadHandle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method != http.MethodPost {
		http.Error(w, "method must be POST", http.StatusInternalServerError)
		return
	}

	//парсинг html формы

	r.ParseMultipartForm(10 << 20)

	//получение из html формы файла,
	//для этого надо открыть код html и найти в мультиформе type="file" name="myFile"

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	//получение данных из файла

	fileContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file content", http.StatusInternalServerError)
		return
	}

	result, err := service.ConvMorse(string(fileContent)) // результат конвертации
	if err != nil {
		http.Error(w, "Failed to convert file content", http.StatusInternalServerError)
	}

	curTime := time.Now().UTC().Format("2006-01-02_15-04-05") //строка с временем создания файла, для названия
	ext := filepath.Ext(header.Filename)

	filename := fmt.Sprintf("%s_result_%s%s", header.Filename, curTime, ext) // название файла

	//создание файла с результатом
	resultFilepath := filepath.Join("..", filename) //путь для сохранения нового файла "путь/имя_файла"

	resultFile, err := os.Create(resultFilepath) //создание файла - результата обработки
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer resultFile.Close()

	//запись результата обработки в созданный для него файл
	_, err = fmt.Fprint(resultFile, result)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, result)

}
