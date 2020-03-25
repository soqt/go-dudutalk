package schema

import (
	"soqt/go-dudutalk/internal/authenticating"

	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

type Server struct {
	Authenticating authenticating.Service
}

func (s *Server) Schema() *graphql.Schema {
	schema := schemabuilder.NewSchema()
	s.registerQuery(schema)
	s.registerMutation(schema)

	s.registerUser(schema)
	return schema.MustBuild()
}
