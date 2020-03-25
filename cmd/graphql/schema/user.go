package schema

import (
	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"soqt/go-dudutalk/internal/user"
)

func (s *Server) registerUser(schema *schemabuilder.Schema) {
	schema.Object("User", user.User{})
}
