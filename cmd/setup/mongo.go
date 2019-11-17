package setup

import (
	"aaa/internal/connections"
	"aaa/internal/store"
	"aaa/tools/env"
)

var (
	mongoAddress     = env.String("AAA_MONGO_ADDR", "mongodb://localhost:27017", "mongo address")
	mongoPingTimeout = env.Duration("AAA_MONGO_PING_TIMEOUT", 2, "mongo ping timeout")
	mongoDatabase    = env.String("AAA_MONGO_DATABASE", "default", "mongo database")
)

func Mongo() {
	mongoClient, err := connections.ConnectToMongo(mongoAddress, mongoPingTimeout)
	if err != nil {
		panic(err)
	}

	db := mongoClient.Database(mongoDatabase)

	s := store.NewStore()
	s.Person = store.NewPersonStore(db)
}
