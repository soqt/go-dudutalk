package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"soqt/go-dudutalk/cmd/graphql/schema"
	"soqt/go-dudutalk/internal/authenticating"
	"soqt/go-dudutalk/internal/storage/prisma"
	"soqt/go-dudutalk/pkg/gc"
	"soqt/go-dudutalk/pkg/jwt"
)

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlHandler(s *schema.Server) gin.HandlerFunc {
	graphqlSchema := s.Schema()
	introspection.AddIntrospectionToSchema(graphqlSchema)

	h := graphql.HTTPHandler(graphqlSchema)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// GetGinEngine register resolvers to http API routes
func GetGinEngine(server *schema.Server, ginMode string) *gin.Engine {
	gin.SetMode(ginMode)
	r := gin.Default()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.GET("/", playgroundHandler())
	r.Use(jwt.JWTMiddleware(false, os.Getenv("JWT_SECRET")))
	r.Use(gc.GinContextToContextMiddleware())
	r.POST("/graphql", graphqlHandler(server))
	return r
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Storage initiation
	prismaStorage, err := prisma.NewStorage()

	authS := authenticating.NewService(prismaStorage)

	server := &schema.Server{
		Authenticating: authS,
	}
	err = GetGinEngine(server, gin.DebugMode).Run()
	if err != nil {
		panic("Server Not Running")
	}
}
