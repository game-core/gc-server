package loginBonus

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	loginBonusProto "github.com/game-core/gc-server/api/game/presentation/proto/loginBonus"
	"github.com/game-core/gc-server/api/game/presentation/proto/loginBonus/userLoginBonus"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/times"
	loginBonusService "github.com/game-core/gc-server/pkg/domain/model/loginBonus"
	userLoginBonusModel "github.com/game-core/gc-server/pkg/domain/model/loginBonus/userLoginBonus"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
)

func TestLoginBonusUsecase_NewLoginBonusUsecase(t *testing.T) {
	type args struct {
		loginBonusService  loginBonusService.LoginBonusService
		transactionService transactionService.TransactionService
	}
	tests := []struct {
		name string
		args args
		want LoginBonusUsecase
	}{
		{
			name: "正常",
			args: args{
				loginBonusService:  nil,
				transactionService: nil,
			},
			want: &loginBonusUsecase{
				loginBonusService:  nil,
				transactionService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLoginBonusUsecase(tt.args.loginBonusService, tt.args.transactionService)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginBonusUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginBonusUsecase_Receive(t *testing.T) {
	type fields struct {
		loginBonusService  func(ctrl *gomock.Controller) loginBonusService.LoginBonusService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *loginBonusProto.LoginBonusReceiveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *loginBonusProto.LoginBonusReceiveResponse
		wantErr error
	}{
		{
			name: "正常：受け取りできる",
			fields: fields{
				loginBonusService: func(ctrl *gomock.Controller) loginBonusService.LoginBonusService {
					m := loginBonusService.NewMockLoginBonusService(ctrl)
					m.EXPECT().
						Receive(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&loginBonusService.LoginBonusReceiveRequest{
								UserId:             "0:test",
								MasterLoginBonusId: 1,
							},
						).
						Return(
							&loginBonusService.LoginBonusReceiveResponse{
								UserLoginBonus: &userLoginBonusModel.UserLoginBonus{
									UserId:             "0:WntR-PyhOJeDiE5jodeR",
									MasterLoginBonusId: 1,
									ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
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
				req: &loginBonusProto.LoginBonusReceiveRequest{
					UserId:             "0:test",
					MasterLoginBonusId: 1,
				},
			},
			want: &loginBonusProto.LoginBonusReceiveResponse{
				UserLoginBonus: &userLoginBonus.UserLoginBonus{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
					ReceivedAt:         times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC))),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.UserMysqlBegin",
			fields: fields{
				loginBonusService: func(ctrl *gomock.Controller) loginBonusService.LoginBonusService {
					m := loginBonusService.NewMockLoginBonusService(ctrl)
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
				req: &loginBonusProto.LoginBonusReceiveRequest{
					UserId:             "0:test",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.UserMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.loginBonusService.Receive",
			fields: fields{
				loginBonusService: func(ctrl *gomock.Controller) loginBonusService.LoginBonusService {
					m := loginBonusService.NewMockLoginBonusService(ctrl)
					m.EXPECT().
						Receive(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&loginBonusService.LoginBonusReceiveRequest{
								UserId:             "0:test",
								MasterLoginBonusId: 1,
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
				req: &loginBonusProto.LoginBonusReceiveRequest{
					UserId:             "0:test",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.loginBonusService.Receive", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &loginBonusUsecase{
				loginBonusService:  tt.fields.loginBonusService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Receive(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Receive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Receive() = %v, want %v", got, tt.want)
			}
		})
	}
}
