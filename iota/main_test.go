package main

import (
	"reflect"
	"testing"
)

func TestEnvironments(t *testing.T) {
	tests := []struct {
		name string
		want []Environment
	}{
		{
			name: "all_environments",
			want: []Environment{
				ProductionEnvironment,
				StagingEnvironment,
				TestingEnvironment,
				DevelopmentEnvironment,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := environments()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironment(t *testing.T) {
	tests := []struct {
		name string
		give string
		want Environment
	}{
		{
			name: "production",
			give: "production",
			want: ProductionEnvironment,
		}, {
			name: "staging",
			give: "staging",
			want: StagingEnvironment,
		}, {
			name: "test",
			give: "test",
			want: TestingEnvironment,
		}, {
			name: "development",
			give: "development",
			want: DevelopmentEnvironment,
		}, {
			name: "empty",
			give: "",
			want: InvalidEnvironment,
		}, {
			name: "integration",
			give: "integration",
			want: InvalidEnvironment,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := environment(tt.give)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
