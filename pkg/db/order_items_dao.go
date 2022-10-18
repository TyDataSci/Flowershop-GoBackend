package db

import (
	"Flowershop-GoBackend/pkg/models"
	"context"
	"fmt"
)

func GetOrderItems(paramOrderID int) ([]models.Item, error) {
	var items []models.Item

	if err := db().PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return items, err
	}
	query := `SELECT i.id,t.name,i.description,i.price,i.image 
			  FROM items as i
			  JOIN types as t ON t.id = i.typeid
			  JOIN order_items oi ON i.id = oi.itemid
			  WHERE oi.orderid = $1`
	rows, err := db().QueryContext(context.Background(), query, paramOrderID)
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

func CreateOrderItem(paramOrderID int, paramItemID int) (models.Order_Item, error) {
	var order_item models.Order_Item
	err := db().QueryRowContext(context.Background(),
		"INSERT INTO order_items (orderid, itemid) VALUES($1, $2)RETURNING *",
		paramOrderID, paramItemID).Scan(&order_item.ID, &order_item.OrderID, &order_item.ItemID, &order_item.Removed)
	if err != nil {
		fmt.Println("db.QueryRowContext line 38", err)
		return order_item, err
	}
	fmt.Println(order_item.ID, order_item.OrderID, order_item.ItemID, order_item.Removed)
	return order_item, nil
}

func UpdateOrderItem(paramOrderID int, paramItemID int, paramRemoved bool) error {
	_, err := db().ExecContext(context.Background(), "UPDATE order_items SET removed = $1 WHERE orderid = $2 AND itemid = $3",
		paramRemoved, paramOrderID, paramItemID)
	if err != nil {
		fmt.Println("db.ExecContext", err)
		return err
	}
	println(paramOrderID, "Successfully updated")
	return nil
}
