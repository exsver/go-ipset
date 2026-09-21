package ipset

import (
	"reflect"
	"testing"
)

func TestCreateOptions_GenArgs(t *testing.T) {
	type fields struct {
		Type     string
		Timeout  int
		Counters bool
		Comment  bool
		MaxElem  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name:    "empty type",
			fields:  fields{},
			wantErr: true,
		},
		{
			name: "unknown type",
			fields: fields{
				Type: "hash:unknown",
			},
			wantErr: true,
		},
		{
			name: "hash:ip minimal",
			fields: fields{
				Type: "hash:ip",
			},
			want: []string{"hash:ip"},
		},
		{
			name: "hash:ip with maxelem",
			fields: fields{
				Type:    "hash:ip",
				MaxElem: 65536,
			},
			want: []string{"hash:ip", "maxelem", "65536"},
		},
		{
			name: "hash:net with all options",
			fields: fields{
				Type:     "hash:net",
				Timeout:  300,
				Counters: true,
				Comment:  true,
				MaxElem:  1024,
			},
			want: []string{"hash:net", "maxelem", "1024", "timeout", "300", "counters", "comment"},
		},
		{
			name: "bitmap:ip with timeout and counters",
			fields: fields{
				Type:     "bitmap:ip",
				Timeout:  60,
				Counters: true,
			},
			want: []string{"bitmap:ip", "timeout", "60", "counters"},
		},
		{
			name: "bitmap:ip ignores maxelem",
			fields: fields{
				Type:    "bitmap:ip",
				MaxElem: 100,
			},
			want: []string{"bitmap:ip"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &CreateOptions{
				Type:     tt.fields.Type,
				Timeout:  tt.fields.Timeout,
				Counters: tt.fields.Counters,
				Comment:  tt.fields.Comment,
				MaxElem:  tt.fields.MaxElem,
			}
			got, err := o.GenArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
