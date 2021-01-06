package handler

import (
	"context"
	"fmt"
	"go-mongo/config"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	ID  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string            `json:"title,omitempty" bson:"title,omitempty"`
	Description  string     `json:"description,omitempty" bson:"description,omitempty"`
	Price  int     `json:"price,omitempty" bson:"price,omitempty"`
	CreatedAt time.Time   `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	          //primitive.Timestamp   //more type available on timestamp  	
}

//context
var ctx = func() context.Context {
	return context.Background()
//	return context.WithTimeout(context.Background(), 10*time.Second)
}()


func CreateProduct(c *gin.Context)  {

	db, _ := config.Connect()

	//var reqBody Product 
	product := new(Product)
	if err :=c.BindJSON(&product); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
		}
	  	//t := time.Now()
			//fmt.Println(t.Format(time.ANSIC))
			
		fmt.Println(product.Title)
		
		res,err :=db.Collection("product").InsertOne(ctx, bson.M{"title": product.Title, "description": product.Description, "price": product.Price, "createdAt": time.Now(),
		"updatedAt": time.Now(),})
		
		if config.Error(c, err) {
			return //exit
		}
		//fmt.Println(res.InsertedID)

	_ = db.Collection("product").FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&product)

		c.JSON(200, gin.H{
		"message": "successfully added product",
		"data": product,
		})
	}


	func GetProducts(c *gin.Context)  {
		db, _ := config.Connect()
	
		cur, err := db.Collection("product").Find(ctx, bson.D{})
	
	fmt.Println(cur)
	//  if err != nil {
	// 	log.Fatal(err)
	// 	c.JSON(404, gin.H{
	// 		"error":   true,
	// 		"message": "something went wrong",
	// 	})
	// 	return
	// }	
  if config.Error(c, err) { //hndling with global error 
		return //exit
	}
	defer cur.Close(ctx)

	result := make([]Product, 0)
	for cur.Next(ctx) {
		var row Product
		err := cur.Decode(&row)
		// if err != nil {
		// 	c.JSON(404, gin.H{
		// 		"error":   true,
		// 		"message": "something went wrong",
		// 	})
		// 	return
		// }
		if config.Error(c, err) { //hndling with global error 
			return //exit
		}

		result = append(result, row)
	}
	fmt.Println(result)
			c.JSON(200, gin.H{
			"message": "get all products",
			"data": result,
			})		
}
	
func SingleProduct(c *gin.Context)  {
	db, _ := config.Connect()

	id := c.Param("id")
	fmt.Println(id)
	
	_id, _ := primitive.ObjectIDFromHex(id)
	fmt.Println(_id)

	product := new(Product)

  err := db.Collection("product").FindOne(ctx, bson.M{"_id": _id}).Decode(&product)

  fmt.Println(*product)
	
	if err != nil {
	log.Fatal(err)
	c.JSON(404, gin.H{
		"error":   true,
		"message": "not found",
	})
	return
}

c.JSON(200, gin.H{
	"message": "success",
	"data": product,
})
}

	
func UpdateProduct(c *gin.Context)  {
	db, _ := config.Connect()

	product := new(Product)
	if err :=c.BindJSON(&product); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
		}

	id := c.Param("id")
	_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": _id}

  _,err := db.Collection("product").UpdateOne(ctx, filter, bson.M{"$set": product})
	if err != nil {
		c.JSON(404, gin.H{
			"error":   true,
			"message": "something went wrong",
		})
		return
	}
  err2 := db.Collection("product").FindOne(ctx, filter).Decode(&product)

if config.Error(c, err2) { //hndling with global error 
	return //exit
}

c.JSON(200, gin.H{
	"message": "succesfully updated",
	"data": product,
})

}

func DeleteProduct(c *gin.Context)  {
	db, _ := config.Connect()

	id := c.Param("id")
	_id, _ := primitive.ObjectIDFromHex(id)

  _, err := db.Collection("product").DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
	log.Fatal(err)
	c.JSON(404, gin.H{
		"error":   true,
		"message": "not found",
	})
	return
}



c.JSON(200, gin.H{
	"message": "success",
	"data": gin.H{"_id": id},
})
}


func LatestProducts(c *gin.Context)  {
	db, _ := config.Connect()

	opts := options.Find()
  opts.SetSort(bson.D{{"createdAt", -1}})

	sortCursor, err := db.Collection("product").Find(ctx, bson.D{{"price" , bson.D{{"$gt", 500}}}}, opts)
	
	if err != nil {
	log.Fatal(err)
	c.JSON(404, gin.H{
		"error":   true,
		"message": "something went wrong",
	})
	return
}

var productsSorted []bson.M//short way of returning array object
if err = sortCursor.All(ctx, &productsSorted); err != nil {
    log.Fatal(err)
}
fmt.Println(productsSorted)
c.JSON(200, gin.H{
	"message": "success",
	"data": productsSorted,
})
}

func PriceBasedProducts(c *gin.Context)  {
	db, _ := config.Connect()

	price := c.Query("lowPrice")  
//	price := c.Request.URL.Query().Get("lowPrice")

  i1, err := strconv.Atoi(price)
	 if err != nil {
		log.Fatal(err);
	}
	//fmt.Println(i1)
		upPrice := c.DefaultQuery("upPrice", "10000")//default query value 10 return if Does not find int the req.query
		i2, err := strconv.Atoi(upPrice)
	 if err != nil {
		log.Fatal(err);
	}
	//fmt.Println(upprice)

	
	sortCursor, err := db.Collection("product").Find(ctx, bson.D{{"price" , bson.D{{"$gt", i1},{"$lt", i2}}}})
	
	if err != nil {
	log.Fatal(err)
	c.JSON(404, gin.H{
		"error":   true,
		"message": "something went wrong",
	})
	return
}

var productsPriceBased []bson.M//short way of returning array object
if err = sortCursor.All(ctx, &productsPriceBased); err != nil {
    log.Fatal(err)
}
fmt.Println(productsPriceBased)
c.JSON(200, gin.H{
	"message": "success",
	"data": productsPriceBased,
})
}


func TitleBasedProduct(c *gin.Context)  {
	db, _ := config.Connect()

	var product Product

  if c.BindQuery(&product) == nil { //another query binding method not body json
		fmt.Println(product.Title)
		fmt.Println("====== Only Bind Query String ======")
		
	}
	filter := bson.M{"title": product.Title}

 sortCursor, err := db.Collection("product").Find(ctx, filter)
	
	
	if err != nil {
	log.Fatal(err)
	c.JSON(404, gin.H{
		"error":   true,
		"message": "something went wrong",
	})
	return
}

var productstitleBased []bson.M//short way of returning array object
if err = sortCursor.All(ctx, &productstitleBased); err != nil {
    log.Fatal(err)
}
fmt.Println(productstitleBased)
c.JSON(200, gin.H{
	"message": "success",
	"data": productstitleBased,
})
}