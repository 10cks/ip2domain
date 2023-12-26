package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.117 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.1.1 Safari/605.1.15",
}

func getRandomUserAgent() string {
	rand.Seed(time.Now().UnixNano())
	return userAgents[rand.Intn(len(userAgents))]
}

type Result struct {
	IP           string   `json:"ip"`
	Address      string   `json:"ip138_address"`
	Message      string   `json:"message,omitempty"`
	BindingTimes []string `json:"result_time,omitempty"`
	BindingSites []string `json:"result_site,omitempty"`
}

var addressPattern = regexp.MustCompile(`<h3>(.*?)</h3>`)
var resultTimePattern = regexp.MustCompile(`class="date">(.*?)</span>`)
var resultSitePattern = regexp.MustCompile(`</span><a href="/(.*?)/" target="_blank">`)

// Modify the ip138Spider function

func ip138Spider(ip string) (*Result, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://site.ip138.com/"+ip, nil)
	req.Header.Set("User-Agent", getRandomUserAgent())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	bodyStr := string(body)

	result := &Result{IP: ip}

	addressMatches := addressPattern.FindStringSubmatch(bodyStr)
	if len(addressMatches) > 1 {
		result.Address = strings.TrimSpace(addressMatches[1])
	}

	if strings.Contains(bodyStr, "暂无结果") {
		result.Message = "未查到相关绑定信息！"
		log.Printf("[+]ip:%s\n", ip)
		log.Printf("归属地：%s\n", result.Address)
		log.Println("未查到相关绑定信息！")
	} else {
		resultTimeMatched := resultTimePattern.FindAllStringSubmatch(bodyStr, -1)
		result.BindingTimes = make([]string, len(resultTimeMatched))
		for i, timeText := range resultTimeMatched {
			result.BindingTimes[i] = strings.TrimSpace(timeText[1])
		}

		resultSiteMatched := resultSitePattern.FindAllStringSubmatch(bodyStr, -1)
		result.BindingSites = make([]string, len(resultSiteMatched))
		for i, siteText := range resultSiteMatched {
			result.BindingSites[i] = strings.TrimSpace(siteText[1])
		}

		log.Printf("[+]ip:%s\n", ip)
		log.Printf("归属地：%s\n", result.Address)
		log.Println("已查到相关绑定信息！")
		log.Println("绑定时间：", result.BindingTimes)
		log.Println("绑定网站：", result.BindingSites)
	}

	log.Println(strings.Repeat("-", 25))
	return result, nil
}

func main() {
	start := time.Now()

	filename := flag.String("f", "", "filename")
	outFile := flag.String("o", "result.json", "output filename")
	flag.Parse()

	if len(*filename) == 0 {
		fmt.Println("Usage: IP2Domain -f <filename> [-o <output filename>]")
		return
	}

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	out, err := os.Create(*outFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer out.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result, err := ip138Spider(scanner.Text())
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		b, err := json.Marshal(result)
		if err != nil {
			fmt.Printf("Error marshaling result: %v\n", err)
			continue
		}

		_, err = out.Write(b)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			continue
		}
		out.Write([]byte("\n"))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println("检索完毕")
	fmt.Printf("运行时间: %.2f秒\n", time.Since(start).Seconds())
}
