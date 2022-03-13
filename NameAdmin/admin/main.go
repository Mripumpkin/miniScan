package main

import (
	"encoding/json"
	//"errors"
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	//"strconv"
	"xorm.io/builder"
	"xorm.io/xorm"

	//"github.com/leffss/gee"
	"github.com/stnc/pongo2gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//"html/template"
	//"github.com/go-sql-driver/mysql"
	//"database/sql"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
	_"xorm.io/builder"
	_"xorm.io/xorm"
)

var (
	db *gorm.DB
	engine *xorm.Engine
	Alldata []AllDomain
	newArr []string		//all ip
	ipsql,iddomain,dnsdata,idip,domaintime map[string]interface{}
	ipnames,domainnames,timedate interface{}
	getiptime []string
)


const (
	Title     = "域名IP查询系统"
	TitleMINI = "支持域名与IP查询"
	Msg	=	"域名不合法，请输入正确域名"
	msg	=	"域名IP添加成功"
)

type AllDomain struct {
	timedata string
	updata	string
	domain string
}

//*******************************************************************************8
type AllIp struct {
	timedata string
	updata	string
	domain string
}

type Domain struct {
	gorm.Model
	Name    string `gorm:"unique"`
	IPAddrs []IPAddr
}

type IPAddr struct {
	gorm.Model
	IP       string
	DomainID uint
}

func init() {

	// sqlite dsn
	dsn := "../db/domaininfo.db"

	// init db
	var err error
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("连接数据库异常: ", err)
		//log.Fatal(err)
	}

	engine,err = xorm.NewEngine("sqlite3","../db/domaininfo.db")
	if err != nil {
		fmt.Println("error:",err)
	}

	db.AutoMigrate(&Domain{}, &IPAddr{})
}

