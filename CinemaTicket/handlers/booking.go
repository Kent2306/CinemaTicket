package handlers

import (
	"cinematicket/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBookingForm(c *gin.Context) {
	movieIDStr := c.Param("id")
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	movies, err := models.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var currentMovie *models.Movie
	for _, movie := range movies {
		if movie.ID == movieID {
			currentMovie = &movie
			break
		}
	}

	if currentMovie == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	bookedSeats, err := models.GetBookedSeats(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "booking.html", gin.H{
		"movie":       currentMovie,
		"bookedSeats": bookedSeats,
	})
}

func CreateBooking(c *gin.Context) {
	movieIDStr := c.Param("id")
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	seatNumberStr := c.PostForm("seat_number")
	customerName := c.PostForm("customer_name")
	customerEmail := c.PostForm("customer_email")

	seatNumber, err := strconv.Atoi(seatNumberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seat number"})
		return
	}

	booking := &models.Booking{
		MovieID:       movieID,
		SeatNumber:    seatNumber,
		CustomerName:  customerName,
		CustomerEmail: customerEmail,
	}

	err = models.CreateBooking(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}
