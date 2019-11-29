package solution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"leaf/server/common"
	"leaf/server/database/solution"
)

// 题解
func getInt64Param(req *http.Request, key string) (int64, error) {
	value := req.FormValue(key)
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("query param %s[%s] not is be int err:%w", key, value, err)
	}
	return int64(number), nil
}

func List(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	problemId, err := getInt64Param(req, "problem_id")
	if err != nil {
		return 400, nil, fmt.Errorf("get query param err:%w", err)
	}
	limit, err := getInt64Param(req, "limit")
	if err != nil {
		return 400, nil, fmt.Errorf("get query param err:%w", err)
	}

	offset, err := getInt64Param(req, "offset")
	if err != nil {
		return 400, nil, fmt.Errorf("get query param err:%w", err)
	}
	solutions, err := solution.Solutions(req.Context(), problemId, limit, offset)
	if err != nil {
		return 500, nil, fmt.Errorf("query solution by problem id[%d] limit[%d] offset[%d] err:%w", problemId, limit, offset, err)
	}
	count, err := solution.SolutionCount(req.Context(), problemId)
	if err != nil {
		return 500, nil, fmt.Errorf("query solution count err:%w", err)
	}
	return 200, common.ResponseSolutions{Solutions: solutions, Count: count}, nil
}

func Get(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	value := params.ByName("id")
	solutionId, err := strconv.Atoi(value)
	if err != nil {
		return 400, nil, fmt.Errorf("param solution_id[%s] to int err:%w", value, err)
	}
	res, err := solution.Solution(req.Context(), solutionId)
	if err != nil {
		return 400, nil, fmt.Errorf("query solution by id[%d] err:%w", solutionId, err)
	}
	return 200, res, nil
}

func Insert(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	data := struct {
		ProblemId int64  `json:"problem_id"`
		Title     string `json:"title"`
		Language  string `json:"language"`
		Content   string `json:"content"`
		Caption   string `json:"caption"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return 400, nil, fmt.Errorf("decode request body err:%w", err)
	}
	id, err := solution.InsertSolution(req.Context(), data.ProblemId, data.Title, data.Language, data.Content, data.Caption)
	if err != nil {
		return 400, nil, fmt.Errorf("insert solution err:%w", err)
	}
	return 200, id, nil
}

func Update(w http.ResponseWriter, req *http.Request, params httprouter.Params) (int, interface{}, error) {
	data := struct {
		Id       int64  `json:"id"`
		Title    string `json:"title"`
		Language string `json:"language"`
		Content  string `json:"content"`
		Caption  string `json:"caption"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return 400, nil, fmt.Errorf("decode request body err:%w", err)
	}
	err = solution.UpdateSolution(req.Context(), data.Id, data.Title, data.Language, data.Content, data.Caption)
	if err != nil {
		return 400, nil, fmt.Errorf("update solution by id[%d] err:%w", data.Id, err)
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
	err = solution.DeleteSolution(req.Context(), data.Id)
	if err != nil {
		return 400, nil, fmt.Errorf("delete solution by id[%d] err:%w", data.Id, err)
	}
	return 200, nil, nil
}
