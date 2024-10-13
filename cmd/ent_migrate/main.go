package main

import (
	"context"
	"log"

	"river0825/cleanarchitecture/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx := context.Background()
	CreateGraph(ctx, client)
	QueryGithub(ctx, client)

	// user, _ := CreateCars(context.Background(), client)
	// QueryCarUsers(context.Background(), user)
	// CreateUser(context.Background(), client)
	// QueryUser(context.Background(), client)
}
