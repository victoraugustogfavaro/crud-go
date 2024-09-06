package mongodb

import (
	"context"

	"github.com/victoraugustogfavaro/crud-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// iniciar conexão com o banco
func InitConnection() {
	ctx := context.Background()
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))

	// tratamento de erro
	if err != nil {
		panic(err)
	}

	// pingar pra testar
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	logger.Info("Connection with db successful")
}
