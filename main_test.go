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
		want jsonOUT
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{`lmutil - Copyright (c) 1989-2018 Flexera. All Rights Reserved.
Flexible License Manager status on Thu 10/15/2020 15:45

[Detecting lmgrd processes...]
License server status: 27000@iss.samba.gazpromproject.ru
    License file(s) on iss.samba.gazpromproject.ru: F:\Autodesk\Network License Manager\iss.samba.gazpromproject.ru.lic:

iss.samba.gazpromproject.ru: license server UP (MASTER) v11.16.2

Vendor daemon status (on iss.samba.gazpromproject.ru):

  adskflex: UP v11.16.2
Feature usage info:

Users of 86815AECCOL_T_F:  (Total of 240 licenses issued;  Total of 7 licenses in use)

  "86815AECCOL_T_F" v1.000, vendor: adskflex, expiry: 15-dec-2022
  vendor_string: commercial:extendable
  floating license

    eprus SPB-00-001686 spb-00-001686 (v1.000) (iss.samba.gazpromproject.ru/27000 43953), start Mon 9/28 12:43
    1 GM-007028 GM-007028 (v1.000) (iss.samba.gazpromproject.ru/27000 37583), start Thu 10/1 11:44  (linger: 4263164 / 5487240)

Users of 85788BDSADV_F:  (Total of 16 licenses issued;  Total of 0 licenses in use)
`},
			want: jsonOUT{
				LicenseServer: []licenseServer{
					licenseServer{
						Server:        "27000@iss.samba.gazpromproject.ru",
						ServerStatus:  "UP",
						ServerVersion: "v11.16.2",
						Vendor:        "adskflex",
						VendorStatus:  "UP",
						VendorVersion: "v11.16.2",
						FeatureUsage: []featureUsage{
							featureUsage{
								Feature:    "86815AECCOL_T_F",
								IssuedLics: "240",
								UsedLics:   "7",
								Users: []users{
									users{
										Userid:     "eprus",
										Host:       "SPB-00-001686",
										Display:    "spb-00-001686",
										ServerHost: "iss.samba.gazpromproject.ru",
										ServerPort: "27000",
									},
									users{
										Userid:     "1",
										Host:       "GM-007028",
										Display:    "GM-007028",
										ServerHost: "iss.samba.gazpromproject.ru",
										ServerPort: "27000",
									},
								},
							},
							featureUsage{
								Feature:    "85788BDSADV_F",
								IssuedLics: "16",
								UsedLics:   "0",
							},
						},
					},
				},
			},
		},
		{
			name: "test2",
			args: args{`lmutil - Copyright (c) 1989-2018 Flexera. All Rights Reserved.
Flexible License Manager status on Thu 10/15/2020 15:45

[Detecting lmgrd processes...]
License server status: 27000@iss.samba.gazpromproject.ru
    License file(s) on iss.samba.gazpromproject.ru: F:\Autodesk\Network License Manager\iss.samba.gazpromproject.ru.lic:

iss.samba.gazpromproject.ru: license server UP (MASTER) v11.16.2

Vendor daemon status (on iss.samba.gazpromproject.ru):

  adskflex: UP v11.16.2
Feature usage info:

Users of 86815AECCOL_T_F:  (Total of 240 licenses issued;  Total of 7 licenses in use)

  "86815AECCOL_T_F" v1.000, vendor: adskflex, expiry: 15-dec-2022
  vendor_string: commercial:extendable
  floating license

    eprus SPB-00-001686 spb-00-001686 (v1.000) (iss.samba.gazpromproject.ru/27000 43953), start Mon 9/28 12:43
    1 GM-007028 GM-007028 (v1.000) (iss.samba.gazpromproject.ru/27000 37583), start Thu 10/1 11:44  (linger: 4263164 / 5487240)

Users of 85788BDSADV_F:  (Total of 16 licenses issued;  Total of 0 licenses in use)



----------------------------------------------------------------------------
License server status: 27002@iss.samba.gazpromproject.ru
    License file(s) on iss.samba.gazpromproject.ru: F:\MapInfo\License Server\MILICSERVER.lic:

iss.samba.gazpromproject.ru: license server UP v11.13.0

Vendor daemon status (on iss):

   unisw20: UP v11.13.0
Feature usage info:

Users of MapInfo_License_Server:  (Total of 1 license issued;  Total of 0 licenses in use)

`},
			want: jsonOUT{
				LicenseServer: []licenseServer{
					licenseServer{
						Server:        "27000@iss.samba.gazpromproject.ru",
						ServerStatus:  "UP",
						ServerVersion: "v11.16.2",
						Vendor:        "adskflex",
						VendorStatus:  "UP",
						VendorVersion: "v11.16.2",
						FeatureUsage: []featureUsage{
							featureUsage{
								Feature:    "86815AECCOL_T_F",
								IssuedLics: "240",
								UsedLics:   "7",
								Users: []users{
									users{
										Userid:     "eprus",
										Host:       "SPB-00-001686",
										Display:    "spb-00-001686",
										ServerHost: "iss.samba.gazpromproject.ru",
										ServerPort: "27000",
									},
									users{
										Userid:     "1",
										Host:       "GM-007028",
										Display:    "GM-007028",
										ServerHost: "iss.samba.gazpromproject.ru",
										ServerPort: "27000",
									},
								},
							},
							featureUsage{
								Feature:    "85788BDSADV_F",
								IssuedLics: "16",
								UsedLics:   "0",
							},
						},
					},
					licenseServer{
						Server:        "27002@iss.samba.gazpromproject.ru",
						ServerStatus:  "UP",
						ServerVersion: "v11.13.0",
						Vendor:        "unisw20",
						VendorStatus:  "UP",
						VendorVersion: "v11.13.0",
						FeatureUsage: []featureUsage{
							featureUsage{
								Feature:    "MapInfo_License_Server",
								IssuedLics: "1",
								UsedLics:   "0",
							},
						},
					},
				},
			},
		},
		{
			name: "test lmstat -vd",
			args: args{`lmutil - Copyright (c) 1989-2018 Flexera. All Rights Reserved.
Wed 10/28/2020 10:48

Vendor daemon status (on iss.samba.gazpromproject.ru):

  adskflex: UP v11.16.2


----------------------------------------------------------------------------
Vendor daemon status (on iss.samba.gazpromproject.ru):

   unisw20: UP v11.13.0
`},
			want: jsonOUT{
				LicenseServer: []licenseServer{
					licenseServer{
						Vendor:        "adskflex",
						VendorStatus:  "UP",
						VendorVersion: "v11.16.2",
					},
					licenseServer{

						Vendor:        "unisw20",
						VendorStatus:  "UP",
						VendorVersion: "v11.13.0",
					},
				},
			},
		},
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
		{
			name: "server string empty",
			args: args{``},
			want: licenseServer{},
		},
		{
			name: "server string only server info",
			args: args{`lmutil - Copyright (c) 1989-2018 Flexera. All Rights Reserved.
Flexible License Manager status on Thu 10/15/2020 15:45

[Detecting lmgrd processes...]
License server status: 27000@iss.samba.gazpromproject.ru
	License file(s) on iss.samba.gazpromproject.ru: F:\Autodesk\Network License Manager\iss.samba.gazpromproject.ru.lic:

iss.samba.gazpromproject.ru: license server UP (MASTER) v11.16.2

`},
			want: licenseServer{
				Server:        "27000@iss.samba.gazpromproject.ru",
				ServerStatus:  "UP",
				ServerVersion: "v11.16.2",
			},
		},
		{
			name: "server string 1",
			args: args{`lmutil - Copyright (c) 1989-2018 Flexera. All Rights Reserved.
Flexible License Manager status on Thu 10/15/2020 15:45

[Detecting lmgrd processes...]
License server status: 27000@iss.samba.gazpromproject.ru
	License file(s) on iss.samba.gazpromproject.ru: F:\Autodesk\Network License Manager\iss.samba.gazpromproject.ru.lic:

iss.samba.gazpromproject.ru: license server UP (MASTER) v11.16.2

Vendor daemon status (on iss.samba.gazpromproject.ru):

  adskflex: UP v11.16.2`},
			want: licenseServer{
				Server:        "27000@iss.samba.gazpromproject.ru",
				ServerStatus:  "UP",
				ServerVersion: "v11.16.2",
				Vendor:        "adskflex",
				VendorStatus:  "UP",
				VendorVersion: "v11.16.2",
			},
		},
		{
			name: "server string 2",
			args: args{`License server status: 27002@iss.samba.gazpromproject.ru
			License file(s) on iss.samba.gazpromproject.ru: F:\MapInfo\License Server\MILICSERVER.lic:

iss.samba.gazpromproject.ru: license server UP v11.13.0

Vendor daemon status (on iss):

   unisw20: UP v11.13.0`},
			want: licenseServer{
				Server:        "27002@iss.samba.gazpromproject.ru",
				ServerStatus:  "UP",
				ServerVersion: "v11.13.0",
				Vendor:        "unisw20",
				VendorStatus:  "UP",
				VendorVersion: "v11.13.0",
			},
		},
		{
			name: "processing lmstat -vd",
			args: args{`lmutil - Copyright (c) 1989-2018 Flexera. All Rights Reserved.
Wed 10/28/2020 10:48

Vendor daemon status (on iss.samba.gazpromproject.ru):

  adskflex: UP v11.16.2


----------------------------------------------------------------------------
Vendor daemon status (on iss.samba.gazpromproject.ru):

   unisw20: UP v11.13.0
`},
			want: licenseServer{
				Vendor:        "adskflex",
				VendorStatus:  "UP",
				VendorVersion: "v11.16.2",
			},
		},
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
		{
			name: "get feature with 1 active user",
			args: args{`Users of 87252IDSP_2020_0F:  (Total of 13 licenses issued;  Total of 1 license in use)

"87252IDSP_2020_0F" v1.000, vendor: adskflex, expiry: permanent(no expiration date)
vendor_string: commercial:permanent
floating license

  47011 UII434-NB UII434-NB (v1.0) (iss.samba.gazpromproject.ru/27000 46192), start Thu 10/15 11:20

`},
			want: []featureUsage{
				featureUsage{
					Feature:    "87252IDSP_2020_0F",
					IssuedLics: "13",
					UsedLics:   "1",
					Users: []users{
						users{
							Userid:     "47011",
							Host:       "UII434-NB",
							Display:    "UII434-NB",
							ServerHost: "iss.samba.gazpromproject.ru",
							ServerPort: "27000",
						},
					},
				},
			},
		},
		{
			name: "get several features",
			args: args{`Users of 87089AMECH_PP_2019_0F:  (Total of 240 licenses issued;  Total of 0 licenses in use)

Users of 86839AMECH_PP_2018_0F:  (Total of 240 licenses issued;  Total of 2 licenses in use)

  "86839AMECH_PP_2018_0F" v1.000, vendor: adskflex, expiry: 15-dec-2022
  vendor_string: commercial:extendable
  floating license

	1 GM-007028 GM-007028 (v1.0) (iss.samba.gazpromproject.ru/27000 37683), start Thu 10/1 11:44  (linger: 4263164 / 5487240)
	58000 SKIA011 DESKTOP-0TO69FR (v1.0) (iss.samba.gazpromproject.ru/27000 389), start Wed 10/14 9:04

Users of 86627AMECH_PP_2017_0F:  (Total of 240 licenses issued;  Total of 0 licenses in use)
`},
			want: []featureUsage{
				featureUsage{
					Feature:    "87089AMECH_PP_2019_0F",
					IssuedLics: "240",
					UsedLics:   "0",
				},
				featureUsage{
					Feature:    "86839AMECH_PP_2018_0F",
					IssuedLics: "240",
					UsedLics:   "2",
					Users: []users{
						users{
							Userid:     "1",
							Host:       "GM-007028",
							Display:    "GM-007028",
							ServerHost: "iss.samba.gazpromproject.ru",
							ServerPort: "27000",
						},
						users{
							Userid:     "58000",
							Host:       "SKIA011",
							Display:    "DESKTOP-0TO69FR",
							ServerHost: "iss.samba.gazpromproject.ru",
							ServerPort: "27000",
						},
					},
				},
				featureUsage{
					Feature:    "86627AMECH_PP_2017_0F",
					IssuedLics: "240",
					UsedLics:   "0",
				},
			},
		},
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
		{
			name: "feature 1",
			args: args{`86468IDSP_2016_0F:  (Total of 13 licenses issued;  Total of 0 licenses in use)`},
			want: featureUsage{
				Feature:    "86468IDSP_2016_0F",
				IssuedLics: "13",
				UsedLics:   "0",
			},
		},
		{
			name: "feature 2",
			args: args{`87252IDSP_2020_0F:  (Total of 13 licenses issued;  Total of 1 license in use)`},
			want: featureUsage{
				Feature:    "87252IDSP_2020_0F",
				IssuedLics: "13",
				UsedLics:   "1",
			},
		},
		{
			name: "feature 3 Total of 1 license",
			args: args{`MapInfo_License_Server:  (Total of 1 license issued;  Total of 0 licenses in use)`},
			want: featureUsage{
				Feature:    "MapInfo_License_Server",
				IssuedLics: "1",
				UsedLics:   "0",
			},
		},
		{
			name: "empty feature",
			args: args{``},
			want: featureUsage{
				Feature:    "",
				IssuedLics: "",
				UsedLics:   "",
			},
		},
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
		{
			name: "one user",
			args: args{`    47011 UII434-NB UII434-NB (v1.000) (iss.samba.gazpromproject.ru/27000 46292), start Thu 10/15 11:20`},
			want: []users{
				users{
					Userid:     "47011",
					Host:       "UII434-NB",
					Display:    "UII434-NB",
					ServerHost: "iss.samba.gazpromproject.ru",
					ServerPort: "27000",
				},
			},
		},
		{
			name: "two users with newlines",
			args: args{`
			dgridnev SPB-00-005001 spb-00-005001 (v1.000) (iss.samba.gazpromproject.ru/27000 39867), start Thu 10/15 9:05
			6325 OAPIU036 OAPIU036 (v1.000) (iss.samba.gazpromproject.ru/27000 13856), start Thu 10/15 15:35
		`},
			want: []users{
				users{
					Userid:     "dgridnev",
					Host:       "SPB-00-005001",
					Display:    "spb-00-005001",
					ServerHost: "iss.samba.gazpromproject.ru",
					ServerPort: "27000",
				},
				users{
					Userid:     "6325",
					Host:       "OAPIU036",
					Display:    "OAPIU036",
					ServerHost: "iss.samba.gazpromproject.ru",
					ServerPort: "27000",
				},
			},
		},
		{
			name: "two users with newlines",
			args: args{`dgridnev SPB-00-005001 spb-00-005001 (v1.000) (iss.samba.gazpromproject.ru/27000 39867), start Thu 10/15 9:05
			6325 OAPIU036 OAPIU036 (v1.000) (iss.samba.gazpromproject.ru/27000 13856), start Thu 10/15 15:35
		`},
			want: []users{
				users{
					Userid:     "dgridnev",
					Host:       "SPB-00-005001",
					Display:    "spb-00-005001",
					ServerHost: "iss.samba.gazpromproject.ru",
					ServerPort: "27000",
				},
				users{
					Userid:     "6325",
					Host:       "OAPIU036",
					Display:    "OAPIU036",
					ServerHost: "iss.samba.gazpromproject.ru",
					ServerPort: "27000",
				},
			},
		},
		{
			name: "empty string",
			args: args{``},
			want: []users{
				users{},
			},
		},
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
			name: "empty data",
			args: args{" "},
			want: users{
				Userid:     "",
				Host:       "",
				Display:    "",
				ServerHost: "",
				ServerPort: "",
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

func Test_splitFeatureUsers(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name               string
		args               args
		wantFeatureInfo    string
		wantFeatureDetails string
		wantUsersInfo      string
	}{
		// TODO: Add test cases.
		{
			name:               "split empty slice",
			args:               args{[]string{}},
			wantFeatureInfo:    "",
			wantFeatureDetails: "",
			wantUsersInfo:      "",
		},
		{
			name:               "split slice with one element",
			args:               args{[]string{"one"}},
			wantFeatureInfo:    "one",
			wantFeatureDetails: "",
			wantUsersInfo:      "",
		},
		{
			name:               "split slice with two elements",
			args:               args{[]string{"one", "two"}},
			wantFeatureInfo:    "one",
			wantFeatureDetails: "two",
			wantUsersInfo:      "",
		},
		{
			name:               "split slice with three elements",
			args:               args{[]string{"one", "two", "three"}},
			wantFeatureInfo:    "one",
			wantFeatureDetails: "two",
			wantUsersInfo:      "three",
		},
		{
			name:               "split slice with more elements",
			args:               args{[]string{"one", "two", "three", "4", "5"}},
			wantFeatureInfo:    "one",
			wantFeatureDetails: "two",
			wantUsersInfo:      "three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFeatureInfo, gotFeatureDetails, gotUsersInfo := splitFeatureUsers(tt.args.slice)
			if gotFeatureInfo != tt.wantFeatureInfo {
				t.Errorf("splitFeatureUsers() gotFeatureInfo = %v, want %v", gotFeatureInfo, tt.wantFeatureInfo)
			}
			if gotFeatureDetails != tt.wantFeatureDetails {
				t.Errorf("splitFeatureUsers() gotFeatureDetails = %v, want %v", gotFeatureDetails, tt.wantFeatureDetails)
			}
			if gotUsersInfo != tt.wantUsersInfo {
				t.Errorf("splitFeatureUsers() gotUsersInfo = %v, want %v", gotUsersInfo, tt.wantUsersInfo)
			}
		})
	}
}

