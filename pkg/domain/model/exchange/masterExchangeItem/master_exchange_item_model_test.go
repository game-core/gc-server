package masterExchangeItem

import (
	"reflect"
	"testing"
)

func TestMasterExchangeItems_ExtractMasterExchangeId(t *testing.T) {
	type fields struct {
		MasterExchangeItems MasterExchangeItems
	}
	type args struct{}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				MasterExchangeItems: MasterExchangeItems{
					{
						MasterExchangeItemId: 1,
						MasterExchangeId:     1,
						MasterItemId:         1,
						Name:                 "テスト1",
						Count:                1,
					},
					{
						MasterExchangeItemId: 2,
						MasterExchangeId:     1,
						MasterItemId:         1,
						Name:                 "テスト2",
						Count:                1,
					},
				},
			},
			args: args{},
			want: 1,
		},
		{
			name: "正常：取得できない",
			fields: fields{
				MasterExchangeItems: MasterExchangeItems{},
			},
			args: args{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.MasterExchangeItems.ExtractMasterExchangeId()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractMasterExchangeId() = %v, want %v", got, tt.want)
			}
		})
	}
}
