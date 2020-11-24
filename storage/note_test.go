package storage

import (
	"context"
	"notes/model"
	"reflect"
	"testing"
)

func TestAddNote(t *testing.T) {
	type args struct {
		ctx     context.Context
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddNote(tt.args.ctx, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("AddNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllNotes(t *testing.T) {
	type args struct {
		ctx    context.Context
		offset int
		size   int
	}
	tests := []struct {
		name    string
		args    args
		want    []model.NewNote
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllNotes(tt.args.ctx, tt.args.offset, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllNotes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNoteByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    model.NewNote
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNoteByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNoteByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNoteByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveNote(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveNote(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RemoveNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
