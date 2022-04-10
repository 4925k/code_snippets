package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// func main() {
// 	for _, value := range logs {
// 		x, _ := ParseJson(value)
// 		//CheckErr(err)
// 		for k, v := range x {
// 			if k == "srcip" {

// 				fmt.Printf("key: %v Value: %v\n", k, v)
// 				fmt.Printf("key: %T Value: %T\n", k, v)
// 				fmt.Println(x[k])
// 			}
// 		}

// 	}
// }

func ParseSyslog(msg []byte) (map[string]interface{}, error) {
	finalResult := make(map[string]interface{}, 1)
	splitContent := bytes.Split(msg, []byte(" "))
	prev := ""

	for _, maps := range splitContent {
		temp := bytes.Split(maps, []byte(" "))
		if len(temp) < 2 {
			if len(prev) > 0 {
				finalResult[prev] = finalResult[prev].(string) + " " + strings.ReplaceAll(string(temp[0]), `"`, ``)
			}
			continue
		}
		finalResult[string(temp[0])] = strings.ReplaceAll(string(temp[1]), `"`, ``)
		prev = string(temp[0])
	}
	return finalResult, nil
}

// ParseJson parses the json type logs
// returns key-value pairs as map
func ParseJson(content string) (map[string]string, error) {
	//create temp map to unmarshal logs
	var temp map[string]interface{}
	err := json.Unmarshal([]byte(content), &temp)
	if err != nil {
		log.Println("[ERROR] unmarshal log: ", err)
		return nil, err
	}

	//create map[string]string for further processing
	result := make(map[string]string)

	//convert map[string]interface{} to map[string]string
	for k, v := range temp {
		result[k] = fmt.Sprint(v)
	}

	return result, nil
}

