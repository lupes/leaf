package problem

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func Insert(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	data := struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return 400, nil, fmt.Errorf("decode request body err:%w", err)
	}
	_, err = problem.InsertProblem(req.Context(), data.Title, data.Content)
	if err != nil {
		return 400, nil, fmt.Errorf("insert problem content[%s] err:%w", data.Content, err)
	}
	return 200, nil, nil
}

func Update(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	data := struct {
		ProblemId int64  `json:"problem_id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return 400, nil, fmt.Errorf("decode request body err:%w", err)
	}
	err = problem.UpdateProblem(req.Context(), data.ProblemId, data.Title, data.Content)
	if err != nil {
		return 400, nil, fmt.Errorf("update problem by id[%d] err:%w", data.ProblemId, err)
	}
	return 200, nil, nil
}

func Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	data := struct {
		ProblemId int64 `json:"problem_id"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return 400, nil, fmt.Errorf("decode request body err:%w", err)
	}
	err = problem.DeleteProblem(req.Context(), data.ProblemId)
	if err != nil {
		return 400, nil, fmt.Errorf("delete problem by id[%d] err:%w", data.ProblemId, err)
	}
	return 200, nil, nil
}
