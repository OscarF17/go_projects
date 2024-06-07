package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {

	dsn := "host=localhost user=tienda_user dbname=tienda password=666 port=8080 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

out:
	for true {
		fmt.Println("Welcome to the database system")
		fmt.Println("[0] View all products")
		fmt.Println("[1] View a product")
		fmt.Println("[2] Create a product")
		fmt.Println("[3] Delete a product")
		fmt.Println("[4] Update a product")
		fmt.Println("[5] Exit")
		var input int

		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}
		switch input {
		case 0:
			viewAllProducts(db)
		case 1:
			viewProduct(db)
		case 2:
			createProduct(db)
		case 3:
			deleteProduct(db)
		case 4:
			updateProduct(db)
		case 5:
			break out
		default:
			fmt.Println("Input a valid option")
		}
	}
	fmt.Println("Goodbye...")

	// Example: Executing raw SQL queries
	// executeRawSQL(db)
	// callStoredProcedure(db)
}

func viewAllProducts(db *gorm.DB) {
	rows, err := db.Raw("SELECT * FROM productos ORDER BY id").Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombre string
		var precio float64
		err = rows.Scan(&id, &nombre, &precio)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Price: %f\n", id, nombre, precio)
	}
}

func viewProduct(db *gorm.DB) {
	fmt.Println("Insert product ID to view")
	var id int
	_, _ = fmt.Scan(&id)
	rows, err := db.Raw("SELECT * FROM get_producto(?)", id).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombre string
		var precio float64
		err = rows.Scan(&id, &nombre, &precio)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, nombre)
	}
}

// otra forma de llamar a procedimientos, devuelve un solo valor pero lo atrapo en forma de string
// esto no es conveniente ya que el string debe de ser separado
// puede ser util si no conozco el numero de parametros en la salida
func viewProduct2(db *gorm.DB) {
	rows, err := db.Raw("SELECT get_producto(?)", 1).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var values string
		err = rows.Scan(&values)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Values: %s\n", values)
	}
}

func createProduct(db *gorm.DB) {
	fmt.Println("Insert product name")
	var name string
	_, _ = fmt.Scan(&name)
	fmt.Println("Insert product price")
	var price float32
	_, _ = fmt.Scan(&price)
	rows, err := db.Raw("CALL create_producto($1, $2)", name, price).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Product created")

}

func deleteProduct(db *gorm.DB) {
	fmt.Println("Insert product ID to delete")
	var id int
	_, _ = fmt.Scan(&id)
	rows, err := db.Raw("CALL delete_producto($1);", id).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Printf("Product %d deleted\n", id)
}

func updateProduct(db *gorm.DB) {
	fmt.Println("Insert product ID to ")
	var id int
	_, _ = fmt.Scan(&id)
	fmt.Println("Insert product name")
	var name string
	_, _ = fmt.Scan(&name)
	fmt.Println("Insert product price")
	var price float32
	_, _ = fmt.Scan(&price)
	rows, err := db.Raw("CALL update_producto($1, $2, $3);", id, name, price).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Printf("Product %d updated\n", id)
}
func executeRawSQL(db *gorm.DB) {
	// Example: Simple SELECT query

}

func callStoredProcedure(db *gorm.DB) {
	// Assuming you have a stored procedure named "get_user" that takes an integer and returns user details
}

func queryWithoutStruct(db *gorm.DB) {
	rows, err := db.Raw("SELECT id, name FROM users").Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}
