package main

import (
	"net/url"
	"strings"
)

var badChars = []string{"#", "@", ";"}
var badFileEndings = []string{".gif", ".jpg", ".jpeg", ".svg", ".png", ".ico", ".pdf", ".swf", ".tif",}

/*
	Parses String to URL and parses them to absolute URLs if necessary
*/
func FixUrl(foundLink, link *string) string {
	uri, err := url.Parse(*foundLink)
	if err != nil {
		AddErrCount()
		Error.Printf("  FixUrl() - Parsing Url failed: \n", err)
		return ""
	}
	baseUrl, err := url.Parse(*link)
	if err != nil {
		AddErrCount()
		Error.Printf("  BaseURL ERROR: %s \n", err)
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

/*
	filters URLs for email adresses, Javascript, images, PDFs etc
*/
func CheckUrl(uri *string) bool {
	for _, str := range badFileEndings {
		if strings.HasSuffix(strings.ToLower(*uri), str) {
			Debug.Printf("  Bad File Ending %s for %s", str, uri)
			return false
		}
	}
	for _, str := range badChars {
		if strings.Contains(*uri, str) {
			Debug.Printf("  Bad Char %s for %s", str, *uri)
			return false
		}
	}

	return true
}

/*
	checks if the found URL and the start URL have the same domain = link to the same page = don't leave the start page (startPage)
*/
func CheckHost(uri *string) bool {
	uriUrl, err := url.Parse(*uri)
	if err != nil {
		AddErrCount()
		Error.Printf("  CheckHost() - Url parsing failed: %s", err)
		return false
	}
	//fmt.Printf("Site Host: %s ## START HOST: %s \n", uriUrl.Host, *startHost)
	//if uriUrl.Host == startHost {
	//if strings.HasPrefix(strings.ToLower(*uri), startPage) {
	if strings.Contains(uriUrl.Host, startHostAdd) {
		return true
	} else {
		Debug.Printf("  CheckHost() - Bad Host: %s instead of %s for %s \n", uriUrl.Host, startHost, uriUrl)
		return false
	}

}
