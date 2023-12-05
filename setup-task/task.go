package main

import (
	"clean-architecture-todo-webapp-by-golang/db"
	"fmt"
	"log"
)

func main() {
	dbCnt := db.Newdb()

	defer db.CloseDB(dbCnt)

	_, err := dbCnt.Exec(`CREATE TABLE IF NOT EXISTS todos (
						id INT AUTO_INCREMENT PRIMARY KEY,
						userId INT NOT NULL,
						todo VARCHAR(255) NOT NULL,
						deleted BOOLEAN DEFAULT FALSE NOT NULL,
						created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
						updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
					)`)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("task_data table created successfully")
}
