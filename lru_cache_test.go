package interim

import (
	"container/list"
	"reflect"
	"testing"
)

func Test_newLRUCache(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *lruCache
	}{
		{
			name: "create new LRU Cache",
			args: args{0},
			want: &lruCache{
				capacity: 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newLRUCache(tt.args.capacity); !reflect.DeepEqual(got.capacity, tt.want.capacity) {
				t.Errorf("newLRUCache() = %v, want %v", got.capacity, tt.want.capacity)
			}
		})
	}
}

func Test_lruCache_get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  any
		want1 bool
	}{
		{
			name:  "get hello",
			args:  args{"hello"},
			want:  "world",
			want1: true,
		},
		{
			name:  "get john",
			args:  args{"john"},
			want:  "doe",
			want1: true,
		},
		{
			name:  "get harry",
			args:  args{"harry"},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLRUCache(8)
			l.put("hello", "world")
			l.put("john", "doe")

			got, got1 := l.get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lruCache.get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("lruCache.get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_lruCache_put(t *testing.T) {
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "put 1",
			args: args{"1", "one"},
		},
		{
			name: "put 2",
			args: args{"2", "two"},
		},
		{
			name: "put 3",
			args: args{"3", "three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLRUCache(2)
			l.put(tt.args.key, tt.args.value)
		})
	}
}

func Test_lruCache_remove(t *testing.T) {
	type fields struct {
		capacity int
		cache    map[any]*list.Element
		list     *list.List
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lruCache{
				capacity: tt.fields.capacity,
				cache:    tt.fields.cache,
				list:     tt.fields.list,
			}
			l.remove(tt.args.key)
		})
	}
}