func Test_splitSliceToStrings(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name   string
		args   args
		wantV1 string
		wantV2 string
		wantV3 string
		wantV4 string
		wantV5 string
	}{
		// TODO: Add test cases.
		{
			name:   "split empty slice",
			args:   args{[]string{}},
			wantV1: "",
			wantV2: "",
			wantV3: "",
			wantV4: "",
			wantV5: "",
		},
		{
			name:   "split slice with one element",
			args:   args{[]string{"one"}},
			wantV1: "one",
			wantV2: "",
			wantV3: "",
			wantV4: "",
			wantV5: "",
		},
		{
			name:   "split slice with two elements",
			args:   args{[]string{"one", "two"}},
			wantV1: "one",
			wantV2: "two",
			wantV3: "",
			wantV4: "",
			wantV5: "",
		},
		{
			name:   "split slice with three elements",
			args:   args{[]string{"one", "two", "three"}},
			wantV1: "one",
			wantV2: "two",
			wantV3: "three",
			wantV4: "",
			wantV5: "",
		},
		{
			name:   "split slice with four elements",
			args:   args{[]string{"one", "two", "", "four"}},
			wantV1: "one",
			wantV2: "two",
			wantV3: "",
			wantV4: "four",
			wantV5: "",
		},
		{
			name:   "split slice with 5 elements",
			args:   args{[]string{"one", "two", "3", "four", "five"}},
			wantV1: "one",
			wantV2: "two",
			wantV3: "3",
			wantV4: "four",
			wantV5: "five",
		},
		{
			name:   "split slice with more elements",
			args:   args{[]string{"one", "two", "3", "four", "five", "six"}},
			wantV1: "one",
			wantV2: "two",
			wantV3: "3",
			wantV4: "four",
			wantV5: "five",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV1, gotV2, gotV3, gotV4, gotV5 := splitSliceToStrings(tt.args.slice)
			if gotV1 != tt.wantV1 {
				t.Errorf("splitSliceToStrings() gotV1 = %v, want %v", gotV1, tt.wantV1)
			}
			if gotV2 != tt.wantV2 {
				t.Errorf("splitSliceToStrings() gotV2 = %v, want %v", gotV2, tt.wantV2)
			}
			if gotV3 != tt.wantV3 {
				t.Errorf("splitSliceToStrings() gotV3 = %v, want %v", gotV3, tt.wantV3)
			}
			if gotV4 != tt.wantV4 {
				t.Errorf("splitSliceToStrings() gotV4 = %v, want %v", gotV4, tt.wantV4)
			}
			if gotV5 != tt.wantV5 {
				t.Errorf("splitSliceToStrings() gotV5 = %v, want %v", gotV5, tt.wantV5)
			}
		})
	}
}
