package version_test

import (
	"reflect"
	"testing"

	version "github.com/MaineK00n/go-cisco-version/ios-xe"
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
			name: "Everest-16.5.1",
			args: args{
				ver: "Everest-16.5.1",
			},
			want: version.Version{
				Major:       16,
				Minor:       5,
				Maintenance: "1",
				Release:     "Everest",
			},
		},
		{
			name: "16.5.1a",
			args: args{
				ver: "16.5.1a",
			},
			want: version.Version{
				Major:       16,
				Minor:       5,
				Maintenance: "1a",
			},
		},
		{
			name: "3.16.1aS",
			args: args{
				ver: "3.16.1aS",
			},
			want: version.Version{
				Major:       3,
				Minor:       16,
				Maintenance: "1a",
				Release:     "S",
			},
		},
		{
			name: "3.16.2S",
			args: args{
				ver: "3.16.2S",
			},
			want: version.Version{
				Major:       3,
				Minor:       16,
				Maintenance: "2",
				Release:     "S",
			},
		},
		{
			name: "3.4.1SG",
			args: args{
				ver: "3.4.1SG",
			},
			want: version.Version{
				Major:       3,
				Minor:       4,
				Maintenance: "1",
				Release:     "SG",
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
		Release     string
		Major       int
		Minor       int
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
			name: "Everest-16.5.1 = 16.5.1",
			fields: fields{
				Release:     "Everest",
				Major:       16,
				Minor:       5,
				Maintenance: "1",
			},
			args: args{
				v2: version.Version{
					Major:       16,
					Minor:       5,
					Maintenance: "1",
				},
			},
			want: 0,
		},
		{
			name: "16.5.1a > 16.5.1",
			fields: fields{
				Major:       16,
				Minor:       5,
				Maintenance: "1a",
			},
			args: args{
				v2: version.Version{
					Major:       16,
					Minor:       5,
					Maintenance: "1",
				},
			},
			want: +1,
		},
		{
			name: "3.16.1aS < 16.5.1",
			fields: fields{
				Major:       3,
				Minor:       16,
				Maintenance: "1a",
				Release:     "S",
			},
			args: args{
				v2: version.Version{
					Major:       16,
					Minor:       5,
					Maintenance: "1",
				},
			},
			want: -1,
		},
		{
			name: "3.16.2S > 3.16.1aS",
			fields: fields{
				Major:       3,
				Minor:       16,
				Maintenance: "2",
				Release:     "S",
			},
			args: args{
				v2: version.Version{
					Major:       3,
					Minor:       16,
					Maintenance: "1a",
					Release:     "S",
				},
			},
			want: +1,
		},
		{
			name: "3.4.1SG vs 3.16.2S",
			fields: fields{
				Major:       3,
				Minor:       4,
				Maintenance: "1",
				Release:     "SG",
			},
			args: args{
				v2: version.Version{
					Major:       3,
					Minor:       16,
					Maintenance: "2",
					Release:     "S",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := version.Version{
				Release:     tt.fields.Release,
				Major:       tt.fields.Major,
				Minor:       tt.fields.Minor,
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
		Release     string
		Major       int
		Minor       int
		Maintenance string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Everest-16.5.1",
			fields: fields{
				Release:     "Everest",
				Major:       16,
				Minor:       5,
				Maintenance: "1",
			},
			want: "Everest-16.5.1",
		},
		{
			name: "16.5.1a",
			fields: fields{
				Release:     "",
				Major:       16,
				Minor:       5,
				Maintenance: "1a",
			},
			want: "16.5.1a",
		},
		{
			name: "3.16.1aS",
			fields: fields{
				Release:     "S",
				Major:       3,
				Minor:       16,
				Maintenance: "1a",
			},
			want: "3.16.1aS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := version.Version{
				Release:     tt.fields.Release,
				Major:       tt.fields.Major,
				Minor:       tt.fields.Minor,
				Maintenance: tt.fields.Maintenance,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
