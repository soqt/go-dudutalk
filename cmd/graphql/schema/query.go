package schema

import (
	"context"

	"github.com/samsarahq/thunder/graphql/schemabuilder"
	// "soqt/go-dudutalk/internal/account"
	// "soqt/go-dudutalk/internal/service/userService"
)

func (s *Server) registerQuery(schema *schemabuilder.Schema) {
	object := schema.Query()

	object.FieldFunc("test", func(ctx context.Context) (string, error) {
		info := "hi"
		return info, nil
	})

	// object.FieldFunc("me", func(ctx context.Context)(*account.User, error) {
	// 	userId, err := s.Permission.IsAuthenticated(ctx)
	// 	if err != nil {
	// 		return nil, apolloerror.AuthenticationError("You need to login")
	// 	}
	// 	id := fmt.Sprintf("%v", userId)

	// 	user, err := s.AccountManaging.GetMe(ctx, &id)

	// 	if err == userService.RecordNotFound {
	// 		return nil, apolloerror.AuthenticationError("User not exists")
	// 	}

	// 	return user, nil
	// })
}
