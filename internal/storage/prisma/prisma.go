// Code generated by Prisma CLI (https://github.com/prisma/prisma). DO NOT EDIT.

package prisma

import (
	"context"
	"errors"

	"os"

	"github.com/prisma/prisma-client-lib-go"

	"github.com/machinebox/graphql"
)

var ErrNoResult = errors.New("query returned no result")

func Str(v string) *string { return &v }
func Int32(v int32) *int32 { return &v }
func Bool(v bool) *bool    { return &v }

type BatchPayloadExec struct {
	exec *prisma.BatchPayloadExec
}

func (exec *BatchPayloadExec) Exec(ctx context.Context) (BatchPayload, error) {
	bp, err := exec.exec.Exec(ctx)
	return BatchPayload(bp), err
}

type BatchPayload struct {
	Count int64 `json:"count"`
}

type Aggregate struct {
	Count int64 `json:"count"`
}

type Client struct {
	Client *prisma.Client
}

type Options struct {
	Endpoint string
	Secret   string
}

func New(options *Options, opts ...graphql.ClientOption) *Client {
	endpoint := DefaultEndpoint
	secret := Secret
	if options != nil {
		endpoint = options.Endpoint
		secret = options.Secret
	}
	return &Client{
		Client: prisma.New(endpoint, secret, opts...),
	}
}

func (client *Client) GraphQL(ctx context.Context, query string, variables map[string]interface{}) (map[string]interface{}, error) {
	return client.Client.GraphQL(ctx, query, variables)
}

var DefaultEndpoint = "http://" + os.Getenv("PRISMA_HOST") + ":4466/dudutalk/" + os.Getenv("PRISMA_STAGE")
var Secret = os.Getenv("PRISMA_SECRET")

func (client *Client) Otp(params OtpWhereUniqueInput) *OtpExec {
	ret := client.Client.GetOne(
		nil,
		params,
		[2]string{"OtpWhereUniqueInput!", "Otp"},
		"otp",
		[]string{"id", "phone", "code", "createdAt"})

	return &OtpExec{ret}
}

type OtpsParams struct {
	Where   *OtpWhereInput   `json:"where,omitempty"`
	OrderBy *OtpOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32           `json:"skip,omitempty"`
	After   *string          `json:"after,omitempty"`
	Before  *string          `json:"before,omitempty"`
	First   *int32           `json:"first,omitempty"`
	Last    *int32           `json:"last,omitempty"`
}

func (client *Client) Otps(params *OtpsParams) *OtpExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"OtpWhereInput", "OtpOrderByInput", "Otp"},
		"otps",
		[]string{"id", "phone", "code", "createdAt"})

	return &OtpExecArray{ret}
}

type OtpsConnectionParams struct {
	Where   *OtpWhereInput   `json:"where,omitempty"`
	OrderBy *OtpOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32           `json:"skip,omitempty"`
	After   *string          `json:"after,omitempty"`
	Before  *string          `json:"before,omitempty"`
	First   *int32           `json:"first,omitempty"`
	Last    *int32           `json:"last,omitempty"`
}

// Nodes return just nodes without cursors. It uses the already fetched edges.
func (s *OtpConnection) Nodes() []Otp {
	var nodes []Otp
	for _, edge := range s.Edges {
		nodes = append(nodes, edge.Node)
	}
	return nodes
}

// Nodes return just nodes without cursors, but as a slice of pointers. It uses the already fetched edges.
func (s *OtpConnection) NodesPtr() []*Otp {
	var nodes []*Otp
	for _, edge := range s.Edges {
		item := edge
		nodes = append(nodes, &item.Node)
	}
	return nodes
}

func (client *Client) OtpsConnection(params *OtpsConnectionParams) *OtpConnectionExec {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"OtpWhereInput", "OtpOrderByInput", "Otp"},
		"otpsConnection",
		[]string{"edges", "pageInfo"})

	return &OtpConnectionExec{ret}
}

func (client *Client) User(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.GetOne(
		nil,
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"user",
		[]string{"id", "phone", "nickname", "avatar", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

type UsersParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) Users(params *UsersParams) *UserExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"UserWhereInput", "UserOrderByInput", "User"},
		"users",
		[]string{"id", "phone", "nickname", "avatar", "createdAt", "updatedAt"})

	return &UserExecArray{ret}
}

type UsersConnectionParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

