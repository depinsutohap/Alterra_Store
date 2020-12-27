package main

import (
	// "bytes"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/alterra_store")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Product struct {
		Id         int
		Name 	string
		Category_id int
	}
	type Category struct {
		Id         int
		Name 	string
	}
	type Cart struct {
		Id         int
		Product_id int
		Qty int
		Checkout int
	}
	router := gin.Default()

// GET all products
router.GET("/products", func(c *gin.Context) {
	var (
		product  Product
		products []Product
	)
	rows, err := db.Query("select id, name, category_id from product;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Category_id)
		products = append(products, product)
		if err != nil {
			fmt.Print(err.Error())
		}
	} 
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": products,
		"count":  len(products),
	})
})

	// POST new person details
	router.POST("/addcart", func(c *gin.Context) {
		product_id := c.PostForm("product_id")
		stmt, err := db.Prepare("insert into cart (product_id) values(?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(product_id)

		if err != nil {
			fmt.Print(err.Error())
		}

		defer stmt.Close()
		row := db.QueryRow("select name from product where id = ?;", product_id)
		var product Product
		err = row.Scan(&product.Name)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s successfully added to cart", product.Name),
		})
	})

	// // GET a person detail
	// router.GET("/person/{:id}", func(c *gin.Context) {
	// 	var (
	// 		person Person
	// 		result gin.H
	// 	)
	// 	id := c.Param("id")
	// 	row := db.QueryRow("select id, first_name, last_name from person where id = ?;", id)
	// 	err = row.Scan(&person.Id, &person.First_Name, &person.Last_Name)
	// 	if err != nil {
	// 		// If no results send null
	// 		result = gin.H{
	// 			"result": nil,
	// 			"count":  0,
	// 		}
	// 	} else {
	// 		result = gin.H{
	// 			"result": person,
	// 			"count":  1,
	// 		}
	// 	}
	// 	c.JSON(http.StatusOK, result)
	// })


	// // PUT - update a person details
	// router.PUT("/person", func(c *gin.Context) {
	// 	var buffer bytes.Buffer
	// 	id := c.Query("id")
	// 	first_name := c.PostForm("first_name")
	// 	last_name := c.PostForm("last_name")
	// 	stmt, err := db.Prepare("update person set first_name= ?, last_name= ? where id= ?;")
	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 	}
	// 	_, err = stmt.Exec(first_name, last_name, id)
	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 	}

	// 	// Fastest way to append strings
	// 	buffer.WriteString(first_name)
	// 	buffer.WriteString(" ")
	// 	buffer.WriteString(last_name)
	// 	defer stmt.Close()
	// 	name := buffer.String()
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": fmt.Sprintf("Successfully updated to %s", name),
	// 	})
	// })

	// // Delete resources
	// router.DELETE("/person", func(c *gin.Context) {
	// 	id := c.Query("id")
	// 	stmt, err := db.Prepare("delete from person where id= ?;")
	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 	}
	// 	_, err = stmt.Exec(id)
	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": fmt.Sprintf("Successfully deleted user: %s", id),
	// 	})
	// })

	router.Run(":3000")
}