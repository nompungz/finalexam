package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	fmt.Print("database url :",os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("db connect error :", err)
		return nil, err
	}
	createTb := `
	CREATE TABLE IF NOT EXISTS CUSTOMER(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);
	`
	_, err = db.Exec(createTb)
	return db, err
}
func CreateCustomer(name string, email string, status string) (*sql.Row, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println("db connect error :", err)
		return nil, err
	}
	defer db.Close()
	query := `INSERT INTO CUSTOMER (NAME,EMAIL,STATUS) VALUES ($1,$2,$3) RETURNING ID,NAME,EMAIL,STATUS`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Prepare Statement error:", err)
		return nil, err
	}
	rows := stmt.QueryRow(name, email, status)
	return rows, err
}

func GetCustomerById(id int) (*sql.Row, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println("db connect error :", err)
		return nil, err
	}
	defer db.Close()
	query := `SELECT ID,NAME,EMAIL,STATUS FROM	CUSTOMER WHERE ID = $1`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Prepare Statement error:", err)
		return nil, err
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Println("Query error :", err)
		return nil, err
	}
	return row, err
}

func GetCustomers() (*sql.Rows, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println("db connect error :", err)
		return nil, err
	}
	defer db.Close()
	query := `SELECT ID,NAME,EMAIL,STATUS FROM	CUSTOMER `
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Prepare Statement error:", err)
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Println("Query error :", err)
		return nil, err
	}
	return rows, err
}

func UpdateCustomer(id int, name string, email string, status string) (*sql.Row, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println("db connect error :", err)
		return nil, err
	}
	defer db.Close()
	query := `UPDATE CUSTOMER SET NAME = $2 , EMAIL = $3 ,STATUS = $4 WHERE ID = $1 RETURNING ID,NAME,EMAIL,STATUS`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Prepare Statement error:", err)
		return nil, err
	}
	row := stmt.QueryRow(id, name, email, status)
	return row, err
}
func DeleteCustomer(id int) (error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println("db connect error :", err)
		return err
	}
	defer db.Close()
	query := `DELETE FROM CUSTOMER WHERE ID = $1 `
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Prepare Statement error:", err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("Query error :", err)
		return err
	}
	return err
}
