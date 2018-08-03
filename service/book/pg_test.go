package book

import (
	"context"
	"reflect"
	"testing"

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

	category := domain.Category{}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Errorf("Failed to create (dummy category for create book) with error : %v", err)
	}

	fakeCategoryID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create book successfully",
			args: args{&domain.Book{
				Name:        "this is a book",
				CategoryID:  category.ID,
				Author:      "I am a writer",
				Description: "this is a book",
			}},
			wantErr: false,
		},
		{
			name: "Create book failed, category is not exists",
			args: args{&domain.Book{
				Name:        "this is a book",
				CategoryID:  fakeCategoryID,
				Author:      "I am a writer",
				Description: "this is a book",
			}},
			wantErr: true,
		},
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
	serviceMock := &ServiceMock{
		UpdateFunc: func(_ context.Context, p *domain.Book) (*domain.Book, error) {
			return p, nil
		},
	}

	defaultCtx := context.Background()

	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name       string
		args       args
		wantOutput *domain.Book
		wantErr    error
	}{
		{
			name: "valid book",
			args: args{&domain.Book{
				Name:        "Nameisvalid",
				Author:      "a",
				Description: "aaaaaa",
			}},
			wantOutput: &domain.Book{
				Name:        "Nameisvalid",
				Author:      "a",
				Description: "aaaaaa",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}

			got, err := mw.Update(defaultCtx, tt.args.p)
			if err != nil {
				if tt.wantErr != err {
					t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.wantOutput) {
				t.Errorf("validationMiddleware.Update() = %v, want %v", got, tt.wantOutput)
			}
		})
	}
}

func Test_pgService_Delete(t *testing.T) {
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
			if err := s.Delete(context.Background(), tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
