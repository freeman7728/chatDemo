package test

import (
	"chat/pkg/utils"
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		id       int64
		username string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "1", args: args{id: 1, username: "123123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.GenerateToken(tt.args.id, tt.args.username)
			fmt.Println(got, err)
		})
	}
}
