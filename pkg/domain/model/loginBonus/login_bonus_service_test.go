package loginBonus

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
	"github.com/game-core/gc-server/pkg/domain/model/item"
	"github.com/game-core/gc-server/pkg/domain/model/item/userItemBox"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/masterLoginBonus"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/masterLoginBonusItem"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/masterLoginBonusSchedule"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/userLoginBonus"
)

func TestNewItemService_NewItemService(t *testing.T) {
	type args struct {
		itemService                             item.ItemService
		eventService                            event.EventService
		userLoginBonusMysqlRepository           userLoginBonus.UserLoginBonusMysqlRepository
		masterLoginBonusMysqlRepository         masterLoginBonus.MasterLoginBonusMysqlRepository
		masterLoginBonusItemMysqlRepository     masterLoginBonusItem.MasterLoginBonusItemMysqlRepository
		masterLoginBonusScheduleMysqlRepository masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want LoginBonusService
	}{
		{
			name: "正常",
			args: args{
				itemService:                             nil,
				eventService:                            nil,
				userLoginBonusMysqlRepository:           nil,
				masterLoginBonusMysqlRepository:         nil,
				masterLoginBonusItemMysqlRepository:     nil,
				masterLoginBonusScheduleMysqlRepository: nil,
			},
			want: &loginBonusService{
				itemService:                             nil,
				eventService:                            nil,
				userLoginBonusMysqlRepository:           nil,
				masterLoginBonusMysqlRepository:         nil,
				masterLoginBonusItemMysqlRepository:     nil,
				masterLoginBonusScheduleMysqlRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLoginBonusService(
				tt.args.itemService,
				tt.args.eventService,
				tt.args.userLoginBonusMysqlRepository,
				tt.args.masterLoginBonusMysqlRepository,
				tt.args.masterLoginBonusItemMysqlRepository,
				tt.args.masterLoginBonusScheduleMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginBonusService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewItemService_Receive(t *testing.T) {
	type fields struct {
		itemService                             func(ctrl *gomock.Controller) item.ItemService
		eventService                            func(ctrl *gomock.Controller) event.EventService
		userLoginBonusMysqlRepository           func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository
		masterLoginBonusMysqlRepository         func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository
		masterLoginBonusItemMysqlRepository     func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository
		masterLoginBonusScheduleMysqlRepository func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		now time.Time
		req *LoginBonusReceiveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *LoginBonusReceiveResponse
		wantErr error
	}{
		{
			name: "正常：受け取りできる(初回)",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Receive(
							nil,
							nil,
							&item.ItemReceiveRequest{
								UserId: "0:WntR-PyhOJeDiE5jodeR",
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
										UserId:       "0:WntR-PyhOJeDiE5jodeR",
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
									Name:          "イベント",
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
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								MasterLoginBonusId: 1,
								MasterEventId:      1,
								Name:               "ログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									MasterLoginBonusScheduleId: 1,
									MasterLoginBonusId:         1,
									Step:                       0,
									Name:                       "ステップ0",
								},
								{
									MasterLoginBonusScheduleId: 2,
									MasterLoginBonusId:         1,
									Step:                       1,
									Name:                       "ステップ1",
								},
								{
									MasterLoginBonusScheduleId: 3,
									MasterLoginBonusId:         1,
									Step:                       2,
									Name:                       "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									MasterLoginBonusItemId:     1,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "ログインボーナスアイテム2",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
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
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
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
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want: &LoginBonusReceiveResponse{
				UserLoginBonus: &userLoginBonus.UserLoginBonus{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
					ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：受け取りできる(２回目以降受け取り)",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Receive(
							nil,
							nil,
							&item.ItemReceiveRequest{
								UserId: "0:WntR-PyhOJeDiE5jodeR",
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
										UserId:       "0:WntR-PyhOJeDiE5jodeR",
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
									Name:          "イベント",
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
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								MasterLoginBonusId: 1,
								MasterEventId:      1,
								Name:               "ログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									MasterLoginBonusScheduleId: 1,
									MasterLoginBonusId:         1,
									Step:                       0,
									Name:                       "ステップ0",
								},
								{
									MasterLoginBonusScheduleId: 2,
									MasterLoginBonusId:         1,
									Step:                       1,
									Name:                       "ステップ1",
								},
								{
									MasterLoginBonusScheduleId: 3,
									MasterLoginBonusId:         1,
									Step:                       2,
									Name:                       "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									MasterLoginBonusItemId:     1,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "ログインボーナスアイテム2",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
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
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want: &LoginBonusReceiveResponse{
				UserLoginBonus: &userLoginBonus.UserLoginBonus{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
					ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：failed to s.masterLoginBonusMysqlRepository.Find: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) event.EventService {
					m := event.NewMockEventService(ctrl)
					return m
				},
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.masterLoginBonusMysqlRepository.Find: test"),
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
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								MasterLoginBonusId: 1,
								MasterEventId:      1,
								Name:               "ログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getEvent: failed to s.eventService.Get: test"),
		},
		{
			name: "異常：failed to s.getEvent: outside the event period",
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
									Name:          "イベント",
									ResetHour:     9,
									IntervalHour:  24,
									RepeatSetting: true,
									StartAt:       time.Date(2023, 1, 10, 9, 0, 0, 0, time.UTC),
									EndAt:         nil,
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								MasterLoginBonusId: 1,
								MasterEventId:      1,
								Name:               "ログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getEvent: outside the event period"),
		},
		{
			name: "異常：failed to s.getSchedule: failed to s.masterLoginBonusScheduleMysqlRepository.FindListByMasterLoginBonusId: test",
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
									Name:          "イベント",
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
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								MasterLoginBonusId: 1,
								MasterEventId:      1,
								Name:               "ログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getSchedule: failed to s.masterLoginBonusScheduleMysqlRepository.FindListByMasterLoginBonusId: test"),
		},
		{
			name: "異常：failed to s.masterLoginBonusItemMysqlRepository.FindListByMasterLoginBonusScheduleId: test",
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
									Name:          "イベント",
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
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								MasterLoginBonusId: 1,
								MasterEventId:      1,
								Name:               "ログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									MasterLoginBonusScheduleId: 1,
									MasterLoginBonusId:         1,
									Step:                       0,
									Name:                       "ステップ0",
								},
								{
									MasterLoginBonusScheduleId: 2,
									MasterLoginBonusId:         1,
									Step:                       1,
									Name:                       "ステップ1",
								},
								{
									MasterLoginBonusScheduleId: 3,
									MasterLoginBonusId:         1,
									Step:                       2,
									Name:                       "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.masterLoginBonusItemMysqlRepository.FindListByMasterLoginBonusScheduleId: test"),
		},
		{
			name: "異常：failed to s.getUser: failed to s.userLoginBonusMysqlRepository.FindOrNil: test",
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
									Name:          "イベント",
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
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								MasterLoginBonusId: 1,
								MasterEventId:      1,
								Name:               "ログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									MasterLoginBonusScheduleId: 1,
									MasterLoginBonusId:         1,
									Step:                       0,
									Name:                       "ステップ0",
								},
								{
									MasterLoginBonusScheduleId: 2,
									MasterLoginBonusId:         1,
									Step:                       1,
									Name:                       "ステップ1",
								},
								{
									MasterLoginBonusScheduleId: 3,
									MasterLoginBonusId:         1,
									Step:                       2,
									Name:                       "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									MasterLoginBonusItemId:     1,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "ログインボーナスアイテム2",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
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
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getUser: failed to s.userLoginBonusMysqlRepository.FindOrNil: test"),
		},
		{
			name: "異常：failed to s.getUser: already received",
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
									Name:          "イベント",
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
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								MasterLoginBonusId: 1,
								MasterEventId:      1,
								Name:               "ログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									MasterLoginBonusScheduleId: 1,
									MasterLoginBonusId:         1,
									Step:                       0,
									Name:                       "ステップ0",
								},
								{
									MasterLoginBonusScheduleId: 2,
									MasterLoginBonusId:         1,
									Step:                       1,
									Name:                       "ステップ1",
								},
								{
									MasterLoginBonusScheduleId: 3,
									MasterLoginBonusId:         1,
									Step:                       2,
									Name:                       "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									MasterLoginBonusItemId:     1,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "ログインボーナスアイテム2",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
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
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getUser: already received"),
		},
		{
			name: "異常：failed to s.receive: failed to s.itemService.Receive: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Receive(
							nil,
							nil,
							&item.ItemReceiveRequest{
								UserId: "0:WntR-PyhOJeDiE5jodeR",
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
									Name:          "イベント",
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
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								MasterLoginBonusId: 1,
								MasterEventId:      1,
								Name:               "ログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									MasterLoginBonusScheduleId: 1,
									MasterLoginBonusId:         1,
									Step:                       0,
									Name:                       "ステップ0",
								},
								{
									MasterLoginBonusScheduleId: 2,
									MasterLoginBonusId:         1,
									Step:                       1,
									Name:                       "ステップ1",
								},
								{
									MasterLoginBonusScheduleId: 3,
									MasterLoginBonusId:         1,
									Step:                       2,
									Name:                       "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									MasterLoginBonusItemId:     1,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "ログインボーナスアイテム2",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							nil,
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.receive: failed to s.itemService.Receive: test"),
		},
		{
			name: "異常：failed to s.update: failed to s.userLoginBonusMysqlRepository.Create: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Receive(
							nil,
							nil,
							&item.ItemReceiveRequest{
								UserId: "0:WntR-PyhOJeDiE5jodeR",
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
										UserId:       "0:WntR-PyhOJeDiE5jodeR",
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
									Name:          "イベント",
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
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								MasterLoginBonusId: 1,
								MasterEventId:      1,
								Name:               "ログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									MasterLoginBonusScheduleId: 1,
									MasterLoginBonusId:         1,
									Step:                       0,
									Name:                       "ステップ0",
								},
								{
									MasterLoginBonusScheduleId: 2,
									MasterLoginBonusId:         1,
									Step:                       1,
									Name:                       "ステップ1",
								},
								{
									MasterLoginBonusScheduleId: 3,
									MasterLoginBonusId:         1,
									Step:                       2,
									Name:                       "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									MasterLoginBonusItemId:     1,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "ログインボーナスアイテム2",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
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
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
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
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.update: failed to s.userLoginBonusMysqlRepository.Create: test"),
		},
		{
			name: "異常：failed to s.update: failed to s.userLoginBonusMysqlRepository.Update: test",
			fields: fields{
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Receive(
							nil,
							nil,
							&item.ItemReceiveRequest{
								UserId: "0:WntR-PyhOJeDiE5jodeR",
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
										UserId:       "0:WntR-PyhOJeDiE5jodeR",
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
									Name:          "イベント",
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
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								MasterLoginBonusId: 1,
								MasterEventId:      1,
								Name:               "ログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									MasterLoginBonusScheduleId: 1,
									MasterLoginBonusId:         1,
									Step:                       0,
									Name:                       "ステップ0",
								},
								{
									MasterLoginBonusScheduleId: 2,
									MasterLoginBonusId:         1,
									Step:                       1,
									Name:                       "ステップ1",
								},
								{
									MasterLoginBonusScheduleId: 3,
									MasterLoginBonusId:         1,
									Step:                       2,
									Name:                       "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									MasterLoginBonusItemId:     1,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "ログインボーナスアイテム2",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
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
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.update: failed to s.userLoginBonusMysqlRepository.Update: test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &loginBonusService{
				itemService:                             tt.fields.itemService(ctrl),
				eventService:                            tt.fields.eventService(ctrl),
				userLoginBonusMysqlRepository:           tt.fields.userLoginBonusMysqlRepository(ctrl),
				masterLoginBonusMysqlRepository:         tt.fields.masterLoginBonusMysqlRepository(ctrl),
				masterLoginBonusItemMysqlRepository:     tt.fields.masterLoginBonusItemMysqlRepository(ctrl),
				masterLoginBonusScheduleMysqlRepository: tt.fields.masterLoginBonusScheduleMysqlRepository(ctrl),
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
