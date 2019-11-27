package problem

import (
	"context"
	"testing"
)

func TestPostGraphQL(t *testing.T) {
	tests := []struct {
		titleSlug string
	}{
		{"set-mismatch"},
		{"maximum-average-subarray-i"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got, err := PostGraphQL(context.TODO(), tt.titleSlug)
			if err != nil {
				t.Errorf("PostGraphQL() error = %v", err)
				return
			}
			t.Logf("%+v\n", got)
		})
	}
}

func TestGetTitleSlugFromUrl(t *testing.T) {
	tests := []struct {
		str string
	}{
		{"https://leetcode-cn.com/problems/largest-rectangle-in-histogram/"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got, err := GetTitleSlugFromUrl(tt.str)
			if err != nil {
				t.Errorf("GetTitleSlugFromUrl() error = %v", err)
				return
			}
			t.Logf(got)
		})
	}
}
