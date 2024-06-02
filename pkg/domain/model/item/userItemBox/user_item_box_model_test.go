package userItemBox

import (
	"reflect"
	"testing"
)

func TestUserItemBoxes_SetExistingItemMaps(t *testing.T) {
	type fields struct {
		UserItemBoxes UserItemBoxes
	}
	type args struct{}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int64]bool
	}{
		{
			name: "正常：無期限",
			fields: fields{
				UserItemBoxes: UserItemBoxes{
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
			},
			args: args{},
			want: map[int64]bool{
				1: true,
				2: true,
				3: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.UserItemBoxes.SetExistingItemMaps()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExistingItemMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
