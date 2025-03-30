package handlers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandle(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		data, err := os.ReadFile("..\\index.html")
		if err != nil {
			log.Fatal(err)
		}
		res.Write(data)
	} else if req.Method != http.MethodGet {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func UploadHandle(w http.ResponseWriter, r *http.Request) {
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

	result := service.ConvMorse(string(fileContent)) // результат конвертации

	//fmt.Println(result)

	//создаем директорию "result" для хранения файлов с результатом
	err = os.Mkdir("..//result", 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}

	//создаем файл с результатом и сохраняем в директорию "result"

	curTime := time.Now().Local().Format("2006-01-02_15-04-05") //строка с временем создания файла, для названия
	ext := filepath.Ext(header.Filename)

	filename := fmt.Sprintf("%s_result_%s%s", header.Filename, curTime, ext) // название файла

	//создание файла с результатом
	resultFilepath := filepath.Join("..//result", filename) //путь для сохранения нового файла "путь/имя_файла"

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

	//тестирование выходных данных
	/*fmt.Println(fileContent)
	testing, err := os.ReadFile(resultFilepath)
	if err != nil {
		return
	}
	fmt.Println(string(testing))*/

}
