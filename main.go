package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	exitOK                = 0
	exitFail              = 1
	serverSeparator       = "----------------------------------------------------------------------------"
	featureUsageSeparator = "Feature usage info:"
	featuresSeparator     = "Users of "
)

type json struct {
	LicenseServer []licenseServer `json:"license_server"`
}
type licenseServer struct {
	Server        string         `json:"server"`
	ServerStatus  string         `json:"server_status"`
	ServerVersion string         `json:"server_version"`
	Vendor        string         `json:"vendor"`
	VendorStatus  string         `json:"vendor_status"`
	VendorVersion string         `json:"vendor_version"`
	FeatureUsage  []featureUsage `json:"feature_usage"`
}
type featureUsage struct {
	Feature    string `json:"feature"`
	IssuedLics string `json:"issued_lics"`
	UsedLics   string `json:"used_lics"`
	Users      []users
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
	var JSON json
	for scanner.Scan() {
		flexlmStats = flexlmStats + scanner.Text() + "\n"
	}
	JSON = getLicenseServersInfo(flexlmStats)
	// TODO !!! marshal JSON struct to JSON object

	fmt.Fprint(stdout, JSON)

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
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
	return slice[0], slice[1]
}

// splitFeatureUsers - Split feature, feature details and users data to separate strings
func splitFeatureUsers(slice []string) (featureInfo, featureDetails, usersInfo string) {
	return slice[0], slice[1], slice[2]
}

func getLicenseServersInfo(flexlmStats string) json {
	var serversFullInfo []string
	var json json
	// Split info by server
	serversFullInfo = splitdata(flexlmStats, serverSeparator)

	// Split data for server info and feature usage info
	for i, data := range serversFullInfo {
		slice := splitdata(data, featureUsageSeparator)
		server, feat := splitTwoValues(slice)
		json.LicenseServer = append(json.LicenseServer, parseServerInfo(server))
		json.LicenseServer[i].FeatureUsage = getFeatureData(feat)

	}
	return json
}

// parseServerInfo - parse server info block
func parseServerInfo(serverInfo string) licenseServer {
	var licenseServer licenseServer

	return licenseServer
}

// getFeatureData - get data from featire usage info
func getFeatureData(flexlmStats string) []featureUsage {
	var featuresUsage []featureUsage
	var features []string
	// Split data by features. String "Users of " will be deleted.
	features = splitdata(flexlmStats, featuresSeparator)

	// feture with users (data) will be processed
	for i, data := range features {
		// TODO split feature data and active users data
		slice := splitdata(data, "\n\n")
		featureInfo, _, usersInfo := splitFeatureUsers(slice)
		featuresUsage = append(featuresUsage, parseFeatureData(featureInfo))
		featuresUsage[i].Users = getUsersData(usersInfo)

	}

	return featuresUsage
}

func parseFeatureData(featureData string) featureUsage {
	var featureUsage featureUsage
	var i1, i2 int
	featureData = strings.Trim(featureData, "\n ")
	// Get feature number
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
	usersData = strings.Trim(usersData, "\n ")
	// split usersData by users
	usersSlice = splitdata(usersData, "\n    ")

	// call parser for each user
	for _, data := range usersSlice {
		users = append(users, parseUserData(data))
	}
	return users
}

func parseUserData(userData string) users {
	var users users
	var serverHost, serverPort string
	userData = strings.Trim(userData, "\n ")
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
