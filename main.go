package main

import (
	"fmt"
	"github.com/j-kk/go-graphql/graph/dtb"
	"github.com/j-kk/go-graphql/graph/model"
	"os"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == "" {
		fmt.Fprintf(os.Stderr, "database_url unset")
		os.Exit(1)
	}
	db, err := dtb.InitDB(dbUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot connect to dtb")
		os.Exit(1)
	}

	defer db.CloseDB()

	income := 100
	//year := 1990
	interests := []string{"Nowe", "Klocki", "Lego"}
	gender := model.GenderM
	var u1 = model.User{Income: &income, Gender: &gender, Interests: interests}
	err = db.AddUser(&u1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dtb addUser, %v", err)
		os.Exit(1)
	} else {
		fmt.Printf("OK\n")
	}
}
