package multiplication_test

import (
	"context"
	"dummy/proto"
	mul "dummy/service/multiplication"
	"reflect"
	"testing"
)

func TestMultiplicationServer_Multiply(t *testing.T) {
	type fields struct {
		UnimplementedMultiplicationServer proto.UnimplementedMultiplicationServer
	}
	type args struct {
		c     context.Context
		input *proto.Input
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.Output
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "test_1",
			fields: fields{},
			args: args{
				c:     context.Background(),
				input: &proto.Input{A: 10, B: 15},
			},
			want:    &proto.Output{SuccessStatus: 1, Res: 150},
			wantErr: false,
		},
		{
			name:   "test_2",
			fields: fields{},
			args: args{
				c:     context.Background(),
				input: &proto.Input{A: 12, B: 12},
			},
			want:    &proto.Output{SuccessStatus: 1, Res: 144},
			wantErr: false,
		},
		{
			name:   "test_3",
			fields: fields{},
			args: args{
				c:     context.Background(),
				input: &proto.Input{A: 27, B: 12},
			},
			want:    &proto.Output{SuccessStatus: 1, Res: 324},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &mul.MultiplicationServer{
				UnimplementedMultiplicationServer: tt.fields.UnimplementedMultiplicationServer,
			}
			got, err := server.Multiply(tt.args.c, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiplicationServer.Multiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplicationServer.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
