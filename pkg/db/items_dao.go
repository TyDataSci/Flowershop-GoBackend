package db

import (
	"Flowershop-GoBackend/pkg/models"
	"context"
	"fmt"
)

func GetItem(paramItemID int) (models.Item, error) {
	var item models.Item
	query := `SELECT i.id,t.name,i.description,i.price,i.image 
			 FROM items as i 
			 JOIN types as t ON t.id = i.typeid
			 WHERE i.id = $1`

	if err := db().PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return item, err
	}
	row := db().QueryRowContext(context.Background(), query,
		paramItemID)
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return item, err
	}

	if err := row.Scan(&item.ID, &item.Type, &item.Description, &item.Price, &item.Image); err != nil {
		fmt.Println("row.Scan", err)
		return item, err
	}

	fmt.Printf("Successfully returned item %v\n", item.ID)
	return item, nil

}

func GetItems() ([]models.Item, error) {
	var items []models.Item

	if err := db().PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return items, err
	}
	query := `SELECT i.id,t.name,i.description,i.price,i.image 
			 FROM items as i 
			 JOIN types as t ON t.id = i.typeid`
	rows, err := db().QueryContext(context.Background(), query)
	if err != nil {
		fmt.Println("db.QueryContext", err)
		return items, err
	}
	defer func() {
		_ = rows.Close()
	}()

	if rows.Err() != nil {
		fmt.Println("row.Err()", err)
		return items, rows.Err()
	}

	for rows.Next() {
		var nextItem models.Item
		if err := rows.Scan(&nextItem.ID, &nextItem.Type, &nextItem.Description, &nextItem.Price, &nextItem.Image); err != nil {
			fmt.Println("row.Scan", err)
			return items, err
		}
		items = append(items, nextItem)
	}

	fmt.Printf("%v Items returned\n", len(items))
	return items, nil
}

func CreateItem(paramType string, paramDescription string, paramPrice string, paramImage string) error {
	var typeExists bool
	row := db().QueryRowContext(context.Background(), "SELECT EXISTS(SELECT id FROM types WHERE name = $1)",
		paramType)
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return err
	}

	if err := row.Scan(&typeExists); err != nil {
		fmt.Println("row.Scan", err)
		return err
	}
	if !typeExists {
		_, err := db().ExecContext(context.Background(),
			"INSERT INTO types (name) VALUES($1)", paramType)
		if err != nil {
			fmt.Println("db.QueryRowContext", err)
			return err
		}

	}

	_, err := db().ExecContext(context.Background(),
		"INSERT INTO items (typeid, description, price, image) VALUES((SELECT id FROM types WHERE name = $1), $2, $3, $4)",
		paramType, paramDescription, paramPrice, paramImage)
	if err != nil {
		fmt.Println("db.QueryRowContext", err)
		return err
	}
	fmt.Printf("%v Item created\n", paramDescription)
	return nil
}

func UpdateItem(paramItem models.Item) error {
	_, err := db().ExecContext(context.Background(),
		"UPDATE order_items SET typeid = (SELECT id FROM types WHERE name = $1), description = $2, price = $3, image = $4 WHERE id = $5",
		paramItem.Type, paramItem.Description, paramItem.Price, paramItem.Image, paramItem.ID)
	if err != nil {
		fmt.Println("db.ExecContext", err)
		return err
	}
	println("%v successfully updated \n", paramItem.Description)
	return nil
}
