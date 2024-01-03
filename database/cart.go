package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't decode product")
	ErrUserIdIsNotValid   = errors.New("this user id is not valid")
	ErrCantUpdateUser     = errors.New("can't add this rpodict to the cart")
	ErrCantRemoveItemCart = errors.New("can't remove this item from the cart")
	ErrCantGetItem        = errors.New("unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("can't update the purchase")
)

func AddProductToCart() error {

}

func RemoveCartItem() {

}

func BuyItemFromCart() error {

}

func InstantBuyer() {

}
