package schema

import (
	"context"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"soqt/go-dudutalk/internal/user"
	"soqt/go-dudutalk/pkg/apolloerror"
)

func (s *Server) registerMutation(schema *schemabuilder.Schema) {
	object := schema.Mutation()
	type AuthPayload struct {
		User  *user.User
		Token string
	}


	type OtpRequestInput struct {
		Phone string
	}

	type OtpPayload struct {
		Success bool
	}

	object.FieldFunc("requestOtp", func(ctx context.Context, args struct{ Input OtpRequestInput }) (*OtpPayload, error) {
		err := s.Authenticating.IssueOTP(ctx, &args.Input.Phone)
		if err != nil {
			return  nil, err
		}
		payload := OtpPayload{
			Success: true,
		}
		return &payload, nil
	})

	type LoginInput struct {
		Phone string
		Code string
	}

	object.FieldFunc("login", func(ctx context.Context, args struct{ Input LoginInput }) (*AuthPayload, error) {
		input := args.Input

		user, jwt, err := s.Authenticating.Login(ctx, input.Phone, input.Code)
		if err != nil {
			return nil, apolloerror.AuthenticationError("login failed" + err.Error())
		}

		payload := AuthPayload{
			User:  user,
			Token: *jwt,
		}

		return &payload, nil
	})
}
