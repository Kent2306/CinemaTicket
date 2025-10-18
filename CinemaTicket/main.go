package main

import (
	"cinematicket/database"
	"cinematicket/handlers"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	err := database.InitDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	// Статические файлы
	router.Static("/static", "./static")

	// Шаблоны с пользовательскими функциями
	router.SetFuncMap(template.FuncMap{
		"seq": func(start, end int) []int {
			var result []int
			for i := start; i <= end; i++ {
				result = append(result, i)
			}
			return result
		},
		"contains": func(slice []int, item int) bool {
			for _, v := range slice {
				if v == item {
					return true
				}
			}
			return false
		},
	})
	router.LoadHTMLGlob("templates/*")

	// Маршруты
	router.GET("/", handlers.GetMovies)
	router.GET("/add-movie", handlers.ShowAddMovieForm)
	router.POST("/add-movie", handlers.AddMovie)
	router.GET("/booking/:id", handlers.ShowBookingForm)
	router.POST("/booking/:id", handlers.CreateBooking)

	// Запуск сервера
	println("Сервер запущен на http://localhost:8080")
	router.Run(":8080")
}
