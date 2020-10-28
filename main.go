package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	exitOK                = 0
	exitFail              = 1
	serverSeparator       = "----------------------------------------------------------------------------"
	serverInfoDelimiter   = "[Detecting lmgrd processes...]\n"
	featureUsageSeparator = "Feature usage info:"
	featuresSeparator     = "Users of "
)

type jsonOUT struct {
	LicenseServer []licenseServer `json:"license_server"`
}
type licenseServer struct {
	Server        string         `json:"server,omitempty"`
	ServerStatus  string         `json:"server_status,omitempty"`
	ServerVersion string         `json:"server_version,omitempty"`
	Vendor        string         `json:"vendor,omitempty"`
	VendorStatus  string         `json:"vendor_status,omitempty"`
	VendorVersion string         `json:"vendor_version,omitempty"`
	FeatureUsage  []featureUsage `json:"feature_usage,omitempty"`
}
type featureUsage struct {
	Feature    string  `json:"feature"`
	IssuedLics string  `json:"issued_lics"`
	UsedLics   string  `json:"used_lics"`
	Users      []users `json:"users,omitempty"`
}
type users struct {
	Userid         string `json:"userid"`
	Host           string `json:"host"`
	Display        string `json:"display"`
	FeatureVersion string `json:"feature_version"`
	ServerHost     string `json:"server_host"`
	ServerPort     string `json:"server_port"`
}

