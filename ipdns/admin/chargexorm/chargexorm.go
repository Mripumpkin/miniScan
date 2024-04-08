package main

import (
	//"container/list"
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	_"reflect"
	"regexp"

	//"strconv"
	_ "strings"
	"time"
	"xorm.io/builder"
	"xorm.io/xorm"
)


var engine *xorm.Engine
var Alldata []AllDomain
//var value interface{}
var ipsql,iddomain,dnsdata,idip map[string]interface{}
var valueint int
var newArr []interface{}
var allTime map[interface{}]interface{}
var v1,v2 interface{}
var getiptime []interface{}

type IpDomain  struct {
	ip interface{}
	domain interface{}
	timedate time.Time
}

type AllDomain struct {
	timedata string
	updata	string
	domain string
}


func init() {
	var err error
	engine,err = xorm.NewEngine("sqlite3","../../db/domaininfo.db")
	if err != nil {
		fmt.Println("error1:",err)
	}
}

//IP
func IPSql(ip string) (reslut map[string]interface{}){
	ip_results, err := engine.QueryInterface(builder.Select("created_at","updated_at","ip","domain_id").From("ip_addrs"))
	if err != nil {
		fmt.Println("error2:",err)
	}
	for _, result := range ip_results {
		if result["ip"] == ip {
			//fmt.Println(result)
			ipsql = result
			break
		}
	}
	//fmt.Println(ipsql)
	return ipsql
}

func IdDomain(id interface{}) (result interface{}) {
	domains_results, err := engine.QueryInterface(builder.Select("name","id","created_at","updated_at",).From("domains"))
	if err != nil {
		fmt.Println("error:",err)
	}
	//fmt.Println(reflect.TypeOf(domains_results[1]["id"]))
	for _,dresult  := range domains_results {
		//fmt.Println(dresult["id"])
		//fmt.Println(reflect.TypeOf(dresult["id"]))
		if dresult["id"] == id {
			//fmt.Println(dresult)
			//fmt.Println(reflect.TypeOf(dresult))
			iddomain = dresult
			break
			//fmt.Println(a)
			//return a
		}
	}
	fmt.Println(iddomain)
	return iddomain["name"]
}

//查域名id------------------------------------------------------------------------------
func DnsSql(dns interface{}) (result map[string]interface{}) {
	domains_results, err := engine.QueryInterface(builder.Select("name","id","created_at","updated_at",).From("domains"))
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _,dresult := range domains_results {
		if dresult["name"] == dns {
			dnsdata = dresult
			break
		}
	}
	//fmt.Println(dnsdata)
	return dnsdata
}

//-----------------------------------------------------------------------------------------------------
func IDIP(id interface{}) (result interface{}) {
	ip_results, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at",).From("ip_addrs"))
	if err != nil {
		fmt.Println("error:",err)
	}
	//fmt.Println(ip_results)
	//ip_string := strconv.FormatInt(id, 10)
	for _,ip_result  := range ip_results {
		//fmt.Println(reflect.TypeOf(idresult["domain_id"]))
		//fmt.Println(idresult["domain_id"])
		if ip_result["domain_id"] == id {
			//fmt.Println(idresult)
			//fmt.Println(reflect.TypeOf(idresult))
			idip = ip_result
			break
		}
	}
	//fmt.Println(idip)
	return idip["ip"]
}


