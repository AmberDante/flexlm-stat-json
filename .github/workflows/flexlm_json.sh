#!/bin/bash
# Get FlexLM stats for zabbix LLD

IFS=$'\n'
JSON="{\"data\":["
SEP=""
flexToJSON=/usr/lib/zabbix/externalscripts/flexlm-stat-json

# Server discovery
# Key: srvdiscovery
if [ "$2" = "srvdiscovery" ]
then
get=`zabbix_get -s $1 -k system.run["lmutil lmstat -lm -vd"] | $flexToJSON | sed 's/{"license_server":/{"data":/gm' | sed 's/{"server":/{"{#SERVER}":/gm' | sed 's/"vendor":/"{#VENDOR}":/gm'`
echo $get
fi

# Features discovery
# Key: featdiscovery
if [ "$2" = "featdiscovery" ]
then
get=`zabbix_get -s $1 -k system.run["lmutil lmstat -a"] | grep "Users of" | sed 's/\(Users of \|: .*\)//g'`
for feature in $get
do
JSON=$JSON"$SEP{\"{#FEATURE}\":\"$feature\"}"
SEP=","
done
JSON=$JSON"]}"
echo $JSON
fi

# Get all data in JSON
# Key: getlics
if [ "$2" = "getlics" ]
then
zabbix_get -s $1 -k system.run["lmutil lmstat -a"] | $flexToJSON
fi
