package main

import (
	"subscriptions/core/error_capture"
	"subscriptions/core/mongo"
	"time"
)

func main() {
	defer error_capture.CaptureDefer()

	settings := mongo.
		NewSettings("mongodb://localhost:27017").
		WithTimeout(time.Second * 10)

	manager := mongo.NewManager(settings)

	println(manager.Random)
	//client, err := mongo.Connect(
	//	context.TODO(),
	//	options.Client().ApplyURI(uri))
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()
	//
	//coll := client.
	//	Database("wesubscriptions").
	//	Collection("movies")
	//
	//title := "Back to the Future"
	//
	//var result bson.M
	//err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).
	//	Decode(&result)
	//if err == mongo.ErrNoDocuments {
	//	fmt.Printf("No document was found with the title %s\n", title)
	//	return
	//}
	//if err != nil {
	//	panic(err)
	//}
	//jsonData, err := json.MarshalIndent(result, "", "    ")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", jsonData)
	//
	//println("Hello")
}
