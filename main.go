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

// splitServerFeature - Split server and features info to couple of string
func splitServerFeature(slice []string) (serverInfo string, featuresInfo string) {
	return slice[0], slice[1]
}

func getLicenseServersInfo(flexlmStats string) json {
	var serversFullInfo []string
	var json json
	// Split info by server
	serversFullInfo = splitdata(flexlmStats, serverSeparator)

	// Split data for server info and feature usage info
	for i, data := range serversFullInfo {
		slice := splitdata(data, featureUsageSeparator)
		server, feat := splitServerFeature(slice)
		json.LicenseServer[i] = parseServerInfo(server)
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
	// TODO Split data by features by "Users of ".

	// TODO split feature data and active users data
	return featuresUsage
}

func parseFeatureData(featureData string) featureUsage {
	var featureUsage featureUsage

	return featureUsage
}

func getUsersData(usersData string) []users {
	var users []users
	// TODO split data by users

	// TODO call parser for each user
	return users
}

func parseUserData(userData string) users {
	var users users
	// TODO parse user data to struct users
	return users
}