//MAP -> STRING
func MapToJson(param map[string]interface{}) string{
	dataType , _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

//查询IP
func IPSql(ip string) (reslut map[string]interface{}){
	ipResults, err := engine.QueryInterface(builder.Select("created_at","updated_at","ip","domain_id").From("ip_addrs"))
	if err != nil {
		fmt.Println("error2:",err)
	}
	for _, result := range ipResults {
		if result["ip"] == ip {
			//fmt.Println(result)
			ipsql = result
			break
		}
	}
	//fmt.Println(ipsql)
	return ipsql
}

//查询DOMAIN
func IdDomain(id interface{}) (result interface{}) {
	domainsResults, err := engine.QueryInterface(builder.Select("name","id","created_at","updated_at").From("domains"))
	if err != nil {
		fmt.Println("error:",err)
	}
	for _,Dresult  := range domainsResults {
		if Dresult["id"] == id {
			iddomain = Dresult
			break
		}
	}
	fmt.Println(iddomain)
	return iddomain["name"]
}

//查域名id------------------------------------------------------------------------------
func DnsSql(dns interface{}) (result map[string]interface{}) {
	domainsResults, err := engine.QueryInterface(builder.Select("name","id","created_at","updated_at").From("domains").OrderBy("created_at"))
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _,Dresult := range domainsResults {
		if Dresult["name"] == dns {
			dnsdata = Dresult
			break
		}
	}
	//fmt.Println(dnsdata)
	return dnsdata
}

//查询id-ip -----------------------------------------------------------------------------------------------------
func IDIP(id interface{}) (result interface{}) {
	ipResults, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at").From("ip_addrs"))
	if err != nil {
		fmt.Println("error:",err)
	}
	for _,ipResult  := range ipResults {
		if ipResult["domain_id"] == id {
			idip = ipResult
			break
		}
	}
	return idip["ip"]
}

//ip creat time ** dict ***************************************************************************************************
func IpTime(id interface{}) (result map[string]string){
	allresults, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at").From("ip_addrs"))
	if err != nil {
		fmt.Println("error:", err)
	}
	AllTimeip := make(map[string]string)
	for _,allresult := range allresults {
		if allresult["domain_id"] == id {
			v1 := allresult["created_at"].(string)
			v2 := allresult["ip"].(string)
			AllTimeip[v2] = v1
		}
	}
	return AllTimeip
}

func TimeIp(id interface{}) (result map[string]string){
	allResults, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at").From("ip_addrs"))
	if err != nil {
		fmt.Println("error:", err)
	}
	AllTimeip := make(map[string]string)
	for _,allresult := range allResults {
		if allresult["domain_id"] == id {
			v1 := allresult["created_at"].(string)
			str := v1[0 : len(v1)-16]
			reg := regexp.MustCompile("T")
			processedString := reg.ReplaceAllString(str, " ")
			v2 := allresult["ip"].(string)
			AllTimeip[processedString] = v2
		}
	}
	return AllTimeip
}

func ALLIPtime(dns string) (result map[string]string){
	value := DnsSql(dns)["id"]
	return IpTime(value)
}

func ALLtimeIP(dns string) (result map[string]string){
	value := DnsSql(dns)["id"]
	return TimeIp(value)
}
//************************************************************************************************
func GetIpTimes(id interface{}) (result []string){
	allResults, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at").From("ip_addrs"))
	if err != nil {
		fmt.Println("error:", err)
	}
	for _,allresult := range allResults {
		if allresult["domain_id"] == id {
			v1 := allresult["created_at"].(string)
			getiptime = append(getiptime,v1)
			fmt.Println(getiptime)
		}
	}
	return getiptime
}

func GetIpTime(dns string) (result []string){
	value := DnsSql(dns)["id"]
	//fmt.Println(value)
	//fmt.Print(value, ",", ok)
	return GetIpTimes(value)
	//fmt.Println()
}

//*************************************************************************************************************
//all ip and time ** touple  ***************************************************************************************************
func IdAllIp(id interface{}) (result []string){
	allresults, err := engine.QueryInterface(builder.Select("ip","domain_id","created_at","updated_at").From("ip_addrs").OrderBy("created_at"))
	if err != nil {
		fmt.Println("error:", err)
	}
	NewArr := newArr
	for _,allresult := range allresults {
		if allresult["domain_id"] == id {
			//ip1 := allresult["ip"].(string)
			v1 := allresult["created_at"].(string)
			str := v1[0 : len(v1)-16]
			reg := regexp.MustCompile("T")
			processedString := reg.ReplaceAllString(str, " ")
			NewArr = append(NewArr, processedString)
		}
	}
	//fmt.Println(NewArr)
	return NewArr
}

func DNSALLIP(dns string) (result interface{}){
	value := DnsSql(dns)["id"]
	return IdAllIp(value)
}


//查询IP-DNS creat time******************************************************************************************
func DomainCreatTime(id interface{}) (result interface{}) {
	domainsResults, err := engine.QueryInterface(builder.Select("name","id","created_at","updated_at").From("domains"))
	if err != nil {
		fmt.Println("error:",err)
	}
	for _,dresult  := range domainsResults {
		if dresult["id"] == id {
			domaintime = dresult
			break
		}
	}
	fmt.Println(domaintime)
	return domaintime["created_at"]
}

func DomainTime(ip string)(result interface{}){
	value := IPSql(ip)["domain_id"]
	v1 := DomainCreatTime(value).(string)
	str := v1[0 : len(v1)-16]
	reg := regexp.MustCompile("T")
	processedString := reg.ReplaceAllString(str, " ")
	return processedString
}

//查询DNS-IP
func DNSIP(dns string) (result interface{}){
	value := DnsSql(dns)["id"]
	return IDIP(value)
	//fmt.Println()
}

//查询IP-DNSr
func IPDNS(ip string)(result interface{}){
	value := IPSql(ip)["domain_id"]

	return IdDomain(value)
}

//GetIndex 获取首页
func GetIndex(c *gin.Context) {
	query := c.Query("q")
	// 检查查询字符串是 IP/domain
	if query == "" {
		// 展示首页
		// Call the HTML method of the Context to render a template
		c.HTML(http.StatusOK, "index.html",
			pongo2.Context{
				"title":      Title,
				"title_mini": TitleMINI,
			})
		return
	}

	ipaddr := net.ParseIP(query)
	if ipaddr != nil {
		// 通过 IP 查询域名
		ipresult := db.Where("ip = ?", ipaddr.String()).Find(&IPAddr{})
		//Domainresult := db.Where("ip = ?",ipaddr.String()).Find(&Domain{}))
		if ipresult.RowsAffected == 0 {
			// 没有找到
			c.HTML(http.StatusOK, "index.html", pongo2.Context{
				"title":      Title,
				"title_mini": TitleMINI,
				"data":       query,
				"code":       404,
			})
			return
		} else {
			// 找到域名
			ipnames = query
			domainnames = IPDNS(query)
			timedate = DomainTime(query)
			log.Println("查询到数据")
			c.HTML(http.StatusOK, "index.html",
				pongo2.Context{
					"title":      Title,
					"title_mini": TitleMINI,
					//"data":       ipname,
					"ipcode":	  1,
					"code":       200,
					"ip":		  ipnames,
					"domain":	  domainnames,
					"time":		  timedate,
				})
		}
		return
	}

	// 使用域名关联查询 IP
	domainResult := db.Where("name = ?", query).Find(&Domain{})
	if DnsSql(query) == nil || domainResult.RowsAffected ==0 {
		log.Printf("find domain worning name: %v\n", query)
		// 没有找到
		c.HTML(http.StatusOK, "index.html", pongo2.Context{
			"title":      Title,
			"title_mini": TitleMINI,
			"data":       "",
			"code":       404,
		})
	} else {
		// 找到IP
		//var alldomain map[]AllDomain

		//b := Get_IpTime(query)
		log.Println("查询到数据")
		c.HTML(http.StatusOK, "index.html",
			pongo2.Context{
				"title":      Title,
				"title_mini": TitleMINI,
				"code":       200,
				"ipcode":	  0,
				"data":		  ALLtimeIP(query),
				"domain":	  query,
				"times":	  DNSALLIP(query),
			})
	}
	return
}

//-----------------------------------------------------------------------------------------------------------
// GetDomain 获取添加域名页面
func GetDomains(c *gin.Context) {
	pathURL := strings.Trim(c.Request.URL.Path, "/")
	nameTPL := fmt.Sprintf("%s.html", pathURL)
	code := 0
	log.Println("namt tpl: ", nameTPL)
	c.HTML(http.StatusOK, nameTPL, pongo2.Context{
		"title": "添加域名",
		"data":  "",
		"code":  code,
		"Msg":   "",
	})
}

// AddDomain 添加新域名
func AddDomains(c *gin.Context) {
	name := strings.TrimSpace(c.PostForm("name"))
	now := time.Now()

	pathURL := strings.Trim(c.Request.URL.Path, "/")
	nameTPL := fmt.Sprintf("%s.html", pathURL)
	//Title := "添加域名"

	// 检查域名合法性
	if name == "" {
		data := ""
		code := http.StatusNotFound
		c.HTML(http.StatusOK, nameTPL, pongo2.Context{
			"title": "添加域名失败",
			"data":  data,
			"code":  code,
			"Msg":   Msg,
		})
		return
	}

	re := regexp.MustCompile(`(?im)^(?:[_a-z0-9](?:[_a-z0-9-]{0,61}[a-z0-9])?\.)+(?:[a-z](?:[a-z0-9-]{0,61}[a-z0-9])?)?$`)
	names := re.FindAllString(name, -1)
	if len(names) < 1 {
		data := ""
		code := http.StatusNotFound
		c.HTML(http.StatusOK, nameTPL, pongo2.Context{
			"title": "添加域名失败",
			"data":  data,
			"code":  code,
			"Msg":   Msg,
		})
		return
	}

	domain := new(Domain)
	domain.CreatedAt = now
	domain.UpdatedAt = now
	domain.Name = names[0]

	// 添加域名
	log.Printf("add domain name: %v\n", name)
	db.Model(&Domain{}).Create(domain)

	//msg := fmt.Sprintf("成功添加新域名: %s", name)
	code := http.StatusOK
	c.HTML(http.StatusOK, nameTPL, pongo2.Context{
		"title": "添加域名成功",
		"data":  "",
		"code":  code,
		"Msg":   msg,
	})
}

/*
//GetShowData 获取域名/IP页面
func GetShowData(c *gin.Context) {
	pathURL := strings.Trim(c.Request.URL.Path, "/")
	nameTPL := fmt.Sprintf("%s.html", pathURL)
	log.Println("namt tpl: ", nameTPL)
	//Db := db
	c.HTML(http.StatusOK, nameTPL, pongo2.Context{
		"title": "查询结果",
		"data":  "",
		"code": "" ,
	})
}

func GetDomainFormDB() (domains []Domain, err error) {
	result := db.Model(&Domain{}).Find(&domains)
	if result.Error != nil {
		errMsg := fmt.Sprintf("查询异常: %s", result.Error)
		return nil, errors.New(errMsg)
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	//fmt.Println(domains)
	return domains, nil
}

func GetAddrByDomain(domain string) (ipaddrs []string, err error) {
	ipaddrs, err = net.LookupHost(domain)
	if err != nil {
		return nil, err
	}
	return ipaddrs, nil
}
*/

//查询ALL------------------------------------------------------------------------------------------------------
func AllDatas() (result []AllDomain){
	results ,err := engine.QueryInterface(builder.Select("name","id","created_at","updated_at").From("domains"))
	if err != nil {
		fmt.Println("Error:",err)
	}
	alldata := Alldata
	alldomain := AllDomain{}
	for _, allresult := range results {
		//域名---------------------------
		alldomain.domain = allresult["name"].(string)
		//ip-----------------------------
		//DNSIP(result["name"])
		//alldomain.ip = result[""]
		//创建时间
		v1 := allresult["updated_at"].(string)
		str1 := v1[0 : len(v1)-16]
		reg1 := regexp.MustCompile("T")
		processedString1 := reg1.ReplaceAllString(str1, " ")
		alldomain.updata = processedString1
		//更新时间------------------------
		v2 := allresult["created_at"].(string)
		str2 := v2[0 : len(v1)-16]
		reg2 := regexp.MustCompile("T")
		processedString2 := reg2.ReplaceAllString(str2, " ")
		alldomain.timedata = processedString2
		alldata = append(alldata, alldomain)
	}
	return  alldata
}

//ALLData
func ShowAllData(c *gin.Context) {
	pathURL := strings.Trim(c.Request.URL.Path, "/")
	nameTPL := fmt.Sprintf("%s.html", pathURL)
	log.Println("namt tpl: ", nameTPL)
	domains := AllDatas()
	//for v := range domains {
	//	fmt.Println(v)
	//}
	code := http.StatusOK
	c.HTML(http.StatusOK,nameTPL, pongo2.Context{
		"title": 	"域名IP查询系统",
		"code": 	code,
		"result": 	domains,
		"data":		"",
	})
}

//删除ip及dns----------------------------------------------------------------------------------
func DelSelect(dns string) (result Domain) {
	var Dns Domain
	Dns.Name = dns
	db.Debug().Delete(Dns)
	return Domain{}
}

func DelShow(c *gin.Context) {
	dns := c.Param("dns")
	c.JSON(200, gin.H{"type" : "delete","dns": dns})
	// 检查字符串domain
	db.AutoMigrate(&Domain{})
	//删除dns
	DelSelect(dns)
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	r.HTMLRender = pongo2gin.TemplatePath("templates")
	// router
	r.GET("/", GetIndex)
	r.POST("/",GetIndex)
	r.GET("/domain", GetDomains)
	r.POST("/domain", AddDomains)
	r.GET("/showdata", ShowAllData)
	//r.POST("/showdata", ShowAllData)
	r.GET("/showdata/:dns", DelShow)
	r.POST("/showdata/:dns", DelShow)

	log.Fatal(r.Run("0.0.0.0:8888"))
}


