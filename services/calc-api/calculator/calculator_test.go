package calculator_test

import (
	"context"
	"errors"
	"io"
	gmrpc "superdecimal/gmicro/pkg/proto"
	gmrpcmock "superdecimal/gmicro/pkg/proto/mock"

	"superdecimal/gmicro/services/calc-api/calculator"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
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
			logger, _ := zap.NewDevelopment()
			srv := calculator.NewServer(logger)
			ctx := context.Background()
			request := gmrpc.AddRequest{A: tt.a, B: tt.b}

			resp, err := srv.Add(ctx, &request)
			assert.NoError(t, err)
			assert.Equal(t, tt.res, resp.GetResult())
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int32
		expect func(
			stream *gmrpcmock.MockCalculatorAPI_SumServer,
			nums []int32,
			res int32,
		)
		res int32
		err bool
	}{
		{
			name: "success",
			nums: []int32{10, 25, 100},
			expect: func(
				stream *gmrpcmock.MockCalculatorAPI_SumServer,
				nums []int32,
				res int32,
			) {
				for _, n := range nums {
					stream.EXPECT().Recv().Return(&gmrpc.Integer{Num: n}, nil)
				}
				stream.EXPECT().Recv().Return(nil, io.EOF)
				stream.EXPECT().SendAndClose(&gmrpc.SumResponse{Result: res})
			},
			res: 135,
		},
		{
			name: "fail",
			expect: func(
				stream *gmrpcmock.MockCalculatorAPI_SumServer,
				nums []int32,
				res int32,
			) {
				stream.EXPECT().Recv().Return(nil, errors.New("random erro"))
			},
			res: 135,
			err: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			logger, _ := zap.NewDevelopment()
			srv := calculator.NewServer(logger)

			stream := gmrpcmock.NewMockCalculatorAPI_SumServer(ctrl)
			tt.expect(stream, tt.nums, tt.res)

			err := srv.Sum(stream)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
