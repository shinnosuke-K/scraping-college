package main

import "testing"

func TestGetPageNum(t *testing.T) {
	type args struct {
		itemNum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{itemNum: 1},
			want: 1,
		},
		{
			name: "5",
			args: args{itemNum: 5},
			want: 1,
		},
		{
			name: "10",
			args: args{itemNum: 10},
			want: 1,
		},
		{
			name: "11",
			args: args{itemNum: 11},
			want: 2,
		},
		{
			name: "20",
			args: args{itemNum: 20},
			want: 2,
		},
		{
			name: "30",
			args: args{itemNum: 30},
			want: 3,
		},
		{
			name: "100",
			args: args{itemNum: 100},
			want: 10,
		},
		{
			name: "101",
			args: args{itemNum: 101},
			want: 11,
		},
		{
			name: "110",
			args: args{itemNum: 110},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPageNum(tt.args.itemNum); got != tt.want {
				t.Errorf("GetPageNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckStatus(t *testing.T) {
	type args struct {
		statusCode int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "StatusOK",
			args:    args{statusCode: 200},
			wantErr: false,
		},
		{
			name:    "StatusContinue",
			args:    args{statusCode: 100},
			wantErr: true,
		},
		{
			name:    "StatusMultipleChoices",
			args:    args{statusCode: 300},
			wantErr: true,
		},
		{
			name:    "StatusBadRequest",
			args:    args{statusCode: 400},
			wantErr: true,
		},
		{
			name:    "StatusInternalServerError",
			args:    args{statusCode: 500},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckStatus(tt.args.statusCode); (err != nil) != tt.wantErr {
				t.Errorf("CheckStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