func AllData() {
	results ,err := engine.QueryInterface(builder.Select("name","id","created_at","updated_at",).From("domains"))
	if err != nil {
		fmt.Println("Errot:",err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
//------------------------------------------------------------------------------------------------
func DNSIP(dns string) (result interface{}){
	//domains_results, err := engine.QueryInterface(builder.Select("created_at","updated_at","id","name").From("domains"))
	//	DnsSql(dns)
	//fmt.Println(reflect.TypeOf(DnsSql(dns)["id"]))
	//DnsSql(dns)["id"]
	//	fmt.Println(DnsSql(dns)["id"]
	//return DnsSql(dns)["id"]
	value := DnsSql(dns)["id"]
	fmt.Println(value)
	//fmt.Print(value, ",", ok)
	return IDIP(value)
	//fmt.Println()
}

//-------------------------------------------------------------------------------------------------
func IPDNS(ip string)(result interface{}){
	//fmt.Println(reflect.TypeOf(IPSql(ip)["domain_id"]))
	value := IPSql(ip)["domain_id"]
	//fmt.Println(reflect.TypeOf(value))
	//value_int, err := strconv.ParseInt(value, 10, 64)
	//if err != nil {
		//fmt.Println("Error:",err)
	//}
	//fmt.Println(value)
	return IdDomain(value)
}

//*********************************************************************************************
func IdAllIp(id interface{}) (result interface{}){
	allresults, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at",).From("ip_addrs"))
	if err != nil {
		fmt.Println("error:", err)
	}
	NewArr := newArr
	for _,allresult := range allresults {
		if allresult["domain_id"] == id {
			v := allresult["ip"]
			NewArr = append(NewArr,v)
		}
	}
	return NewArr
}

func DNSALLIP(dns string) (result interface{}){
	value := DnsSql(dns)["id"]
	fmt.Println(value)
	//fmt.Print(value, ",", ok)
	return IdAllIp(value)
	//fmt.Println()
}

func MapToJson(param map[string]interface{}) string{
	dataType , _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func interface2String(inter interface{}) {
    switch inter.(type) {
    case string:
        fmt.Println("string", inter.(string))
        break
    case int:
        fmt.Println("int", inter.(int))
        break
    case float64:
        fmt.Println("float64", inter.(float64))
        break
    }
}

//ip creat time ***************************************************************************************************
func IpTime(id interface{}) (result map[interface{}]interface{}){
	allresults, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at",).From("ip_addrs"))
	if err != nil {
		fmt.Println("error:", err)
	}
	AllTimeip := make(map[interface{}]interface{})
	for _,allresult := range allresults {
		if allresult["domain_id"] == id {
			v1 = allresult["created_at"]
			//fmt.Println(v1)
			v2 = allresult["ip"]
			//fmt.Println(v2)
			AllTimeip[v1] = v2
			//fmt.Println(ALLIPtime)
		}
	}
	return AllTimeip
}

func GetIpTime(id interface{}) (result []interface{}){
	allresults, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at",).From("ip_addrs"))
	if err != nil {
		fmt.Println("error:", err)
	}
	for _,allresult := range allresults {
		if allresult["domain_id"] == id {
			v1 = allresult["created_at"].(string)
			getiptime = append(getiptime,v1)
			fmt.Println(getiptime)
		}
	}
	return getiptime
}

func Get_IpTime(dns string) (result []interface{}){
	value := DnsSql(dns)["id"]
	//fmt.Println(value)
	//fmt.Print(value, ",", ok)
	return GetIpTime(value)
	//fmt.Println()
}

func IptoTime(id interface{}) (result map[interface{}]interface{}){
	allresults, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at",).From("ip_addrs"))
	if err != nil {
		fmt.Println("error:", err)
	}
	AllTimeip := make(map[interface{}]interface{})
	for _,allresult := range allresults {
		if allresult["domain_id"] == id {
			v1 = allresult["created_at"]
			//fmt.Println(v1)
			v2 = allresult["ip"]
			//fmt.Println(v2)
			AllTimeip[v1] = v2
			//fmt.Println(ALLIPtime)
		}
	}
	return AllTimeip
}

func ALLIPtime(dns string) (result map[interface{}]interface{}){
	value := DnsSql(dns)["id"]
	//fmt.Println(value)
	//fmt.Print(value, ",", ok)
	return IpTime(value)
	//fmt.Println()
}

//all ip and time ** touple  ***************************************************************************************************
func IdAllIp2(id interface{}) (result interface{}){
	allresults, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at",).From("ip_addrs"))
	if err != nil {
		fmt.Println("error:", err)
	}
	NewArr := newArr
	for _,allresult := range allresults {
		if allresult["domain_id"] == id {
			ip1 := allresult["ip"].(string)
			time1 := allresult["created_at"].(string)
			var s string
			s += ip1 + "   :   " +  time1
			NewArr = append(NewArr, s)
		}
	}
	fmt.Println(NewArr)
	return NewArr
}

func DNSALLIP2(dns string) (result interface{}){
	value := DnsSql(dns)["id"]
	return IdAllIp2(value)
}

func IdAllIp4(id interface{}) (result []string){
	allresults, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at",).From("ip_addrs").OrderBy("created_at"))
	if err != nil {
		fmt.Println("error:", err)
	}
	var newArr []string
	NewArr := newArr
	for _,allresult := range allresults {
		if allresult["domain_id"] == id {
			//ip1 := allresult["ip"].(string)
			v3 := allresult["created_at"].(string)
			str := v3[0 : len(v3)-16]
			reg := regexp.MustCompile("T")
			processedString := reg.ReplaceAllString(str, " ")
			NewArr = append(NewArr, processedString)
		}
	}
	//fmt.Println(NewArr)
	return NewArr
}

func DNSALLIP4(dns string) (result []string){
	value := DnsSql(dns)["id"]
	return IdAllIp4(value)
}

func AllDatas() (result []AllDomain){
	results ,err := engine.QueryInterface(builder.Select("name","id","created_at","updated_at",).From("domains"))
	if err != nil {
		fmt.Println("Errot:",err)
	}
	alldata := Alldata
	alldomain := AllDomain{}
	for _, result := range results {
		//域名---------------------------
		alldomain.domain = result["name"].(string)
		//ip-----------------------------
		//DNSIP(result["name"])
		//alldomain.ip = result[""]
		//创建时间
		v1 := result["updated_at"].(string)
		str1 := v1[0 : len(v1)-16]
		reg1 := regexp.MustCompile("T")
		processedString1 := reg1.ReplaceAllString(str1, " ")
		alldomain.updata = processedString1
		//更新时间------------------------
		v2 := result["created_at"].(string)
		str2 := v2[0 : len(v1)-16]
		reg2 := regexp.MustCompile("T")
		processedString2 := reg2.ReplaceAllString(str2, " ")
		alldomain.timedata = processedString2
		alldata = append(alldata, alldomain)
	}
	return  alldata
}

func main() {
	for _,a := range AllDatas() {
		fmt.Println(a)
		fmt.Println(a.timedata)
	}
}



