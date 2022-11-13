package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	os.Setenv("DATABASE_URL", "postgres://postgres:123@localhost:5432/go_dev")

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}

type User struct {
	Id       int
	Name     string
	Language string
}

func GetUser(conn *pgx.Conn, userId int) {
	var user User
	row := conn.QueryRow(
		context.Background(),
		"select id, name, language from users where id=$1",
		userId,
	)
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Language,
	)
	fmt.Printf("User info: %+v\n", user)
	if err != nil {
		panic(err)
	}
}

func GetUsers(conn *pgx.Conn) {
	rows, err := conn.Query(
		context.Background(),
		"select name, weight from widgets where",
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Language)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", user)
	}
}

func InsertUser(conn *pgx.Conn, user User) {
	var id int
	err := conn.QueryRow(
		context.Background(),
		"insert into users(name, language) values($1, $2)",
		user.Name, user.Language,
	).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("new user id: %d\n", id)
}

func DeleteUser(conn *pgx.Conn, userId int) {
	tag, err := conn.Exec(
		context.Background(),
		"delete from users where id=$1",
		userId,
	)
	fmt.Println("rows updated", tag.RowsAffected())
	if err != nil {
		panic(err)
	}
}
