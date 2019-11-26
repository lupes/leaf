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
	problems := [][3]string{
		{"最优除法", "https://leetcode-cn.com/problems/optimal-division/", `给定一组正整数，相邻的整数之间将会进行浮点除法操作。例如， [2,3,4] -> 2 / 3 / 4 。

但是，你可以在任意位置添加任意数目的括号，来改变算数的优先级。你需要找出怎么添加括号，才能得到最大的结果，并且返回相应的字符串格式的表达式。你的表达式不应该含有冗余的括号。`},
		{"砖墙", "https://leetcode-cn.com/problems/brick-wall/", `你的面前有一堵方形的、由多行砖块组成的砖墙。 这些砖块高度相同但是宽度不同。你现在要画一条自顶向下的、穿过最少砖块的垂线。

砖墙由行的列表表示。 每一行都是一个代表从左至右每块砖的宽度的整数列表。

如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。你需要找出怎样画才能使这条线穿过的砖块数量最少，并且返回穿过的砖块数量。

你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。`},
	}
	for _, problem := range problems {
		id, err := InsertProblem(ctx, problem[0], problem[1], problem[2])
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
