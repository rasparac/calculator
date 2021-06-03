package calculator

import (
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/require"
)

func TestCache_Get(t *testing.T) {
	type fields struct {
		cache *cache.Cache
	}
	type args struct {
		key string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     float64
		want1    bool
		setKey   string
		setValue float64
	}{
		{
			name: "it should found values in cache",
			fields: fields{
				cache: cache.New(1*time.Minute, 1*time.Minute),
			},
			args: args{
				key: "test",
			},
			setKey:   "test",
			setValue: 1,
			want:     1,
			want1:    true,
		},
		{
			name: "it should not found values in cache",
			fields: fields{
				cache: cache.New(1*time.Minute, 1*time.Minute),
			},
			args: args{
				key: "test",
			},
			setKey:   "test2",
			setValue: 1,
			want:     0,
			want1:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				cache: tt.fields.cache,
			}

			c.Set(tt.setKey, tt.setValue)

			got, got1 := c.Get(tt.args.key)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.want1, got1)
		})
	}
}
