package main

import (
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestAA(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatal("not equal")
	}
	connStr := "user=postgres dbname=postgres password=secret sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	var res string
	err = db.Get(&res, "SELECT version()")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)

}
