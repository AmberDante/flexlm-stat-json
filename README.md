# flexlm-stat-json
Convert lmutil lmstat -a to JSON for usage in zabbix stats

It can help you to get all necessary statistics of FlexLM with one Item(Servers, features with issued and used lics, active users). Then zabbix can parse it with **JSONPath** to Dependent Items. It's mush more faster then use hundreds of External Items. (read the Warnning on [zabbix docs](ttps://www.zabbix.com/documentation/4.0/manual/config/items/itemtypes/external))


## How to use

copy External Script **flexlm_json.sh** to `/usr/lib/zabbix/externalscripts`. It's default folder for zabbix external scripts. It can be defined by zabbix config file.  
copy **flexlm-stat-json** to `/usr/lib/zabbix/externalscripts` or to the host where FlexLM runs. If you copy **flexlm-stat-json** to the host where FlexLM runs you have to rewrite **flexlm_json.sh** for that changes.  

External script which take data to STDIN of **flexlm-stat-json** can look like thath examples


Example 1:
<pre>zabbix-get -s 1$ -k system.run["lmutil lmstat -a"] | flexlm-stat-json</pre>

Example 2:
<pre>zabbix-get -s 1$ -k system.run["lmutil lmstat -a | flexlm-stat-json"]</pre>

On the host where FlexLM runs you have to set PATH to **flexlm-stat-json** or use absolute path to it

# Prerequisite
zabbix version > 4.0.11