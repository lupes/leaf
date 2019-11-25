package problem

import (
	"context"
	"testing"

	"leaf/log"
)

var ctx = context.TODO()

func TestMain(m *testing.M) {
	log.Init()
	m.Run()
}

func TestProblems(t *testing.T) {
	got, err := Problems(ctx, 10, 0)
	if err != nil {
		t.Errorf("Problems() error = %v", err)
		return
	}
	t.Logf("%+v\n", got)
}

func TestProblemCount(t *testing.T) {
	got, err := ProblemCount(ctx)
	if err != nil {
		t.Errorf("ProblemCount() error = %v", err)
		return
	}
	t.Logf("%+v\n", got)
}

func TestProblem(t *testing.T) {
	got, err := Problem(ctx, 2)
	if err != nil {
		t.Errorf("Problem() error = %v", err)
		return
	}
	t.Logf("%+v\n", got)
}

func TestInsertProblem(t *testing.T) {
	problems := [][2]string{
		{"title", "this is a problem"},
		{"题目名称", "这是一个问题"},
	}
	for _, problem := range problems {
		id, err := InsertProblem(ctx, problem[0], problem[1])
		if err != nil {
			t.Errorf("InsertProblem() error = %v", err)
			return
		}
		t.Logf("%+v\n", id)
	}
}

func TestUpdateProblem(t *testing.T) {
	title := "更新题目名称"
	content := "update this is a problem"
	err := UpdateProblem(ctx, 2, title, content)
	if err != nil {
		t.Errorf("UpdateProblem() error = %v", err)
		return
	}
}

func TestDeleteProblem(t *testing.T) {
	err := DeleteProblem(ctx, 1)
	if err != nil {
		t.Errorf("DeleteProblem() error = %v", err)
		return
	}
}
