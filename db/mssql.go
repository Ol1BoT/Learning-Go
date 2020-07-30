package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	id    int
	title string
	name  string
)

type communityMember struct {
	ID    int
	Title string
	Name  string
}

func main() {
	fmt.Println("Database Connection to MSSQL")

	server := "redacted server"
	port := 1433 // MSSQL port
	user := "redacted user"
	password := "redatected password"
	database := "redacted database"

	// sql.Open("placeholder", "placeholder")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)

	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatalln("DB Connection failed", err.Error())
	}

	defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println("Ping Failed")
	// }

	// stmt, err := db.Prepare("select ID,NameExternal from dbo.community where id=52804")

	community := make([]communityMember, 0)

	stmt, err := db.Prepare("select Title,ID,NameExternal from dbo.community where surname='Mignot' order by title desc")
	if err != nil {
		log.Fatalln("Failed to prepapre SQL statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatalln("Failed to return rows from SQL statement")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&title, &id, &name)
		if err != nil {
			log.Fatal(err)
		}
		community = append(community, communityMember{
			id,
			title,
			name,
		})
		// fmt.Println(title, id, name)
	}

	// fmt.Println(community)
	// fmt.Printf("%+v\n", community)
	fmt.Println(prettyPrint(community))

}

func prettyPrint(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	return string(s)
}
