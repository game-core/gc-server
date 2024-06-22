package userExchange

import (
	"reflect"
	"testing"
	"time"
)

func TestNewUserExchange_CheckResetAt(t *testing.T) {
	type fields struct {
		UserExchange *UserExchange
	}
	type args struct {
		now          time.Time
		intervalHour int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "正常：リセット有",
			fields: fields{
				UserExchange: &UserExchange{
					UserId:           "0:test",
					MasterExchangeId: 1,
					ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				now:          time.Date(2023, 1, 2, 9, 0, 0, 1, time.UTC),
				intervalHour: 24,
			},
			want: true,
		},
		{
			name: "正常：リセット無し",
			fields: fields{
				UserExchange: &UserExchange{
					UserId:           "0:test",
					MasterExchangeId: 1,
					ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				now:          time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				intervalHour: 24,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.UserExchange.CheckResetAt(tt.args.now, tt.args.intervalHour)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckResetAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserExchange_CreateResetAt(t *testing.T) {
	type fields struct {
		UserExchange *UserExchange
	}
	type args struct {
		now          time.Time
		startAt      time.Time
		resetHour    int32
		intervalHour int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		{
			name: "正常：リセットする",
			fields: fields{
				UserExchange: &UserExchange{
					UserId:           "0:test",
					MasterExchangeId: 1,
					ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				now:          time.Date(2023, 1, 2, 9, 0, 0, 1, time.UTC),
				startAt:      time.Date(2022, 12, 1, 9, 0, 0, 0, time.UTC),
				resetHour:    9,
				intervalHour: 24,
			},
			want: time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
		},
		{
			name: "正常：リセットする",
			fields: fields{
				UserExchange: &UserExchange{
					UserId:           "0:test",
					MasterExchangeId: 1,
					ResetAt:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				now:          time.Date(2023, 1, 3, 8, 0, 0, 0, time.UTC),
				startAt:      time.Date(2022, 12, 1, 9, 0, 0, 0, time.UTC),
				resetHour:    9,
				intervalHour: 24,
			},
			want: time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
		},
		{
			name: "正常：nil",
			fields: fields{
				UserExchange: nil,
			},
			args: args{
				now:          time.Date(2023, 1, 2, 9, 0, 0, 1, time.UTC),
				startAt:      time.Date(2022, 12, 1, 9, 0, 0, 0, time.UTC),
				resetHour:    9,
				intervalHour: 24,
			},
			want: time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
		},
		{
			name: "正常：nil",
			fields: fields{
				UserExchange: nil,
			},
			args: args{
				now:          time.Date(2023, 1, 2, 9, 0, 0, 1, time.UTC),
				startAt:      time.Date(2022, 12, 1, 9, 0, 0, 0, time.UTC),
				resetHour:    10,
				intervalHour: 24,
			},
			want: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "正常：nil",
			fields: fields{
				UserExchange: nil,
			},
			args: args{
				now:          time.Date(2023, 1, 2, 9, 0, 0, 1, time.UTC),
				startAt:      time.Date(2022, 12, 1, 9, 0, 0, 0, time.UTC),
				resetHour:    8,
				intervalHour: 24,
			},
			want: time.Date(2023, 1, 1, 8, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.UserExchange.CreateResetAt(tt.args.now, tt.args.startAt, tt.args.resetHour, tt.args.intervalHour)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateResetAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
