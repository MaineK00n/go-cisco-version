package version_test

import (
	"reflect"
	"testing"

	version "github.com/MaineK00n/go-cisco-version/ios-xr"
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
			name: "4.3.2",
			args: args{
				ver: "4.3.2",
			},
			want: version.Version{
				Major:   4,
				Minor:   3,
				Release: 2,
			},
		},
		{
			name: "24.1.1",
			args: args{
				ver: "24.1.1",
			},
			want: version.Version{
				Major:   24,
				Minor:   1,
				Release: 1,
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
		Major   int
		Minor   int
		Release int
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
			name: "4.3.2 = 4.3.2",
			fields: fields{
				Major:   4,
				Minor:   3,
				Release: 2,
			},
			args: args{
				v2: version.Version{
					Major:   4,
					Minor:   3,
					Release: 2,
				},
			},
			want: 0,
		},
		{
			name: "4.3.2 < 4.3.3",
			fields: fields{
				Major:   4,
				Minor:   3,
				Release: 2,
			},
			args: args{
				v2: version.Version{
					Major:   4,
					Minor:   3,
					Release: 3,
				},
			},
			want: -1,
		},
		{
			name: "24.1.1 > 4.3.2",
			fields: fields{
				Major:   24,
				Minor:   1,
				Release: 1,
			},
			args: args{
				v2: version.Version{
					Major:   4,
					Minor:   3,
					Release: 2,
				},
			},
			want: +1,
		},
		{
			name: "24.1.1 < 24.1.2",
			fields: fields{
				Major:   24,
				Minor:   1,
				Release: 1,
			},
			args: args{
				v2: version.Version{
					Major:   24,
					Minor:   1,
					Release: 2,
				},
			},
			want: -1,
		},
		{
			name: "25.1.1 > 24.1.2",
			fields: fields{
				Major:   25,
				Minor:   1,
				Release: 1,
			},
			args: args{
				v2: version.Version{
					Major:   24,
					Minor:   1,
					Release: 2,
				},
			},
			want: +1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := version.Version{
				Major:   tt.fields.Major,
				Minor:   tt.fields.Minor,
				Release: tt.fields.Release,
			}
			if got := v1.Compare(tt.args.v2); got != tt.want {
				t.Errorf("Version.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_String(t *testing.T) {
	type fields struct {
		Major   int
		Minor   int
		Release int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "4.3.2",
			fields: fields{
				Major:   4,
				Minor:   3,
				Release: 2,
			},
			want: "4.3.2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := version.Version{
				Major:   tt.fields.Major,
				Minor:   tt.fields.Minor,
				Release: tt.fields.Release,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
