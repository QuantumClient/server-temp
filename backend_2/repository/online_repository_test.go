package repository

import (
	"github.com/google/uuid"
	"quantumclient.org/backend/v2/models"
	"reflect"
	"testing"
)


func TestOnlineRepo_Add(t *testing.T) {
	o := models.NewOnline(uuid.New(), models.User{})
	r := NewOnlineRepo()
	err := r.Add(o)
	if err != nil {
		t.Errorf("OnlineRepo.Add() error = %v", err)
		return
	}
	if len(r.Online) == 0 {
		t.Errorf("OnlineRepo.Add() error = %v", err)
		return
	}


}

func TestOnlineRepo_Delete(t *testing.T) {
	type fields struct {
		Online []*models.Online
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := OnlineRepo{
				Online: tt.fields.Online,
			}
			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnlineRepo_Get(t *testing.T) {
	type fields struct {
		Online []*models.Online
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Online
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := OnlineRepo{
				Online: tt.fields.Online,
			}
			got, err := r.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOnlineRepo_GetAll(t *testing.T) {
	r := NewOnlineRepo()
	r.Online = append(r.Online, models.NewOnline(uuid.New(), models.User{}))
	got, err := r.GetAll()
	if err != nil {
		t.Errorf("OnlineRepo.GetAll() error = %v", err)
		return
	}
	if len(got) == 0 {
		t.Errorf("OnlineRepo.GetAll() error = %v", err)
		return
	}
}

func TestOnlineRepo_delete(t *testing.T) {
	type fields struct {
		Online []*models.Online
	}
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := OnlineRepo{
				Online: tt.fields.Online,
			}
			if err := r.delete(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnlineRepo_findIndex(t *testing.T) {

}
