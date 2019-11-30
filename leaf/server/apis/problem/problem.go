package problem

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"

	"leaf/server/common"
	"leaf/server/database/problem"
)

// 问题
func getInt64Param(req *http.Request, key string) (int64, error) {
	value := req.FormValue(key)
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("query param %s[%s] not is be int err:%w", key, value, err)
	}
	return int64(number), nil
}

func List(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	limit, err := getInt64Param(req, "limit")
	if err != nil {
		return 400, nil, fmt.Errorf("get query param err:%w", err)
	}

	offset, err := getInt64Param(req, "offset")
	if err != nil {
		return 400, nil, fmt.Errorf("get query param err:%w", err)
	}

	problems, err := problem.Problems(req.Context(), limit, offset)
	if err != nil {
		return 500, nil, fmt.Errorf("query problems limit[%d] offset[%d] err:%w", limit, offset, err)
	}
	count, err := problem.ProblemCount(req.Context())
	if err != nil {
		return 500, nil, fmt.Errorf("query problems count err:%w", err)
	}
	return 200, common.ResponseProblems{Problems: problems, Count: count}, nil
}

func Get(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	value := params.ByName("id")
	problemId, err := strconv.Atoi(value)
	if err != nil {
		return 400, nil, fmt.Errorf("param problem_id[%s] to int err:%w", value, err)
	}
	res, err := problem.Problem(req.Context(), problemId)
	if err != nil {
		return 400, nil, fmt.Errorf("query problem by id[%d] err:%w", problemId, err)
	}
	return 200, res, nil
}

func AddByLeetcodeUrl(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	data := struct {
		Url string `json:"url"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return 400, nil, fmt.Errorf("decode request body err:%w", err)
	}

	slug, err := GetTitleSlugFromUrl(data.Url)
	if err != nil {
		return 400, nil, fmt.Errorf("get title slug err:%w", err)
	}

	res, err := PostGraphQL(req.Context(), slug)
	if err != nil {
		return 400, nil, fmt.Errorf("get leetcode data err:%w", err)
	}

	question := res.Data.Question

	var topics []string
	for _, topic := range question.TopicTags {
		topics = append(topics, topic.TranslatedName)
	}

	id, err := problem.InsertProblem(req.Context(), question.TranslatedTitle, data.Url,
		strings.Join(topics, ";"), question.Difficulty, question.TranslatedContent)
	if err != nil {
		return 400, nil, fmt.Errorf("insert problem err:%w", err)
	}
	return 200, id, nil
}

func Insert(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	data := struct {
		Title      string `json:"title"`
		Url        string `json:"url"`
		Topics     string `json:"topics"`
		Difficulty string `json:"difficulty"`
		Content    string `json:"content"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return 400, nil, fmt.Errorf("decode request body err:%w", err)
	}
	id, err := problem.InsertProblem(req.Context(), data.Title, data.Url, data.Topics, data.Difficulty, data.Content)
	if err != nil {
		return 400, nil, fmt.Errorf("insert problem content[%s] err:%w", data.Content, err)
	}
	return 200, id, nil
}

func Update(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	data := struct {
		Id         int64  `json:"id"`
		Title      string `json:"title"`
		Difficulty string `json:"difficulty"`
		Topics     string `json:"topics"`
		Url        string `json:"url"`
		Content    string `json:"content"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return 400, nil, fmt.Errorf("decode request body err:%w", err)
	}
	err = problem.UpdateProblem(req.Context(), data.Id, data.Title, data.Difficulty, data.Topics, data.Url, data.Content)
	if err != nil {
		return 400, nil, fmt.Errorf("update problem by id[%d] err:%w", data.Id, err)
	}
	return 200, nil, nil
}

func Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	data := struct {
		Id int64 `json:"id"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return 400, nil, fmt.Errorf("decode request body err:%w", err)
	}
	err = problem.DeleteProblem(req.Context(), data.Id)
	if err != nil {
		return 400, nil, fmt.Errorf("delete problem by id[%d] err:%w", data.Id, err)
	}
	return 200, nil, nil
}
