package apolloerror

import "github.com/vektah/gqlparser/gqlerror"

func ApolloError(message string, code string, properties *map[string]interface{}) error {
	extension := make(map[string]interface{})
	var invalidArgs []string
	if properties != nil {
		for k := range *properties {
			invalidArgs = append(invalidArgs, k)
		}
	}
	extension["code"] = code
	extension["invalidArgs"] = invalidArgs

	err := gqlerror.Error{
		Message:    message,
		Extensions: extension,
	}
	return &err
}

func SyntaxError(message string) error {
	msg := message
	if message == "" {
		msg = "GraphQL request syntax error"
	}
	return ApolloError(msg, "GRAPHQL_PARSE_FAILED", nil)
}

func ValidationError(message string) error {
	msg := message
	if message == "" {
		msg = "validation error"
	}
	return ApolloError(msg, "GRAPHQL_VALIDATION_FAILED", nil)
}

func AuthenticationError(message string) error {
	msg := message
	if message == "" {
		msg = "unauthenticated"
	}
	return ApolloError(msg, "UNAUTHENTICATED", nil)
}

func ForbiddenError(message string) error {
	msg := message
	if message == "" {
		msg = "forbidden request"
	}
	return ApolloError(msg, "FORBIDDEN", nil)
}

func PersistedQueryNotFoundError() error {
	return ApolloError("Persisted query not found", "PERSISTED_QUERY_NOT_FOUND", nil)
}

func PersistedQueryNotSupportedError() error {
	return ApolloError("Persisted query not support", "PERSISTED_QUERY_NOT_SUPPORTED", nil)
}

func UserInputError(message string, properties *map[string]interface{}) error {
	msg := message
	if message == "" {
		msg = "input error"
	}
	return ApolloError(msg, "BAD_USER_INPUT", properties)
}
