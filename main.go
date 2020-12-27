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

	// POST add product to cart
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

// GET all products in cart
router.GET("/cart", func(c *gin.Context) {
	var (
		product  Product
		products []Product
		cart	Cart
	)
	rows, err := db.Query("select product_id from cart where checkout != 1;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&cart.Product_id)
		row := db.QueryRow("select id, name, category_id from product where id = ?;", cart.Product_id)
		err = row.Scan(&product.Id, &product.Name, &product.Category_id)
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

	// POST remove product in cart
	router.POST("/removecart", func(c *gin.Context) {
		product_id := c.PostForm("product_id")
		stmt, err := db.Prepare("delete from cart where product_id = ?;")
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
			"message": fmt.Sprintf("%s successfully removed from cart", product.Name),
		})
	})

	// POST checkout product in cart
	router.POST("/checkoutcart", func(c *gin.Context) {
		row := db.QueryRow("select COUNT(id) from cart where checkout != 1")
		var count int
		row.Scan(&count)
		if count <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Please Add Product to Cart first"),
			})
		}else{
			stmt, err := db.Prepare("UPDATE cart SET checkout = ?;")
			if err != nil {
				fmt.Print(err.Error())
			}
			_, err = stmt.Exec(1)
	
			if err != nil {
				fmt.Print(err.Error())
			}
	
			defer stmt.Close()
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("You successfully checkout please continue to the payment section"),
			})
		}
	})

	router.Run(":3000")
}