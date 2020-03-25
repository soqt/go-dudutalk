package prisma

import (
	"context"
	"soqt/go-dudutalk/internal/user"

	"github.com/pkg/errors"
)

func (s *Storage) GetUserById(ctx context.Context, id *string) (*user.User, error) {
	u, err := s.Prisma.User(UserWhereUniqueInput{
		ID: id,
	}).Exec(ctx)

	if err != nil && err == ErrNoResult {
		if err == ErrNoResult {
			return nil, errors.Wrap(user.ErrRecordNotFound, "prisma.user.GetUserById")
		}
		return nil, err
	}

	ru := s.userReducer(u)

	return ru, nil
}

func (s *Storage) GetUserByPhone(ctx context.Context, phone *string) (*user.User, error) {
	u, err := s.Prisma.User(UserWhereUniqueInput{
		Phone: phone,
	}).Exec(ctx)

	if err != nil {
		if err == ErrNoResult {
			return nil, errors.Wrap(user.ErrRecordNotFound, "prisma.user.GetUserByPhone")
		}
		return nil, errors.Wrap(err, "prisma.user.GetUserByPhone")
	}

	ru := s.userReducer(u)

	return ru, nil
}

func (s *Storage) CreateUser(ctx context.Context, input *user.UserInput) (*user.User, error) {
	u, err := s.Prisma.CreateUser(UserCreateInput{
		Phone: input.Phone,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	user := s.userReducer(u)
	return user, nil
}

func (s *Storage) userReducer(u *User) *user.User {
	ru := user.User{
		Id:        u.ID,
		Phone:     u.Phone,
		Nickname:  u.Nickname,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return &ru
}
