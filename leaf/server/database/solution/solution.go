package solution

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"leaf/server/common"
	"leaf/server/database"
)

const (
	selectSolution      = "select id, problem_id, title, language, content, caption, created_time, updated_time from solution where problem_id=? and deleted_time is null order by id limit ?, ?"
	selectSolutionCount = "select count(*) from solution where problem_id=? and deleted_time is null"
	selectSolutionById  = "select id, problem_id, title, language, content, caption, created_time, updated_time from solution where id=? and deleted_time is null"
	insertSolution      = "insert into solution(problem_id, title, language, content, caption, created_time, updated_time) values (?, ?, ?, ?, ?, ?, ?)"
	updateSolution      = "update solution set title=?, language=?, content=?, caption=?, updated_time=? where id=? and deleted_time is null"
	deleteSolution      = "update solution set deleted_time=? where id=? and deleted_time is null"
)

func Solutions(ctx context.Context, problemId, limit, offset int64) ([]common.Solution, error) {
	db := database.GetDB()
	rows, err := db.QueryContext(ctx, selectSolution, problemId, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("query solution err:%w", err)
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("close rows error:%s", err)
		}
	}()
	var solutions []common.Solution
	for rows.Next() {
		solution := common.Solution{}
		err := rows.Scan(&solution.Id, &solution.ProblemId, &solution.Title, &solution.Language, &solution.Content, &solution.Caption, &solution.CreatedTime, &solution.UpdatedTime)
		if err != nil {
			return nil, fmt.Errorf("scan solution err:%w", err)
		}
		solutions = append(solutions, solution)
	}
	return solutions, nil
}

func SolutionCount(ctx context.Context, problemId int64) (int, error) {
	var count int
	row := database.GetDB().QueryRowContext(ctx, selectSolutionCount, problemId)
	err := row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("scan problem solution count err:%w", err)
	}
	return count, nil
}

func Solution(ctx context.Context, id int) (common.Solution, error) {
	var solution common.Solution
	row := database.GetDB().QueryRowContext(ctx, selectSolutionById, id)
	err := row.Scan(&solution.Id, &solution.ProblemId, &solution.Title, &solution.Language, &solution.Content, &solution.Caption, &solution.CreatedTime, &solution.UpdatedTime)
	if err != nil {
		return solution, fmt.Errorf("scan solution err:%w", err)
	}
	return solution, nil
}

func InsertSolution(ctx context.Context, problemId int64, title, language, content, caption string) (int64, error) {
	db := database.GetDB()
	res, err := db.ExecContext(ctx, insertSolution, problemId, title, language, content, caption, time.Now(), time.Now())
	if err != nil {
		return 0, fmt.Errorf("insert to db err:%w", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("get insert solution affected err:%w", err)
	}
	if affected != 1 {
		return 0, fmt.Errorf("insert solution affected not one")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get insert solution id err:%w", err)
	}
	return id, nil
}

func UpdateSolution(ctx context.Context, id int64, title, language, content, caption string) error {
	db := database.GetDB()
	res, err := db.ExecContext(ctx, updateSolution, title, language, content, caption, time.Now(), id)
	if err != nil {
		return fmt.Errorf("update solution err:%w", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("update solution affected err:%w", err)
	}
	if affected != 1 {
		return fmt.Errorf("update solution affected not one")
	}
	return nil
}

func DeleteSolution(ctx context.Context, id int64) error {
	db := database.GetDB()
	res, err := db.ExecContext(ctx, deleteSolution, time.Now(), id)
	if err != nil {
		return fmt.Errorf("get insert solution id err:%w", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("get insert solution affected err:%w", err)
	}
	if affected != 1 {
		return fmt.Errorf("insert solution affected not one")
	}
	return nil
}
