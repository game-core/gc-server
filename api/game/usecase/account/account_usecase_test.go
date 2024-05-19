package account

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"

	accountServer "github.com/game-core/gc-server/api/game/presentation/server/account"
	"github.com/game-core/gc-server/api/game/presentation/server/account/userAccount"
	"github.com/game-core/gc-server/internal/errors"
	accountService "github.com/game-core/gc-server/pkg/domain/model/account"
	userAccountModel "github.com/game-core/gc-server/pkg/domain/model/account/userAccount"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
)

func TestAccountUsecase_NewAccountUsecase(t *testing.T) {
	type args struct {
		accountService     accountService.AccountService
		transactionService transactionService.TransactionService
	}
	tests := []struct {
		name string
		args args
		want AccountUsecase
	}{
		{
			name: "正常",
			args: args{
				accountService:     nil,
				transactionService: nil,
			},
			want: &accountUsecase{
				accountService:     nil,
				transactionService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAccountUsecase(tt.args.accountService, tt.args.transactionService)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUsecase_Get(t *testing.T) {
	type fields struct {
		accountService     func(ctrl *gomock.Controller) accountService.AccountService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *accountServer.AccountGetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *accountServer.AccountGetResponse
		wantErr error
	}{
		{
			name: "正常：確認できる",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) accountService.AccountService {
					m := accountService.NewMockAccountService(ctrl)
					m.EXPECT().
						Get(
							gomock.Any(),
							&accountService.AccountGetRequest{
								UserId: "0:test",
							},
						).
						Return(
							&accountService.AccountGetResponse{
								UserAccount: &userAccountModel.UserAccount{
									UserId:   "0:test",
									Name:     "test_user_account",
									Password: "test",
									LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &accountServer.AccountGetRequest{
					UserId: "0:test",
				},
			},
			want: &accountServer.AccountGetResponse{
				UserAccount: &userAccount.UserAccount{
					UserId:   "0:test",
					Name:     "test_user_account",
					Password: "test",
					LoginAt:  timestamppb.New(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
					LogoutAt: timestamppb.New(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.accountService.Get",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) accountService.AccountService {
					m := accountService.NewMockAccountService(ctrl)
					m.EXPECT().
						Get(
							gomock.Any(),
							&accountService.AccountGetRequest{
								UserId: "0:test",
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &accountServer.AccountGetRequest{
					UserId: "0:test",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.accountService.Get", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &accountUsecase{
				accountService:     tt.fields.accountService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Get(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUsecase_Create(t *testing.T) {
	type fields struct {
		accountService     func(ctrl *gomock.Controller) accountService.AccountService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *accountServer.AccountCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *accountServer.AccountCreateResponse
		wantErr error
	}{
		{
			name: "正常：ユーザーアカウントが作成できる",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) accountService.AccountService {
					m := accountService.NewMockAccountService(ctrl)
					m.EXPECT().
						CreateUserId(
							nil,
						).
						Return(
							"0:test",
							nil,
						)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							&accountService.AccountCreateRequest{
								UserId: "0:test",
								Name:   "test_user_account",
							},
						).
						Return(
							&accountService.AccountCreateResponse{
								UserAccount: &userAccountModel.UserAccount{
									UserId:   "0:test",
									Name:     "test_user_account",
									Password: "test",
									LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						UserMysqlBegin(
							gomock.Any(),
							"0:test",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						UserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &accountServer.AccountCreateRequest{
					Name: "test_user_account",
				},
			},
			want: &accountServer.AccountCreateResponse{
				UserAccount: &userAccount.UserAccount{
					UserId:   "0:test",
					Name:     "test_user_account",
					Password: "test",
					LoginAt:  timestamppb.New(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
					LogoutAt: timestamppb.New(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.accountService.CreateUserId",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) accountService.AccountService {
					m := accountService.NewMockAccountService(ctrl)
					m.EXPECT().
						CreateUserId(
							nil,
						).
						Return(
							"",
							errors.NewTestError(),
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &accountServer.AccountCreateRequest{
					Name: "test_user_account",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.accountService.CreateUserId", errors.NewTestError()),
		},
		{
			name: "異常：s.transactionService.UserMysqlBegin",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) accountService.AccountService {
					m := accountService.NewMockAccountService(ctrl)
					m.EXPECT().
						CreateUserId(
							nil,
						).
						Return(
							"0:test",
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						UserMysqlBegin(
							gomock.Any(),
							"0:test",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &accountServer.AccountCreateRequest{
					Name: "test_user_account",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.UserMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.accountService.Create",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) accountService.AccountService {
					m := accountService.NewMockAccountService(ctrl)
					m.EXPECT().
						CreateUserId(
							nil,
						).
						Return(
							"0:test",
							nil,
						)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							&accountService.AccountCreateRequest{
								UserId: "0:test",
								Name:   "test_user_account",
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						UserMysqlBegin(
							gomock.Any(),
							"0:test",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						UserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &accountServer.AccountCreateRequest{
					Name: "test_user_account",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.accountService.Create", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &accountUsecase{
				accountService:     tt.fields.accountService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Create(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUsecase_Login(t *testing.T) {
	type fields struct {
		accountService     func(ctrl *gomock.Controller) accountService.AccountService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *accountServer.AccountLoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *accountServer.AccountLoginResponse
		wantErr error
	}{
		{
			name: "正常：ログインできる",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) accountService.AccountService {
					m := accountService.NewMockAccountService(ctrl)
					m.EXPECT().
						Login(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&accountService.AccountLoginRequest{
								UserId:   "0:test",
								Password: "test",
							},
						).
						Return(
							&accountService.AccountLoginResponse{
								Token: "token",
								UserAccount: &userAccountModel.UserAccount{
									UserId:   "0:test",
									Name:     "test_user_account",
									Password: "test",
									LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						UserMysqlBegin(
							gomock.Any(),
							"0:test",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						UserRedisBegin().
						Return(
							nil,
						)
					m.EXPECT().
						UserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						)
					m.EXPECT().
						UserRedisEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &accountServer.AccountLoginRequest{
					UserId:   "0:test",
					Password: "test",
				},
			},
			want: &accountServer.AccountLoginResponse{
				Token: "token",
				UserAccount: &userAccount.UserAccount{
					UserId:   "0:test",
					Name:     "test_user_account",
					Password: "test",
					LoginAt:  timestamppb.New(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
					LogoutAt: timestamppb.New(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.UserMysqlBegin",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) accountService.AccountService {
					m := accountService.NewMockAccountService(ctrl)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						UserMysqlBegin(
							gomock.Any(),
							"0:test",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &accountServer.AccountLoginRequest{
					UserId:   "0:test",
					Password: "test",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.UserMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.accountService.Login",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) accountService.AccountService {
					m := accountService.NewMockAccountService(ctrl)
					m.EXPECT().
						Login(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&accountService.AccountLoginRequest{
								UserId:   "0:test",
								Password: "test",
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						UserMysqlBegin(
							gomock.Any(),
							"0:test",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						UserRedisBegin().
						Return(
							nil,
						)
					m.EXPECT().
						UserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						)
					m.EXPECT().
						UserRedisEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &accountServer.AccountLoginRequest{
					UserId:   "0:test",
					Password: "test",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.accountService.Login", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &accountUsecase{
				accountService:     tt.fields.accountService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Login(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
