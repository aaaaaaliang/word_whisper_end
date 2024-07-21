package fronted

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"strconv"
	"word_whisper_end/internal/service"

	"word_whisper_end/api/fronted/v1"
)

func (c *HomeControllerV1) GetHomeInfo(ctx context.Context, req *v1.GetHomeInfoReq) (res *v1.GetHomeInfoRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := r.Header.Get("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("转换错误:", err)
		return
	}
	lw, rw, lastId, err := service.Home().GetHomeInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	return &v1.GetHomeInfoRes{
		WordsToLearn:  strconv.Itoa(lw),
		WordsToReview: strconv.Itoa(rw),
		LastId:        strconv.Itoa(lastId),
	}, nil
}
