package common

import (
	"time"
)

type Problem struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

type Problems []Problem

type Solution struct {
	Id          int64     `json:"id"`
	ProblemId   int64     `json:"problem_id"`
	Title       string    `json:"title"`
	Language    string    `json:"language"`
	Content     string    `json:"content"`
	Caption     string    `json:"caption"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

type Solutions []Solution

type ResponseProblems struct {
	Problems Problems `json:"problems"`
	Count    int      `json:"count"`
}

type ResponseSolutions struct {
	Solutions Solutions `json:"solutions"`
	Count     int       `json:"count"`
}

type ResponseCommon struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
