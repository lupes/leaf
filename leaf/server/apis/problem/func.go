package problem

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type GraphQLRequest struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		TitleSlug string `json:"titleSlug"`
	} `json:"variables"`
	Query string `json:"query"`
}

type GraphQLResponse struct {
	Data struct {
		Question struct {
			QuestionFrontendId string `json:"question_frontend_id"`
			TitleSlug          string `json:"titleSlug"`
			TranslatedTitle    string `json:"translatedTitle"`
			TranslatedContent  string `json:"translatedContent"`
			Difficulty         string `json:"difficulty"`
			TopicTags          []struct {
				TranslatedName string `json:"translatedName"`
			}
		}
	} `json:"data"`
}

func GetTitleSlugFromUrl(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", fmt.Errorf("GraphQL data 格式化异常:%w", err)
	}
	path := strings.Trim(u.Path, "/")
	arr := strings.Split(path, "/")
	if len(arr) != 2 {
		return "", fmt.Errorf("leetcode url 格式异常:%s", str)
	}
	return arr[1], nil
}

func PostGraphQL(ctx context.Context, titleSlug string) (*GraphQLResponse, error) {
	data := GraphQLRequest{
		OperationName: "questionData",
		Variables: struct {
			TitleSlug string `json:"titleSlug"`
		}{
			TitleSlug: titleSlug,
		},
		Query: "query questionData($titleSlug: String!) {  question(titleSlug: $titleSlug) {    questionFrontendId\n    titleSlug\n    translatedTitle\n    translatedContent\n    difficulty\n    topicTags {translatedName}}}",
	}
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("GraphQL data 格式化异常:%w", err)
	}
	response, err := http.Post("https://leetcode-cn.com/graphql/", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("获取问题数据异常:%w", err)
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("获取问题数据请求不是成功状态:%d", response.StatusCode)
	}

	res := &GraphQLResponse{}
	err = json.NewDecoder(response.Body).Decode(res)
	if err != nil {
		return nil, fmt.Errorf("解析数据异常:%w", err)
	}
	return res, nil
}
