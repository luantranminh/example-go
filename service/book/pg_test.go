package book

import (
	"context"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	testutil "github.com/luantranminh/example-go/config/database/pg/util"
	"github.com/luantranminh/example-go/domain"
)

func Test_pgService_Create(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	type args struct {
		p *domain.Book
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
			s := &pgService{
				db: testDB,
			}
			if err := s.Create(context.Background(), tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pgService_Update(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
		p   *domain.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.Update(tt.args.in0, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgService_Find(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
		p   *domain.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.Find(tt.args.in0, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgService_FindAll(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.FindAll(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgService_Delete(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
		p   *domain.Book
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
			s := &pgService{
				db: tt.fields.db,
			}
			if err := s.Delete(tt.args.in0, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
