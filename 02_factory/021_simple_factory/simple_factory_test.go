package factory

import (
	"reflect"
	"testing"
)

func TestNewRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct{
		name string
		args args
		want IRuleConfigparser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRuleConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}