func main() {
	err := run(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
	os.Exit(exitOK)
}

func run(stdin io.Reader, stdout io.Writer) error {
	scanner := bufio.NewScanner(stdin)
	var flexlmStats string
	var preJSON jsonOUT
	for scanner.Scan() {
		flexlmStats = flexlmStats + scanner.Text() + "\n"
	}
	preJSON = getLicenseServersInfo(flexlmStats)
	// TODO !!! marshal JSON struct to JSON object
	JSONtoOUT, err := createJSON(preJSON)
	if err != nil {
		return err
	}
	fmt.Fprint(stdout, string(JSONtoOUT))

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func createJSON(s jsonOUT) (string, error) {
	jsonM, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonM), nil
}

func splitdata(s string, sep string) []string {
	if sep == "" {
		var onestring []string
		onestring = append(onestring, s)
		return onestring
	}
	return strings.Split(s, sep)
}

// TODO Refactor splitTwoValues. Function have to be replaced by method for structs
// splitTwoValues - Split server and features info to couple of string
func splitTwoValues(slice []string) (v1, v2 string) {
	switch len(slice) {
	case 0:
		return "", ""
	case 1:
		return slice[0], ""
	default:
		return slice[0], slice[1]
	}

}

// splitFeatureUsers - Split feature, feature details and users data to separate strings
func splitFeatureUsers(slice []string) (featureInfo, featureDetails, usersInfo string) {
	switch len(slice) {
	case 0:
		return "", "", ""
	case 1:
		return slice[0], "", ""
	case 2:
		return slice[0], slice[1], ""
	default:
		return slice[0], slice[1], slice[2]
	}
}

func getLicenseServersInfo(flexlmStats string) jsonOUT {
	var serversFullInfo []string
	var jsonOUT jsonOUT
	// Split info by server
	serversFullInfo = splitdata(flexlmStats, serverSeparator)

	// Split data for server info and feature usage info
	for i, data := range serversFullInfo {
		slice := splitdata(data, featureUsageSeparator)
		server, feat := splitTwoValues(slice)
		jsonOUT.LicenseServer = append(jsonOUT.LicenseServer, parseServerInfo(server))
		if len(feat) > 0 {
			jsonOUT.LicenseServer[i].FeatureUsage = getFeatureData(feat)
		}

	}
	return jsonOUT
}

// parseServerInfo - parse server info block
func parseServerInfo(serverInfo string) licenseServer {
	var licenseServer licenseServer
	var i1, i2 int
	// Trim unnecessary data
	serverInfo = strings.Trim(serverInfo, "\n ")
	i1 = strings.Index(serverInfo, serverInfoDelimiter)
	if i1 >= 0 {
		i1 = i1 + len(serverInfoDelimiter)
		serverInfo = serverInfo[i1:]
	}

	// Split data by strings
	slice := strings.Split(serverInfo, "\n\n")
	if len(slice[0]) == 0 {
		slice = slice[1:]
	}

	// Get server name
	i1 = strings.Index(slice[0], ": ") + 2
	i2 = strings.Index(slice[0], "\n")
	licenseServer.Server = slice[0][i1:i2]
	// Get server status
	i1 = strings.Index(slice[1], ": license server ") + len(": license server ")
	i2 = i1 + strings.Index(slice[1][i1:], " ")
	licenseServer.ServerStatus = slice[1][i1:i2]
	// Get server version
	i1 = strings.LastIndex(slice[1], " ") + 1
	licenseServer.ServerVersion = slice[1][i1:]
	// Get vendor data
	slice[3] = strings.Trim(slice[3], "\n ")
	vendorData := strings.Split(slice[3], " ")
	for i, v := range vendorData {
		vendorData[i] = strings.Trim(v, ": \t")
	}
	if len(vendorData) == 3 {
		licenseServer.Vendor = vendorData[0]
		licenseServer.VendorStatus = vendorData[1]
		licenseServer.VendorVersion = vendorData[2]
	}

	return licenseServer
}

// getFeatureData - get data from featire usage info
func getFeatureData(flexlmStats string) []featureUsage {
	var featuresUsage []featureUsage
	var features []string
	var featureInfo, usersInfo string
	flexlmStats = strings.Trim(flexlmStats, "\n \t")
	// Split data by features. String "Users of " will be deleted.
	features = splitdata(flexlmStats, featuresSeparator)
	if len(features[0]) == 0 {
		features = features[1:]
	}

	// feture with users (data) will be processed
	for i, data := range features {
		data = strings.Trim(data, "\n \t")
		// split feature data and active users data
		slice := splitdata(data, "\n\n")
		if len(slice) > 1 {
			featureInfo, _, usersInfo = splitFeatureUsers(slice)
		}
		if len(slice) == 1 {
			featureInfo = slice[0]
			usersInfo = ""
		}
		featuresUsage = append(featuresUsage, parseFeatureData(featureInfo))
		if len(usersInfo) > 0 {
			featuresUsage[i].Users = getUsersData(usersInfo)
		}

	}

	return featuresUsage
}

func parseFeatureData(featureData string) featureUsage {
	var featureUsage featureUsage
	var i1, i2 int
	featureData = strings.Trim(featureData, "\n \t")
	// Get feature number
	if len(featureData) == 0 {
		return featureUsage
	}
	i2 = strings.Index(featureData, ":")
	// TODO check -1 return
	featureUsage.Feature = featureData[i1:i2]
	// Issued licenses
	i1 = i2
	i1 = i1 + strings.Index(featureData[i1:], "(Total of ") + len("(Total of ")
	i2 = i1 + strings.Index(featureData[i1:], " license")
	featureUsage.IssuedLics = featureData[i1:i2]
	// Used license
	i1 = i2 + strings.Index(featureData[i2:], ";  Total of ") + len(";  Total of ")
	i2 = i1 + strings.Index(featureData[i1:], " license")
	featureUsage.UsedLics = featureData[i1:i2]
	return featureUsage
}

func getUsersData(usersData string) []users {
	var users []users

	var usersSlice []string
	// Cut leading and trailing \n and spaces
	usersData = strings.Trim(usersData, "\n \t")
	// split usersData by users
	usersSlice = splitdata(usersData, "\n")

	// call parser for each user
	for _, data := range usersSlice {
		users = append(users, parseUserData(data))
	}
	return users
}

func parseUserData(userData string) users {
	var users users
	var serverHost, serverPort string
	userData = strings.Trim(userData, "\n \t")
	if len(userData) == 0 {
		return users
	}
	//parse user data to struct users
	// TODO check -1 return
	slice := strings.Split(userData, " ")
	// Cuts special simbols
	for i, v := range slice {
		slice[i] = strings.Trim(v, "(),")
	}
	users.Userid = slice[0]
	users.Host = slice[1]
	users.Display = slice[2]
	serverHost, serverPort = splitTwoValues(strings.Split(slice[4], "/"))
	users.ServerHost = serverHost
	users.ServerPort = serverPort

	return users
}
