package health

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/health/commonHealth"
	"github.com/game-core/gc-server/pkg/domain/model/health/masterHealth"
)

func TestNewHealthService_NewHealthService(t *testing.T) {
	type args struct {
		commonHealthMysqlRepository commonHealth.CommonHealthMysqlRepository
		masterHealthMysqlRepository masterHealth.MasterHealthMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want HealthService
	}{
		{
			name: "正常",
			args: args{
				commonHealthMysqlRepository: nil,
				masterHealthMysqlRepository: nil,
			},
			want: &healthService{
				commonHealthMysqlRepository: nil,
				masterHealthMysqlRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHealthService(
				tt.args.commonHealthMysqlRepository,
				tt.args.masterHealthMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHealthService_Check(t *testing.T) {
	type fields struct {
		commonHealthMysqlRepository func(ctrl *gomock.Controller) commonHealth.CommonHealthMysqlRepository
		masterHealthMysqlRepository func(ctrl *gomock.Controller) masterHealth.MasterHealthMysqlRepository
	}
	type args struct {
		ctx context.Context
		req *HealthCheckRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *HealthCheckResponse
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				commonHealthMysqlRepository: func(ctrl *gomock.Controller) commonHealth.CommonHealthMysqlRepository {
					m := commonHealth.NewMockCommonHealthMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&commonHealth.CommonHealth{
								HealthId:         1,
								Name:             "test",
								CommonHealthType: commonHealth.CommonHealthType_CommonSuccess,
							},
							nil,
						)
					return m
				},
				masterHealthMysqlRepository: func(ctrl *gomock.Controller) masterHealth.MasterHealthMysqlRepository {
					m := masterHealth.NewMockMasterHealthMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterHealth.MasterHealth{
								HealthId:         1,
								Name:             "test",
								MasterHealthType: masterHealth.MasterHealthType_MasterSuccess,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &HealthCheckRequest{
					HealthId:         1,
					Name:             "test",
					CommonHealthType: commonHealth.CommonHealthType_CommonSuccess,
					MasterHealthType: masterHealth.MasterHealthType_MasterSuccess,
				},
			},
			want: &HealthCheckResponse{
				CommonHealth: &commonHealth.CommonHealth{
					HealthId:         1,
					Name:             "test",
					CommonHealthType: commonHealth.CommonHealthType_CommonSuccess,
				},
				MasterHealth: &masterHealth.MasterHealth{
					HealthId:         1,
					Name:             "test",
					MasterHealthType: masterHealth.MasterHealthType_MasterSuccess,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常： failed to s.commonHealthMysqlRepository.Find: test",
			fields: fields{
				commonHealthMysqlRepository: func(ctrl *gomock.Controller) commonHealth.CommonHealthMysqlRepository {
					m := commonHealth.NewMockCommonHealthMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewError("test"),
						)
					return m
				},
				masterHealthMysqlRepository: func(ctrl *gomock.Controller) masterHealth.MasterHealthMysqlRepository {
					m := masterHealth.NewMockMasterHealthMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &HealthCheckRequest{
					HealthId:         1,
					Name:             "test",
					CommonHealthType: commonHealth.CommonHealthType_CommonSuccess,
					MasterHealthType: masterHealth.MasterHealthType_MasterSuccess,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.commonHealthMysqlRepository.Find: test"),
		},
		{
			name: "異常： failed to s.masterHealthMysqlRepository.Find: test",
			fields: fields{
				commonHealthMysqlRepository: func(ctrl *gomock.Controller) commonHealth.CommonHealthMysqlRepository {
					m := commonHealth.NewMockCommonHealthMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&commonHealth.CommonHealth{
								HealthId:         1,
								Name:             "test",
								CommonHealthType: commonHealth.CommonHealthType_CommonSuccess,
							},
							nil,
						)
					return m
				},
				masterHealthMysqlRepository: func(ctrl *gomock.Controller) masterHealth.MasterHealthMysqlRepository {
					m := masterHealth.NewMockMasterHealthMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewError("test"),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &HealthCheckRequest{
					HealthId:         1,
					Name:             "test",
					CommonHealthType: commonHealth.CommonHealthType_CommonSuccess,
					MasterHealthType: masterHealth.MasterHealthType_MasterSuccess,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.masterHealthMysqlRepository.Find: test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &healthService{
				commonHealthMysqlRepository: tt.fields.commonHealthMysqlRepository(ctrl),
				masterHealthMysqlRepository: tt.fields.masterHealthMysqlRepository(ctrl),
			}

			got, err := s.Check(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
