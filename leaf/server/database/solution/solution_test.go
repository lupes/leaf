package solution

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

func TestSolutions(t *testing.T) {
	got, err := Solutions(ctx, 2, 10, 0)
	if err != nil {
		t.Errorf("Solutions() error = %v", err)
		return
	}
	t.Logf("%+v", got)
}

func TestSolutionCount(t *testing.T) {
	got, err := SolutionCount(ctx, 2)
	if err != nil {
		t.Errorf("SolutionCount() error = %v", err)
		return
	}
	t.Logf("%+v", got)
}

func TestSolution(t *testing.T) {
	got, err := Solution(ctx, 2)
	if err != nil {
		t.Errorf("Solution() error = %v", err)
		return
	}
	t.Logf("%+v", got)
}

func TestInsertSolution(t *testing.T) {
	title := "Hello World"
	language := "golang"
	content := `func main() {

}`
	caption := "这是一个示例"
	got, err := InsertSolution(ctx, 2, title, language, content, caption)
	if err != nil {
		t.Errorf("InsertSolution() error = %v", err)
		return
	}
	t.Logf("%d\n", got)
}

func TestUpdateSolution(t *testing.T) {
	title := "Hello World"
	language := "golang"
	content := `func main() {
	fmt.Println("Hello World")
}`
	caption := "这是一个示例"

	err := UpdateSolution(ctx, 2, title, language, content, caption)
	if err != nil {
		t.Errorf("UpdateSolution() error = %v", err)
	}
}

func TestDeleteSolution(t *testing.T) {
	err := DeleteSolution(ctx, 1)
	if err != nil {
		t.Errorf("DeleteSolution() error = %v", err)
	}
}
