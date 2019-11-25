package common

import (
	"time"
)

type Problem struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

type Solution struct {
	Id          int64     `json:"id"`
	ProblemId   int64     `json:"problem_id"`
	Language    string    `json:"language"`
	Content     string    `json:"content"`
	Caption     string    `json:"caption"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

type ResponseProblems struct {
	Problems []Problem `json:"problems"`
	Count    int       `json:"count"`
}

type ResponseSolutions struct {
	Solutions []Solution `json:"solutions"`
	Count     int        `json:"count"`
}

type ResponseCommon struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
