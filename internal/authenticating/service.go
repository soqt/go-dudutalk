package authenticating

import (
	"context"
	"fmt"
	"os"
	"soqt/go-dudutalk/internal/otp"
	"strconv"

	"github.com/nyaruka/phonenumbers"
	errs "github.com/pkg/errors"

	"soqt/go-dudutalk/internal/user"
	aplerror "soqt/go-dudutalk/pkg/apolloerror"
	"soqt/go-dudutalk/pkg/auth"
	"soqt/go-dudutalk/pkg/jwt"
)

//go:generate mockgen -destination=./mock_user.go -package=authenticating -self_package=soqt/go-dudutalk/internal/service/authenticating soqt/go-dudutalk/internal/service/authenticating Repository

type Repository interface {
	GetUserById(ctx context.Context, id *string) (*user.User, error)
	GetUserByPhone(ctx context.Context, phone *string) (*user.User, error)
	CreateUser(ctx context.Context, input *user.UserInput) (*user.User, error)
	GetOTPByPhone(ctx context.Context, phone *string) (*string, error)
	SaveOTP(ctx context.Context, phone *string, otp *string) error
	DeleteOTPByPhone(ctx context.Context, phone *string) error
}

type Service interface {
	GetUserById(ctx context.Context, id *string) (*user.User, error)
	IssueOTP(ctx context.Context, phone *string) (err error)
	Login(ctx context.Context, phone string, code string) (user *user.User, token *string, err error)
}

type service struct {
	userRepo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetUserById(ctx context.Context, id *string) (*user.User, error)  {
	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type ErrorResponse struct {
	Err string
}

func genJwt(id string) string {
	secret := os.Getenv("JWT_SECRET")
	token, _ := jwt.JWTGen(secret, id)
	return token
}

// validPhoneNumber 测试是否是正确的中国电话号码
func validPhoneNumber(phone string) (string, bool) {
	if phone == "" {
		return "", false
	}
	num, err := phonenumbers.Parse(phone, "CN")
	if !phonenumbers.IsValidNumberForRegion(num, "CN") || err != nil {
		return "", false
	}
	pn := strconv.Itoa(int(num.GetNationalNumber()))

	return pn, true
}

func (s *service) IssueOTP(ctx context.Context, phone *string) (error) {
	pn, ok := validPhoneNumber(*phone)

	if !ok {
		fields := make(map[string]interface{})
		fields["PhoneNumber"] = pn
		return aplerror.UserInputError("请输入正确的电话号码", &fields)
	}

	otp := auth.GenOTP(4)
	//alisms.SendAuthOTP(pn, otp)
	s.userRepo.SaveOTP(ctx, phone, &otp)
	fmt.Println(otp)
	return nil
}

func (s *service) Login(ctx context.Context, phone string, code string) (u *user.User, token *string, err error) {
	c, err := s.userRepo.GetOTPByPhone(ctx, &phone)
	if err != nil {
		if err == otp.ErrRecordNotFound {
			return nil, nil, aplerror.UserInputError("短信验证码不存在", nil)
		}
		return nil, nil, errs.Wrap(err, "authenticating.service.Login GetOTP")
	}

	if *c != code {
		fields := make(map[string]interface{})
		fields["otp"] = code
		return nil, nil, aplerror.UserInputError("短信验证码错误", &fields)
	}


	u, err = s.userRepo.GetUserByPhone(ctx, &phone)

	if err != nil && errs.Cause(err) == user.ErrRecordNotFound {
		u, err = s.userRepo.CreateUser(ctx, &user.UserInput{
			Phone:     phone,
		})
	}

	if err != nil {
		return nil, nil, errs.Wrap(err, "authenticating.service.Login")
	}

	jwt := genJwt(u.Id)
	s.userRepo.DeleteOTPByPhone(ctx, &phone)
	return u, &jwt, nil
}
