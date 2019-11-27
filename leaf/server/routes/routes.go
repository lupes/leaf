package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"

	"leaf/server/apis"
	"leaf/server/apis/problem"
	"leaf/server/apis/solution"
	"leaf/server/common"
)

type Handler func(w http.ResponseWriter, req *http.Request, params httprouter.Params) (status int, res interface{}, err error)

func toResponse(handler Handler) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		var res common.ResponseCommon
		start := time.Now()
		status, data, err := handler(w, req, params)
		if err != nil {
			logrus.Errorf("%s", err)
			res = common.ResponseCommon{
				Code:    0,
				Message: err.Error(),
				Data:    nil,
			}
		} else {
			res = common.ResponseCommon{
				Code:    1,
				Message: "success",
				Data:    data,
			}
		}
		body, err := json.Marshal(res)
		if err != nil {
			logrus.Errorf("%s", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, err = bytes.NewBuffer(body).WriteTo(w)
		if err != nil {
			logrus.Errorf("%s", err)
		}
		logrus.Infof("use time:%s", time.Now().Sub(start))
	}
}

func InitRoutes(router *httprouter.Router) {
	router.GET("/api/ping", toResponse(apis.Pong))
	router.GET("/api/test", toResponse(apis.Test))

	router.GET("/api/v1/problem", toResponse(problem.List))
	router.GET("/api/v1/problem/:id", toResponse(problem.Get))
	router.POST("/api/v1/problem", toResponse(problem.Insert))
	router.POST("/api/v1/problem/url", toResponse(problem.AddByLeetcodeUrl))
	router.PUT("/api/v1/problem", toResponse(problem.Update))
	router.DELETE("/api/v1/problem", toResponse(problem.Delete))

	router.GET("/api/v1/solution", toResponse(solution.List))
	router.GET("/api/v1/solution/:id", toResponse(solution.Get))
	router.POST("/api/v1/solution", toResponse(solution.Insert))
	router.PUT("/api/v1/solution", toResponse(solution.Update))
	router.DELETE("/api/v1/solution", toResponse(solution.Delete))
}
