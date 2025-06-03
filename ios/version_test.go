package version_test

import (
	"reflect"
	"testing"

	version "github.com/MaineK00n/go-cisco-version/ios"
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
			name: "15.0",
			args: args{
				ver: "15.0",
			},
			wantErr: true,
		},
		{
			name: "15.0(1)",
			args: args{
				ver: "15.0(1)",
			},
			want: version.Version{
				Major:   15,
				Minor:   0,
				Feature: "1",
			},
		},
		{
			name: "15.0(1a)",
			args: args{
				ver: "15.0(1a)",
			},
			want: version.Version{
				Major:   15,
				Minor:   0,
				Feature: "1a",
			},
		},
		{
			name: "15.0(1)M",
			args: args{
				ver: "15.0(1)M",
			},
			want: version.Version{
				Major:   15,
				Minor:   0,
				Feature: "1",
				Release: "M",
			},
		},
		{
			name: "15.0(1)M1",
			args: args{
				ver: "15.0(1)M1",
			},
			want: version.Version{
				Major:       15,
				Minor:       0,
				Feature:     "1",
				Release:     "M",
				Maintenance: "1",
			},
		},
		{
			name: "15.0(1)SY1a",
			args: args{
				ver: "15.0(1)SY1a",
			},
			want: version.Version{
				Major:       15,
				Minor:       0,
				Feature:     "1",
				Release:     "SY",
				Maintenance: "1a",
			},
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
		Feature     string
		Release     string
		Maintenance string
	}
	type args struct {
		v2 version.Version
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "15.0(1) = 15.0(1)",
			fields: fields{
				Major:   15,
				Minor:   0,
				Feature: "1",
			},
			args: args{
				v2: version.Version{
					Major:   15,
					Minor:   0,
					Feature: "1",
				},
			},
			want: 0,
		},
		{
			name: "15.0(1)M = 15.0(1)M",
			fields: fields{
				Major:   15,
				Minor:   0,
				Feature: "1",
				Release: "M",
			},
			args: args{
				v2: version.Version{
					Major:   15,
					Minor:   0,
					Feature: "1",
					Release: "M",
				},
			},
			want: 0,
		},
		{
			name: "15.0(1)M1 = 15.0(1)M1",
			fields: fields{
				Major:       15,
				Minor:       0,
				Feature:     "1",
				Release:     "M",
				Maintenance: "1",
			},
			args: args{
				v2: version.Version{
					Major:       15,
					Minor:       0,
					Feature:     "1",
					Release:     "M",
					Maintenance: "1",
				},
			},
			want: 0,
		},
		{
			name: "15.0(1) < 15.0(2)",
			fields: fields{
				Major:   15,
				Minor:   0,
				Feature: "1",
			},
			args: args{
				v2: version.Version{
					Major:   15,
					Minor:   0,
					Feature: "2",
				},
			},
			want: -1,
		},
		{
			name: "15.0(1)M1 > 15.0(1)M",
			fields: fields{
				Major:       15,
				Minor:       0,
				Feature:     "1",
				Release:     "M",
				Maintenance: "1",
			},
			args: args{
				v2: version.Version{
					Major:   15,
					Minor:   0,
					Feature: "1",
					Release: "M",
				},
			},
			want: +1,
		},
		{
			name: "15.0(1)M1 < 15.0(1)M1a",
			fields: fields{
				Major:       15,
				Minor:       0,
				Feature:     "1",
				Release:     "M",
				Maintenance: "1",
			},
			args: args{
				v2: version.Version{
					Major:       15,
					Minor:       0,
					Feature:     "1",
					Release:     "M",
					Maintenance: "1a",
				},
			},
			want: -1,
		},
		{
			name: "15.0(1) vs 15.0(1)M",
			fields: fields{
				Major:   15,
				Minor:   0,
				Feature: "1",
			},
			args: args{
				v2: version.Version{
					Major:   15,
					Minor:   0,
					Feature: "1",
					Release: "M",
				},
			},
			wantErr: true,
		},
		{
			name: "15.0(1)M vs 15.0(1)T",
			fields: fields{
				Major:   15,
				Minor:   0,
				Feature: "1",
				Release: "M",
			},
			args: args{
				v2: version.Version{
					Major:   15,
					Minor:   0,
					Feature: "1",
					Release: "T",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := version.Version{
				Major:       tt.fields.Major,
				Minor:       tt.fields.Minor,
				Feature:     tt.fields.Feature,
				Release:     tt.fields.Release,
				Maintenance: tt.fields.Maintenance,
			}
			got, err := v1.Compare(tt.args.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Version.Compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Version.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_String(t *testing.T) {
	type fields struct {
		Major       int
		Minor       int
		Feature     string
		Release     string
		Maintenance string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "15.0(1)",
			fields: fields{
				Major:   15,
				Minor:   0,
				Feature: "1",
			},
			want: "15.0(1)",
		},
		{
			name: "15.0(1)M",
			fields: fields{
				Major:   15,
				Minor:   0,
				Feature: "1",
				Release: "M",
			},
			want: "15.0(1)M",
		},
		{
			name: "15.0(1)M1",
			fields: fields{
				Major:       15,
				Minor:       0,
				Feature:     "1",
				Release:     "M",
				Maintenance: "1",
			},
			want: "15.0(1)M1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := version.Version{
				Major:       tt.fields.Major,
				Minor:       tt.fields.Minor,
				Feature:     tt.fields.Feature,
				Release:     tt.fields.Release,
				Maintenance: tt.fields.Maintenance,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
