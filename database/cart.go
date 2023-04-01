package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can't find the cart item")
	ErrCantDecodeProducts = errors.New("can't decode the cart item")
	ErrUserIdInvalid      = errors.New("the user ID is not invalid")
	ErrCantUpdateUser     = errors.New("cannot add this product to the cart")
	ErrCantRemoveCartItem = errors.New("cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("not able to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstanBuyer() {

}
