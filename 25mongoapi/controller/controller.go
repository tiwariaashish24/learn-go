package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodb/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://kumarashishicould_db_user:LxFl2ZKePaYa6fof@learngolang.w0t9sr3.mongodb.net/?retryWrites=true&w=majority&appName=learngolang"
const dbName = "netflix"
const colName = "watchlist"

//most inportant

var collection *mongo.Collection //you are declaring a variable named collection that stores a pointer to a mongo.Collection value.

func init() {
	clientOption := options.Client().ApplyURI(connectionString).SetServerSelectionTimeout(5 * time.Second) //cotnrol mongodb operatios

	//connect mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Mongodb Connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collectin instance is ready")
}

func ErrorHandling(err error) {
	if err != nil {
		panic(err)
	}
}

//MONGODB - helper file
//inserti 1 movie

func insertOneMovie(movie model.Netflix) error { //you can import netflix from modelfile
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		return err
	}
	fmt.Println("Inserted 1 movie in db with id", inserted.InsertedID)
	return nil
}

// update one movie
func UpdateOneMovie(movieId string) error {
	_id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"Watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("modify count", result.ModifiedCount)
	return nil
}

func ErrorFatel(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// delete one movie
func deleteOneMovie(movieID string) error {
	_id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": _id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	fmt.Println("Movie got deleted with count: ", deleteCount)
	return nil
}

// delete all record in mongodb
func deleteAllMovie() (int64, error) {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	fmt.Println("The Number of Movie deleted", deleteResult.DeletedCount)
	return deleteResult.DeletedCount, nil
}

//all the movie form database

func getAllTheMovive() ([]primitive.M, error) {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	var movies []primitive.M

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)

	}
	return movies, nil
}

func GetMyAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //use x-wwwform-urlencode in the place of json
	allMovie, err := getAllTheMovive()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get movies: " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(allMovie)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow_method", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	err := insertOneMovie(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to insert movie: " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(movie)
}

func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-control-Allow-Methods", "POST")

	params := mux.Vars(r)
	err := UpdateOneMovie(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update movie: " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")

	params := mux.Vars(r)

	err := deleteOneMovie(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete movie: " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")

	count, err := deleteAllMovie()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete all movies: " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(count)

}

/*
1. context is used to control MongoDB operations.
2. mongo.Collection represents a MongoDB collection.
3. options.Client().ApplyURI() sets the MongoDB connection URL.
4. mongo.Connect() creates a MongoDB client.
5. ErrorHandling() checks errors and stops the program if an error occurs.
6. client.Database(dbName) selects the database.
7. Collection(colName) selects the collection.
8. collection stores the collection reference for database operations.
9. init() runs automatically when the package initializes.
10. Use client.Ping() to check whether MongoDB is reachable.
11. ObjectIDFromHex() convert stinng into objectId  in mongodb
12. bson.M — Order doesn't matter
13.	bson.D — Order matters
*/
