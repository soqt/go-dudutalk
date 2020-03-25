package prisma


import (
	"context"
	"soqt/go-dudutalk/internal/otp"
	"soqt/go-dudutalk/internal/user"

	"github.com/pkg/errors"
)

func (s *Storage) SaveOTP(ctx context.Context, phone *string, otp *string) error {
	_, err := s.Prisma.UpsertOtp(OtpUpsertParams{
		Where:  OtpWhereUniqueInput{
			Phone: phone,
		},
		Create: OtpCreateInput{
			Phone: *phone,
			Code: *otp,
		},
		Update: OtpUpdateInput{
			Code: otp,
		},
	}).Exec(ctx)


	if err != nil {
		if err == ErrNoResult {
			return errors.Wrap(user.ErrRecordNotFound, "prisma.otp.SaveOTP")
		}
		return errors.Wrap(err, "prisma.otp.SaveOTP")
	}
	return nil
}

func (s *Storage) GetOTPByPhone(ctx context.Context, phone *string) (*string, error)  {
	code, err := s.Prisma.Otp(OtpWhereUniqueInput{
		Phone: phone,
	}).Exec(ctx)
	if err != nil {
		if err == ErrNoResult {
			return nil, errors.Wrap(otp.ErrRecordNotFound, "prisma.otp.GetOTPByPhone")
		}
		return nil, errors.Wrap(err, "prisma.otp.GetOTPByPhone")
	}
	return &code.Code, nil
}

func (s *Storage) DeleteOTPByPhone(ctx context.Context, phone *string) error  {
	_, err := s.Prisma.DeleteOtp(OtpWhereUniqueInput{
		Phone: phone,
	}).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "prisma.otp.DeleteOTPByPhone")
	}
	return  nil
}