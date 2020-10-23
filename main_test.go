package main

import (
	"reflect"
	"testing"
)

func Test_getLicenseServersInfo(t *testing.T) {
	type args struct {
		flexlmStats string
	}
	tests := []struct {
		name string
		args args
		want json
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLicenseServersInfo(tt.args.flexlmStats); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLicenseServersInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseServerInfo(t *testing.T) {
	type args struct {
		serverInfo string
	}
	tests := []struct {
		name string
		args args
		want licenseServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseServerInfo(tt.args.serverInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseServerInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFeatureData(t *testing.T) {
	type args struct {
		flexlmStats string
	}
	tests := []struct {
		name string
		args args
		want []featureUsage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFeatureData(tt.args.flexlmStats); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFeatureData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFeatureData(t *testing.T) {
	type args struct {
		featureData string
	}
	tests := []struct {
		name string
		args args
		want featureUsage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFeatureData(tt.args.featureData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFeatureData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUsersData(t *testing.T) {
	type args struct {
		usersData string
	}
	tests := []struct {
		name string
		args args
		want []users
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUsersData(tt.args.usersData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUsersData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseUserData(t *testing.T) {
	type args struct {
		userData string
	}
	tests := []struct {
		name string
		args args
		want users
	}{
		// TODO: Add test cases.
		{
			name: "user 1",
			args: args{`    yablontseva 03-VSNIKOLAEVA THINKPAD-T530 (v1.000) (iss.samba.gazpromproject.ru/27000 11645), start Mon 10/12 9:02`},
			want: users{
				Userid:     "yablontseva",
				Host:       "03-VSNIKOLAEVA",
				Display:    "THINKPAD-T530",
				ServerHost: "iss.samba.gazpromproject.ru",
				ServerPort: "27000",
			},
		},
		{
			name: "user 2",
			args: args{`1 GM-007028 GM-007028 (v1.000) (iss.samba.gazpromproject.ru/27000 37583), start Thu 10/1 11:44  (linger: 4263164 / 5487240)`},
			want: users{
				Userid:     "1",
				Host:       "GM-007028",
				Display:    "GM-007028",
				ServerHost: "iss.samba.gazpromproject.ru",
				ServerPort: "27000",
			},
		},
		{
			name: "user 3",
			args: args{`
			eprus SPB-00-001686 spb-00-001686 (v1.000) (iss.samba.gazpromproject.ru/27000 43953), start Mon 9/28 12:43`},
			want: users{
				Userid:     "eprus",
				Host:       "SPB-00-001686",
				Display:    "spb-00-001686",
				ServerHost: "iss.samba.gazpromproject.ru",
				ServerPort: "27000",
			},
		},
		{
			name: "user 4",
			args: args{`    eprus SPB-00-001686 spb-00-001686 (v1.000) (iss.samba.gazpromproject.ru/27000 43953), start Mon 9/28 12:43`},
			want: users{
				Userid:     "eprus",
				Host:       "SPB-00-001686",
				Display:    "spb-00-001686",
				ServerHost: "iss.samba.gazpromproject.ru",
				ServerPort: "27000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseUserData(tt.args.userData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseUserData() = %v, want %v", got, tt.want)
			}
		})
	}
}