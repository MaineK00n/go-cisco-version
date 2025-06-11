package version_test

import (
	"reflect"
	"testing"

	version "github.com/MaineK00n/go-cisco-version/nx-os"
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
			name: "6.2(8b)",
			args: args{
				ver: "6.2(8b)",
			},
			want: version.Version{
				Major:       6,
				Minor:       2,
				Maintenance: "8b",
			},
		},
		{
			name: "7.1(3)N1(2)",
			args: args{
				ver: "7.1(3)N1(2)",
			},
			want: version.Version{
				Major:               7,
				Minor:               1,
				Maintenance:         "3",
				Platform:            "N",
				PlatformMinor:       1,
				PlatformMaintenance: "2",
			},
		},
		{
			name: "7.3(0)DX(1)",
			args: args{
				ver: "7.3(0)DX(1)",
			},
			want: version.Version{
				Major:               7,
				Minor:               3,
				Maintenance:         "0",
				Platform:            "DX",
				PlatformMinor:       0,
				PlatformMaintenance: "1",
			},
		},
		{
			name: "5.2(1)SM3(1.1a)",
			args: args{
				ver: "5.2(1)SM3(1.1a)",
			},
			want: version.Version{
				Major:               5,
				Minor:               2,
				Maintenance:         "1",
				Platform:            "SM",
				PlatformMinor:       3,
				PlatformMaintenance: "1.1a",
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
		Major               int
		Minor               int
		Maintenance         string
		Platform            string
		PlatformMinor       int
		PlatformMaintenance string
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
			name: "6.2(8b) = 6.2(8b)",
			fields: fields{
				Major:       6,
				Minor:       2,
				Maintenance: "8b",
			},
			args: args{
				v2: version.Version{
					Major:       6,
					Minor:       2,
					Maintenance: "8b",
				},
			},
			want: 0,
		},
		{
			name: "6.2(8b) < 6.3(0)",
			fields: fields{
				Major:       6,
				Minor:       2,
				Maintenance: "8b",
			},
			args: args{
				v2: version.Version{
					Major:       6,
					Minor:       3,
					Maintenance: "0",
				},
			},
			want: -1,
		},
		{
			name: "6.2(8b) < 7.1(3)N1(2)",
			fields: fields{
				Major:       6,
				Minor:       2,
				Maintenance: "8b",
			},
			args: args{
				v2: version.Version{
					Major:               7,
					Minor:               1,
					Maintenance:         "3",
					Platform:            "N",
					PlatformMinor:       1,
					PlatformMaintenance: "2",
				},
			},
			want: -1,
		},
		{
			name: "7.1(3)N1(2) = 7.1(3)N1(2)",
			fields: fields{
				Major:               7,
				Minor:               1,
				Maintenance:         "3",
				Platform:            "N",
				PlatformMinor:       1,
				PlatformMaintenance: "2",
			},
			args: args{
				v2: version.Version{
					Major:               7,
					Minor:               1,
					Maintenance:         "3",
					Platform:            "N",
					PlatformMinor:       1,
					PlatformMaintenance: "2",
				},
			},
			want: 0,
		},
		{
			name: "7.1(3)N1(2) > 7.1(3)N(3)",
			fields: fields{
				Major:               7,
				Minor:               1,
				Maintenance:         "3",
				Platform:            "N",
				PlatformMinor:       1,
				PlatformMaintenance: "2",
			},
			args: args{
				v2: version.Version{
					Major:               7,
					Minor:               1,
					Maintenance:         "3",
					Platform:            "N",
					PlatformMinor:       0,
					PlatformMaintenance: "3",
				},
			},
			want: +1,
		},
		{
			name: "7.1(3)N1(2) vs 7.1(3)D1(2)",
			fields: fields{
				Major:               7,
				Minor:               1,
				Maintenance:         "3",
				Platform:            "N",
				PlatformMinor:       1,
				PlatformMaintenance: "2",
			},
			args: args{
				v2: version.Version{
					Major:               7,
					Minor:               1,
					Maintenance:         "3",
					Platform:            "D",
					PlatformMinor:       0,
					PlatformMaintenance: "3",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := version.Version{
				Major:               tt.fields.Major,
				Minor:               tt.fields.Minor,
				Maintenance:         tt.fields.Maintenance,
				Platform:            tt.fields.Platform,
				PlatformMinor:       tt.fields.PlatformMinor,
				PlatformMaintenance: tt.fields.PlatformMaintenance,
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
		Major               int
		Minor               int
		Maintenance         string
		Platform            string
		PlatformMinor       int
		PlatformMaintenance string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "6.2(8b)",
			fields: fields{
				Major:       6,
				Minor:       2,
				Maintenance: "8b",
			},
			want: "6.2(8b)",
		},
		{
			name: "7.3(0)DX(1)",
			fields: fields{
				Major:               7,
				Minor:               3,
				Maintenance:         "0",
				Platform:            "DX",
				PlatformMinor:       0,
				PlatformMaintenance: "1",
			},
			want: "7.3(0)DX(1)",
		},
		{
			name: "7.1(3)N1(2)",
			fields: fields{
				Major:               7,
				Minor:               1,
				Maintenance:         "3",
				Platform:            "N",
				PlatformMinor:       1,
				PlatformMaintenance: "2",
			},
			want: "7.1(3)N1(2)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := version.Version{
				Major:               tt.fields.Major,
				Minor:               tt.fields.Minor,
				Maintenance:         tt.fields.Maintenance,
				Platform:            tt.fields.Platform,
				PlatformMinor:       tt.fields.PlatformMinor,
				PlatformMaintenance: tt.fields.PlatformMaintenance,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
