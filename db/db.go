package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:anatsaya@tcp(localhost:3306)/ticket2_shoppingcart")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(8)

	err = CreateTicketListsTable()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	err = CreateDiscountsTable()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

}

func CreateTicketListsTable() error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		CREATE TABLE IF NOT EXISTS ticketLists (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title TEXT NOT NULL,
			img TEXT NOT NULL,
			price INT NOT NULL,
			descriptionEng TEXT NOT NULL,
			descriptionThai TEXT NOT NULL
		);
	`

	_, err := DB.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("could not create ticketLists table: %w", err)
	}

	return nil
}

func CreateDiscountsTable() error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		CREATE TABLE IF NOT EXISTS discounts (
			id INT AUTO_INCREMENT PRIMARY KEY,
			code TEXT NOT NULL,
			discount INT NOT NULL,
			type TEXT NOT NULL
		);
	`

	_, err := DB.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("could not create discounts table: %w", err)
	}

	return nil
}
