package mongodb

import (
	"context"
	"os"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// variáveis global
var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

// construtor da conexão com o banco de dados
func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	// pegando o conteúdo do .env
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_USER_DB)

	client, err := mongo.Connect(options.Client().ApplyURI(mongodb_uri))

	// tratamento de erro
	if err != nil {
		return nil, err
	}

	// pingar pra testar
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	// instanciando banco de dados
	return client.Database(mongodb_database), nil
}
