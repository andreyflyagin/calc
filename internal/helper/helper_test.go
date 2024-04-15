package helper

import (
	"calc/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPairsToStrings(t *testing.T) {
	type args struct {
		pairs []model.Pair
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ok",
			args: args{
				pairs: []model.Pair{{
					Base:  "USDT",
					Quote: "USD",
				}, {
					Base:  "EUR",
					Quote: "ETH",
				}},
			},
			want: []string{"USDT/USD", "EUR/ETH"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, PairsToStrings(tt.args.pairs))
		})
	}
}
