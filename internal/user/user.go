package user

import (
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrRecordNotFound    = errors.New("record not found")
	ErrInternal          = errors.New("internal error")
)

type UserInput struct {
	Phone    string
	Code     string
	Nickname *string
}

type User struct {
	Id        string
	Phone     string
	Avatar    string
	Nickname  *string
	CreatedAt string
	UpdatedAt string
}

func (u *User) ChangeAvatar(avatar string) {
	u.Avatar = avatar
}

func (u *User) ChangeNickname(nickname *string) {
	u.Nickname = nickname
}

func NewAccount(input UserInput) (*User, error) {
	// 查询input.code 是否一致
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, fmt.Errorf("password Encryption failed")
	// }
	return &User{
		Phone: input.Phone,
	}, nil
}