var syslog = []string{`<189>date=2022-02-23 time=11:23:49 devname="CORE-FW-PRI" devid="FG200ETK18916930" logid="0001000014" type="traffic" subtype="local" level="notice" vd="root" eventtime=1645594730335227814 tz="+0545" srcip=10.10.6.115 srcname="DESKTOP-NPM2H6L" srcport=137 srcintf="StorageVLAN" srcintfrole="lan" dstip=10.10.6.255 dstport=137 dstintf="unknown0" dstintfrole="undefined" sessionid=146550708 proto=17 action="deny" policyid=0 policytype="local-in-policy" service="udp/137" dstcountry="Reserved" srccountry="Reserved" trandisp="noop" app="netbios forward" duration=0 sentbyte=0 rcvdbyte=0 sentpkt=0 appcat="unscanned" osname="Windows" srcswversion="10" mastersrcmac="30:9c:23:a5:b2:6d" srcmac="30:9c:23:a5:b2:6d" srcserver=0`,
	`<189>date=2022-02-23 time=11:23:49 devname="CORE-FW-PRI" devid="FG200ETK18916930" logid="0001000014" type="traffic" subtype="local" level="notice" vd="root" eventtime=1645594730345226209 tz="+0545" srcip=10.10.6.107 srcname="DESKTOP-2A688GO" srcport=62807 srcintf="StorageVLAN" srcintfrole="lan" dstip=10.10.6.1 dstport=8013 dstintf="root" dstintfrole="undefined" sessionid=146550715 proto=6 action="deny" policyid=0 policytype="local-in-policy" service="tcp/8013" dstcountry="Reserved" srccountry="Reserved" trandisp="noop" app="Endpoint Control Registration" duration=0 sentbyte=0 rcvdbyte=0 sentpkt=0 appcat="unscanned" crscore=5 craction=262144 crlevel="low" osname="Windows" srcswversion="10" mastersrcmac="94:c6:91:5c:10:04" srcmac="94:c6:91:5c:10:04" srcserver=0`,
	`<189>date=2022-02-23 time=11:23:49 devname="CORE-FW-PRI" devid="FG200ETK18916930" logid="0000000013" type="traffic" subtype="forward" level="notice" vd="root" eventtime=1645594730345232842 tz="+0545" srcip=10.10.6.75 srcname="DESKTOP-FH3GB5R" srcport=53263 srcintf="StorageVLAN" srcintfrole="lan" dstip=104.16.124.96 dstport=443 dstintf="VIANET" dstintfrole="wan" srcuuid="736230f0-0550-51e9-8926-2cb0e17b5adc" dstuuid="736230f0-0550-51e9-8926-2cb0e17b5adc" poluuid="0de121bc-cd84-51e9-72b1-3ec9b64f9756" sessionid=146550485 proto=6 action="close" policyid=18 policytype="policy" service="HTTPS" dstcountry="United States" srccountry="Reserved" trandisp="snat" transip=103.41.173.115 transport=53263 duration=1 sentbyte=819 rcvdbyte=3590 sentpkt=8 rcvdpkt=10 shapingpolicyid=27 shapersentname="StorageTrafficShaper" shaperdropsentbyte=0 shaperrcvdname="StorageTrafficShaper" shaperdroprcvdbyte=0 shaperperipname="LowBandwidth_IP" shaperperipdropbyte=0 vwlid=2 vwlquality="Seq_num(4), alive, selected" appcat="unscanned" srchwvendor="Samsung" devtype="Tablet" srcfamily="Galaxy" osname="Windows" srchwversion="Tab" srcswversion="10" mastersrcmac="94:c6:91:5c:10:1c" srcmac="94:c6:91:5c:10:1c" srcserver=0`,
	`<189>date=2022-02-23 time=11:23:50 devname="CORE-FW-PRI" devid="FG200ETK18916930" logid="0001000014" type="traffic" subtype="local" level="notice" vd="root" eventtime=1645594730355226053 tz="+0545" srcip=10.10.6.107 srcname="DESKTOP-2A688GO" srcport=62808 srcintf="StorageVLAN" srcintfrole="lan" dstip=10.10.6.1 dstport=8013 dstintf="root" dstintfrole="undefined" sessionid=146550719 proto=6 action="deny" policyid=0 policytype="local-in-policy" service="tcp/8013" dstcountry="Reserved" srccountry="Reserved" trandisp="noop" app="Endpoint Control Registration" duration=0 sentbyte=0 rcvdbyte=0 sentpkt=0 appcat="unscanned" crscore=5 craction=262144 crlevel="low" osname="Windows" srcswversion="10" mastersrcmac="94:c6:91:5c:10:04" srcmac="94:c6:91:5c:10:04" srcserver=0`,
	`<189>date=2022-02-23 time=11:23:50 devname="CORE-FW-PRI" devid="FG200ETK18916930" logid="0000000013" type="traffic" subtype="forward" level="notice" vd="root" eventtime=1645594730355229115 tz="+0545" srcip=10.10.6.115 srcname="DESKTOP-NPM2H6L" srcport=51966 srcintf="StorageVLAN" srcintfrole="lan" dstip=34.250.67.152 dstport=443 dstintf="VIANET" dstintfrole="wan" srcuuid="736230f0-0550-51e9-8926-2cb0e17b5adc" dstuuid="736230f0-0550-51e9-8926-2cb0e17b5adc" poluuid="0de121bc-cd84-51e9-72b1-3ec9b64f9756" sessionid=146549206 proto=6 action="client-rst" policyid=18 policytype="policy" service="HTTPS" dstcountry="Ireland" srccountry="Reserved" trandisp="snat" transip=103.41.173.115 transport=51966 duration=7 sentbyte=1055 rcvdbyte=4611 sentpkt=10 shapingpolicyid=27 shapersentname="StorageTrafficShaper" shaperdropsentbyte=0 shaperrcvdname="StorageTrafficShaper" shaperdroprcvdbyte=0 shaperperipname="LowBandwidth_IP" shaperperipdropbyte=0 vwlid=2 vwlquality="Seq_num(4), alive, selected" appcat="unscanned" osname="Windows" srcswversion="10" mastersrcmac="30:9c:23:a5:b2:6d" srcmac="30:9c:23:a5:b2:6d" srcserver=0`,
	`<189>date=2022-02-23 time=11:23:50 devname="CORE-FW-PRI" devid="FG200ETK18916930" logid="0000000013" type="traffic" subtype="forward" level="notice" vd="root" eventtime=1645594730365225992 tz="+0545" srcip=10.11.0.148 srcname="HUAWEI_P30_lite-39d19b83b" srcport=57050 srcintf="SSID_GUEST" srcintfrole="lan" dstip=52.85.235.88 dstport=443 dstintf="VIANET" dstintfrole="wan" srcuuid="736230f0-0550-51e9-8926-2cb0e17b5adc" dstuuid="736230f0-0550-51e9-8926-2cb0e17b5adc" poluuid="ec37db68-cd61-51e9-6ab6-57b2185c0187" sessionid=146535701 proto=6 action="client-rst" policyid=16 policytype="policy" service="HTTPS" dstcountry="United States" srccountry="Reserved" trandisp="snat" transip=103.41.173.115 transport=57050 duration=76 sentbyte=2934 rcvdbyte=8885 sentpkt=15 shapingpolicyid=5 shapersentname="GuestTrafficShaper" shaperdropsentbyte=0 shaperrcvdname="GuestTrafficShaper" shaperdroprcvdbyte=0 shaperperipname="LowBandwidth_IP" shaperperipdropbyte=0 vwlid=2 vwlquality="Seq_num(4), alive, selected" appcat="unscanned" srchwvendor="Huawei" osname="Android" srcswversion="10" mastersrcmac="34:79:16:da:c9:a6" srcmac="34:79:16:da:c9:a6" srcserver=0`,
	`<189>date=2022-02-23 time=11:23:50 devname="CORE-FW-PRI" devid="FG200ETK18916930" logid="0000000013" type="traffic" subtype="forward" level="notice" vd="root" eventtime=1645594730365227900 tz="+0545" srcip=10.11.0.148 srcname="HUAWEI_P30_lite-39d19b83b" srcport=57048 srcintf="SSID_GUEST" srcintfrole="lan" dstip=52.85.235.88 dstport=443 dstintf="VIANET" dstintfrole="wan" srcuuid="736230f0-0550-51e9-8926-2cb0e17b5adc" dstuuid="736230f0-0550-51e9-8926-2cb0e17b5adc" poluuid="ec37db68-cd61-51e9-6ab6-57b2185c0187" sessionid=146535571 proto=6 action="client-rst" policyid=16 policytype="policy" service="HTTPS" dstcountry="United States" srccountry="Reserved" trandisp="snat" transip=103.41.173.115 transport=57048 duration=76 sentbyte=2978 rcvdbyte=7605 sentpkt=18 shapingpolicyid=5 shapersentname="GuestTrafficShaper" shaperdropsentbyte=0 shaperrcvdname="GuestTrafficShaper" shaperdroprcvdbyte=0 shaperperipname="LowBandwidth_IP" shaperperipdropbyte=0 vwlid=2 vwlquality="Seq_num(4), alive, selected" appcat="unscanned" srchwvendor="Huawei" osname="Android" srcswversion="10" mastersrcmac="34:79:16:da:c9:a6" srcmac="34:79:16:da:c9:a6" srcserver=0`,
	`<189>date=2022-02-23 time=11:23:50 devname="CORE-FW-PRI" devid="FG200ETK18916930" logid="0001000014" type="traffic" subtype="local" level="notice" vd="root" eventtime=1645594730385234956 tz="+0545" srcip=10.10.6.141 srcname="DESKTOP-7678LAD" srcport=50981 srcintf="StorageVLAN" srcintfrole="lan" dstip=10.10.6.255 dstport=51007 dstintf="unknown0" dstintfrole="undefined" sessionid=146550728 proto=17 action="deny" policyid=0 policytype="local-in-policy" service="udp/51007" dstcountry="Reserved" srccountry="Reserved" trandisp="noop" duration=0 sentbyte=0 rcvdbyte=0 sentpkt=0 appcat="unscanned" osname="Windows" srcswversion="10" mastersrcmac="98:ee:cb:cd:94:1f" srcmac="98:ee:cb:cd:94:1f" srcserver=0`}

