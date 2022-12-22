package fireutil

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirebaseClient *firestore.Client

func InitFirebaseApp(configBytes []byte) *firestore.Client {
	opt := option.WithCredentialsJSON(configBytes)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	return client
}
