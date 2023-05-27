package mongo

import (
	"context"
	"signupservice/src/models"
)

// Use for the add user to mongo db database.
func AddVerifyToMongo(verify models.VerifyCodeTemplate) {
	collection := getCollectionVerify()
	collection.InsertOne(context.Background(), verify)
}
