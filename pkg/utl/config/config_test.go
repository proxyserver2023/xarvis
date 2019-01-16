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
		{
			name:    "Fail on wrong file format",
			path:    "testdata/config.invalid.yaml",
			wantErr: true,
		},
		{
			name: "Success",
			path: "testdata/config.testdata.yaml",
			wantData: &Configuration{
				DB: &Database{
					PSN:        "postgres://postgres:postgres@postgres",
					LogQueries: true,
					Timeout:    20,
				},
				Server: &Server{
					Port:         ":8080",
					Debug:        true,
					ReadTimeout:  15,
					WriteTimeout: 20,
				},
				JWT: &JWT{
					Secret:           "testing",
					Duration:         10,
					RefreshDuration:  10,
					MaxRefresh:       144,
					SigningAlgorithm: "HS384",
				},
				App: &Application{
					MinPasswordStr: 3,
					SwaggerUIPath:  "assets/swagger",
				},
			},
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