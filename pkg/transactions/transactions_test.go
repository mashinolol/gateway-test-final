package transactions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGroupTransactions(t *testing.T) {
	type args struct {
		transactions []*Transaction
		interval     string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Transaction
		wantErr bool
	}{
		{
			name: "Group by day",
			args: args{
				transactions: []*Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
					{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
				interval: "day",
			},
			want: []*Transaction{
				{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
				{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
			},
			wantErr: false,
		},
		{
			name: "Group by hour",
			args: args{
				transactions: []*Transaction{
					{Value: 100, Timestamp: time.Unix(1616006400, 0)},
					{Value: 200, Timestamp: time.Unix(1616009999, 0)},
					{Value: 150, Timestamp: time.Unix(1616010000, 0)},
				},
				interval: "hour",
			},
			want: []*Transaction{
				{Value: 100, Timestamp: time.Unix(1616006400, 0)},
				{Value: 150, Timestamp: time.Unix(1616010000, 0)},
			},
			wantErr: false,
		},
		{
			name: "Group by month",
			args: args{
				transactions: []*Transaction{
					{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
					{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
					{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
					{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
					{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
				},
				interval: "month",
			},
			want: []*Transaction{
				{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
			},
			wantErr: false,
		},
		{
			name: "Group by week",
			args: args{
				transactions: []*Transaction{
					{Value: 100, Timestamp: time.Unix(1614800000, 0)},
					{Value: 200, Timestamp: time.Unix(1614886400, 0)},
					{Value: 150, Timestamp: time.Unix(1615070400, 0)},
				},
				interval: "week",
			},
			want: []*Transaction{
				{Value: 150, Timestamp: time.Unix(1615070400, 0)},
			},
			wantErr: false,
		},
		{
			name: "Invalid interval",
			args: args{
				transactions: []*Transaction{
					{Value: 100, Timestamp: time.Unix(1616006400, 0)},
					{Value: 200, Timestamp: time.Unix(1616009999, 0)},
					{Value: 150, Timestamp: time.Unix(1616010000, 0)},
				},
				interval: "invalid_interval",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Group(tt.args.transactions, tt.args.interval)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				require.Equal(t, tt.want, result)
			}
		})
	}
}
