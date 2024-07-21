package fronted

import (
	"context"
	"strconv"
	v1 "word_whisper_end/api/fronted/v1"
	"word_whisper_end/internal/service"
)

func (c *HomeControllerV1) GetWordsByLastIdReq(ctx context.Context, req *v1.GetWordsByLastIdReq) (res *v1.GetWordsByLastIdRes, err error) {
	lastId, err := strconv.Atoi(req.LastId)
	if err != nil {
		return nil, err
	}
	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, err
	}
	wordsData, err := service.Home().GetWordsInfo(ctx, lastId, userId)
	if err != nil {
		return nil, err
	}

	var wordList []v1.Word
	for _, word := range wordsData {
		wordList = append(wordList, v1.Word{
			WordId:  strconv.Itoa(word.Id),
			English: word.Word,
			Chinese: word.Translation,
		})
	}

	res = &v1.GetWordsByLastIdRes{
		Data: wordList,
	}
	return res, nil
}