// Nodes return just nodes without cursors. It uses the already fetched edges.
func (s *UserConnection) Nodes() []User {
	var nodes []User
	for _, edge := range s.Edges {
		nodes = append(nodes, edge.Node)
	}
	return nodes
}

// Nodes return just nodes without cursors, but as a slice of pointers. It uses the already fetched edges.
func (s *UserConnection) NodesPtr() []*User {
	var nodes []*User
	for _, edge := range s.Edges {
		item := edge
		nodes = append(nodes, &item.Node)
	}
	return nodes
}

func (client *Client) UsersConnection(params *UsersConnectionParams) *UserConnectionExec {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"UserWhereInput", "UserOrderByInput", "User"},
		"usersConnection",
		[]string{"edges", "pageInfo"})

	return &UserConnectionExec{ret}
}

func (client *Client) CreateOtp(params OtpCreateInput) *OtpExec {
	ret := client.Client.Create(
		params,
		[2]string{"OtpCreateInput!", "Otp"},
		"createOtp",
		[]string{"id", "phone", "code", "createdAt"})

	return &OtpExec{ret}
}

type OtpUpdateParams struct {
	Data  OtpUpdateInput      `json:"data"`
	Where OtpWhereUniqueInput `json:"where"`
}

func (client *Client) UpdateOtp(params OtpUpdateParams) *OtpExec {
	ret := client.Client.Update(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[3]string{"OtpUpdateInput!", "OtpWhereUniqueInput!", "Otp"},
		"updateOtp",
		[]string{"id", "phone", "code", "createdAt"})

	return &OtpExec{ret}
}

type OtpUpdateManyParams struct {
	Data  OtpUpdateManyMutationInput `json:"data"`
	Where *OtpWhereInput             `json:"where,omitempty"`
}

func (client *Client) UpdateManyOtps(params OtpUpdateManyParams) *BatchPayloadExec {
	exec := client.Client.UpdateMany(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[2]string{"OtpUpdateManyMutationInput!", "OtpWhereInput"},
		"updateManyOtps")
	return &BatchPayloadExec{exec}
}

type OtpUpsertParams struct {
	Where  OtpWhereUniqueInput `json:"where"`
	Create OtpCreateInput      `json:"create"`
	Update OtpUpdateInput      `json:"update"`
}

func (client *Client) UpsertOtp(params OtpUpsertParams) *OtpExec {
	uparams := &prisma.UpsertParams{
		Where:  params.Where,
		Create: params.Create,
		Update: params.Update,
	}
	ret := client.Client.Upsert(
		uparams,
		[4]string{"OtpWhereUniqueInput!", "OtpCreateInput!", "OtpUpdateInput!", "Otp"},
		"upsertOtp",
		[]string{"id", "phone", "code", "createdAt"})

	return &OtpExec{ret}
}

func (client *Client) DeleteOtp(params OtpWhereUniqueInput) *OtpExec {
	ret := client.Client.Delete(
		params,
		[2]string{"OtpWhereUniqueInput!", "Otp"},
		"deleteOtp",
		[]string{"id", "phone", "code", "createdAt"})

	return &OtpExec{ret}
}

func (client *Client) DeleteManyOtps(params *OtpWhereInput) *BatchPayloadExec {
	exec := client.Client.DeleteMany(params, "OtpWhereInput", "deleteManyOtps")
	return &BatchPayloadExec{exec}
}

