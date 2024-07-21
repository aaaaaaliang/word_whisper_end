// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"word_whisper_end/internal/model/entity"
)

type (
	IHome interface {
		GetHomeInfo(ctx context.Context, userId int) (learnWords int, reviewWords int, lastId int, err error)
		GetWordsInfo(ctx context.Context, lastId int, userId int) (words []entity.Words, err error)
		XingHuoSDK(ctx context.Context, query string) (answer string, err error)
	}
)

var (
	localHome IHome
)

func Home() IHome {
	if localHome == nil {
		panic("implement not found for interface IHome, forgot register?")
	}
	return localHome
}

func RegisterHome(i IHome) {
	localHome = i
}
