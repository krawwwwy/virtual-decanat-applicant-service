package main

import (
	"applicant-service/utils"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing the connection: %v", err)
		}
	}(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Println("Error parsing template: %v", err)
		} else {
			tmpl.Execute(w, nil)
		}
	})
	fmt.Println("Successfully run conteiner")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
