package user

import (
	"context"
	"fmt"
	"word_whisper_end/internal/dao"
	"word_whisper_end/internal/model/do"
	"word_whisper_end/internal/model/entity"
	"word_whisper_end/internal/service"
	"word_whisper_end/utility"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) CreateUser(ctx context.Context, username string, password string) (userId int64, err error) {
	count, err := dao.Users.Ctx(ctx).Where("username = ?", username).Count()
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, fmt.Errorf("The username already exists")
	}
	hashPassword := utility.MD5Hash(password)
	userId, err = dao.Users.Ctx(ctx).Data(do.Users{
		Username: username,
		Password: hashPassword,
	}).InsertAndGetId()
	if err != nil {
		return 0, err
	}
	_, err = dao.UserProgress.Ctx(ctx).Data(do.UserProgress{
		UserId:        userId,
		LastWordId:    1,
		WordsToLearn:  15,
		WordsToReview: 0,
	}).InsertAndGetId()
	if err != nil {
		return
	}
	return userId, nil
}

func (s *sUser) QueryUser(ctx context.Context, username string, password string) (token string, userId int, err error) {
	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("username = ?", username).Scan(&user)
	if err != nil {
		return "", 0, err
	}
	hashPassword := utility.MD5Hash(password)
	if user.Password != hashPassword {
		return "", 0, fmt.Errorf("password wrong")
	}
	token, err = utility.GenerateJWT(userId)
	if err != nil {
		return "", 0, err
	}
	return token, user.Id, nil
}
