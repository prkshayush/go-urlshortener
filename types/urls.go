package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// struct for taking aceepting long url as json string
type ShortUrlBody struct{
	LonguRL string `json:"long_url"`
}

// database model struct
// omitempty that neans this file can be empty or mongoDB can automatically fill it
// format for new url that would be generated
//for personal uses createdAt and expiredAt
type UrlDB struct {
	ID 		 	primitive.ObjectID 	`bson:"_id, omitempty"`
	UrlCode  	string 				`bson:"url_code"`
	LongUrl  	string 				`bsom:"long_url"`
	ShortUrl 	string 				`bson:"short_url"`
	CreatedAt 	int64 				`bson:"created_at"`
	ExpireAt 	int64 				`bson:"expire_at"`
}
