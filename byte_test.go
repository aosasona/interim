package interim

import (
	"testing"
)

func Test_encodeToByte(t *testing.T) {
	type args struct {
		data any
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "encode byte",
			args:    args{data: "hi"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := encodeToByte(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeToByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_decodeFromByte(t *testing.T) {
	type args struct {
		encoded []byte
	}
	target := ""
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "decode bytes",
			args: args{
				encoded: []byte{5, 12, 0, 2, 104, 105},
			},
			want:    "hi",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := decodeFromByte(tt.args.encoded, &target); (err != nil) != tt.wantErr || target != tt.want {
				t.Errorf("decodeFromByte() error = %v, wantErr %v -- decoded: %v, expected: %v", err, tt.wantErr, target, tt.want)
			}
		})
	}
}
