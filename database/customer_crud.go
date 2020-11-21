package database

import (
	"log"

	"github.com/Arashage/gofinal/customer"
)

func CreateCustomer() {
	createScript := `
	CREATE TABLE IF NOT EXISTS customer (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);
	`
	_, err := DB.Exec(createScript)

	if err != nil {
		log.Fatal("Can't create table customer.", err)
	}

}

func InsertCustomer(c *customer.Customer) error {
	insertScript := `
	INSERT INTO customer (name, email, status)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	row := DB.QueryRow(insertScript, c.Name, c.Email, c.Status)
	var id int
	err := row.Scan(&id)

	if err != nil {
		log.Fatal("Can't scan ID.", err)
	}

	c.ID = id

	return err
}

func UpdateCustomer(c *customer.Customer) error {
	updateScript := `
	UPDATE customer 
		SET name = $2, email = $3, status = $4 
		WHERE id = $1;
	`

	stmt, err := DB.Prepare(updateScript)
	if err != nil {
		log.Fatal("Can't prepare update statement", err)
		return err
	}

	_, err2 := stmt.Exec(c.ID, c.Name, c.Email, c.Status)
	if err2 != nil {
		log.Fatal("Can't execute update", err2)
	}
	return err2

}

func DeleteCustomer(id int) error {
	deleteScript := `
	DELETE FROM customer 
		WHERE id = $1;
	`

	stmt, err := DB.Prepare(deleteScript)
	if err != nil {
		log.Fatal("Can't prepare delete statement", err)
		return err
	}

	_, err2 := stmt.Exec(id)
	if err2 != nil {
		log.Fatal("Can't execute update", err2)
	}
	return err2
}

func GetAllCustomer() ([]customer.Customer, error) {
	queryAll := "SELECT id, name, email, status FROM customer"

	stmt, err := DB.Prepare(queryAll)
	if err != nil {
		log.Fatal("Can't prepare query all for customer.", err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Can't query all for customer.", err)
		return nil, err
	}

	customers := []customer.Customer{}
	for rows.Next() {
		c := customer.Customer{}
		err = rows.Scan(&c.ID, &c.Name, &c.Email, &c.Status)
		if err != nil {
			log.Fatal("Can't scan row into customer", err)
		}
		customers = append(customers, c)
	}

	return customers, err

}

func GetCustomerByID(id int) (customer.Customer, error) {
	queryAll := "SELECT id, name, email, status FROM customer WHERE id =$1"
	c := customer.Customer{}

	stmt, err := DB.Prepare(queryAll)
	if err != nil {
		log.Fatal("Can't prepare query by ID for customer.", err)
		return c, err
	}

	row := stmt.QueryRow(id)
	err = row.Scan(&c.ID, &c.Name, &c.Email, &c.Status)
	if err != nil {
		log.Fatal("Can't scan row into customer", err)
	}

	return c, err

}
