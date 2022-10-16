package db

import (
	"Flowershop-GoBackend/pkg/models"
	"context"
	"fmt"
)

func GetUserSession(paramToken string) (models.Session, error) {
	var userSession models.Session

	if err := db().PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return userSession, err
	}
	row := db().QueryRowContext(context.Background(), "SELECT * FROM sessions WHERE token = $1",
		paramToken)
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return userSession, err
	}
	if err := row.Scan(&userSession.Token, &userSession.UserID, &userSession.OrderID, &userSession.Expiry); err != nil {
		fmt.Println("row.Scan", err)
		return userSession, err
	}

	fmt.Printf("userid: %v, orderid: %v, expiry: %v\n", userSession.UserID, userSession.OrderID, userSession.Expiry)
	return userSession, nil
}

func GetUserLastSession(paramUserID int) (models.Session, error) {
	var userSession models.Session

	if err := db().PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return userSession, err
	}
	row := db().QueryRowContext(context.Background(), "SELECT * FROM sessions WHERE userid = $1 ORDER BY expiry DESC",
		paramUserID)
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return userSession, err
	}
	if err := row.Scan(&userSession.Token, &userSession.UserID, &userSession.OrderID, &userSession.Expiry); err != nil {
		fmt.Println("row.Scan", err)
		return userSession, err
	}

	fmt.Printf("userid: %v, orderid: %v, expiry: %v\n", userSession.UserID, userSession.OrderID, userSession.Expiry)
	return userSession, nil
}

func CreateSession(session models.Session) (models.Session, error) {
	_, err := db().ExecContext(context.Background(),
		"INSERT INTO sessions(token,userid,orderid,expiry) VALUES($1,$2,$3,$4)",
		session.Token,
		session.UserID,
		session.OrderID,
		session.Expiry)
	if err != nil {
		fmt.Println("db.QueryRowContext", err)
		return session, err
	}
	fmt.Println("Token: ", session.Token, session.UserID, session.OrderID, session.Expiry)
	return session, nil
}

func UpdateSessionIDs(session models.Session) error {
	_, err := db().ExecContext(context.Background(),
		"UPDATE sessions SET userid=$1, orderid=$2,expiry=$3 WHERE token = $4",
		session.UserID,
		session.OrderID,
		session.Expiry,
		session.Token)
	if err != nil {
		fmt.Println("db.QueryRowContext", err)
		return err
	}
	fmt.Println("Token: ", session.Token, session.UserID, session.OrderID, session.Expiry)
	return nil
}
