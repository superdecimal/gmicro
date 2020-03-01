package calculator_test

import (
	"context"
	gmrpc "superdecimal/gmicro/pkg/proto"
	"superdecimal/gmicro/services/calc-api/calculator"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int32
		b    int32
		res  int32
	}{
		{
			name: "success",
			a:    5,
			b:    5,
			res:  10,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			srv := calculator.NewServer()
			ctx := context.Background()
			request := gmrpc.AddRequest{A: tt.a, B: tt.b}

			resp, err := srv.Add(ctx, &request)
			assert.NoError(t, err)
			assert.Equal(t, tt.res, resp.GetResult())
		})
	}
}
