package fronted

import (
	"context"
	v1 "word_whisper_end/api/fronted/v1"
	"word_whisper_end/internal/service"
)

func (c *HomeControllerV1) SearchReq(ctx context.Context, req *v1.SearchReq) (res *v1.SearchRes, err error) {
	answer, err := service.Home().XingHuoSDK(ctx, req.Query)
	if err != nil {
		return nil, err
	}
	return &v1.SearchRes{Answer: answer}, nil
}
