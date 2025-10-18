package models

import (
	"cinematicket/database"
	"time"
)

type Booking struct {
	ID            int       `json:"id"`
	MovieID       int       `json:"movie_id"`
	SeatNumber    int       `json:"seat_number"`
	CustomerName  string    `json:"customer_name"`
	CustomerEmail string    `json:"customer_email"`
	BookingTime   time.Time `json:"booking_time"`
	Status        string    `json:"status"`
}

func CreateBooking(booking *Booking) error {
	query := `
        INSERT INTO bookings (movie_id, seat_number, customer_name, customer_email)
        VALUES (?, ?, ?, ?)
    `

	result, err := database.DB.Exec(
		query, booking.MovieID, booking.SeatNumber,
		booking.CustomerName, booking.CustomerEmail,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	booking.ID = int(id)
	return nil
}

func GetBookedSeats(movieID int) ([]int, error) {
	query := "SELECT seat_number FROM bookings WHERE movie_id = ? AND status = 'active'"

	rows, err := database.DB.Query(query, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookedSeats []int
	for rows.Next() {
		var seat int
		err := rows.Scan(&seat)
		if err != nil {
			return nil, err
		}
		bookedSeats = append(bookedSeats, seat)
	}

	return bookedSeats, nil
}
