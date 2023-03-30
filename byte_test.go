package interim

import (
	"reflect"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeToByte(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeToByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeFromByte(t *testing.T) {
	type args struct {
		encoded []byte
		target  any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := decodeFromByte(tt.args.encoded, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("decodeFromByte() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
