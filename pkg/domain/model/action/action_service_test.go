package action

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/pointers"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterAction"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionRun"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionStep"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionTrigger"
	"github.com/game-core/gc-server/pkg/domain/model/action/userAction"
)

func TestNewActionService_NewActionService(t *testing.T) {
	type args struct {
		masterActionMysqlRepository        masterAction.MasterActionMysqlRepository
		masterActionRunMysqlRepository     masterActionRun.MasterActionRunMysqlRepository
		masterActionStepMysqlRepository    masterActionStep.MasterActionStepMysqlRepository
		masterActionTriggerMysqlRepository masterActionTrigger.MasterActionTriggerMysqlRepository
		userActionMysqlRepository          userAction.UserActionMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want ActionService
	}{
		{
			name: "正常",
			args: args{
				masterActionMysqlRepository:        nil,
				masterActionRunMysqlRepository:     nil,
				masterActionStepMysqlRepository:    nil,
				masterActionTriggerMysqlRepository: nil,
				userActionMysqlRepository:          nil,
			},
			want: &actionService{
				masterActionMysqlRepository:        nil,
				masterActionRunMysqlRepository:     nil,
				masterActionStepMysqlRepository:    nil,
				masterActionTriggerMysqlRepository: nil,
				userActionMysqlRepository:          nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewActionService(
				tt.args.masterActionMysqlRepository,
				tt.args.masterActionRunMysqlRepository,
				tt.args.masterActionStepMysqlRepository,
				tt.args.masterActionTriggerMysqlRepository,
				tt.args.userActionMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewActionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActionService_GetListMaster(t *testing.T) {
	type fields struct {
		masterActionMysqlRepository        func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository
		masterActionRunMysqlRepository     func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository
		masterActionStepMysqlRepository    func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository
		masterActionTriggerMysqlRepository func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository
		userActionMysqlRepository          func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ActionGetListMasterResponse
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterAction.MasterActions{
								{
									MasterActionId:          1,
									Name:                    "テストアクション1",
									MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
									MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
									TargetId:                nil,
									TriggerMasterActionId:   nil,
									Expiration:              nil,
								},
								{
									MasterActionId:          2,
									Name:                    "テストアクション2",
									MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
									MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
									TargetId:                pointers.Int64ToPointer(2),
									TriggerMasterActionId:   pointers.Int64ToPointer(1),
									Expiration:              pointers.Int32ToPointer(24),
								},
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionRun.MasterActionRuns{
								{
									MasterActionRunId: 1,
									Name:              "テストアクション1",
									MasterActionId:    1,
								},
								{
									MasterActionRunId: 2,
									Name:              "テストアクション2",
									MasterActionId:    2,
								},
							},
							nil,
						)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionStep.MasterActionSteps{
								{
									MasterActionStepId:   1,
									Name:                 "無",
									MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
								},
								{
									MasterActionStepId:   2,
									Name:                 "チュートリアル突破",
									MasterActionStepEnum: masterActionStep.MasterActionStepEnum_PassedTutorial,
								},
							},
							nil,
						)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionTrigger.MasterActionTriggers{
								{
									MasterActionTriggerId:   1,
									Name:                    "無期限",
									MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								},
								{
									MasterActionTriggerId:   2,
									Name:                    "期限あり",
									MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Discontinuation,
								},
							},
							nil,
						)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want: &ActionGetListMasterResponse{
				MasterActions: masterAction.MasterActions{
					{
						MasterActionId:          1,
						Name:                    "テストアクション1",
						MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
						MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
						TargetId:                nil,
						TriggerMasterActionId:   nil,
						Expiration:              nil,
					},
					{
						MasterActionId:          2,
						Name:                    "テストアクション2",
						MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
						MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
						TargetId:                pointers.Int64ToPointer(2),
						TriggerMasterActionId:   pointers.Int64ToPointer(1),
						Expiration:              pointers.Int32ToPointer(24),
					},
				},
				MasterActionRuns: masterActionRun.MasterActionRuns{
					{
						MasterActionRunId: 1,
						Name:              "テストアクション1",
						MasterActionId:    1,
					},
					{
						MasterActionRunId: 2,
						Name:              "テストアクション2",
						MasterActionId:    2,
					},
				},
				MasterActionSteps: masterActionStep.MasterActionSteps{
					{
						MasterActionStepId:   1,
						Name:                 "無",
						MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					},
					{
						MasterActionStepId:   2,
						Name:                 "チュートリアル突破",
						MasterActionStepEnum: masterActionStep.MasterActionStepEnum_PassedTutorial,
					},
				},
				MasterActionTriggers: masterActionTrigger.MasterActionTriggers{
					{
						MasterActionTriggerId:   1,
						Name:                    "無期限",
						MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
					},
					{
						MasterActionTriggerId:   2,
						Name:                    "期限あり",
						MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Discontinuation,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterActionMysqlRepository.FindList",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
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
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterActionMysqlRepository.FindList", errors.NewTestError()),
		},
		{
			name: "異常：s.masterActionRunMysqlRepository.FindList",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterAction.MasterActions{
								{
									MasterActionId:          1,
									Name:                    "テストアクション1",
									MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
									MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
									TargetId:                nil,
									TriggerMasterActionId:   nil,
									Expiration:              nil,
								},
								{
									MasterActionId:          2,
									Name:                    "テストアクション2",
									MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
									MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
									TargetId:                pointers.Int64ToPointer(2),
									TriggerMasterActionId:   pointers.Int64ToPointer(1),
									Expiration:              pointers.Int32ToPointer(24),
								},
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
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
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterActionRunMysqlRepository.FindList", errors.NewTestError()),
		},
		{
			name: "異常：s.masterActionStepMysqlRepository.FindList",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterAction.MasterActions{
								{
									MasterActionId:          1,
									Name:                    "テストアクション1",
									MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
									MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
									TargetId:                nil,
									TriggerMasterActionId:   nil,
									Expiration:              nil,
								},
								{
									MasterActionId:          2,
									Name:                    "テストアクション2",
									MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
									MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
									TargetId:                pointers.Int64ToPointer(2),
									TriggerMasterActionId:   pointers.Int64ToPointer(1),
									Expiration:              pointers.Int32ToPointer(24),
								},
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionRun.MasterActionRuns{
								{
									MasterActionRunId: 1,
									Name:              "テストアクション1",
									MasterActionId:    1,
								},
								{
									MasterActionRunId: 2,
									Name:              "テストアクション2",
									MasterActionId:    2,
								},
							},
							nil,
						)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
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
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterActionStepMysqlRepository.FindList", errors.NewTestError()),
		},
		{
			name: "正常：取得できる",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterAction.MasterActions{
								{
									MasterActionId:          1,
									Name:                    "テストアクション1",
									MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
									MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
									TargetId:                nil,
									TriggerMasterActionId:   nil,
									Expiration:              nil,
								},
								{
									MasterActionId:          2,
									Name:                    "テストアクション2",
									MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
									MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
									TargetId:                pointers.Int64ToPointer(2),
									TriggerMasterActionId:   pointers.Int64ToPointer(1),
									Expiration:              pointers.Int32ToPointer(24),
								},
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionRun.MasterActionRuns{
								{
									MasterActionRunId: 1,
									Name:              "テストアクション1",
									MasterActionId:    1,
								},
								{
									MasterActionRunId: 2,
									Name:              "テストアクション2",
									MasterActionId:    2,
								},
							},
							nil,
						)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionStep.MasterActionSteps{
								{
									MasterActionStepId:   1,
									Name:                 "無",
									MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
								},
								{
									MasterActionStepId:   2,
									Name:                 "チュートリアル突破",
									MasterActionStepEnum: masterActionStep.MasterActionStepEnum_PassedTutorial,
								},
							},
							nil,
						)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
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
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterActionTriggerMysqlRepository.FindList", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &actionService{
				masterActionMysqlRepository:        tt.fields.masterActionMysqlRepository(ctrl),
				masterActionRunMysqlRepository:     tt.fields.masterActionRunMysqlRepository(ctrl),
				masterActionStepMysqlRepository:    tt.fields.masterActionStepMysqlRepository(ctrl),
				masterActionTriggerMysqlRepository: tt.fields.masterActionTriggerMysqlRepository(ctrl),
				userActionMysqlRepository:          tt.fields.userActionMysqlRepository(ctrl),
			}

			got, err := s.GetListMaster(tt.args.ctx)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetListMaster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListMaster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActionService_Check(t *testing.T) {
	type fields struct {
		masterActionMysqlRepository        func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository
		masterActionRunMysqlRepository     func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository
		masterActionStepMysqlRepository    func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository
		masterActionTriggerMysqlRepository func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository
		userActionMysqlRepository          func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository
	}
	type args struct {
		ctx context.Context
		now time.Time
		req *ActionCheckRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "正常：確認できる（TargetIdがある場合）",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(1),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：確認できる（TargetIdがない場合）",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnum(
							nil,
							masterActionStep.MasterActionStepEnum_None,
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                nil,
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             nil,
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：アクションが存在しない（TargetIdがある場合）",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(1),
						).
						Return(
							nil,
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(1),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：アクションが存在しない（TargetIdがない場合）",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnum(
							nil,
							masterActionStep.MasterActionStepEnum_None,
						).
						Return(
							nil,
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             nil,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.getMasterAction: failed to s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnumAndTargetId",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(1),
				},
			},
			wantErr: errors.NewMethodError("s.getMasterAction: failed to s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnumAndTargetId", errors.NewTestError()),
		},
		{
			name: "異常：s.getUserAction: failed to s.userActionMysqlRepository.Find",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
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
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(1),
				},
			},
			wantErr: errors.NewMethodError("s.getUserAction: failed to s.userActionMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.getMasterAction: failed to s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnum",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnum(
							nil,
							masterActionStep.MasterActionStepEnum_None,
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             nil,
				},
			},
			wantErr: errors.NewMethodError("s.getMasterAction: failed to s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnum", errors.NewTestError()),
		},
		{
			name: "異常：s.getUserAction: failed to s.userActionMysqlRepository.Find",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnum(
							nil,
							masterActionStep.MasterActionStepEnum_None,
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                nil,
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
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
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             nil,
				},
			},
			wantErr: errors.NewMethodError("s.getUserAction: failed to s.userActionMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.getUserAction: expiration date has expired",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              pointers.Int32ToPointer(24),
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 2, 1, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(1),
				},
			},
			wantErr: errors.NewError("failed to s.getUserAction: expiration date has expired"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &actionService{
				masterActionMysqlRepository:        tt.fields.masterActionMysqlRepository(ctrl),
				masterActionRunMysqlRepository:     tt.fields.masterActionRunMysqlRepository(ctrl),
				masterActionStepMysqlRepository:    tt.fields.masterActionStepMysqlRepository(ctrl),
				masterActionTriggerMysqlRepository: tt.fields.masterActionTriggerMysqlRepository(ctrl),
				userActionMysqlRepository:          tt.fields.userActionMysqlRepository(ctrl),
			}

			err := s.Check(tt.args.ctx, tt.args.now, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestActionService_Run(t *testing.T) {
	type fields struct {
		masterActionMysqlRepository        func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository
		masterActionRunMysqlRepository     func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository
		masterActionStepMysqlRepository    func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository
		masterActionTriggerMysqlRepository func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository
		userActionMysqlRepository          func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		now time.Time
		req *ActionRunRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "正常：実行できる（TargetIdがある場合）",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{},
							nil,
						)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：実行できる（TargetIdがない場合）",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnum(
							nil,
							masterActionStep.MasterActionStepEnum_None,
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{},
							nil,
						)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             nil,
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：実行できる（トリガーアクションがない場合）",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{},
							nil,
						)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：実行できる（トリガーアクションがDiscontinuationの場合）",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Discontinuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              pointers.Int32ToPointer(24),
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{},
							nil,
						)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：アクションがない場合",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							nil,
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：実行できる（ActionRunが存在する場合）",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{
								{
									MasterActionRunId: 3,
									Name:              "アクションRun3",
									MasterActionId:    3,
								},
								{
									MasterActionRunId: 4,
									Name:              "アクションRun4",
									MasterActionId:    4,
								},
							},
							nil,
						)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(3),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 3,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 3,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(4),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 4,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 4,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：実行できる（期限切れにより再実行された場合）",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Discontinuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              pointers.Int32ToPointer(32),
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{},
							nil,
						)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2022, 12, 30, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.getMasterAction: failed to s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnumAndTargetId",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.getMasterAction: failed to s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnumAndTargetId", errors.NewTestError()),
		},
		{
			name: "異常：s.checkTriggerUserAction: failed to s.masterActionMysqlRepository.Find）",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
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
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.checkTriggerUserAction: failed to s.masterActionMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.deleteTriggerUserAction: failed to s.userActionMysqlRepository.Delete",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Discontinuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              pointers.Int32ToPointer(24),
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
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
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.deleteTriggerUserAction: failed to s.userActionMysqlRepository.Delete", errors.NewTestError()),
		},
		{
			name: "異常：s.run: failed to s.update: failed to s.userActionMysqlRepository.FindOrNil",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
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
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.run: failed to s.update: failed to s.userActionMysqlRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：s.run: failed to s.update: failed to s.userActionMysqlRepository.Create",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
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
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.run: failed to s.update: failed to s.userActionMysqlRepository.Create", errors.NewTestError()),
		},
		{
			name: "異常：s.getMasterAction: failed to s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnum",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnum(
							nil,
							masterActionStep.MasterActionStepEnum_None,
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             nil,
				},
			},
			wantErr: errors.NewMethodError("s.getMasterAction: failed to s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnum", errors.NewTestError()),
		},
		{
			name: "異常：s.checkTriggerUserAction: failed to s.getUserAction: failed to s.userActionMysqlRepository.Find",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
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
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.checkTriggerUserAction: failed to s.getUserAction: failed to s.userActionMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.checkTriggerUserAction: failed to s.getUserAction: expiration date has expired",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Discontinuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              pointers.Int32ToPointer(24),
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 3, 1, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewError("failed to s.checkTriggerUserAction: failed to s.getUserAction: expiration date has expired"),
		},
		{
			name: "異常：s.deleteTriggerUserAction: MasterActionTriggerType does not exist",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: 999,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              pointers.Int32ToPointer(24),
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewError("failed to s.deleteTriggerUserAction: MasterActionTriggerType does not exist"),
		},
		{
			name: "異常：s.run: failed to s.masterActionRunMysqlRepository.FindListByMasterActionId",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterActionId(
							nil,
							int64(2),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.run: failed to s.masterActionRunMysqlRepository.FindListByMasterActionId", errors.NewTestError()),
		},
		{
			name: "異常：s.run: failed to s.userActionMysqlRepository.Create: failed to s.userActionMysqlRepository.FindOrNil",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              nil,
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{
								{
									MasterActionRunId: 3,
									Name:              "アクションRun3",
									MasterActionId:    3,
								},
								{
									MasterActionRunId: 4,
									Name:              "アクションRun4",
									MasterActionId:    4,
								},
							},
							nil,
						)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(3),
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
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.run: failed to s.userActionMysqlRepository.Create: failed to s.userActionMysqlRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：s.run: failed to s.update: failed to s.userActionMysqlRepository.Update",
			fields: fields{
				masterActionMysqlRepository: func(ctrl *gomock.Controller) masterAction.MasterActionMysqlRepository {
					m := masterAction.NewMockMasterActionMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNilByMasterActionStepEnumAndTargetId(
							nil,
							masterActionStep.MasterActionStepEnum_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          2,
								Name:                    "テストアクション2",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Continuation,
								TargetId:                pointers.Int64ToPointer(2),
								TriggerMasterActionId:   pointers.Int64ToPointer(1),
								Expiration:              nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								MasterActionId:          1,
								Name:                    "テストアクション1",
								MasterActionStepEnum:    masterActionStep.MasterActionStepEnum_None,
								MasterActionTriggerEnum: masterActionTrigger.MasterActionTriggerEnum_Discontinuation,
								TargetId:                pointers.Int64ToPointer(1),
								TriggerMasterActionId:   nil,
								Expiration:              pointers.Int32ToPointer(32),
							},
							nil,
						)
					return m
				},
				masterActionRunMysqlRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunMysqlRepository {
					m := masterActionRun.NewMockMasterActionRunMysqlRepository(ctrl)
					return m
				},
				masterActionStepMysqlRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepMysqlRepository {
					m := masterActionStep.NewMockMasterActionStepMysqlRepository(ctrl)
					return m
				},
				masterActionTriggerMysqlRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerMysqlRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerMysqlRepository(ctrl)
					return m
				},
				userActionMysqlRepository: func(ctrl *gomock.Controller) userAction.UserActionMysqlRepository {
					m := userAction.NewMockUserActionMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2022, 12, 30, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
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
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:               "0:test",
					MasterActionStepEnum: masterActionStep.MasterActionStepEnum_None,
					TargetId:             pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.run: failed to s.update: failed to s.userActionMysqlRepository.Update", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &actionService{
				masterActionMysqlRepository:        tt.fields.masterActionMysqlRepository(ctrl),
				masterActionRunMysqlRepository:     tt.fields.masterActionRunMysqlRepository(ctrl),
				masterActionStepMysqlRepository:    tt.fields.masterActionStepMysqlRepository(ctrl),
				masterActionTriggerMysqlRepository: tt.fields.masterActionTriggerMysqlRepository(ctrl),
				userActionMysqlRepository:          tt.fields.userActionMysqlRepository(ctrl),
			}

			err := s.Run(tt.args.ctx, tt.args.tx, tt.args.now, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
