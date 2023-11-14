package main

import "testing"

func TestGetRepoName(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "valid-repoUrl-1",
			args:    args{url: "https://github.com/123/repo"},
			want:    "repo",
			wantErr: false,
		},
		{
			name:    "valid-repoUrl-2",
			args:    args{url: "https://github.com/123/repo.git"},
			want:    "repo",
			wantErr: false,
		},
		{
			name:    "valid-repoUrl-3",
			args:    args{url: "git@github.com:123/repo.git"},
			want:    "repo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRepoName(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepoName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRepoName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFolderExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "valid-folder",
			args:    args{path: "./config"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "path-not-exists",
			args:    args{path: "./null"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "invalid-file",
			args:    args{path: "./main.go"},
			want:    true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FolderExists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FolderExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FolderExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}
