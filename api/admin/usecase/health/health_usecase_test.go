package health

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	healthProto "github.com/game-core/gc-server/api/admin/presentation/proto/health"
	"github.com/game-core/gc-server/api/admin/presentation/proto/health/adminHealth"
	"github.com/game-core/gc-server/api/admin/presentation/proto/health/commonHealth"
	"github.com/game-core/gc-server/api/admin/presentation/proto/health/masterHealth"
	"github.com/game-core/gc-server/internal/errors"
	healthService "github.com/game-core/gc-server/pkg/domain/model/health"
	adminHealthModel "github.com/game-core/gc-server/pkg/domain/model/health/adminHealth"
	commonHealthModel "github.com/game-core/gc-server/pkg/domain/model/health/commonHealth"
	masterHealthModel "github.com/game-core/gc-server/pkg/domain/model/health/masterHealth"
)

func TestHealthUsecase_NewHealthUsecase(t *testing.T) {
	type args struct {
		healthService healthService.HealthService
	}
	tests := []struct {
		name string
		args args
		want HealthUsecase
	}{
		{
			name: "正常",
			args: args{
				healthService: nil,
			},
			want: &healthUsecase{
				healthService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHealthUsecase(tt.args.healthService)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHealthUsecase_Check(t *testing.T) {
	type fields struct {
		healthService func(ctrl *gomock.Controller) healthService.HealthService
	}
	type args struct {
		ctx context.Context
		req *healthProto.HealthCheckRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *healthProto.HealthCheckResponse
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				healthService: func(ctrl *gomock.Controller) healthService.HealthService {
					m := healthService.NewMockHealthService(ctrl)
					m.EXPECT().
						Check(
							gomock.Any(),
							&healthService.HealthCheckRequest{
								HealthId: 1,
							},
						).
						Return(
							&healthService.HealthCheckResponse{
								AdminHealth: &adminHealthModel.AdminHealth{
									HealthId:        1,
									Name:            "health",
									AdminHealthEnum: adminHealthModel.AdminHealthEnum_AdminSuccess,
								},
								CommonHealth: &commonHealthModel.CommonHealth{
									HealthId:         1,
									Name:             "health",
									CommonHealthEnum: commonHealthModel.CommonHealthEnum_CommonSuccess,
								},
								MasterHealth: &masterHealthModel.MasterHealth{
									HealthId:         1,
									Name:             "health",
									MasterHealthEnum: masterHealthModel.MasterHealthEnum_MasterSuccess,
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &healthProto.HealthCheckRequest{
					HealthId: 1,
				},
			},
			want: &healthProto.HealthCheckResponse{
				AdminHealth: &adminHealth.AdminHealth{
					HealthId:        1,
					Name:            "health",
					AdminHealthEnum: adminHealth.AdminHealthEnum_AdminSuccess,
				},
				CommonHealth: &commonHealth.CommonHealth{
					HealthId:         1,
					Name:             "health",
					CommonHealthEnum: commonHealth.CommonHealthEnum_CommonSuccess,
				},
				MasterHealth: &masterHealth.MasterHealth{
					HealthId:         1,
					Name:             "health",
					MasterHealthEnum: masterHealth.MasterHealthEnum_MasterSuccess,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常： failed to s.healthService.Check: test",
			fields: fields{
				healthService: func(ctrl *gomock.Controller) healthService.HealthService {
					m := healthService.NewMockHealthService(ctrl)
					m.EXPECT().
						Check(
							gomock.Any(),
							&healthService.HealthCheckRequest{
								HealthId: 1,
							},
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
				req: &healthProto.HealthCheckRequest{
					HealthId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.healthService.Check: test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &healthUsecase{
				healthService: tt.fields.healthService(ctrl),
			}

			got, err := u.Check(tt.args.ctx, tt.args.req)
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
