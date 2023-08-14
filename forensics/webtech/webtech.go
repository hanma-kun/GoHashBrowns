package webtech

import (
	"io"
	"net/http"
	"regexp"
	"strings"
)

func normalizeURL(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	return url
}
// IdentifyTechnologies identifies web technologies from a target URL
func IdentifyTechnologies(targetURL string) (map[string]string, error) {
	technologies := make(map[string]string)

	targetURL = normalizeURL(targetURL)
	
	response, err := http.Get(targetURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Identify web server from response headers
	webServer := response.Header.Get("Server")
	if webServer != "" {
		technologies["Web Server"] = webServer
	}

	// Read response body into a byte buffer
	var bodyBuilder strings.Builder
	_, err = io.Copy(&bodyBuilder, response.Body)
	if err != nil {
		return nil, err
	}

	bodyString := bodyBuilder.String()

	regexIdentifiers := map[string]string{
			// Content Management Systems (CMS)
			"WordPress":       `<meta name="generator" content="WordPress (\d+\.\d+\.?\d*)"`,
			"Drupal":          `Drupal (\d+\.\d+)`,
			"Joomla":          `Joomla! (\d+\.\d+)`,
			"Magento":         `Magento/(\d+\.\d+\.\d+)`,
			"Shopify":         `cdn.shopify.com`,
			"CMS Made Simple": `CMS Made Simple (\d+\.\d+)`,
			"TYPO3":           `TYPO3/(\d+\.\d+)`,
			"PrestaShop":      `PrestaShop/(\d+\.\d+)`,
			"OpenCart":        `OpenCart/(\d+\.\d+)`,
			"phpBB":           `phpBB (\d+\.\d+\.\d+)`,
			"MyBB":            `MyBB (\d+\.\d+\.\d+)`,
			"vBulletin":       `vBulletin (\d+\.\d+\.\d+)`,
			"bbPress":         `bbPress (\d+\.\d+\.\d+)`,
		
			// Web Servers
			"nginx":           `nginx/(\d+\.\d+\.\d+)`,
			"Apache":          `Apache/(\d+\.\d+\.\d+)`,
			"Nginx":           `nginx/(\d+\.\d+\.\d+)`,
			"IIS":             `microsoft-iis/(\d+\.\d+)`,
			"LiteSpeed":       `LiteSpeed`,
			"Caddy":           `Caddy/(\d+\.\d+\.?\d+)`,
		
			// JavaScript Libraries/Frameworks
			"jQuery":          `jQuery v(\d+\.\d+\.\d+)`,
			"React":           `React (\d+\.\d+\.\d+)`,
			"Angular":         `angular\.version\.full\s*=\s*'(\d+\.\d+\.\d+)'`,
			"Vue.js":          `Vue\.js v(\d+\.\d+\.\d+)`,
			"Bootstrap":       `Bootstrap (\d+\.\d+\.\d+)`,
		
			// Analytics
			"Google Analytics": `UA-\d+-\d+`,
		
			// Frameworks and Platforms
			"Node.js":         `nodejs/(\d+\.\d+\.\d+)`,
			"Django":          `Django/(\d+\.\d+)`,
			"Laravel":         `Laravel/(\d+\.\d+\.\d+)`,
			"Symfony":         `Symfony/(\d+\.\d+)`,
			"Express":         `Express/(\d+\.\d+\.\d+)`,
			"Flask":           `Flask/(\d+\.\d+\.\d+)`,
			"Ruby on Rails":   `Rails (\d+\.\d+\.\d+)`,
			"ASP.NET":         `ASP\.NET (\d+\.\d+)`,
			"Spring Boot":     `Spring Boot (\d+\.\d+\.\d+)`,
			
		}		

	for identifier, regexPattern := range regexIdentifiers {
		re := regexp.MustCompile(regexPattern)
		match := re.FindStringSubmatch(bodyString)
		if len(match) > 1 {
			technologies[identifier] = match[1]
		}
	}

	return technologies, nil
}