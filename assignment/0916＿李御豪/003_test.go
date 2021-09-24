package main

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestTeacherDAO_CreateTeacher(t *testing.T) {
	type fields struct {
		mq MessageQueue
		db *gorm.DB
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Teacher
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := TeacherDAO{
				mq: tt.fields.mq,
				db: tt.fields.db,
			}
			got, err := dao.CreateTeacher(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("TeacherDAO.CreateTeacher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TeacherDAO.CreateTeacher() = %v, want %v", got, tt.want)
			}
		})
	}
}