func (client *Client) CreateUser(params UserCreateInput) *UserExec {
	ret := client.Client.Create(
		params,
		[2]string{"UserCreateInput!", "User"},
		"createUser",
		[]string{"id", "phone", "nickname", "avatar", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

type UserUpdateParams struct {
	Data  UserUpdateInput      `json:"data"`
	Where UserWhereUniqueInput `json:"where"`
}

func (client *Client) UpdateUser(params UserUpdateParams) *UserExec {
	ret := client.Client.Update(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[3]string{"UserUpdateInput!", "UserWhereUniqueInput!", "User"},
		"updateUser",
		[]string{"id", "phone", "nickname", "avatar", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

type UserUpdateManyParams struct {
	Data  UserUpdateManyMutationInput `json:"data"`
	Where *UserWhereInput             `json:"where,omitempty"`
}

func (client *Client) UpdateManyUsers(params UserUpdateManyParams) *BatchPayloadExec {
	exec := client.Client.UpdateMany(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[2]string{"UserUpdateManyMutationInput!", "UserWhereInput"},
		"updateManyUsers")
	return &BatchPayloadExec{exec}
}

type UserUpsertParams struct {
	Where  UserWhereUniqueInput `json:"where"`
	Create UserCreateInput      `json:"create"`
	Update UserUpdateInput      `json:"update"`
}

func (client *Client) UpsertUser(params UserUpsertParams) *UserExec {
	uparams := &prisma.UpsertParams{
		Where:  params.Where,
		Create: params.Create,
		Update: params.Update,
	}
	ret := client.Client.Upsert(
		uparams,
		[4]string{"UserWhereUniqueInput!", "UserCreateInput!", "UserUpdateInput!", "User"},
		"upsertUser",
		[]string{"id", "phone", "nickname", "avatar", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

func (client *Client) DeleteUser(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.Delete(
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"deleteUser",
		[]string{"id", "phone", "nickname", "avatar", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

func (client *Client) DeleteManyUsers(params *UserWhereInput) *BatchPayloadExec {
	exec := client.Client.DeleteMany(params, "UserWhereInput", "deleteManyUsers")
	return &BatchPayloadExec{exec}
}

type OtpOrderByInput string

const (
	OtpOrderByInputIDAsc         OtpOrderByInput = "id_ASC"
	OtpOrderByInputIDDesc        OtpOrderByInput = "id_DESC"
	OtpOrderByInputPhoneAsc      OtpOrderByInput = "phone_ASC"
	OtpOrderByInputPhoneDesc     OtpOrderByInput = "phone_DESC"
	OtpOrderByInputCodeAsc       OtpOrderByInput = "code_ASC"
	OtpOrderByInputCodeDesc      OtpOrderByInput = "code_DESC"
	OtpOrderByInputCreatedAtAsc  OtpOrderByInput = "createdAt_ASC"
	OtpOrderByInputCreatedAtDesc OtpOrderByInput = "createdAt_DESC"
)

type MutationType string

const (
	MutationTypeCreated MutationType = "CREATED"
	MutationTypeUpdated MutationType = "UPDATED"
	MutationTypeDeleted MutationType = "DELETED"
)

type UserOrderByInput string

const (
	UserOrderByInputIDAsc         UserOrderByInput = "id_ASC"
	UserOrderByInputIDDesc        UserOrderByInput = "id_DESC"
	UserOrderByInputPhoneAsc      UserOrderByInput = "phone_ASC"
	UserOrderByInputPhoneDesc     UserOrderByInput = "phone_DESC"
	UserOrderByInputNicknameAsc   UserOrderByInput = "nickname_ASC"
	UserOrderByInputNicknameDesc  UserOrderByInput = "nickname_DESC"
	UserOrderByInputAvatarAsc     UserOrderByInput = "avatar_ASC"
	UserOrderByInputAvatarDesc    UserOrderByInput = "avatar_DESC"
	UserOrderByInputCreatedAtAsc  UserOrderByInput = "createdAt_ASC"
	UserOrderByInputCreatedAtDesc UserOrderByInput = "createdAt_DESC"
	UserOrderByInputUpdatedAtAsc  UserOrderByInput = "updatedAt_ASC"
	UserOrderByInputUpdatedAtDesc UserOrderByInput = "updatedAt_DESC"
)

type UserCreateInput struct {
	ID       *string `json:"id,omitempty"`
	Phone    string  `json:"phone"`
	Nickname *string `json:"nickname,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}

type OtpUpdateInput struct {
	Phone *string `json:"phone,omitempty"`
	Code  *string `json:"code,omitempty"`
}

type OtpWhereUniqueInput struct {
	ID    *string `json:"id,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

type OtpWhereInput struct {
	ID                 *string         `json:"id,omitempty"`
	IDNot              *string         `json:"id_not,omitempty"`
	IDIn               []string        `json:"id_in,omitempty"`
	IDNotIn            []string        `json:"id_not_in,omitempty"`
	IDLt               *string         `json:"id_lt,omitempty"`
	IDLte              *string         `json:"id_lte,omitempty"`
	IDGt               *string         `json:"id_gt,omitempty"`
	IDGte              *string         `json:"id_gte,omitempty"`
	IDContains         *string         `json:"id_contains,omitempty"`
	IDNotContains      *string         `json:"id_not_contains,omitempty"`
	IDStartsWith       *string         `json:"id_starts_with,omitempty"`
	IDNotStartsWith    *string         `json:"id_not_starts_with,omitempty"`
	IDEndsWith         *string         `json:"id_ends_with,omitempty"`
	IDNotEndsWith      *string         `json:"id_not_ends_with,omitempty"`
	Phone              *string         `json:"phone,omitempty"`
	PhoneNot           *string         `json:"phone_not,omitempty"`
	PhoneIn            []string        `json:"phone_in,omitempty"`
	PhoneNotIn         []string        `json:"phone_not_in,omitempty"`
	PhoneLt            *string         `json:"phone_lt,omitempty"`
	PhoneLte           *string         `json:"phone_lte,omitempty"`
	PhoneGt            *string         `json:"phone_gt,omitempty"`
	PhoneGte           *string         `json:"phone_gte,omitempty"`
	PhoneContains      *string         `json:"phone_contains,omitempty"`
	PhoneNotContains   *string         `json:"phone_not_contains,omitempty"`
	PhoneStartsWith    *string         `json:"phone_starts_with,omitempty"`
	PhoneNotStartsWith *string         `json:"phone_not_starts_with,omitempty"`
	PhoneEndsWith      *string         `json:"phone_ends_with,omitempty"`
	PhoneNotEndsWith   *string         `json:"phone_not_ends_with,omitempty"`
	Code               *string         `json:"code,omitempty"`
	CodeNot            *string         `json:"code_not,omitempty"`
	CodeIn             []string        `json:"code_in,omitempty"`
	CodeNotIn          []string        `json:"code_not_in,omitempty"`
	CodeLt             *string         `json:"code_lt,omitempty"`
	CodeLte            *string         `json:"code_lte,omitempty"`
	CodeGt             *string         `json:"code_gt,omitempty"`
	CodeGte            *string         `json:"code_gte,omitempty"`
	CodeContains       *string         `json:"code_contains,omitempty"`
	CodeNotContains    *string         `json:"code_not_contains,omitempty"`
	CodeStartsWith     *string         `json:"code_starts_with,omitempty"`
	CodeNotStartsWith  *string         `json:"code_not_starts_with,omitempty"`
	CodeEndsWith       *string         `json:"code_ends_with,omitempty"`
	CodeNotEndsWith    *string         `json:"code_not_ends_with,omitempty"`
	CreatedAt          *string         `json:"createdAt,omitempty"`
	CreatedAtNot       *string         `json:"createdAt_not,omitempty"`
	CreatedAtIn        []string        `json:"createdAt_in,omitempty"`
	CreatedAtNotIn     []string        `json:"createdAt_not_in,omitempty"`
	CreatedAtLt        *string         `json:"createdAt_lt,omitempty"`
	CreatedAtLte       *string         `json:"createdAt_lte,omitempty"`
	CreatedAtGt        *string         `json:"createdAt_gt,omitempty"`
	CreatedAtGte       *string         `json:"createdAt_gte,omitempty"`
	And                []OtpWhereInput `json:"AND,omitempty"`
	Or                 []OtpWhereInput `json:"OR,omitempty"`
	Not                []OtpWhereInput `json:"NOT,omitempty"`
}

type UserWhereInput struct {
	ID                    *string          `json:"id,omitempty"`
	IDNot                 *string          `json:"id_not,omitempty"`
	IDIn                  []string         `json:"id_in,omitempty"`
	IDNotIn               []string         `json:"id_not_in,omitempty"`
	IDLt                  *string          `json:"id_lt,omitempty"`
	IDLte                 *string          `json:"id_lte,omitempty"`
	IDGt                  *string          `json:"id_gt,omitempty"`
	IDGte                 *string          `json:"id_gte,omitempty"`
	IDContains            *string          `json:"id_contains,omitempty"`
	IDNotContains         *string          `json:"id_not_contains,omitempty"`
	IDStartsWith          *string          `json:"id_starts_with,omitempty"`
	IDNotStartsWith       *string          `json:"id_not_starts_with,omitempty"`
	IDEndsWith            *string          `json:"id_ends_with,omitempty"`
	IDNotEndsWith         *string          `json:"id_not_ends_with,omitempty"`
	Phone                 *string          `json:"phone,omitempty"`
	PhoneNot              *string          `json:"phone_not,omitempty"`
	PhoneIn               []string         `json:"phone_in,omitempty"`
	PhoneNotIn            []string         `json:"phone_not_in,omitempty"`
	PhoneLt               *string          `json:"phone_lt,omitempty"`
	PhoneLte              *string          `json:"phone_lte,omitempty"`
	PhoneGt               *string          `json:"phone_gt,omitempty"`
	PhoneGte              *string          `json:"phone_gte,omitempty"`
	PhoneContains         *string          `json:"phone_contains,omitempty"`
	PhoneNotContains      *string          `json:"phone_not_contains,omitempty"`
	PhoneStartsWith       *string          `json:"phone_starts_with,omitempty"`
	PhoneNotStartsWith    *string          `json:"phone_not_starts_with,omitempty"`
	PhoneEndsWith         *string          `json:"phone_ends_with,omitempty"`
	PhoneNotEndsWith      *string          `json:"phone_not_ends_with,omitempty"`
	Nickname              *string          `json:"nickname,omitempty"`
	NicknameNot           *string          `json:"nickname_not,omitempty"`
	NicknameIn            []string         `json:"nickname_in,omitempty"`
	NicknameNotIn         []string         `json:"nickname_not_in,omitempty"`
	NicknameLt            *string          `json:"nickname_lt,omitempty"`
	NicknameLte           *string          `json:"nickname_lte,omitempty"`
	NicknameGt            *string          `json:"nickname_gt,omitempty"`
	NicknameGte           *string          `json:"nickname_gte,omitempty"`
	NicknameContains      *string          `json:"nickname_contains,omitempty"`
	NicknameNotContains   *string          `json:"nickname_not_contains,omitempty"`
	NicknameStartsWith    *string          `json:"nickname_starts_with,omitempty"`
	NicknameNotStartsWith *string          `json:"nickname_not_starts_with,omitempty"`
	NicknameEndsWith      *string          `json:"nickname_ends_with,omitempty"`
	NicknameNotEndsWith   *string          `json:"nickname_not_ends_with,omitempty"`
	Avatar                *string          `json:"avatar,omitempty"`
	AvatarNot             *string          `json:"avatar_not,omitempty"`
	AvatarIn              []string         `json:"avatar_in,omitempty"`
	AvatarNotIn           []string         `json:"avatar_not_in,omitempty"`
	AvatarLt              *string          `json:"avatar_lt,omitempty"`
	AvatarLte             *string          `json:"avatar_lte,omitempty"`
	AvatarGt              *string          `json:"avatar_gt,omitempty"`
	AvatarGte             *string          `json:"avatar_gte,omitempty"`
	AvatarContains        *string          `json:"avatar_contains,omitempty"`
	AvatarNotContains     *string          `json:"avatar_not_contains,omitempty"`
	AvatarStartsWith      *string          `json:"avatar_starts_with,omitempty"`
	AvatarNotStartsWith   *string          `json:"avatar_not_starts_with,omitempty"`
	AvatarEndsWith        *string          `json:"avatar_ends_with,omitempty"`
	AvatarNotEndsWith     *string          `json:"avatar_not_ends_with,omitempty"`
	CreatedAt             *string          `json:"createdAt,omitempty"`
	CreatedAtNot          *string          `json:"createdAt_not,omitempty"`
	CreatedAtIn           []string         `json:"createdAt_in,omitempty"`
	CreatedAtNotIn        []string         `json:"createdAt_not_in,omitempty"`
	CreatedAtLt           *string          `json:"createdAt_lt,omitempty"`
	CreatedAtLte          *string          `json:"createdAt_lte,omitempty"`
	CreatedAtGt           *string          `json:"createdAt_gt,omitempty"`
	CreatedAtGte          *string          `json:"createdAt_gte,omitempty"`
	UpdatedAt             *string          `json:"updatedAt,omitempty"`
	UpdatedAtNot          *string          `json:"updatedAt_not,omitempty"`
	UpdatedAtIn           []string         `json:"updatedAt_in,omitempty"`
	UpdatedAtNotIn        []string         `json:"updatedAt_not_in,omitempty"`
	UpdatedAtLt           *string          `json:"updatedAt_lt,omitempty"`
	UpdatedAtLte          *string          `json:"updatedAt_lte,omitempty"`
	UpdatedAtGt           *string          `json:"updatedAt_gt,omitempty"`
	UpdatedAtGte          *string          `json:"updatedAt_gte,omitempty"`
	And                   []UserWhereInput `json:"AND,omitempty"`
	Or                    []UserWhereInput `json:"OR,omitempty"`
	Not                   []UserWhereInput `json:"NOT,omitempty"`
}

type UserWhereUniqueInput struct {
	ID    *string `json:"id,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

type OtpCreateInput struct {
	ID    *string `json:"id,omitempty"`
	Phone string  `json:"phone"`
	Code  string  `json:"code"`
}

type OtpSubscriptionWhereInput struct {
	MutationIn                 []MutationType              `json:"mutation_in,omitempty"`
	UpdatedFieldsContains      *string                     `json:"updatedFields_contains,omitempty"`
	UpdatedFieldsContainsEvery []string                    `json:"updatedFields_contains_every,omitempty"`
	UpdatedFieldsContainsSome  []string                    `json:"updatedFields_contains_some,omitempty"`
	Node                       *OtpWhereInput              `json:"node,omitempty"`
	And                        []OtpSubscriptionWhereInput `json:"AND,omitempty"`
	Or                         []OtpSubscriptionWhereInput `json:"OR,omitempty"`
	Not                        []OtpSubscriptionWhereInput `json:"NOT,omitempty"`
}

type OtpUpdateManyMutationInput struct {
	Phone *string `json:"phone,omitempty"`
	Code  *string `json:"code,omitempty"`
}

type UserSubscriptionWhereInput struct {
	MutationIn                 []MutationType               `json:"mutation_in,omitempty"`
	UpdatedFieldsContains      *string                      `json:"updatedFields_contains,omitempty"`
	UpdatedFieldsContainsEvery []string                     `json:"updatedFields_contains_every,omitempty"`
	UpdatedFieldsContainsSome  []string                     `json:"updatedFields_contains_some,omitempty"`
	Node                       *UserWhereInput              `json:"node,omitempty"`
	And                        []UserSubscriptionWhereInput `json:"AND,omitempty"`
	Or                         []UserSubscriptionWhereInput `json:"OR,omitempty"`
	Not                        []UserSubscriptionWhereInput `json:"NOT,omitempty"`
}

type UserUpdateInput struct {
	Phone    *string `json:"phone,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}

type UserUpdateManyMutationInput struct {
	Phone    *string `json:"phone,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}

type UserEdgeExec struct {
	exec *prisma.Exec
}

func (instance *UserEdgeExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "phone", "nickname", "avatar", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

func (instance UserEdgeExec) Exec(ctx context.Context) (*UserEdge, error) {
	var v UserEdge
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserEdgeExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserEdgeExecArray struct {
	exec *prisma.Exec
}

func (instance UserEdgeExecArray) Exec(ctx context.Context) ([]UserEdge, error) {
	var v []UserEdge
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var UserEdgeFields = []string{"cursor"}

type UserEdge struct {
	Node   User   `json:"node"`
	Cursor string `json:"cursor"`
}

type PageInfoExec struct {
	exec *prisma.Exec
}

func (instance PageInfoExec) Exec(ctx context.Context) (*PageInfo, error) {
	var v PageInfo
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PageInfoExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PageInfoExecArray struct {
	exec *prisma.Exec
}

func (instance PageInfoExecArray) Exec(ctx context.Context) ([]PageInfo, error) {
	var v []PageInfo
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var PageInfoFields = []string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}

type UserPreviousValuesExec struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExec) Exec(ctx context.Context) (*UserPreviousValues, error) {
	var v UserPreviousValues
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserPreviousValuesExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserPreviousValuesExecArray struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExecArray) Exec(ctx context.Context) ([]UserPreviousValues, error) {
	var v []UserPreviousValues
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var UserPreviousValuesFields = []string{"id", "phone", "nickname", "avatar", "createdAt", "updatedAt"}

type UserPreviousValues struct {
	ID        string  `json:"id"`
	Phone     string  `json:"phone"`
	Nickname  *string `json:"nickname,omitempty"`
	Avatar    *string `json:"avatar,omitempty"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type OtpConnectionExec struct {
	exec *prisma.Exec
}

func (instance *OtpConnectionExec) PageInfo() *PageInfoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PageInfo"},
		"pageInfo",
		[]string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"})

	return &PageInfoExec{ret}
}

func (instance *OtpConnectionExec) Edges() *OtpEdgeExecArray {
	edges := instance.exec.Client.GetMany(
		instance.exec,
		nil,
		[3]string{"OtpWhereInput", "OtpOrderByInput", "OtpEdge"},
		"edges",
		[]string{"cursor"})

	nodes := edges.Client.GetOne(
		edges,
		nil,
		[2]string{"", "Otp"},
		"node",
		OtpFields)

	return &OtpEdgeExecArray{nodes}
}

func (instance *OtpConnectionExec) Aggregate(ctx context.Context) (*Aggregate, error) {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "AggregateOtp"},
		"aggregate",
		[]string{"count"})

	var v Aggregate
	_, err := ret.Exec(ctx, &v)
	return &v, err
}

func (instance OtpConnectionExec) Exec(ctx context.Context) (*OtpConnection, error) {
	edges, err := instance.Edges().Exec(ctx)
	if err != nil {
		return nil, err
	}

	pageInfo, err := instance.PageInfo().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &OtpConnection{
		Edges:    edges,
		PageInfo: *pageInfo,
	}, nil
}

func (instance OtpConnectionExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type OtpConnectionExecArray struct {
	exec *prisma.Exec
}

func (instance OtpConnectionExecArray) Exec(ctx context.Context) ([]OtpConnection, error) {
	var v []OtpConnection
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var OtpConnectionFields = []string{}

type OtpConnection struct {
	PageInfo PageInfo  `json:"pageInfo"`
	Edges    []OtpEdge `json:"edges"`
}

type UserConnectionExec struct {
	exec *prisma.Exec
}

func (instance *UserConnectionExec) PageInfo() *PageInfoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PageInfo"},
		"pageInfo",
		[]string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"})

	return &PageInfoExec{ret}
}

func (instance *UserConnectionExec) Edges() *UserEdgeExecArray {
	edges := instance.exec.Client.GetMany(
		instance.exec,
		nil,
		[3]string{"UserWhereInput", "UserOrderByInput", "UserEdge"},
		"edges",
		[]string{"cursor"})

	nodes := edges.Client.GetOne(
		edges,
		nil,
		[2]string{"", "User"},
		"node",
		UserFields)

	return &UserEdgeExecArray{nodes}
}

func (instance *UserConnectionExec) Aggregate(ctx context.Context) (*Aggregate, error) {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "AggregateUser"},
		"aggregate",
		[]string{"count"})

	var v Aggregate
	_, err := ret.Exec(ctx, &v)
	return &v, err
}

func (instance UserConnectionExec) Exec(ctx context.Context) (*UserConnection, error) {
	edges, err := instance.Edges().Exec(ctx)
	if err != nil {
		return nil, err
	}

	pageInfo, err := instance.PageInfo().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &UserConnection{
		Edges:    edges,
		PageInfo: *pageInfo,
	}, nil
}

func (instance UserConnectionExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserConnectionExecArray struct {
	exec *prisma.Exec
}

func (instance UserConnectionExecArray) Exec(ctx context.Context) ([]UserConnection, error) {
	var v []UserConnection
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var UserConnectionFields = []string{}

type UserConnection struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Edges    []UserEdge `json:"edges"`
}

type OtpSubscriptionPayloadExec struct {
	exec *prisma.Exec
}

func (instance *OtpSubscriptionPayloadExec) Node() *OtpExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "Otp"},
		"node",
		[]string{"id", "phone", "code", "createdAt"})

	return &OtpExec{ret}
}

func (instance *OtpSubscriptionPayloadExec) PreviousValues() *OtpPreviousValuesExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "OtpPreviousValues"},
		"previousValues",
		[]string{"id", "phone", "code", "createdAt"})

	return &OtpPreviousValuesExec{ret}
}

func (instance OtpSubscriptionPayloadExec) Exec(ctx context.Context) (*OtpSubscriptionPayload, error) {
	var v OtpSubscriptionPayload
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance OtpSubscriptionPayloadExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type OtpSubscriptionPayloadExecArray struct {
	exec *prisma.Exec
}

func (instance OtpSubscriptionPayloadExecArray) Exec(ctx context.Context) ([]OtpSubscriptionPayload, error) {
	var v []OtpSubscriptionPayload
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var OtpSubscriptionPayloadFields = []string{"mutation", "updatedFields"}

type OtpSubscriptionPayload struct {
	Mutation      MutationType `json:"mutation"`
	Node          *Otp         `json:"node,omitempty"`
	UpdatedFields []string     `json:"updatedFields,omitempty"`
}

type OtpExec struct {
	exec *prisma.Exec
}

func (instance OtpExec) Exec(ctx context.Context) (*Otp, error) {
	var v Otp
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance OtpExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type OtpExecArray struct {
	exec *prisma.Exec
}

func (instance OtpExecArray) Exec(ctx context.Context) ([]Otp, error) {
	var v []Otp
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var OtpFields = []string{"id", "phone", "code", "createdAt"}

type Otp struct {
	ID        string `json:"id"`
	Phone     string `json:"phone"`
	Code      string `json:"code"`
	CreatedAt string `json:"createdAt"`
}

type OtpEdgeExec struct {
	exec *prisma.Exec
}

func (instance *OtpEdgeExec) Node() *OtpExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "Otp"},
		"node",
		[]string{"id", "phone", "code", "createdAt"})

	return &OtpExec{ret}
}

func (instance OtpEdgeExec) Exec(ctx context.Context) (*OtpEdge, error) {
	var v OtpEdge
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance OtpEdgeExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type OtpEdgeExecArray struct {
	exec *prisma.Exec
}

func (instance OtpEdgeExecArray) Exec(ctx context.Context) ([]OtpEdge, error) {
	var v []OtpEdge
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var OtpEdgeFields = []string{"cursor"}

type OtpEdge struct {
	Node   Otp    `json:"node"`
	Cursor string `json:"cursor"`
}

type UserExec struct {
	exec *prisma.Exec
}

func (instance UserExec) Exec(ctx context.Context) (*User, error) {
	var v User
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserExecArray struct {
	exec *prisma.Exec
}

func (instance UserExecArray) Exec(ctx context.Context) ([]User, error) {
	var v []User
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var UserFields = []string{"id", "phone", "nickname", "avatar", "createdAt", "updatedAt"}

type User struct {
	ID        string  `json:"id"`
	Phone     string  `json:"phone"`
	Nickname  *string `json:"nickname,omitempty"`
	Avatar    *string `json:"avatar,omitempty"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type OtpPreviousValuesExec struct {
	exec *prisma.Exec
}

func (instance OtpPreviousValuesExec) Exec(ctx context.Context) (*OtpPreviousValues, error) {
	var v OtpPreviousValues
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance OtpPreviousValuesExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type OtpPreviousValuesExecArray struct {
	exec *prisma.Exec
}

func (instance OtpPreviousValuesExecArray) Exec(ctx context.Context) ([]OtpPreviousValues, error) {
	var v []OtpPreviousValues
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var OtpPreviousValuesFields = []string{"id", "phone", "code", "createdAt"}

type OtpPreviousValues struct {
	ID        string `json:"id"`
	Phone     string `json:"phone"`
	Code      string `json:"code"`
	CreatedAt string `json:"createdAt"`
}

type UserSubscriptionPayloadExec struct {
	exec *prisma.Exec
}

func (instance *UserSubscriptionPayloadExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "phone", "nickname", "avatar", "createdAt", "updatedAt"})

	return &UserExec{ret}
}

func (instance *UserSubscriptionPayloadExec) PreviousValues() *UserPreviousValuesExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "UserPreviousValues"},
		"previousValues",
		[]string{"id", "phone", "nickname", "avatar", "createdAt", "updatedAt"})

	return &UserPreviousValuesExec{ret}
}

func (instance UserSubscriptionPayloadExec) Exec(ctx context.Context) (*UserSubscriptionPayload, error) {
	var v UserSubscriptionPayload
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserSubscriptionPayloadExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserSubscriptionPayloadExecArray struct {
	exec *prisma.Exec
}

func (instance UserSubscriptionPayloadExecArray) Exec(ctx context.Context) ([]UserSubscriptionPayload, error) {
	var v []UserSubscriptionPayload
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var UserSubscriptionPayloadFields = []string{"mutation", "updatedFields"}

type UserSubscriptionPayload struct {
	Mutation      MutationType `json:"mutation"`
	Node          *User        `json:"node,omitempty"`
	UpdatedFields []string     `json:"updatedFields,omitempty"`
}
