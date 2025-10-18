package handlers

import (
	"cinematicket/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	movies, err := models.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"movies": movies,
	})
}

func ShowAddMovieForm(c *gin.Context) {
	halls, err := models.GetAllHalls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "add_movie.html", gin.H{
		"halls": halls,
	})
}

func AddMovie(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	posterURL := c.PostForm("poster_url")
	durationStr := c.PostForm("duration")
	priceStr := c.PostForm("price")
	hallIDStr := c.PostForm("hall_id")
	showTimeStr := c.PostForm("show_time")

	duration, _ := strconv.Atoi(durationStr)
	price, _ := strconv.ParseFloat(priceStr, 64)
	hallID, _ := strconv.Atoi(hallIDStr)
	showTime, _ := time.Parse("2006-01-02T15:04", showTimeStr)

	movie := &models.Movie{
		Title:       title,
		Description: description,
		PosterURL:   posterURL,
		Duration:    duration,
		Price:       price,
		HallID:      hallID,
		ShowTime:    showTime,
	}

	err := models.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}
