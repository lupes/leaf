package problem

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"leaf/server/common"
	"leaf/server/database"
)

const (
	selectProblem      = "select id, title, url, topics, difficulty, content, created_time, updated_time from problem where deleted_time is null order by id limit ?, ?"
	selectProblemCount = "select count(*) from problem where deleted_time is null"
	selectProblemById  = "select id, title, url, topics, difficulty, content, created_time, updated_time from problem where id = ? and deleted_time is null"
	insertProblem      = "insert into problem(title, url, topics, difficulty, content, created_time, updated_time) values (?, ?, ?, ?, ?, ?)"
	updateProblem      = "update problem set title=?, url=?, topics=?, difficulty=?, content=?, updated_time = ? where id = ? and deleted_time is null"
	deleteProblem      = "update problem set deleted_time = ? where id = ? and deleted_time is null"
)

func Problems(ctx context.Context, limit, offset int64) (common.Problems, error) {
	db := database.GetDB()
	rows, err := db.QueryContext(ctx, selectProblem, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("query problem err:%w", err)
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("close rows error:%s", err)
		}
	}()
	var problems common.Problems
	for rows.Next() {
		problem := common.Problem{}
		err := rows.Scan(&problem.Id, &problem.Title, &problem.Url, &problem.Topics,
			&problem.Difficulty, &problem.Content, &problem.CreatedTime, &problem.UpdatedTime)
		if err != nil {
			return nil, fmt.Errorf("query problem err:%w", err)
		}
		problems = append(problems, problem)
	}
	return problems, nil
}

func ProblemCount(ctx context.Context) (int, error) {
	var count int
	row := database.GetDB().QueryRowContext(ctx, selectProblemCount)
	err := row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("scan problem err:%w", err)
	}
	return count, nil
}

func Problem(ctx context.Context, problemId int) (common.Problem, error) {
	var problem common.Problem
	row := database.GetDB().QueryRowContext(ctx, selectProblemById, problemId)
	err := row.Scan(&problem.Id, &problem.Title, &problem.Url, &problem.Topics,
		&problem.Difficulty, &problem.Content, &problem.CreatedTime, &problem.UpdatedTime)
	if err != nil {
		return problem, fmt.Errorf("scan problem err:%w", err)
	}
	return problem, nil
}

func InsertProblem(ctx context.Context, title, url, topics, difficulty, content string) (int64, error) {
	db := database.GetDB()
	res, err := db.ExecContext(ctx, insertProblem, title, url, topics, difficulty, content, time.Now(), time.Now())
	if err != nil {
		return 0, fmt.Errorf("insert to db err:%w", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("get insert problem affected err:%w", err)
	}
	if affected != 1 {
		return 0, fmt.Errorf("insert problem affected not one")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get insert problem id err:%w", err)
	}
	return id, nil
}

func UpdateProblem(ctx context.Context, problemId int64, title, url, content string) error {
	db := database.GetDB()
	res, err := db.ExecContext(ctx, updateProblem, title, url, content, time.Now(), problemId)
	if err != nil {
		return fmt.Errorf("update problem err:%w", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("update problem affected err:%w", err)
	}
	if affected != 1 {
		return fmt.Errorf("update problem affected not one")
	}
	return nil
}

func DeleteProblem(ctx context.Context, problemId int64) error {
	db := database.GetDB()
	res, err := db.ExecContext(ctx, deleteProblem, time.Now(), problemId)
	if err != nil {
		return fmt.Errorf("delete problem id[%d] err:%w", problemId, err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("get delete problem affected err:%w", err)
	}
	if affected != 1 {
		return fmt.Errorf("delete affected not one")
	}
	return nil
}
