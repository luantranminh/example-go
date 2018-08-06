package loan

import (
	"context"
	"testing"
	"time"

	testutil "github.com/luantranminh/example-go/config/database/pg/util"
	"github.com/luantranminh/example-go/domain"
)

var fakeUserID = domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")
var fakeBookID = domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")
var fakeLoanID = domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

func Test_pgService_Create(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Create dummy book failed by error %v", err)
	}

	user := domain.User{}
	err = testDB.Create(&user).Error
	if err != nil {
		t.Fatalf("Create dummy user failed by error %v", err)
	}

	type args struct {
		p *domain.Loan
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success create loan",
			args: args{&domain.Loan{
				BookID: book.ID,
				UserID: user.ID,
				To:     parseTime("2014-11-12T11:45:26.371Z"),
			}},
			wantErr: false,
		},
		{
			name: "book not exists,",
			args: args{&domain.Loan{
				BookID: fakeBookID,
				UserID: user.ID,
				To:     parseTime("2014-11-12T11:45:26.371Z"),
			}},
			wantErr: true,
		},
		{
			name: "user not exists,",
			args: args{&domain.Loan{
				BookID: book.ID,
				UserID: fakeUserID,
				To:     parseTime("2014-11-12T11:45:26.371Z"),
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
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Create dummy book failed by error %v", err)
	}

	book1 := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Create dummy book failed by error %v", err)
	}

	user := domain.User{}
	err = testDB.Create(&user).Error
	if err != nil {
		t.Fatalf("Create dummy user failed by error %v", err)
	}

	loan := domain.Loan{
		UserID: user.ID,
		BookID: book1.ID,
		To:     parseTime("2014-11-12T11:45:26.371Z"),
	}
	err = testDB.Create(&loan).Error
	if err != nil {
		t.Fatalf("Create dummy action user borrow a book failed by error %v", err)
	}

	type args struct {
		p *domain.Loan
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Loan
		wantErr bool
	}{
		{
			name: "success update loan",
			args: args{&domain.Loan{
				Model:  domain.Model{ID: loan.ID},
				BookID: book.ID,
				UserID: user.ID,
				To:     parseTime("2014-11-12T11:45:26.371Z"),
			}},
			want: &domain.Loan{
				Model:  domain.Model{ID: loan.ID},
				BookID: book.ID,
				UserID: user.ID,
				To:     parseTime("2014-11-12T11:45:26.371Z"),
			},
			wantErr: false,
		},
		{
			name: "loan not exists,",
			args: args{&domain.Loan{
				Model:  domain.Model{ID: fakeLoanID},
				BookID: fakeBookID,
				UserID: user.ID,
				To:     parseTime("2014-11-12T11:45:26.371Z"),
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			got, err := s.Update(context.Background(), tt.args.p)
			if err != nil {
				if tt.wantErr == false {
					t.Errorf("pgService.Update() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}
			}

			if got != nil && tt.want != nil {
				if got.BookID != tt.want.BookID {
					t.Errorf("pgService.Update() got = %v, want %v", got.BookID, tt.want.BookID)
				}
				if got.UserID != tt.want.UserID {
					t.Errorf("pgService.Update() got = %v, want %v", got.BookID, tt.want.BookID)
				}
				if got.ID != tt.want.ID {
					t.Errorf("pgService.Update() got = %v, want %v", got.ID, tt.want.ID)
				}
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
		p *domain.Loan
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

func parseTime(str string) *time.Time {
	time, err := time.Parse(time.RFC3339, "2014-11-12T11:45:26.371Z")

	if err != nil {
		panic(err)
	}

	return &time
}
