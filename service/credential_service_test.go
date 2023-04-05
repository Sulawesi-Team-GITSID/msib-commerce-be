package service

import (
	"backend-service/configuration/config"
	"backend-service/configuration/repository"
	"context"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	cfg, err := config.NewConfig("./../.env")
	config.CheckError(err)

	db = config.OpenDatabase(cfg)
}

func TestLogin(t *testing.T) {
	repo := repository.NewCredentialRepository(db)
	svc := NewCredentialService(repo)

	type args struct {
		email    string
		password string
	}

	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		{
			"Success",
			args{
				email:    "test@test.com",
				password: "12345678",
			},
			func() {

			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			if _, err := svc.Login(context.Background(), tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("[TestLogin] error: %v, wantErr: %v", err, tt.wantErr)
			}
		})
	}
}
