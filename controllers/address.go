package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/erdincmutlu/ecommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			c.Abort()
			return
		}

		address, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal Server Error")
		}

		var addresses models.Address
		addresses.AddressID = primitive.NewObjectID()
		err = c.BindJSON(&addresses)
		if err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		matchFilter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$address"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$address_id"},
			{Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}},
		}
		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{matchFilter, unwind, group})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
		}

		var addressInfo []bson.M
		err = pointCursor.All(ctx, &addressInfo)
		if err != nil {
			panic(err)
		}

		var size int32
		for _, addressNo := range addressInfo {
			count := addressNo["count"]
			size = count.(int32)
		}
		if size >= 2 {
			c.IndentedJSON(http.StatusBadRequest, "Not allowed")
		}

		filter := bson.D{primitive.E{Key: "_id", Value: address}}
		update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Println(err.Error())
		}

		ctx.Done()
	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid"})
			c.Abort()
			return
		}

		userIDH, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal Server Error")
		}

		var editAddress models.Address
		err = c.BindJSON(&editAddress)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: userIDH}}
		update := bson.D{{Key: "$set", Value: bson.D{
			primitive.E{Key: "address.0.house_name", Value: editAddress.House},
			{Key: "address.0.street_name", Value: editAddress.Street},
			{Key: "address.0.city_name", Value: editAddress.City},
			{Key: "address.0.pin_code", Value: editAddress.PinCode},
		}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong")
			return
		}
		ctx.Done()
		c.IndentedJSON(http.StatusOK, "Successfully updated the home address")
	}
}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid"})
			c.Abort()
			return
		}

		userIDH, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal Server Error")
		}

		var editAddress models.Address
		err = c.BindJSON(&editAddress)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: userIDH}}
		update := bson.D{{Key: "$set", Value: bson.D{
			primitive.E{Key: "address.1.house_name", Value: editAddress.House},
			{Key: "address.1.street_name", Value: editAddress.Street},
			{Key: "address.1.city_name", Value: editAddress.City},
			{Key: "address.1.pin_code", Value: editAddress.PinCode},
		}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong")
			return
		}
		ctx.Done()
		c.IndentedJSON(http.StatusOK, "Successfully updated the work address")
	}
}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid search index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)
		userIDH, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal Server Error")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: userIDH}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "wrong command")
			return
		}

		ctx.Done()
		c.IndentedJSON(http.StatusOK, "Successfully deleted")
	}
}
