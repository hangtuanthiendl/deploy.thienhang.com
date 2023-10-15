package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"regexp"

	// "log"

	// "os"
	// "strings"
	"time"
)

var (
	interval = 3600
	// services []DdnsService
	// logger service.Logger

	regex = regexp.MustCompile("(?m)[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}")
)

var urls = []string{
	"myipinfo.net",
	"myip.dnsomatic.com",
	"icanhazip.com",
	"checkip.dyndns.org",
	"www.myipnumber.com",
	"checkmyip.com",
	"myexternalip.com",
	"www.ipchicken.com",
	"ipecho.net/plain",
	"bot.whatismyipaddress.com",
	"smart-ip.net/myip",
	"checkip.amazonaws.com",
	"www.checkip.org",
}

var domains = []string{
	"https://N2Lr5q1p6gYwGwdU:bq060q5wgiYGeYAo@domains.google.com/nic/update?hostname=dev.thienhang.com&myip=",
}

func getExternalIP() string {
	var currentIP net.IP
	ip := ""
	for _, i := range rand.Perm(len(urls)) {
		url := "http://" + urls[i]

		content, err := GetResponse(url, "", "")
		if err != nil {
			// logger.Errorf("%s - %s", url, err)
			continue
		}

		ip = regex.FindString(content)

		if currentIP = net.ParseIP(ip); currentIP != nil {
			return ip
		}
	}

	return ip
}

func main() {
	fmt.Println("Hello!")
	ipx := getExternalIP()
	for i, d := range domains {
		fmt.Println(i)
		UpdateIP(d, ipx)
	}
	ticker := time.NewTicker(60 * time.Second)
	for range ticker.C {
		ipx := getExternalIP()
		for i, d := range domains {
			fmt.Println(i)
			UpdateIP(d, ipx)
		}
	}

}

// GetResponse returns the content at the url address
func GetResponse(url string, login string, password string) (string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	if login != "" && password != "" {
		request.SetBasicAuth(login, password)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func UpdateIP(url, ipx string) {
	url = url + ipx
	fmt.Println("Hi", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
	return
}
