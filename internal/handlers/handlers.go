package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"s6-final/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	htmlForm, err := os.ReadFile("index.html")
	if err != nil {
		log.Printf("ошибка при чтении файла index.html: %v", err)
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlForm)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("ошибка при получении файла: %v", err)
		http.Error(w, "ошибка при получении", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("ошибка при чтении файла: %v", err)
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	result, err := service.MorseOrTextRecognition(data)
	if err != nil {
		log.Printf("ошибка при распознавании Морзе или текста: %v", err)
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("%s%s", time.Now().UTC().Format("02.01.2006_15:04:05"), filepath.Ext(handler.Filename))
	localFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("ошибка при открытии файла для записи: %v", err)
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	if _, err := localFile.Write([]byte(result)); err != nil {
		log.Printf("ошибка при записи в файл: %v", err)
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	log.Printf("файл успешно загружен и сохранен как: %s", fileName)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(result))
}
