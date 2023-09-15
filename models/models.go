package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID primitive.ObjectID 	`json:"_id" bson:"_id"`
	FirstName *string 		`json:"firstname" validate:"required, min=2, max=30"` 
	LastName *string 		`json:"lastname" validate:"required, min=2, max=30"`
	Password *string 		`json:"password" validate:"required, min=6"`
	Email *string 			`json:"email" validate:"required"`
	PhoneNumber *string 	`json:"phonenumber" validate:"required"`
	Token *string 			`json:"token"`
	RefreshToken *string 	`json:"refreshtoken"`
	Created_At time.Time 	`json:"createdat"`
	Updated_At time.Time 	`json:"updatedat"`
	User_ID string 			`json:"userid"`
	UserCart []ProductUser  `json:"usercart" bson:"usercart"`
	Address_Details []Address `json:"address" bson:"address"`
	Order_Status []Order	`json:"orders" bson:"orders"`
}

type Product struct{
	Product_ID primitive.ObjectID 	`bson:"_id"`
	Product_Name *string 			`json:"productname"`
	Price *uint64 					`json:"price"`
	Rating *uint8 					`json:"rating"`
	Image *string 					`json:"image"`
}

type ProductUser struct{
	Product_ID primitive.ObjectID 	`bson:"_id"`
	Product_Name *string 			`json:"productname" bson:"productname"`
	Price int 						`json:"price" bson:"price"`
	Rating *uint 					`json:"rating" bson:"rating"`
	Image *string 					`json:"image" bson:"image"`
}

type Address struct{
	Address_ID primitive.ObjectID 	`bson:"_id"`
	House *string 					`json:"house_name" bson:"house_name"`
	Street *string 					`json:"street_name" bson:"street_name"`
	City *string 					`json:"city_name" bson:"city_name"`
	State *string 					`json:"state_name" bson:"state_name"`
	Pincode *string 				`json:"pin_code" bson:"pin_code"`
}

type Order struct{
	Order_ID primitive.ObjectID 	`bson:"_id"`
	Order_Cart []ProductUser		`json:"order_list" bson:"order_list"`
	Ordered_At time.Time 			`json:"ordered_at" bson:"ordered_at"`
	Price float64 					`json:"price" bson:"price"`
	Discount *int					`json:"discount" bson:"discount"`
	Payment_Method Payment			`json:"payment_method" bson:"payment_method"`
}
type Payment struct{
	Digital bool 
	COD bool
}