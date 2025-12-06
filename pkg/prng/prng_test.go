package prng_test

import (
	"testing"

	"github.com/GeekchanskiY/pet-project/pkg/prng"
)

func Test_generator_Generate(t *testing.T) {
	type fields struct {
		seed string
	}
	type args struct {
		pos int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint8
		wantErr bool
	}{
		{
			name:    "test zero value 1",
			fields:  fields{seed: "hello world"},
			args:    args{pos: 0},
			want:    155,
			wantErr: false,
		},
		{
			name:    "test zero value 2",
			fields:  fields{seed: "hello world 2"},
			args:    args{pos: 0},
			want:    185,
			wantErr: false,
		},
		{
			name:    "test zero value 3",
			fields:  fields{seed: "amogus"},
			args:    args{pos: 0},
			want:    180,
			wantErr: false,
		},
		{
			name:    "test first value 1",
			fields:  fields{seed: "amogus"},
			args:    args{pos: 1},
			want:    48,
			wantErr: false,
		},
		{
			name:    "test first value 2",
			fields:  fields{seed: "hello world"},
			args:    args{pos: 1},
			want:    2,
			wantErr: false,
		},
		{
			name:    "test first value 3",
			fields:  fields{seed: "hello world 2"},
			args:    args{pos: 1},
			want:    71,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := prng.NewGenerator(tt.fields.seed)
			if err != nil && !tt.wantErr {
				t.Fatalf("NewGenerator() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if got := g.Generate(tt.args.pos); got != tt.want {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
