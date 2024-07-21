package home

import (
	"context"
	"word_whisper_end/internal/dao"
	"word_whisper_end/internal/model/entity"
	"word_whisper_end/internal/service"
	"word_whisper_end/utility"
)

type sHome struct {
}

func init() {
	service.RegisterHome(New())
}

func New() service.IHome {
	return &sHome{}
}

func (s *sHome) GetHomeInfo(ctx context.Context, userId int) (learnWords int, reviewWords int, lastId int, err error) {
	var userProgress entity.UserProgress
	err = dao.UserProgress.Ctx(ctx).Where("user_id = ?", userId).Scan(&userProgress)
	if err != nil {
		return
	}
	return userProgress.WordsToLearn, userProgress.WordsToReview, userProgress.LastWordId, nil
}

func (s *sHome) GetWordsInfo(ctx context.Context, lastId int, userId int) (words []entity.Words, err error) {
	var userProgress entity.UserProgress
	err = dao.UserProgress.Ctx(ctx).Where("user_id = ?", userId).Scan(&userProgress)
	if err != nil {
		return nil, err
	}
	err = dao.Words.Ctx(ctx).
		Where("id > ?", lastId).
		Order("id asc").
		Limit(userProgress.WordsToLearn).
		Scan(&words)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (s *sHome) XingHuoSDK(ctx context.Context, query string) (answer string, err error) {
	answer, err = utility.QueryAI(query)
	if err != nil {
		return "", err
	}
	return answer, nil
}
