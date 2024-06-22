package exchange

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	exchangeProto "github.com/game-core/gc-server/api/game/presentation/proto/exchange"
	"github.com/game-core/gc-server/api/game/presentation/proto/exchange/userExchange"
	"github.com/game-core/gc-server/api/game/presentation/proto/exchange/userExchangeItem"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/times"
	exchangeService "github.com/game-core/gc-server/pkg/domain/model/exchange"
	userExchangeModel "github.com/game-core/gc-server/pkg/domain/model/exchange/userExchange"
	userExchangeItemModel "github.com/game-core/gc-server/pkg/domain/model/exchange/userExchangeItem"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
)

func TestExchangeUsecase_NewExchangeUsecase(t *testing.T) {
	type args struct {
		exchangeService    exchangeService.ExchangeService
		transactionService transactionService.TransactionService
	}
	tests := []struct {
		name string
		args args
		want ExchangeUsecase
	}{
		{
			name: "正常",
			args: args{
				exchangeService:    nil,
				transactionService: nil,
			},
			want: &exchangeUsecase{
				exchangeService:    nil,
				transactionService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewExchangeUsecase(tt.args.exchangeService, tt.args.transactionService)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExchangeUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeUsecase_Update(t *testing.T) {
	type fields struct {
		exchangeService    func(ctrl *gomock.Controller) exchangeService.ExchangeService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *exchangeProto.ExchangeUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *exchangeProto.ExchangeUpdateResponse
		wantErr error
	}{
		{
			name: "正常：受け取りできる",
			fields: fields{
				exchangeService: func(ctrl *gomock.Controller) exchangeService.ExchangeService {
					m := exchangeService.NewMockExchangeService(ctrl)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&exchangeService.ExchangeUpdateRequest{
								UserId:           "0:test",
								MasterExchangeId: 1,
							},
						).
						Return(
							&exchangeService.ExchangeUpdateResponse{
								UserExchange: &userExchangeModel.UserExchange{
									UserId:           "0:WntR-PyhOJeDiE5jodeR",
									MasterExchangeId: 1,
									ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
								},
								UserExchangeItems: userExchangeItemModel.UserExchangeItems{
									{
										UserId:               "0:WntR-PyhOJeDiE5jodeR",
										MasterExchangeId:     1,
										MasterExchangeItemId: 1,
										Count:                9,
									},
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
				req: &exchangeProto.ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want: &exchangeProto.ExchangeUpdateResponse{
				UserExchange: &userExchange.UserExchange{
					UserId:           "0:WntR-PyhOJeDiE5jodeR",
					MasterExchangeId: 1,
					ResetAt:          times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC))),
				},
				UserExchangeItems: userExchangeItem.UserExchangeItems{
					{
						UserId:               "0:WntR-PyhOJeDiE5jodeR",
						MasterExchangeId:     1,
						MasterExchangeItemId: 1,
						Count:                9,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.UserMysqlBegin",
			fields: fields{
				exchangeService: func(ctrl *gomock.Controller) exchangeService.ExchangeService {
					m := exchangeService.NewMockExchangeService(ctrl)
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
				req: &exchangeProto.ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.UserMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.exchangeService.Update",
			fields: fields{
				exchangeService: func(ctrl *gomock.Controller) exchangeService.ExchangeService {
					m := exchangeService.NewMockExchangeService(ctrl)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&exchangeService.ExchangeUpdateRequest{
								UserId:           "0:test",
								MasterExchangeId: 1,
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
				req: &exchangeProto.ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.exchangeService.Update", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &exchangeUsecase{
				exchangeService:    tt.fields.exchangeService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Update(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeUsecase_Receive(t *testing.T) {
	type fields struct {
		exchangeService    func(ctrl *gomock.Controller) exchangeService.ExchangeService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *exchangeProto.ExchangeReceiveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *exchangeProto.ExchangeReceiveResponse
		wantErr error
	}{
		{
			name: "正常：受け取りできる",
			fields: fields{
				exchangeService: func(ctrl *gomock.Controller) exchangeService.ExchangeService {
					m := exchangeService.NewMockExchangeService(ctrl)
					m.EXPECT().
						Receive(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&exchangeService.ExchangeReceiveRequest{
								UserId:               "0:test",
								MasterExchangeItemId: 1,
								Count:                1,
							},
						).
						Return(
							&exchangeService.ExchangeReceiveResponse{
								UserExchange: &userExchangeModel.UserExchange{
									UserId:           "0:WntR-PyhOJeDiE5jodeR",
									MasterExchangeId: 1,
									ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
								},
								UserExchangeItem: &userExchangeItemModel.UserExchangeItem{
									UserId:               "0:WntR-PyhOJeDiE5jodeR",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                9,
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
				req: &exchangeProto.ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want: &exchangeProto.ExchangeReceiveResponse{
				UserExchange: &userExchange.UserExchange{
					UserId:           "0:WntR-PyhOJeDiE5jodeR",
					MasterExchangeId: 1,
					ResetAt:          times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC))),
				},
				UserExchangeItem: &userExchangeItem.UserExchangeItem{
					UserId:               "0:WntR-PyhOJeDiE5jodeR",
					MasterExchangeId:     1,
					MasterExchangeItemId: 1,
					Count:                9,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.UserMysqlBegin",
			fields: fields{
				exchangeService: func(ctrl *gomock.Controller) exchangeService.ExchangeService {
					m := exchangeService.NewMockExchangeService(ctrl)
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
				req: &exchangeProto.ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.UserMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.exchangeService.Receive",
			fields: fields{
				exchangeService: func(ctrl *gomock.Controller) exchangeService.ExchangeService {
					m := exchangeService.NewMockExchangeService(ctrl)
					m.EXPECT().
						Receive(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&exchangeService.ExchangeReceiveRequest{
								UserId:               "0:test",
								MasterExchangeItemId: 1,
								Count:                1,
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
				req: &exchangeProto.ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.exchangeService.Receive", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &exchangeUsecase{
				exchangeService:    tt.fields.exchangeService(ctrl),
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
