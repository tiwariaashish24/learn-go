package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

/*
1. Netflix struct stores movie or series details.
2. ObjectID represents MongoDB's document ID.
3. Title stores the movie or series name.
4. Type stores movie or series.
5. Watched tracks watched status.
6. Seasons stores a list of seasons.
7. Episodes stores episode details.
8. bson tags define MongoDB field names.
9. json tags define JSON field names.
10. Each Netflix struct represents one document, not the whole collection.
*/
