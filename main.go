package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	API_SENDGRID_COM = "https://api.sendgrid.com"
)

var today = time.Now().Format("2006-01-02")

var dateOption = flag.String("date", "1", "Retrieve the timestamp of the invalid email records. It will return a date in a MySQL timestamp format - YYYY-MM-DD HH:MM:SS")
var daysOption = flag.String("d", "", "Number of days in the past for which to retrieve invalid emails (includes today)")
var startDateOption = flag.String("s", today, "The start of the date range for which to retrieve invalid emails. Date must be in YYYY-MM-DD format ")
var endDateOption = flag.String("e", today, "The end of the date range for which to retrieve invalid emails. Date must be in YYYY-MM-DD format ")
var limitOption = flag.String("l", "", "Optional field to limit the number of results returned.")
var offsetOption = flag.String("o", "", "Optional beginning point in the list to retrieve from.")
var bounceTypeOption = flag.String("t", "", "Hard or Soft. Choose the type of bounce to search for.")
var email = flag.String("email", "", "Optional email to search.")

type SendGrid struct {
	auth       Auth
	date       string
	days       string
	startDate  string
	endDate    string
	limit      string
	offset     string
	bounceType string
	email      string
}

type Auth struct {
	ApiUser string
	ApiKey  string
}

func (sg *SendGrid) buildParams() url.Values {
	values := url.Values{}
	if sg.auth.ApiUser != "" {
		values.Add("api_user", sg.auth.ApiUser)
	}
	if sg.auth.ApiKey != "" {
		values.Add("api_key", sg.auth.ApiKey)
	}
	if sg.date != "" {
		values.Add("date", sg.date)
	}
	if sg.days != "" {
		values.Add("days", sg.days)
	}
	if sg.startDate != "" {
		values.Add("start_date", sg.startDate)
	}
	if sg.endDate != "" {
		values.Add("end_date", sg.endDate)
	}
	if sg.limit != "" {
		values.Add("limit", sg.limit)
	}
	if sg.offset != "" {
		values.Add("offset", sg.offset)
	}
	if sg.bounceType != "" {
		values.Add("type", sg.bounceType)
	}
	if sg.email != "" {
		values.Add("email", sg.email)
	}
	return values
}

func (sg *SendGrid) getStatistics() string {
	api := API_SENDGRID_COM + "/api/stats.get.json"
	return httpGet(api, sg.buildParams())
}

func (sg *SendGrid) getBounce() string {
	api := API_SENDGRID_COM + "/api/bounces.get.json"
	return httpGet(api, sg.buildParams())
}

func (sg *SendGrid) getBlocks() string {
	api := API_SENDGRID_COM + "/api/blocks.get.json"
	return httpGet(api, sg.buildParams())
}

func (sg *SendGrid) getInvalid() string {
	api := API_SENDGRID_COM + "/api/invalidemails.get.json"
	return httpGet(api, sg.buildParams())
}

func (sg *SendGrid) getSpam() string {
	api := API_SENDGRID_COM + "/api/spamreports.get.json"
	return httpGet(api, sg.buildParams())
}

func httpGet(api string, queryString url.Values) string {
	client := &http.Client{
		Timeout: time.Duration(3 * time.Second),
	}

	req, _ := http.NewRequest("GET", api, nil)
	req.URL.RawQuery = queryString.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func httpPost(api string, postData url.Values) string {
	client := &http.Client{
		Timeout: time.Duration(3 * time.Second),
	}

	req, _ := http.NewRequest("POST", api, strings.NewReader(postData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func main() {
	flag.Parse()

	api_user := os.Getenv("SENDGRID_API_USER")
	api_key := os.Getenv("SENDGRID_API_KEY")
	if api_user == "" || api_key == "" {
		fmt.Println("env SENDGRID_API_USER and SENDGRID_API_KEY are required")
		return
	}

	auth := &Auth{
		ApiUser: api_user,
		ApiKey:  api_key,
	}

	sg := &SendGrid{
		auth:       *auth,
		date:       *dateOption,
		days:       *daysOption,
		startDate:  *startDateOption,
		endDate:    *endDateOption,
		limit:      *limitOption,
		offset:     *offsetOption,
		bounceType: *bounceTypeOption,
		email:      *email,
	}

	if flag.NArg() < 1 {
		fmt.Println("statistics|bounce|block|invalid|spam")
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case "stat":
		fmt.Println(sg.getStatistics())
	case "bounce":
		fmt.Println(sg.getBounce())
	case "block":
		fmt.Println(sg.getBlocks())
	case "invalid":
		fmt.Println(sg.getInvalid())
	case "spam":
		fmt.Println(sg.getSpam())
	}
}
