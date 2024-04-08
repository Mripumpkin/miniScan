package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	//"time"

	//"net"
	//"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

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
	// mysql dsn
	// dsn := "root:db-q5n2g@tcp(db:3306)/domaininfo?charset=utf8mb4&parseTime=True&loc=Local"

	// sqlite dsn
	dsn := "domaininfo.db"

	// init db
	var err error
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("连接数据库异常: ", err)
	}

}

// GetDomainFormDB 从数据库获取域名
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


//GetAddrByDomain 解析域名
func GetAddrByDomain(domain string) (ipaddrs []string, err error) {
	ipaddrs, err = net.LookupHost(domain)
	if err != nil {
		return nil, err
	}

	return ipaddrs, nil

}
/*
func SaveDomainIP(domain Domain, ipaddrs []string) (err error) {

	newIPaddrs := make([]*IPAddr, 0)
	now := time.Now()

	for _, ip := range ipaddrs {
		ipaddr := new(IPAddr)
		ipaddr.CreatedAt = now
		ipaddr.IP = ip
		ipaddr.DomainID = domain.ID
		newIPaddrs = append(newIPaddrs, ipaddr)
	}

	result := db.Create(newIPaddrs)

	if result.Error != nil {
		log.Println("添加IP到数据库异常: ", err)
		return
	}

	log.Println("添加IP到数据库成功: ", domain.Name, ipaddrs)
	return

}
*/
func main() {

	// 从数据库获取域名
	domains, err := GetDomainFormDB()
	if err != nil {
		log.Println("查询异常: ", err)
		return
	}
	//fmt.Println(domains)
	for _, domain := range domains {
		//fmt.Println(domain.Name)
		//fmt.Println(domain,ipaddrs)
		// 解析域名
		ipaddrs, err := GetAddrByDomain(domain.Name)
		if err != nil {
			//log.Println("查询域名失败: ", domain.Name)
			continue
		}
		fmt.Println(domain.Name,ipaddrs)
	}
 /*
		// 保存域名-ip信息到数据库
		SaveDomainIP(domain, ipaddrs)
	}

	log.Println("解析域名完成.")
*/
}
