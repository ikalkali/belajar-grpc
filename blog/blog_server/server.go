package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type User struct {
	ID     int64
	Name   string
	Emails []string
}

type Story struct {
	ID       int64
	Title    string
	AuthorID int64
	Author   *User `bun:"rel:belongs-to"`
}

func main() {
	dsn := "postgres://postgres:root@localhost:5432/grpc-blog?sslmode=disable"
	ctx := context.Background()

	// sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn), pgdriver))
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	err := db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to Postgre: %v", err)
	}
	fmt.Println("Connected to Postgres!")

	db.RegisterModel((*User)(nil), (*Story)(nil))
	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	if err := fixture.Load(ctx, os.DirFS("./blog/blog_server"), "fixture.yaml"); err != nil {
		panic(err)
	}

	users := make([]User, 0)
	if err := db.NewSelect().Model(&users).OrderExpr("id ASC").Scan(ctx); err != nil {
		panic(err)
	}
	fmt.Printf("all users: %v\n\n", users)

	defer sqldb.Close()
}
