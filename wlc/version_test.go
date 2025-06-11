package version_test

import (
	"reflect"
	"testing"

	version "github.com/MaineK00n/go-cisco-version/wlc"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    version.Version
		wantErr bool
	}{
		{
			name: "7.0.1",
			args: args{ver: "7.0.1"},
			want: version.Version{Major: 7, Minor: 0, Maintenance: 1},
		},
		{
			name: "7.0(1)",
			args: args{ver: "7.0(1)"},
			want: version.Version{Major: 7, Minor: 0, Maintenance: 1},
		},
		{
			name: "7.0.1.1",
			args: args{ver: "7.0.1.1"},
			want: version.Version{Major: 7, Minor: 0, Maintenance: 1, Build: 1},
		},
		{
			name: "7.0(1)1",
			args: args{ver: "7.0(1)1"},
			want: version.Version{Major: 7, Minor: 0, Maintenance: 1, Build: 1},
		},
		{
			name: "7.0(1.1)",
			args: args{ver: "7.0(1.1)"},
			want: version.Version{Major: 7, Minor: 0, Maintenance: 1, Build: 1},
		},
		{
			name:    "7.0",
			args:    args{ver: "7.0"},
			wantErr: true,
		},
		{
			name:    "7.0(1)1a",
			args:    args{ver: "7.0(1)1a"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := version.NewVersion(tt.args.ver)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_Compare(t *testing.T) {
	type fields struct {
		Major       int
		Minor       int
		Maintenance int
		Build       int
	}
	type args struct {
		v2 version.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "7.0.1 = 7.0.1",
			fields: fields{Major: 7, Minor: 0, Maintenance: 1},
			args:   args{v2: version.Version{Major: 7, Minor: 0, Maintenance: 1}},
			want:   0,
		},
		{
			name:   "7.0.1.1 = 7.0.1.1",
			fields: fields{Major: 7, Minor: 0, Maintenance: 1, Build: 1},
			args:   args{v2: version.Version{Major: 7, Minor: 0, Maintenance: 1, Build: 1}},
			want:   0,
		},
		{
			name:   "7.0.1 < 7.0.2",
			fields: fields{Major: 7, Minor: 0, Maintenance: 1},
			args:   args{v2: version.Version{Major: 7, Minor: 0, Maintenance: 2}},
			want:   -1,
		},
		{
			name:   "7.0.1.2 > 7.0.1.1",
			fields: fields{Major: 7, Minor: 0, Maintenance: 1, Build: 2},
			args:   args{v2: version.Version{Major: 7, Minor: 0, Maintenance: 1, Build: 1}},
			want:   +1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := version.Version{
				Major:       tt.fields.Major,
				Minor:       tt.fields.Minor,
				Maintenance: tt.fields.Maintenance,
				Build:       tt.fields.Build,
			}
			if got := v1.Compare(tt.args.v2); got != tt.want {
				t.Errorf("Version.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_String(t *testing.T) {
	type fields struct {
		Major       int
		Minor       int
		Maintenance int
		Build       int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "7.0.1",
			fields: fields{Major: 7, Minor: 0, Maintenance: 1},
			want:   "7.0.1.0",
		},
		{
			name:   "7.0.1.1",
			fields: fields{Major: 7, Minor: 0, Maintenance: 1, Build: 1},
			want:   "7.0.1.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := version.Version{
				Major:       tt.fields.Major,
				Minor:       tt.fields.Minor,
				Maintenance: tt.fields.Maintenance,
				Build:       tt.fields.Build,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
