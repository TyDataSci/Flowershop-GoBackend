package db

import (
	"Flowershop-GoBackend/pkg/models"
	"context"
	"fmt"
)

func GetOrder(paramUserID int) (models.Order, error) {
	var order models.Order

	if err := db().PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return order, err
	}
	row := db().QueryRowContext(context.Background(), "SELECT * FROM orders WHERE userid = $1",
		paramUserID)
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return order, err
	}

	if err := row.Scan(&order.ID, &order.Date, &order.UserID, &order.Delivery, &order.Completed); err != nil {
		fmt.Println("row.Scan", err)
		return order, err
	}

	fmt.Printf("userid: %v, date: %v, userid: %v, delivery: %v, completed: %v\n", order.ID, order.Date, order.UserID, order.Delivery, order.Completed)
	return order, nil
}

func GetIncompletOrder(paramUserID int) (models.Order, error) {
	var order models.Order

	if err := db().PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return order, err
	}
	row := db().QueryRowContext(context.Background(), "SELECT * FROM orders WHERE userid = $1 AND completed = false",
		paramUserID)
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return order, err
	}

	if err := row.Scan(&order.ID, &order.Date, &order.UserID, &order.Delivery, &order.Completed); err != nil {
		fmt.Println("row.Scan", err)
		return order, err
	}

	fmt.Printf("orderid: %v, date: %v, userid: %v, delivery: %v, completed: %v\n", order.ID, order.Date, order.UserID, order.Delivery, order.Completed)
	return order, nil
}

func CreateOrder(paramUserID int) (models.Order, error) {
	var order models.Order
	err := db().QueryRowContext(context.Background(),
		"INSERT INTO orders (userid) VALUES($1)RETURNING *",
		paramUserID).Scan(&order.ID, &order.Date, &order.UserID, &order.Delivery, &order.Completed)
	if err != nil {
		fmt.Println("db.QueryRowContext line 38", err)
		return order, err
	}
	fmt.Println(order.ID, order.Date, order.UserID, order.Delivery, order.Completed)
	return order, nil
}

func UpdateOrder(paramOrder models.Order) error {
	_, err := db().ExecContext(context.Background(), "UPDATE orders SET delivery = $1, completed = $2 WHERE id = $3",
		paramOrder.Delivery, paramOrder.Completed, paramOrder.ID)
	if err != nil {
		fmt.Println("db.ExecContext", err)
		return err
	}
	println(paramOrder.ID, "Successfully updated")
	return nil
}
