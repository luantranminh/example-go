package loan

import (
	"context"
	"reflect"
	"testing"

	"github.com/luantranminh/example-go/domain"
)

func Test_validationMiddleware_Create(t *testing.T) {
	serviceMock := &ServiceMock{
		CreateFunc: func(_ context.Context, p *domain.Loan) error {
			return nil
		},
	}
	defaultCtx := context.Background()
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
			mw := validationMiddleware{
				Service: serviceMock,
			}
			if err := mw.Create(defaultCtx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validationMiddleware_Update(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx  context.Context
		loan *domain.Loan
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Loan
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			got, err := mw.Update(tt.args.ctx, tt.args.loan)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validationMiddleware.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validationMiddleware_Delete(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx  context.Context
		loan *domain.Loan
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
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			if err := mw.Delete(tt.args.ctx, tt.args.loan); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
