// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package fronted

import (
	"context"

	"word_whisper_end/api/fronted/v1"
)

type IUserV1 interface {
	SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error)
	SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error)
}

type IHomeV1 interface {
	GetHomeInfo(ctx context.Context, req *v1.GetHomeInfoReq) (res *v1.GetHomeInfoRes, err error)
	GetWordsByLastIdReq(ctx context.Context, req *v1.GetWordsByLastIdReq) (res *v1.GetWordsByLastIdRes, err error)
	SearchReq(ctx context.Context, req *v1.SearchReq) (res *v1.SearchRes, err error)
}
