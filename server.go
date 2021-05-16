package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/medwing/insights-service/db"
	"github.com/medwing/insights-service/env"
	"github.com/medwing/insights-service/graph"
	"github.com/medwing/insights-service/graph/generated"
	"gorm.io/gorm"
)

const defaultPort = ":8081"

// Defining the Graphql handler
func graphqlHandler(db *gorm.DB) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func HealthGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()
	pgCli := db.PgClient{
		Config: &db.PgConfig{
			Host:     env.PostgresHost(),
			Port:     env.PostgresPort(),
			User:     env.PostgresUser(),
			Password: env.PostgresPassword(),
			DBName:   env.PostgresDBName(),
			Ssl:      env.PostgresSSL(),
		},
	}
	if err := pgCli.Connect(); err != nil {
		panic(err)
	}
	// db, err := database.GetDatabase()
	// if err != nil {
	// 	panic(fmt.Printf("Unable to connect to database, %v", err))
	// }
	// defer db.Close()
	r.POST("/query", graphqlHandler(pgCli.Orm()))
	r.GET("/healthz", HealthGET)
	r.GET("/", playgroundHandler())

	r.Run(defaultPort)
}
