package exchange

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/event"
	"github.com/game-core/gc-server/pkg/domain/model/event/masterEvent"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/masterExchange"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/masterExchangeCost"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/masterExchangeItem"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchange"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchangeItem"
	"github.com/game-core/gc-server/pkg/domain/model/item"
	"github.com/game-core/gc-server/pkg/domain/model/item/userItemBox"
)

func TestExchangeService_NewExchangeService(t *testing.T) {
	type args struct {
		itemService                       item.ItemService
		eventService                      event.EventService
		masterExchangeMysqlRepository     masterExchange.MasterExchangeMysqlRepository
		masterExchangeCostMysqlRepository masterExchangeCost.MasterExchangeCostMysqlRepository
		masterExchangeItemMysqlRepository masterExchangeItem.MasterExchangeItemMysqlRepository
		userExchangeMysqlRepository       userExchange.UserExchangeMysqlRepository
		userExchangeItemMysqlRepository   userExchangeItem.UserExchangeItemMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want ExchangeService
	}{
		{
			name: "正常",
			args: args{
				itemService:                       nil,
				eventService:                      nil,
				masterExchangeMysqlRepository:     nil,
				masterExchangeCostMysqlRepository: nil,
				masterExchangeItemMysqlRepository: nil,
				userExchangeMysqlRepository:       nil,
				userExchangeItemMysqlRepository:   nil,
			},
			want: &exchangeService{
				itemService:                       nil,
				eventService:                      nil,
				masterExchangeMysqlRepository:     nil,
				masterExchangeCostMysqlRepository: nil,
				masterExchangeItemMysqlRepository: nil,
				userExchangeMysqlRepository:       nil,
				userExchangeItemMysqlRepository:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewExchangeService(
				tt.args.itemService,
				tt.args.eventService,
				tt.args.masterExchangeMysqlRepository,
				tt.args.masterExchangeCostMysqlRepository,
				tt.args.masterExchangeItemMysqlRepository,
				tt.args.userExchangeMysqlRepository,
				tt.args.userExchangeItemMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExchangeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeService_Update(t *testing.T) {
	type fields struct {
		itemService                       func(ctrl *gomock.Controller) item.ItemService
		eventService                      func(ctrl *gomock.Controller) event.EventService
		masterExchangeMysqlRepository     func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository
		masterExchangeCostMysqlRepository func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository
		masterExchangeItemMysqlRepository func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository
		userExchangeMysqlRepository       func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository
		userExchangeItemMysqlRepository   func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		now time.Time
		req *ExchangeUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ExchangeUpdateResponse
		wantErr error
	}{
		{
			name: "正常：更新できる場合",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeItem.MasterExchangeItems{
								{
									MasterExchangeItemId: 1,
									MasterExchangeId:     1,
									MasterItemId:         1,
									Name:                 "アイテム1",
									Count:                10,
								},
								{
									MasterExchangeItemId: 2,
									MasterExchangeId:     1,
									MasterItemId:         2,
									Name:                 "アイテム2",
									Count:                10,
								},
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						CreateList(
							nil,
							nil,
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                10,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                10,
								},
							},
						).
						Return(
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                10,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                10,
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want: &ExchangeUpdateResponse{
				UserExchange: &userExchange.UserExchange{
					UserId:           "0:test",
					MasterExchangeId: 1,
					ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				},
				UserExchangeItems: userExchangeItem.UserExchangeItems{
					{
						UserId:               "0:test",
						MasterExchangeId:     1,
						MasterExchangeItemId: 1,
						Count:                10,
					},
					{
						UserId:               "0:test",
						MasterExchangeId:     1,
						MasterExchangeItemId: 2,
						Count:                10,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：更新できる場合",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeItem.MasterExchangeItems{
								{
									MasterExchangeItemId: 1,
									MasterExchangeId:     1,
									MasterItemId:         1,
									Name:                 "アイテム1",
									Count:                10,
								},
								{
									MasterExchangeItemId: 2,
									MasterExchangeId:     1,
									MasterItemId:         2,
									Name:                 "アイテム2",
									Count:                10,
								},
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByUserIdAndMasterExchangeId(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                0,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                0,
								},
							},
							nil,
						)
					m.EXPECT().
						DeleteList(
							nil,
							nil,
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                0,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                0,
								},
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						CreateList(
							nil,
							nil,
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                10,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                10,
								},
							},
						).
						Return(
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                10,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                10,
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want: &ExchangeUpdateResponse{
				UserExchange: &userExchange.UserExchange{
					UserId:           "0:test",
					MasterExchangeId: 1,
					ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
				},
				UserExchangeItems: userExchangeItem.UserExchangeItems{
					{
						UserId:               "0:test",
						MasterExchangeId:     1,
						MasterExchangeItemId: 1,
						Count:                10,
					},
					{
						UserId:               "0:test",
						MasterExchangeId:     1,
						MasterExchangeItemId: 2,
						Count:                10,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：更新できる場合",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByUserIdAndMasterExchangeId(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                0,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                0,
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want: &ExchangeUpdateResponse{
				UserExchange: &userExchange.UserExchange{
					UserId:           "0:test",
					MasterExchangeId: 1,
					ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
				},
				UserExchangeItems: userExchangeItem.UserExchangeItems{
					{
						UserId:               "0:test",
						MasterExchangeId:     1,
						MasterExchangeItemId: 1,
						Count:                0,
					},
					{
						UserId:               "0:test",
						MasterExchangeId:     1,
						MasterExchangeItemId: 2,
						Count:                0,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：failed to s.masterExchangeMysqlRepository.Find: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.masterExchangeMysqlRepository.Find: test"),
		},
		{
			name: "異常：failed to s.getEvent: failed to s.eventService.Get: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getEvent: failed to s.eventService.Get: test"),
		},
		{
			name: "異常: failed to s.getEvent: outside the event period",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 8, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getEvent: outside the event period"),
		},
		{
			name: "異常：failed to s.reset: failed to s.userExchangeMysqlRepository.Find: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.reset: failed to s.userExchangeMysqlRepository.Find: test"),
		},
		{
			name: "異常：failed to s.reset: failed to s.userExchangeMysqlRepository.Create: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.reset: failed to s.userExchangeMysqlRepository.Create: test"),
		},
		{
			name: "異常：failed to s.reset: failed to s.setUserExchangeItemModels: failed to s.masterExchangeItemMysqlRepository.FindListByMasterExchangeId: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeId(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.reset: failed to s.setUserExchangeItemModels: failed to s.masterExchangeItemMysqlRepository.FindListByMasterExchangeId: test"),
		},
		{
			name: "異常：failed to s.reset: failed to s.setUserExchangeItemModels: failed to s.userExchangeItemMysqlRepository.CreateList: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeItem.MasterExchangeItems{
								{
									MasterExchangeItemId: 1,
									MasterExchangeId:     1,
									MasterItemId:         1,
									Name:                 "アイテム1",
									Count:                10,
								},
								{
									MasterExchangeItemId: 2,
									MasterExchangeId:     1,
									MasterItemId:         2,
									Name:                 "アイテム2",
									Count:                10,
								},
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						CreateList(
							nil,
							nil,
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                10,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                10,
								},
							},
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
				tx:  nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.reset: failed to s.setUserExchangeItemModels: failed to s.userExchangeItemMysqlRepository.CreateList: test"),
		},
		{
			name: "異常：failed to s.reset: failed to s.userExchangeItemMysqlRepository.DeleteList: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByUserIdAndMasterExchangeId(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                0,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                0,
								},
							},
							nil,
						)
					m.EXPECT().
						DeleteList(
							nil,
							nil,
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                0,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                0,
								},
							},
						).
						Return(
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.reset: failed to s.userExchangeItemMysqlRepository.DeleteList: test"),
		},
		{
			name: "異常：failed to s.reset: failed to s.userExchangeMysqlRepository.Update: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByUserIdAndMasterExchangeId(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                0,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                0,
								},
							},
							nil,
						)
					m.EXPECT().
						DeleteList(
							nil,
							nil,
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                0,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                0,
								},
							},
						).
						Return(
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.reset: failed to s.userExchangeMysqlRepository.Update: test"),
		},
		{
			name: "異常：failed to s.reset: failed to s.createUserExchangeItemModels: failed to s.masterExchangeItemMysqlRepository.FindListByMasterExchangeId: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeId(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByUserIdAndMasterExchangeId(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                0,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                0,
								},
							},
							nil,
						)
					m.EXPECT().
						DeleteList(
							nil,
							nil,
							userExchangeItem.UserExchangeItems{
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 1,
									Count:                0,
								},
								{
									UserId:               "0:test",
									MasterExchangeId:     1,
									MasterExchangeItemId: 2,
									Count:                0,
								},
							},
						).
						Return(
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.reset: failed to s.createUserExchangeItemModels: failed to s.masterExchangeItemMysqlRepository.FindListByMasterExchangeId: test"),
		},
		{
			name: "異常：failed to s.reset: failed to s.userExchangeItemMysqlRepository.FindListByUserIdAndMasterExchangeId: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "交換テスト",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByUserIdAndMasterExchangeId(
							nil,
							"0:test",
							int64(1),
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
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeUpdateRequest{
					UserId:           "0:test",
					MasterExchangeId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.reset: failed to s.userExchangeItemMysqlRepository.FindListByUserIdAndMasterExchangeId: test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &exchangeService{
				itemService:                       tt.fields.itemService(ctrl),
				eventService:                      tt.fields.eventService(ctrl),
				masterExchangeMysqlRepository:     tt.fields.masterExchangeMysqlRepository(ctrl),
				masterExchangeCostMysqlRepository: tt.fields.masterExchangeCostMysqlRepository(ctrl),
				masterExchangeItemMysqlRepository: tt.fields.masterExchangeItemMysqlRepository(ctrl),
				userExchangeMysqlRepository:       tt.fields.userExchangeMysqlRepository(ctrl),
				userExchangeItemMysqlRepository:   tt.fields.userExchangeItemMysqlRepository(ctrl),
			}

			got, err := s.Update(tt.args.ctx, tt.args.tx, tt.args.now, tt.args.req)
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

func TestExchangeService_Receive(t *testing.T) {
	type fields struct {
		itemService                       func(ctrl *gomock.Controller) item.ItemService
		eventService                      func(ctrl *gomock.Controller) event.EventService
		masterExchangeMysqlRepository     func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository
		masterExchangeCostMysqlRepository func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository
		masterExchangeItemMysqlRepository func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository
		userExchangeMysqlRepository       func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository
		userExchangeItemMysqlRepository   func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		now time.Time
		req *ExchangeReceiveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ExchangeReceiveResponse
		wantErr error
	}{
		{
			name: "正常:受け取れる場合",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Consume(
							nil,
							nil,
							time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
							&item.ItemConsumeRequest{
								UserId: "0:test",
								Items: item.Items{
									{
										MasterItemId: 2,
										Count:        1,
									},
								},
							},
						).
						Return(
							&item.ItemConsumeResponse{
								UserItemBoxes: userItemBox.UserItemBoxes{
									{
										UserId:       "0:test",
										MasterItemId: 2,
										Count:        1,
									},
								},
							},
							nil,
						)
					m.EXPECT().
						Receive(
							nil,
							nil,
							time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
							&item.ItemReceiveRequest{
								UserId: "0:test",
								Items: item.Items{
									{
										MasterItemId: 1,
										Count:        1,
									},
								},
							},
						).
						Return(
							&item.ItemReceiveResponse{
								UserItemBoxes: userItemBox.UserItemBoxes{
									{
										UserId:       "0:test",
										MasterItemId: 1,
										Count:        1,
									},
								},
							},
							nil,
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "テスト交換",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeItemId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeCost.MasterExchangeCosts{
								{
									MasterExchangeCostId: 1,
									MasterExchangeItemId: 1,
									MasterItemId:         2,
									Name:                 "テスト消費アイテム1",
									Count:                1,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchangeItem.MasterExchangeItem{
								MasterExchangeItemId: 1,
								MasterExchangeId:     1,
								MasterItemId:         1,
								Name:                 "テストアイテム1",
								Count:                10,
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchangeItem.UserExchangeItem{
								UserId:               "0:test",
								MasterExchangeId:     1,
								MasterExchangeItemId: 1,
								Count:                10,
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userExchangeItem.UserExchangeItem{
								UserId:               "0:test",
								MasterExchangeId:     1,
								MasterExchangeItemId: 1,
								Count:                9,
							},
						).
						Return(
							&userExchangeItem.UserExchangeItem{
								UserId:               "0:test",
								MasterExchangeId:     1,
								MasterExchangeItemId: 1,
								Count:                9,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want: &ExchangeReceiveResponse{
				UserExchange: &userExchange.UserExchange{
					UserId:           "0:test",
					MasterExchangeId: 1,
					ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
				},
				UserExchangeItem: &userExchangeItem.UserExchangeItem{
					UserId:               "0:test",
					MasterExchangeId:     1,
					MasterExchangeItemId: 1,
					Count:                9,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常:failed to s.getMasterExchangeItemModelsAndMasterExchangeCostModels: failed to s.masterExchangeItemMysqlRepository.Find: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getMasterExchangeItemModelsAndMasterExchangeCostModels: failed to s.masterExchangeItemMysqlRepository.Find: test"),
		},
		{
			name: "異常:failed to s.getMasterExchangeItemModelsAndMasterExchangeCostModels: failed to s.masterExchangeCostMysqlRepository.FindByMasterExchangeItemId: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeItemId(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchangeItem.MasterExchangeItem{
								MasterExchangeItemId: 1,
								MasterExchangeId:     1,
								MasterItemId:         1,
								Name:                 "テストアイテム1",
								Count:                10,
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getMasterExchangeItemModelsAndMasterExchangeCostModels: failed to s.masterExchangeCostMysqlRepository.FindByMasterExchangeItemId: test"),
		},
		{
			name: "異常:failed to s.getMasterExchangeItemModelsAndMasterExchangeCostModels: failed to s.masterExchangeMysqlRepository.Find: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeItemId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeCost.MasterExchangeCosts{
								{
									MasterExchangeCostId: 1,
									MasterExchangeItemId: 1,
									MasterItemId:         2,
									Name:                 "テスト消費アイテム1",
									Count:                1,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchangeItem.MasterExchangeItem{
								MasterExchangeItemId: 1,
								MasterExchangeId:     1,
								MasterItemId:         1,
								Name:                 "テストアイテム1",
								Count:                10,
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getMasterExchangeItemModelsAndMasterExchangeCostModels: failed to s.masterExchangeMysqlRepository.Find: test"),
		},
		{
			name: "異常:failed to s.getMasterExchangeItemModelsAndMasterExchangeCostModels: failed to s.getEvent: failed to s.eventService.Get: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "テスト交換",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeItemId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeCost.MasterExchangeCosts{
								{
									MasterExchangeCostId: 1,
									MasterExchangeItemId: 1,
									MasterItemId:         2,
									Name:                 "テスト消費アイテム1",
									Count:                1,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchangeItem.MasterExchangeItem{
								MasterExchangeItemId: 1,
								MasterExchangeId:     1,
								MasterItemId:         1,
								Name:                 "テストアイテム1",
								Count:                10,
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getMasterExchangeItemModelsAndMasterExchangeCostModels: failed to s.getEvent: failed to s.eventService.Get: test"),
		},
		{
			name: "異常:failed to s.getUserExchangeModelAndUserExchangeItemModel: failed to s.userExchangeItemMysqlRepository.Find: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "テスト交換",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeItemId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeCost.MasterExchangeCosts{
								{
									MasterExchangeCostId: 1,
									MasterExchangeItemId: 1,
									MasterItemId:         2,
									Name:                 "テスト消費アイテム1",
									Count:                1,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchangeItem.MasterExchangeItem{
								MasterExchangeItemId: 1,
								MasterExchangeId:     1,
								MasterItemId:         1,
								Name:                 "テストアイテム1",
								Count:                10,
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
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
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getUserExchangeModelAndUserExchangeItemModel: failed to s.userExchangeItemMysqlRepository.Find: test"),
		},
		{
			name: "異常:failed to s.getUserExchangeModelAndUserExchangeItemModel: over limit",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "テスト交換",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeItemId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeCost.MasterExchangeCosts{
								{
									MasterExchangeCostId: 1,
									MasterExchangeItemId: 1,
									MasterItemId:         2,
									Name:                 "テスト消費アイテム1",
									Count:                1,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchangeItem.MasterExchangeItem{
								MasterExchangeItemId: 1,
								MasterExchangeId:     1,
								MasterItemId:         1,
								Name:                 "テストアイテム1",
								Count:                10,
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchangeItem.UserExchangeItem{
								UserId:               "0:test",
								MasterExchangeId:     1,
								MasterExchangeItemId: 1,
								Count:                0,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getUserExchangeModelAndUserExchangeItemModel: over limit"),
		},
		{
			name: "異常:failed to s.getUserExchangeModelAndUserExchangeItemModel: failed to s.userExchangeMysqlRepository.Find: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "テスト交換",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeItemId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeCost.MasterExchangeCosts{
								{
									MasterExchangeCostId: 1,
									MasterExchangeItemId: 1,
									MasterItemId:         2,
									Name:                 "テスト消費アイテム1",
									Count:                1,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchangeItem.MasterExchangeItem{
								MasterExchangeItemId: 1,
								MasterExchangeId:     1,
								MasterItemId:         1,
								Name:                 "テストアイテム1",
								Count:                10,
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchangeItem.UserExchangeItem{
								UserId:               "0:test",
								MasterExchangeId:     1,
								MasterExchangeItemId: 1,
								Count:                10,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getUserExchangeModelAndUserExchangeItemModel: failed to s.userExchangeMysqlRepository.Find: test"),
		},
		{
			name: "異常:failed to s.consume: failed to s.itemService.Consume: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Consume(
							nil,
							nil,
							time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
							&item.ItemConsumeRequest{
								UserId: "0:test",
								Items: item.Items{
									{
										MasterItemId: 2,
										Count:        1,
									},
								},
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "テスト交換",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeItemId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeCost.MasterExchangeCosts{
								{
									MasterExchangeCostId: 1,
									MasterExchangeItemId: 1,
									MasterItemId:         2,
									Name:                 "テスト消費アイテム1",
									Count:                1,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchangeItem.MasterExchangeItem{
								MasterExchangeItemId: 1,
								MasterExchangeId:     1,
								MasterItemId:         1,
								Name:                 "テストアイテム1",
								Count:                10,
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchangeItem.UserExchangeItem{
								UserId:               "0:test",
								MasterExchangeId:     1,
								MasterExchangeItemId: 1,
								Count:                10,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.consume: failed to s.itemService.Consume: test"),
		},
		{
			name: "異常:failed to s.receive: failed to s.itemService.Receive: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Consume(
							nil,
							nil,
							time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
							&item.ItemConsumeRequest{
								UserId: "0:test",
								Items: item.Items{
									{
										MasterItemId: 2,
										Count:        1,
									},
								},
							},
						).
						Return(
							&item.ItemConsumeResponse{
								UserItemBoxes: userItemBox.UserItemBoxes{
									{
										UserId:       "0:test",
										MasterItemId: 2,
										Count:        1,
									},
								},
							},
							nil,
						)
					m.EXPECT().
						Receive(
							nil,
							nil,
							time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
							&item.ItemReceiveRequest{
								UserId: "0:test",
								Items: item.Items{
									{
										MasterItemId: 1,
										Count:        1,
									},
								},
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "テスト交換",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeItemId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeCost.MasterExchangeCosts{
								{
									MasterExchangeCostId: 1,
									MasterExchangeItemId: 1,
									MasterItemId:         2,
									Name:                 "テスト消費アイテム1",
									Count:                1,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchangeItem.MasterExchangeItem{
								MasterExchangeItemId: 1,
								MasterExchangeId:     1,
								MasterItemId:         1,
								Name:                 "テストアイテム1",
								Count:                10,
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchangeItem.UserExchangeItem{
								UserId:               "0:test",
								MasterExchangeId:     1,
								MasterExchangeItemId: 1,
								Count:                10,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.receive: failed to s.itemService.Receive: test"),
		},
		{
			name: "異常:failed to s.userExchangeItemMysqlRepository.Update: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Consume(
							nil,
							nil,
							time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
							&item.ItemConsumeRequest{
								UserId: "0:test",
								Items: item.Items{
									{
										MasterItemId: 2,
										Count:        1,
									},
								},
							},
						).
						Return(
							&item.ItemConsumeResponse{
								UserItemBoxes: userItemBox.UserItemBoxes{
									{
										UserId:       "0:test",
										MasterItemId: 2,
										Count:        1,
									},
								},
							},
							nil,
						)
					m.EXPECT().
						Receive(
							nil,
							nil,
							time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
							&item.ItemReceiveRequest{
								UserId: "0:test",
								Items: item.Items{
									{
										MasterItemId: 1,
										Count:        1,
									},
								},
							},
						).
						Return(
							&item.ItemReceiveResponse{
								UserItemBoxes: userItemBox.UserItemBoxes{
									{
										UserId:       "0:test",
										MasterItemId: 1,
										Count:        1,
									},
								},
							},
							nil,
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					m.EXPECT().
						Get(
							nil,
							&event.EventGetRequest{
								MasterEventId: 1,
							},
						).
						Return(
							&event.EventGetResponse{
								MasterEvent: &masterEvent.MasterEvent{
									MasterEventId: 1,
									Name:          "交換イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeMysqlRepository: func(ctrl *gomock.Controller) masterExchange.MasterExchangeMysqlRepository {
					m := masterExchange.NewMockMasterExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchange.MasterExchange{
								MasterExchangeId: 1,
								MasterEventId:    1,
								Name:             "テスト交換",
							},
							nil,
						)
					return m
				},
				masterExchangeCostMysqlRepository: func(ctrl *gomock.Controller) masterExchangeCost.MasterExchangeCostMysqlRepository {
					m := masterExchangeCost.NewMockMasterExchangeCostMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterExchangeItemId(
							nil,
							int64(1),
						).
						Return(
							masterExchangeCost.MasterExchangeCosts{
								{
									MasterExchangeCostId: 1,
									MasterExchangeItemId: 1,
									MasterItemId:         2,
									Name:                 "テスト消費アイテム1",
									Count:                1,
								},
							},
							nil,
						)
					return m
				},
				masterExchangeItemMysqlRepository: func(ctrl *gomock.Controller) masterExchangeItem.MasterExchangeItemMysqlRepository {
					m := masterExchangeItem.NewMockMasterExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterExchangeItem.MasterExchangeItem{
								MasterExchangeItemId: 1,
								MasterExchangeId:     1,
								MasterItemId:         1,
								Name:                 "テストアイテム1",
								Count:                10,
							},
							nil,
						)
					return m
				},
				userExchangeMysqlRepository: func(ctrl *gomock.Controller) userExchange.UserExchangeMysqlRepository {
					m := userExchange.NewMockUserExchangeMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchange.UserExchange{
								UserId:           "0:test",
								MasterExchangeId: 1,
								ResetAt:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				userExchangeItemMysqlRepository: func(ctrl *gomock.Controller) userExchangeItem.UserExchangeItemMysqlRepository {
					m := userExchangeItem.NewMockUserExchangeItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userExchangeItem.UserExchangeItem{
								UserId:               "0:test",
								MasterExchangeId:     1,
								MasterExchangeItemId: 1,
								Count:                10,
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userExchangeItem.UserExchangeItem{
								UserId:               "0:test",
								MasterExchangeId:     1,
								MasterExchangeItemId: 1,
								Count:                9,
							},
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
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &ExchangeReceiveRequest{
					UserId:               "0:test",
					MasterExchangeItemId: 1,
					Count:                1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.userExchangeItemMysqlRepository.Update: test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &exchangeService{
				itemService:                       tt.fields.itemService(ctrl),
				eventService:                      tt.fields.eventService(ctrl),
				masterExchangeMysqlRepository:     tt.fields.masterExchangeMysqlRepository(ctrl),
				masterExchangeCostMysqlRepository: tt.fields.masterExchangeCostMysqlRepository(ctrl),
				masterExchangeItemMysqlRepository: tt.fields.masterExchangeItemMysqlRepository(ctrl),
				userExchangeMysqlRepository:       tt.fields.userExchangeMysqlRepository(ctrl),
				userExchangeItemMysqlRepository:   tt.fields.userExchangeItemMysqlRepository(ctrl),
			}

			got, err := s.Receive(tt.args.ctx, tt.args.tx, tt.args.now, tt.args.req)
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
