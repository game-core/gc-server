package item

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/item/masterItem"
	"github.com/game-core/gc-server/pkg/domain/model/item/userItemBox"
	"github.com/game-core/gc-server/pkg/domain/model/rarity/masterRarity"
	"github.com/game-core/gc-server/pkg/domain/model/resource/masterResource"
)

func TestNewItemService_NewItemService(t *testing.T) {
	type args struct {
		userItemBoxMysqlRepository userItemBox.UserItemBoxMysqlRepository
		masterItemMysqlRepository  masterItem.MasterItemMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want ItemService
	}{
		{
			name: "正常",
			args: args{
				userItemBoxMysqlRepository: nil,
				masterItemMysqlRepository:  nil,
			},
			want: &itemService{
				userItemBoxMysqlRepository: nil,
				masterItemMysqlRepository:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewItemService(
				tt.args.userItemBoxMysqlRepository,
				tt.args.masterItemMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemService_Receive(t *testing.T) {
	type fields struct {
		userItemBoxMysqlRepository func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository
		masterItemMysqlRepository  func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		req *ItemReceiveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ItemReceiveResponse
		wantErr error
	}{
		{
			name: "正常：受け取りできる",
			fields: fields{
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								MasterItemId:       1,
								Name:               "アイテム1",
								MasterResourceEnum: masterResource.MasterResourceEnum_Normal,
								MasterRarityEnum:   masterRarity.MasterRarityEnum_N,
								Content:            "ノーマルアイテム1",
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(4),
						).
						Return(
							&masterItem.MasterItem{
								MasterItemId:       4,
								Name:               "アイテム4",
								MasterResourceEnum: masterResource.MasterResourceEnum_Normal,
								MasterRarityEnum:   masterRarity.MasterRarityEnum_N,
								Content:            "ノーマルアイテム4",
							},
							nil,
						)
					return m
				},
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
							"0:testUserId",
						).
						Return(
							userItemBox.UserItemBoxes{
								{
									UserId:       "0:testUserId",
									MasterItemId: 1,
									Count:        1,
								},
								{
									UserId:       "0:testUserId",
									MasterItemId: 2,
									Count:        1,
								},
								{
									UserId:       "0:testUserId",
									MasterItemId: 3,
									Count:        1,
								},
							},
							nil,
						)
					m.EXPECT().
						UpdateList(
							nil,
							nil,
							userItemBox.UserItemBoxes{
								{
									UserId:       "0:testUserId",
									MasterItemId: 1,
									Count:        11,
								},
								{
									UserId:       "0:testUserId",
									MasterItemId: 4,
									Count:        10,
								},
							},
						).
						Return(
							userItemBox.UserItemBoxes{
								{
									UserId:       "0:testUserId",
									MasterItemId: 1,
									Count:        11,
								},
								{
									UserId:       "0:testUserId",
									MasterItemId: 4,
									Count:        10,
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
				req: &ItemReceiveRequest{
					UserId: "0:testUserId",
					Items: Items{
						{
							MasterItemId: 1,
							Count:        10,
						},
						{
							MasterItemId: 4,
							Count:        10,
						},
					},
				},
			},
			want: &ItemReceiveResponse{
				UserItemBoxes: userItemBox.UserItemBoxes{
					{
						UserId:       "0:testUserId",
						MasterItemId: 1,
						Count:        11,
					},
					{
						UserId:       "0:testUserId",
						MasterItemId: 4,
						Count:        10,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：failed to s.checkItems: failed to s.masterItemMysqlRepository.Find: test",
			fields: fields{
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								MasterItemId:       1,
								Name:               "アイテム1",
								MasterResourceEnum: masterResource.MasterResourceEnum_Normal,
								MasterRarityEnum:   masterRarity.MasterRarityEnum_N,
								Content:            "ノーマルアイテム1",
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(4),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ItemReceiveRequest{
					UserId: "0:testUserId",
					Items: Items{
						{
							MasterItemId: 1,
							Count:        10,
						},
						{
							MasterItemId: 4,
							Count:        10,
						},
					},
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.checkItems: failed to s.masterItemMysqlRepository.Find: test"),
		},
		{
			name: "異常：failed to s.userItemMysqlRepository.FindList: test",
			fields: fields{
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								MasterItemId:       1,
								Name:               "アイテム1",
								MasterResourceEnum: masterResource.MasterResourceEnum_Normal,
								MasterRarityEnum:   masterRarity.MasterRarityEnum_N,
								Content:            "ノーマルアイテム1",
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(4),
						).
						Return(
							&masterItem.MasterItem{
								MasterItemId:       4,
								Name:               "アイテム4",
								MasterResourceEnum: masterResource.MasterResourceEnum_Normal,
								MasterRarityEnum:   masterRarity.MasterRarityEnum_N,
								Content:            "ノーマルアイテム4",
							},
							nil,
						)
					return m
				},
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
							"0:testUserId",
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
				req: &ItemReceiveRequest{
					UserId: "0:testUserId",
					Items: Items{
						{
							MasterItemId: 1,
							Count:        10,
						},
						{
							MasterItemId: 4,
							Count:        10,
						},
					},
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.userItemMysqlRepository.FindList: test"),
		},
		{
			name: "異常：failed to s.userItemBoxMysqlRepository.UpdateList: test",
			fields: fields{
				masterItemMysqlRepository: func(ctrl *gomock.Controller) masterItem.MasterItemMysqlRepository {
					m := masterItem.NewMockMasterItemMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterItem.MasterItem{
								MasterItemId:       1,
								Name:               "アイテム1",
								MasterResourceEnum: masterResource.MasterResourceEnum_Normal,
								MasterRarityEnum:   masterRarity.MasterRarityEnum_N,
								Content:            "ノーマルアイテム1",
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(4),
						).
						Return(
							&masterItem.MasterItem{
								MasterItemId:       4,
								Name:               "アイテム4",
								MasterResourceEnum: masterResource.MasterResourceEnum_Normal,
								MasterRarityEnum:   masterRarity.MasterRarityEnum_N,
								Content:            "ノーマルアイテム4",
							},
							nil,
						)
					return m
				},
				userItemBoxMysqlRepository: func(ctrl *gomock.Controller) userItemBox.UserItemBoxMysqlRepository {
					m := userItemBox.NewMockUserItemBoxMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
							"0:testUserId",
						).
						Return(
							userItemBox.UserItemBoxes{
								{
									UserId:       "0:testUserId",
									MasterItemId: 1,
									Count:        1,
								},
								{
									UserId:       "0:testUserId",
									MasterItemId: 2,
									Count:        1,
								},
								{
									UserId:       "0:testUserId",
									MasterItemId: 3,
									Count:        1,
								},
							},
							nil,
						)
					m.EXPECT().
						UpdateList(
							nil,
							nil,
							userItemBox.UserItemBoxes{
								{
									UserId:       "0:testUserId",
									MasterItemId: 1,
									Count:        11,
								},
								{
									UserId:       "0:testUserId",
									MasterItemId: 4,
									Count:        10,
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
				req: &ItemReceiveRequest{
					UserId: "0:testUserId",
					Items: Items{
						{
							MasterItemId: 1,
							Count:        10,
						},
						{
							MasterItemId: 4,
							Count:        10,
						},
					},
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.userItemBoxMysqlRepository.UpdateList: test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &itemService{
				userItemBoxMysqlRepository: tt.fields.userItemBoxMysqlRepository(ctrl),
				masterItemMysqlRepository:  tt.fields.masterItemMysqlRepository(ctrl),
			}

			got, err := s.Receive(tt.args.ctx, tt.args.tx, tt.args.req)
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