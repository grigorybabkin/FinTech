package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func OpenConnectionDB() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func AddURLsDB(fullURL, shortURL string) {

	db := OpenConnectionDB()
	sqlStatement := `INSERT INTO urls (full_url, short_url) VALUES ($1, $2)`
	db.Exec(sqlStatement, fullURL, shortURL)

	defer db.Close()
}

func GetURLDB(query string, args ...interface{}) string {

	db := OpenConnectionDB()
	var responseURL string
	db.QueryRow(query, args...).Scan(&responseURL)

	defer db.Close()
	return responseURL
}

func RowExistsDB(query string, args ...interface{}) bool {

	db := OpenConnectionDB()
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	db.QueryRow(query, args...).Scan(&exists)
	//if err != nil && err != sql.ErrNoRows {
	//glog.Fatalf("error checking if row exists '%s' %v", args, err)
	//}
	defer db.Close()
	return exists
}