var logs = []string{`<14>Feb 28 10:19:09 KBL-PA-3220-FW1 1,2022/02/28 10:19:09,016201014833,TRAFFIC,deny,2304,2022/02/28 10:19:09,192.168.122.231,157.240.16.16,202.63.245.34,157.240.16.16,deny internet,,,facebook-base,vsys1,INTERNAL-BR,SUBISU-INTERNET,ae1,ethernet1/1,Logs to SIEM,2022/02/28 10:19:09,259282,1,59896,80,4985,80,0x400000,tcp,reset-both,438,364,74,4,2022/02/28 10:19:03,0,social-networking,0,14624854793,0x0,192.168.0.0-192.168.255.255,India,0,3,1,policy-deny,0,0,0,0,INTERNET-VSYS,KBL-PA-3220-FW1,from-application,,,0,,0,,N/A,0,0,0,0,5d5b4ac5-6a46-4857-9d50-3bcafc60d282,0`,
	`<14>Feb 28 10:19:09 KBL-PA-3220-FW1 1,2022/02/28 10:19:09,016201014833,TRAFFIC,deny,2304,2022/02/28 10:19:09,192.168.56.244,8.8.8.8,202.79.34.231,8.8.8.8,deny internet,,,dns,vsys1,INTERNAL-BR,WLINK-INTERNET,ae1,ethernet1/2,Logs to SIEM,2022/02/28 10:19:09,288042,1,59924,53,25795,53,0x420000,udp,deny,76,76,0,1,2022/02/28 10:19:03,0,any,0,14624854797,0x0,192.168.0.0-192.168.255.255,United States,0,1,0,policy-deny,0,0,0,0,INTERNET-VSYS,KBL-PA-3220-FW1,from-policy,,,0,,0,,N/A,0,0,0,0,5d5b4ac5-6a46-4857-9d50-3bcafc60d282,0`,
	`<14>Feb 28 10:19:09 KBL-PA-3220-FW1 1,2022/02/28 10:19:09,016201014833,TRAFFIC,start,2304,2022/02/28 10:19:09,10.100.5.60,52.111.240.2,202.63.245.34,52.111.240.2,Office 365-URL rule,kumaribank\sunil.nagarkoti,,ssl,vsys1,INTERNAL-BR,SUBISU-INTERNET,ae1,ethernet1/1,Logs to SIEM,2022/02/28 10:19:09,933271,1,55657,443,52660,443,0x400000,tcp,allow,763,697,66,4,2022/02/28 10:19:03,0,office 365-URL,0,14624854794,0x0,10.0.0.0-10.255.255.255,Singapore,0,3,1,n/a,0,0,0,0,INTERNET-VSYS,KBL-PA-3220-FW1,from-policy,,,0,,0,,N/A,0,0,0,0,b9bffe15-c5ff-4047-8e8a-5817f179ddb2,0`,
	`<14>Feb 28 10:19:09 KBL-PA-3220-FW1 1,2022/02/28 10:19:09,016201014833,TRAFFIC,deny,2304,2022/02/28 10:19:09,192.168.120.11,34.107.221.82,202.63.245.34,34.107.221.82,deny internet,192.168.100.205\kiran.bhandari,,web-browsing,vsys1,INTERNAL-BR,SUBISU-INTERNET,ae1,ethernet1/1,Logs to SIEM,2022/02/28 10:19:09,135109,1,51212,80,24204,80,0x400000,tcp,reset-both,963,481,482,6,2022/02/28 10:19:03,0,computer-and-internet-info,0,14624854796,0x0,192.168.0.0-192.168.255.255,United States,0,3,3,policy-deny,0,0,0,0,INTERNET-VSYS,KBL-PA-3220-FW1,from-application,,,0,,0,,N/A,0,0,0,0,5d5b4ac5-6a46-4857-9d50-3bcafc60d282,0`,
}
