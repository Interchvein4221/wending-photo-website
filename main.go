package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Photo struct {
	Number int
	File   string
}

type PageData struct {
	Title  string
	Date   string
	Couple string
	Photos []Photo
}

// loadPhotos автоматически загружает фотографии из static/photos
func loadPhotos() ([]Photo, error) {
	files, err := os.ReadDir("static/photos")
	if err != nil {
		return nil, err
	}

	// Сортируем файлы по номеру DSCxxxxx
	sort.Slice(files, func(i, j int) bool {
		return photoNumber(files[i].Name()) < photoNumber(files[j].Name())
	})

	photos := make([]Photo, 0, len(files))

	number := 1

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))

		// Берём только фотографии
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			continue
		}

		photos = append(photos, Photo{
			Number: number,
			File:   file.Name(),
		})

		number++
	}

	return photos, nil
}

// photoNumber достаёт номер из имени фотографии.
//
// Например:
//
// DSC03565.jpg → 3565
// DSC03716.jpg → 3716
// DSC04081.jpg → 4081
//
// Это позволяет сортировать фотографии именно по номеру.
func photoNumber(filename string) int {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))

	// Ищем цифры в конце имени
	i := len(name) - 1

	for i >= 0 && name[i] >= '0' && name[i] <= '9' {
		i--
	}

	if i == len(name)-1 {
		// Если цифр в конце нет
		return 0
	}

	number, err := strconv.Atoi(name[i+1:])
	if err != nil {
		return 0
	}

	return number
}

func main() {

	// Загружаем фотографии автоматически
	photos, err := loadPhotos()
	if err != nil {
		log.Fatal("Не удалось загрузить фотографии:", err)
	}

	log.Printf("Загружено фотографий: %d\n", len(photos))

	tmpl := template.Must(
		template.ParseFiles("templates/index.html"),
	)

	data := PageData{
		Title:  "Артур и Анастасия",
		Date:   "14.07.2026",
		Couple: "Артур & Анастасия",
		Photos: photos,
	}

	// Статические файлы:
	// CSS
	// JavaScript
	// фотографии
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// Главная страница
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		err := tmpl.Execute(w, data)

		if err != nil {
			http.Error(
				w,
				"Ошибка сервера",
				http.StatusInternalServerError,
			)

			log.Println(err)
		}
	})

	log.Println("================================")
	log.Println("Свадебная галерея")
	log.Println("Артур & Анастасия")
	log.Println("14.07.2026")
	log.Println("================================")
	log.Println("Сайт запущен:")
	log.Println("http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}