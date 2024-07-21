package fronted

import (
	"context"
	"word_whisper_end/internal/service"

	"word_whisper_end/api/fronted/v1"
)

func (c *UserControllerV1) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {

	userId, err := service.User().CreateUser(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.SignUpRes{UserID: int(userId)}, nil
}
