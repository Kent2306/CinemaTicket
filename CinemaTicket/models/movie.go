package models

import (
	"cinematicket/database"
	"time"
)

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PosterURL   string    `json:"poster_url"`
	Duration    int       `json:"duration"`
	Price       float64   `json:"price"`
	HallID      int       `json:"hall_id"`
	ShowTime    time.Time `json:"show_time"`
	CreatedAt   time.Time `json:"created_at"`
	HallName    string    `json:"hall_name"`
}

func GetAllMovies() ([]Movie, error) {
	query := `
        SELECT m.id, m.title, m.description, m.poster_url, m.duration, 
               m.price, m.hall_id, m.show_time, m.created_at, h.name as hall_name
        FROM movies m
        JOIN halls h ON m.hall_id = h.id
        ORDER BY m.show_time
    `

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID, &movie.Title, &movie.Description, &movie.PosterURL,
			&movie.Duration, &movie.Price, &movie.HallID, &movie.ShowTime,
			&movie.CreatedAt, &movie.HallName,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func CreateMovie(movie *Movie) error {
	query := `
        INSERT INTO movies (title, description, poster_url, duration, price, hall_id, show_time)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `

	result, err := database.DB.Exec(
		query, movie.Title, movie.Description, movie.PosterURL,
		movie.Duration, movie.Price, movie.HallID, movie.ShowTime,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	movie.ID = int(id)
	return nil
}
