package fronted

import (
	"context"
	"word_whisper_end/internal/service"

	"word_whisper_end/api/fronted/v1"
)

func (c *UserControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	token, userId, err := service.User().QueryUser(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.SignInRes{
		Token:  token,
		UserID: userId,
	}, nil
}
