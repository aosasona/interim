package interim

import (
	"reflect"
	"testing"
)

func Test_interim_GetSet(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "get successfully",
			args:    args{"hello"},
			want:    "world",
			wantErr: false,
		},
		{
			name:    "get failed",
			args:    args{"john"},
			want:    "",
			wantErr: true,
		},
	}

	i := New(Config{CacheSize: 8})
	i.Set("hello", "world")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var result string
			err := i.Get(tt.args.key, &result)

			if (err != nil) != tt.wantErr {
				t.Errorf("interim.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("interim.Get() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_interim_Exists(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exists",
			args: args{"exists"},
			want: true,
		},
		{
			name: "doesn't exist",
			args: args{"somekey"},
			want: false,
		},
	}

	i := New(Config{CacheSize: 8})
	i.Set("exists", true)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := i.Exists(tt.args.key); got != tt.want {
				t.Errorf("interim.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_interim_Remove(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no error",
			args: args{"exists"},
			want: true,
		},
	}

	i := New(Config{CacheSize: 8})
	i.Set("exists", true)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := i.Delete(tt.args.key); err != nil {
				t.Errorf("interim.Delete() = %v, want %v", err, tt.want)
			}

			if i.Len() != 0 {
				t.Errorf("interim is not empty, it still contains %d records", i.Len())
			}
		})
	}
}
