package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type propinsi struct {
	ID    int
	NAMA_PROPINSI  string
}

func (item propinsi) String() string {
	return fmt.Sprintf("%d, %s, %s, %d, %s", item.ID, item.NAMA_PROPINSI)
}

func main() {
	db, err := sql.Open("sqlite3", "./TekCloudResponsi.db")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected")
		// var id = 1
		rows, err := db.Query("select * from propinsi ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		var result []propinsi

		for rows.Next() {
			var each = propinsi{}
			var err = rows.Scan(&each.ID, &each.NAMA_PROPINSI)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result = append(result, each)
		}

		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, each := range result {
			fmt.Println(each.String())
		}
	}
	defer db.Close()
}