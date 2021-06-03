package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const MaxResponseCount = "100"
const ParentTicket = "1"

// Issue show response issue
type Issue struct {
	ID        int       `json:"id"`
	ProjectID int       `json:"projectId"`
	IssueKey  string    `json:"issueKey"`
	Summary   string    `json:"summary"`
	IssueType IssueType `json:"issueType"`
}

// IssueType show response issueType
type IssueType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	lambda.Start(handler)
}

func errorExit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func handler() {

	body, err := getTicketsCreatedPreviousBusinessDay()
	errorExit(err)

	text, err := makeText(body)
	errorExit(err)

	err = sendToSlack(text)
	errorExit(err)

	fmt.Println(text)
}

func getTicketsCreatedPreviousBusinessDay() (body []byte, err error) {
	values := url.Values{}
	values.Add("apiKey", os.Getenv("BacklogApiKey"))
	values.Add("projectId[]", os.Getenv("BacklogProjectId"))
	values.Add("createdSince", getPreviousBusinessDay())
	values.Add("createdUntil", getPreviousBusinessDay())
	values.Add("count", MaxResponseCount)
	values.Add("parentChild", ParentTicket)

	requestUrl := os.Getenv("BacklogGetIssuesUrl") + values.Encode()
	body, err = Get(requestUrl)
	if err != nil {
		return body, err
	}
	return body, nil
}

func makeText(body []byte) (text string, err error) {
	var issues []Issue
	err = json.Unmarshal(body, &issues)
	if err != nil {
		return "failed", err
	}

	if len(issues) == 0 {
		text += "前営業日に発番されたチケットはありません。"
	} else {
		text += "新しくチケットが発番されました。" + "\n"
		text += "対象チケット数:" + strconv.Itoa(len(issues)) + "個\n"

		for _, value := range issues {
			text += value.IssueType.Name + ":<" + os.Getenv("BacklogIssuePath") + value.IssueKey + "|"
			text += value.IssueKey + ">\t"
			text += value.Summary + "\n"
		}
	}
	return text, nil
}

func sendToSlack(text string) (err error) {
	data := new(slack.WebhookMessage)
	data.Channel = os.Getenv("SlackChannelName")
	data.Username = os.Getenv("SlackUserName")
	data.Text = text
	data.IconEmoji = os.Getenv("SlackIconEmoji")

	err = slack.PostWebhook(os.Getenv("SlackIncomingWebhooksUrl"), data)
	if err != nil {
		return err
	}
	return
}

func Get(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}

func getPreviousBusinessDay() string {
	day := time.Now()
	if day.Weekday() == time.Monday {
		return day.AddDate(0, 0, -3).Format("2006-01-02")
	} else {
		return day.AddDate(0, 0, -1).Format("2006-01-02")
	}
}
