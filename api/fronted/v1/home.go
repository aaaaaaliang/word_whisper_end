package v1

import "github.com/gogf/gf/v2/frame/g"

type GetHomeInfoReq struct {
	g.Meta `path:"/api/front/home" method:"get" tags:"主页模块"  summary:"初始化页面"`
}

type GetHomeInfoRes struct {
	WordsToLearn  string `json:"wordsToLearn"`
	WordsToReview string `json:"wordsToReview"`
	LastId        string `json:"lastId"`
}

type GetWordsByLastIdReq struct {
	g.Meta `path:"/api/front/home/getWords" method:"post" tags:"主页模块" summary:"获取单词"`
	LastId string `json:"lastId"`
	UserId string `json:"userId"`
}
type GetWordsByLastIdRes struct {
	Data []Word `json:"data"`
}

type Word struct {
	WordId  string `json:"wordId"`
	English string `json:"english"`
	Chinese string `json:"chinese"`
}

type SearchReq struct {
	g.Meta `path:"/api/front/home/search" method:"post" tags:"主页模块" summary:"搜索"`
	Query  string `json:"query"`
}

type SearchRes struct {
	Answer string `json:"answer"`
}
