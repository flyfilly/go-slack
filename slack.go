package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Slack is the data structure used to hold the slack parameters and execute posts
type Slack struct {
	URL     string
	Channel string
	Secret  string
}

// Post is the shape required to post a message to slack
type Post struct {
	Tokent         string        `json:"token"`
	Mrkdwn         bool          `json:"mrkwn"`
	Message        string        `json:"text"`
	Channel        string        `json:"channel"`
	AsUser         string        `json:"as_user"`
	Attachments    []Attachment  `json:"attachments"`
	Blocks         []interface{} `json:"blocks"`
	IconEmoji      string        `json:"icon_emoji"`
	IconURL        string        `json:"icon_url"`
	LinkNames      bool          `json:"link_names"`
	Parse          string        `json:"parse"`
	ReplyBroadcast bool          `json:"reply_broadcast"`
	ThreadTS       float32       `json:"thread_ts"`
	UnfurlMedia    bool          `json:"unfurl_media"`
	Username       string        `json:"username"`
}

// Field is the shape for Attachment fields
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short string `json:"short"`
}

// Attachment is the shape requirement for Post Attachments
type Attachment struct {
	Fallback   string  `json:"fallback"`
	Color      string  `json:"color"`
	Pretext    string  `json:"pretext"`
	AuthorName string  `json:"author_name"`
	AuthorLink string  `json:"author_link"`
	AuthorIcon string  `json:"author_icon"`
	Title      string  `json:"title"`
	TitleLink  string  `json:"title_link"`
	Text       string  `json:"text"`
	Fields     []Field `json:"fields"`
	ImageURL   string  `json:"image_url"`
	ThumbURL   string  `json:"thumb_url"`
	Footer     string  `json:"footer"`
	FooterIcon string  `json:"footer_icon"`
	TS         float32 `json:"ts"`
}

// Response is the shape of the response from slack
type Response struct {
	OK               bool                   `json:"ok"`
	Channel          string                 `json:"channel"`
	TS               string                 `json:"ts"`
	Message          map[string]string      `json:"message"`
	Warning          string                 `json:"warning"`
	ResponseMetaData map[string]interface{} `json:"response_metadata"`
}

// NewSlack returns a new slack instance
func NewSlack(url, channel, secret string) *Slack {
	return &Slack{
		URL:     url,
		Channel: channel,
		Secret:  secret,
	}
}

// DoPost posts a message to slack
func (slack *Slack) DoPost(post Post) (Response, error) {
	slackResponse := Response{}

	if post.Channel == "" {
		post.Channel = slack.Channel
	}

	payload, err := json.Marshal(post)

	fmt.Println(string(payload))

	if err != nil {
		return slackResponse, err
	}

	if req, err := http.NewRequest(
		http.MethodPost,
		slack.URL,
		bytes.NewBuffer(payload),
	); err == nil {
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", ("Bearer " + slack.Secret))

		if response, err := http.DefaultClient.Do(req); err == nil {

			fmt.Println(response.Body)
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)

			if err != nil {
				return slackResponse, err
			}

			if err := json.Unmarshal(body, &slackResponse); err != nil {
				return slackResponse, err
			}
		} else {
			return slackResponse, err
		}
	} else {
		return slackResponse, nil
	}

	return slackResponse, nil
}
