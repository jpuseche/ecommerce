package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindProduct    = errors.New("couldn't find the Product")
	ErrCantDecodeProducts = errors.New("couldn't find the Product")
	ErrUserIdIsNotValid   = errors.New("this User is not valid")
	ErrCantUpdateUser     = errors.New("can't add the Product to the Cart")
	ErrCantRemoveItemCart = errors.New("can't remove this Item from Cart")
	ErrCantGetItem        = errors.New("can't get Item from Cart")
	ErrCantBuyCartItem    = errors.New("can't update Purchase")
)

func AddProductToCart(ctx context.Context, productCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userQueryID string) error {
	return nil
}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
