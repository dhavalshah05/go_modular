package fireutil

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"google.golang.org/api/iterator"
)

func FromDocumentsToSlice[T any](documents *firestore.DocumentIterator) []T {
	var results []T
	for {
		next, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		var result T
		err = next.DataTo(&result)
		if err == nil {
			results = append(results, result)
		}
	}
	return results
}

func FromSnapShotToStruct[T any](snapshot *firestore.DocumentSnapshot) *T {
	var result T
	err := snapshot.DataTo(&result)
	if err != nil {
		return nil
	}
	return &result
}
