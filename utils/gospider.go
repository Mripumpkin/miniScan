package utils

import (
	"net/url"
	"path"
	"regexp"
	"strings"

	"golang.org/x/net/publicsuffix"
)

const SUBRE = `(?i)(([a-zA-Z0-9]{1}|[_a-zA-Z0-9]{1}[_a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1})[.]{1})+` // 子域名正则
var nameStripRE = regexp.MustCompile("(?i)^((20)|(25)|(2b)|(2f)|(3d)|(3a)|(40))+")

var linkFinderRegex = regexp.MustCompile(`(?:"|')(((?:[a-zA-Z]{1,10}://|//)[^"'/]{1,}\.[a-zA-Z]{2,}[^"']{0,})|((?:/|\.\./|\./)[^"'><,;| *()(%%$^/\\\[\]][^"'><,;|()]{1,})|([a-zA-Z0-9_\-/]{1,}/[a-zA-Z0-9_\-/]{1,}\.(?:[a-zA-Z]{1,4}|action)(?:[\?|#][^"|']{0,}|))|([a-zA-Z0-9_\-/]{1,}/[a-zA-Z0-9_\-/]{3,}(?:[\?|#][^"|']{0,}|))|([a-zA-Z0-9_\-]{1,}\.(?:php|asp|aspx|jsp|json|action|html|js|txt|xml|do|shtml|jspx)(?:[\?|#][^"|']{0,}|)))(?:"|')`)

// LinkFinder 匹配链接
func LinkFinder(source string) ([]string, error) {
	var links []string
	if len(source) > 1000000 {
		source = strings.ReplaceAll(source, ";", ";\r\n")
		source = strings.ReplaceAll(source, ",", ",\r\n")
	}
	source = DecodeChars(source)
	match := linkFinderRegex.FindAllStringSubmatch(source, -1)
	for _, m := range match {
		matchGroup1 := FilterNewLines(m[1])
		if matchGroup1 == "" {
			continue
		}
		links = append(links, matchGroup1)
	}
	links = Unique(links)
	return links, nil
}

// DecodeChars 编码字符串
func DecodeChars(s string) string {
	source, err := url.QueryUnescape(s)
	if err == nil {
		s = source
	}
	replacer := strings.NewReplacer(
		`\u002f`, "/",
		`\u0026`, "&",
	)
	s = replacer.Replace(s)
	return s
}

// FilterNewLines
func FilterNewLines(s string) string {
	return regexp.MustCompile(`[\t\r\n]+`).ReplaceAllString(strings.TrimSpace(s), " ")
}

// Unique 切片去重
func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// GetDomain 根据URL获取Domain
func GetDomain(site *url.URL) string {
	domain, err := publicsuffix.EffectiveTLDPlusOne(site.Hostname())
	if err != nil {
		return ""
	}
	return domain
}

// GetSubdomains 根据正则表达式匹配子域名
func GetSubdomains(source, domain string) []string {
	var subs []string
	re := subdomainRegex(domain)
	for _, match := range re.FindAllStringSubmatch(source, -1) {
		subs = append(subs, CleanSubdomain(match[0]))
	}
	return subs
}

func subdomainRegex(domain string) *regexp.Regexp {
	d := strings.Replace(domain, ".", "[.]", -1)
	return regexp.MustCompile(SUBRE + d)
}

func CleanSubdomain(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	s = strings.TrimPrefix(s, "*.")
	s = cleanName(s)
	return s
}

func cleanName(name string) string {
	for {
		if i := nameStripRE.FindStringIndex(name); i != nil {
			name = name[i[1]:]
		} else {
			break
		}
	}
	name = strings.Trim(name, "-")
	if len(name) > 1 && name[0] == '.' {
		name = name[1:]
	}
	return name
}

// InScope 判断URL是否在源内
func InScope(u *url.URL, regexps []*regexp.Regexp) bool {
	for _, r := range regexps {
		if r.MatchString(u.Hostname()) {
			return true
		}
	}
	return false
}

// FixUrl 修复URL
func FixUrl(mainSite *url.URL, nextLoc string, domain string) string {
	nextLocUrl, err := url.Parse(nextLoc)
	if err != nil {
		return ""
	}
	url1 := mainSite.ResolveReference(nextLocUrl).String()
	nextLocUrl1, err := url.Parse(url1)
	if err != nil {
		return ""
	}
	if domain != GetDomain(nextLocUrl1) {
		return ""
	} else {
		return nextLocUrl1.String()
	}
}

// GetExtType 获取URL后缀名
func GetExtType(rawUrl string) string {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	return path.Ext(u.Path)
}
