package data

import (
	"database/sql"
	"log"

	"github.com/ajanthan/go-ecommerce-demo/CartService/model"
	_ "github.com/go-sql-driver/mysql"
)

//CartRepository exposes api for cart data
type CartRepository struct {
	Db *sql.DB
}

func (c *CartRepository) InitRepository(connectionStr string) {
	db, dbConnectionErr := sql.Open("mysql", connectionStr)
	if dbConnectionErr != nil {
		log.Fatal(dbConnectionErr)
	}
	c.Db = db
}
func (c *CartRepository) AddItemToCart(userID string, item model.Item) {
	_, queryErr := c.Db.Query("insert into cart (userID,productID,quantity) values(?,?,?)", userID, item.ProductID, item.Quantity)
	if queryErr != nil {
		log.Fatal(queryErr)
	}
}
func (c *CartRepository) GetCart(userID string) model.Cart {
	log.Println("user in repository ", userID)
	cart := model.Cart{}
	cart.UserID = userID
	rows, queryErr := c.Db.Query("select productID,quantity from cart where userID=?", userID)
	if queryErr != nil {
		log.Fatal(queryErr)
	}
	items := make([]model.Item, 0)
	for rows.Next() {
		item := model.Item{}
		var productID string
		var quantity int
		if err := rows.Scan(&productID, &quantity); err != nil {
			log.Fatal(err)
		}
		item.ProductID = productID
		item.Quantity = quantity
		log.Println("Item in repository ", item)
		items = append(items, item)

	}
	cart.Items = items
	log.Println("Cart in repository ", cart)
	return cart
}
func (c *CartRepository) EmptyCart(userID string) {
	_, queryErr := c.Db.Query("delete from cart where userID=?", userID)
	if queryErr != nil {
		log.Fatal(queryErr)
	}
}