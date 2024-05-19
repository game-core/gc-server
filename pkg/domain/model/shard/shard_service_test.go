package shard

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/shard/masterShard"
)

func TestShardService_NewShardService(t *testing.T) {
	type args struct {
		masterShardMysqlRepository masterShard.MasterShardMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want ShardService
	}{
		{
			name: "正常",
			args: args{
				masterShardMysqlRepository: nil,
			},
			want: &shardService{
				masterShardMysqlRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewShardService(
				tt.args.masterShardMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShardService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShardService_GetShardKey(t *testing.T) {
	type fields struct {
		masterShardMysqlRepository func(ctrl *gomock.Controller) masterShard.MasterShardMysqlRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{
			name: "正常：取得できる（Countが同じ場合）",
			fields: fields{
				masterShardMysqlRepository: func(ctrl *gomock.Controller) masterShard.MasterShardMysqlRepository {
					m := masterShard.NewMockMasterShardMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterShard.MasterShards{
								{
									MasterShardId: 1,
									ShardKey:      "SHARD_1",
									Name:          "name1",
									Count:         1,
								},
								{
									MasterShardId: 2,
									ShardKey:      "SHARD_2",
									Name:          "name2",
									Count:         1,
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    "SHARD_1",
			wantErr: nil,
		},
		{
			name: "正常：取得できる（Countが違う場合）",
			fields: fields{
				masterShardMysqlRepository: func(ctrl *gomock.Controller) masterShard.MasterShardMysqlRepository {
					m := masterShard.NewMockMasterShardMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterShard.MasterShards{
								{
									MasterShardId: 1,
									ShardKey:      "SHARD_1",
									Name:          "name1",
									Count:         2,
								},
								{
									MasterShardId: 2,
									ShardKey:      "SHARD_2",
									Name:          "name2",
									Count:         1,
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    "SHARD_2",
			wantErr: nil,
		},
		{
			name: "異常：s.masterShardMysqlRepository.FindList",
			fields: fields{
				masterShardMysqlRepository: func(ctrl *gomock.Controller) masterShard.MasterShardMysqlRepository {
					m := masterShard.NewMockMasterShardMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
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
			},
			want:    "",
			wantErr: errors.NewMethodError("shards.GetShardKey: failed to s.masterShardMysqlRepository.FindList", errors.NewTestError()),
		},
		{
			name: "異常：common_shard does not exist",
			fields: fields{
				masterShardMysqlRepository: func(ctrl *gomock.Controller) masterShard.MasterShardMysqlRepository {
					m := masterShard.NewMockMasterShardMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterShard.MasterShards{},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    "",
			wantErr: errors.NewError("failed to shards.GetShardKey: common_shard does not exist"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			masterShard.MasterShardInstances = masterShard.NewMasterShards()
			ctrl := gomock.NewController(t)

			s := &shardService{
				masterShardMysqlRepository: tt.fields.masterShardMysqlRepository(ctrl),
			}

			got, err := s.GetShardKey(tt.args.ctx)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetShardKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShardKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
