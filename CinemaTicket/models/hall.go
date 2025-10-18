package models

import "cinematicket/database"

type Hall struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

func GetAllHalls() ([]Hall, error) {
	rows, err := database.DB.Query("SELECT id, name, capacity FROM halls")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var halls []Hall
	for rows.Next() {
		var hall Hall
		err := rows.Scan(&hall.ID, &hall.Name, &hall.Capacity)
		if err != nil {
			return nil, err
		}
		halls = append(halls, hall)
	}

	return halls, nil
}
