package models

import (
	"fmt"

	"example.com/shoppingCart-api/db"
)

type TicketList struct {
	ID              int    `json:"id"`
	Title           string `json:"title" binding:"required"`
	Img             string `json:"img" binding:"required"`
	Price           string `json:"price" binding:"required"`
	DescriptionEng  string `json:"descriptionEng" binding:"required"`
	DescriptionThai string `json:"descriptionThai" binding:"required"`
}

func GetAllTicketLists() ([]TicketList, error) {

	rows, err := db.DB.Query("SELECT id, title, img, price, descriptionEng, descriptionThai FROM ticketLists")

	if err != nil {
		return nil, fmt.Errorf("could not fetch events: %w", err)
	}

	defer rows.Close()

	var ticketLists []TicketList

	for rows.Next() {
		var ticketList TicketList

		if err := rows.Scan(&ticketList.ID, &ticketList.Title, &ticketList.Img, &ticketList.Price, &ticketList.DescriptionEng, &ticketList.DescriptionThai); err != nil {
			return nil, fmt.Errorf("could not fetch ticketLists: %w", err)
		}
		ticketLists = append(ticketLists, ticketList)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return ticketLists, nil

}
