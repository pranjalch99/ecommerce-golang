package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	First_Name      *string            `json:"firstName,omitempty" bson:"firstName,omitempty" validate:"required,min=2,max=30"`
	Last_Name       *string            `json:"lastName,omitempty" bson:"lastName,omitempty" validate:"required,min=2,max=30"`
	Password        *string            `json:"password,omitempty" bson:"password,omitempty" validate:"required,min=6"`
	Email           *string            `json:"email,omitempty" bson:"email,omitempty" validate:"email,required"`
	Phone           *string            `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	Token           *string            `json:"token,omitempty" bson:"token,omitempty"`
	Refresh_Token   *string            `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	Created_At      time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Updated_At      time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	User_ID         *string            `json:"userId,omitempty" bson:"userId,omitempty"`
	UserCart        []ProductUser      `json:"userCart,omitempty" bson:"userCart,omitempty"`
	Address_Details []Address          `json:"address,omitempty" bson:"address,omitempty"`
	Order_Details   []Order            `json:"orders,omitempty" bson:"orders,omitempty"`
}

type Product struct {
	Product_ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Product_Name *string            `json:"productName,omitempty" bson:"productName,omitempty"`
	Price        *uint64            `json:"price,omitempty" bson:"price,omitempty"`
	Rating       *uint8             `json:"rating,omitempty" bson:"rating,omitempty"`
	Image        *string            `json:"image,omitempty" bson:"image,omitempty"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Product_Name *string            `json:"productName,omitempty" bson:"productName,omitempty"`
	Price        int                `json:"price,omitempty" bson:"price,omitempty"`
	Rating       *uint              `json:"rating,omitempty" bson:"rating,omitempty"`
	Image        *string            `json:"image,omitempty" bson:"image,omitempty"`
}

type Address struct {
	Address_ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	House      *string            `json:"house,omitempty" bson:"house,omitempty"`
	Street     *string            `json:"street,omitempty" bson:"street,omitempty"`
	City       *string            `json:"city,omitempty" bson:"city,omitempty"`
	Pincode    *string            `json:"pincode,omitempty" bson:"pincode,omitempty"`
}

type Order struct {
	Order_ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Order_Cart     []ProductUser      `json:"orderCart,omitempty" bson:"orderCart,omitempty"`
	Ordered_At     time.Time          `json:"orderedAt,omitempty" bson:"orderedAt,omitempty"`
	Price          int                `json:"price,omitempty" bson:"price,omitempty"`
	Discount       *int               `json:"discount,omitempty" bson:"discount,omitempty"`
	Payment_Method Payment            `json:"paymentMethod,omitempty" bson:"paymentMethod,omitempty"`
}

type Payment struct {
	Digital bool `json:"digital,omitempty" bson:"digital,omitempty"`
	COD     bool `json:"cod,omitempty" bson:"cod,omitempty"`
}
