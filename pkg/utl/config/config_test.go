package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		name string
		path string
		wantData *Configuration
		wantErr bool
	}{
		{
			name: "If file is not present then fail",
			path: "notExists",
			wantData: nil,
			wantErr: true,
		},
	}


	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T){
			cfg, err := Load(tc.path)
			assert.Equal(t, tc.wantData, cfg)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}