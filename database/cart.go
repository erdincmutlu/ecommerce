package database

import (
	"context"
	"errors"
	"log"

	"github.com/erdincmutlu/ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't decode product")
	ErrUserIdIsNotValid   = errors.New("this user id is not valid")
	ErrCantUpdateUser     = errors.New("can't add this rpodict to the cart")
	ErrCantRemoveItemCart = errors.New("can't remove this item from the cart")
	ErrCantGetItem        = errors.New("unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("can't update the purchase")
)

func AddProductToCart(ctx context.Context, prodCollection *mongo.Collection,
	userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {

	searchFromDB, err := prodCollection.Find(ctx, bson.M{"_id": productID})
	if err != nil {
		log.Println(err.Error())
		return ErrCantFindProduct
	}
	var productCart []models.ProductUser
	err = searchFromDB.All(ctx, &productCart)
	if err != nil {
		log.Println(err.Error())
		return ErrCantDecodeProducts
	}
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err.Error())
		return ErrUserIdIsNotValid
	}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "usercart", Value: bson.D{{Key: "$each", Value: productCart}}}}}}
	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err.Error())
		return ErrCantUpdateUser
	}

	return nil
}

func RemoveCartItem() {

}

func BuyItemFromCart() error {

}

func InstantBuyer() {

}
