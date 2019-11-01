package slack

import (
	"testing"
)

func TestInitSlack(t *testing.T) {
	errors := ""
	expected := Slack{
		URL:     slackURL,
		Secret:  slackSecret,
		Channel: slackChannel,
	}

	actual := NewSlack(
		slackURL,
		slackChannel,
		slackSecret,
	)

	if expected.URL != actual.URL {
		errors += "\n- URL:\n\tExpected:\t" + expected.URL + "\n\tActual:\t\t" + actual.URL
	}

	if expected.Channel != actual.Channel {
		errors += "\n- Channel:\n\tExpected:\t" + expected.Channel + "\n\tActual:\t\t" + actual.Channel
	}

	if expected.Secret != actual.Secret {
		errors += "\n- Secret:\n\tExpected:\t" + expected.Secret + "\n\tActual:\t\t" + actual.Secret
	}

	if errors != "" {
		t.Errorf(errors)
	}
}

func TestPost(t *testing.T) {
	expected := Response{
		OK:      true,
		Channel: "C9LRQPB89",
		Message: map[string]string{
			"type":     "message",
			"subtype":  "bot_message",
			"text":     "I am a test message",
			"username": "Info-Bot",
			"bot_id":   "B9MGYQSUD",
		},
		Warning: "missing_charset",
		ResponseMetaData: map[string]interface{}{
			"warnings": []string{"missing_charset"},
		},
	}

	slack := Slack{
		URL:     slackURL,
		Secret:  slackSecret,
		Channel: slackChannel,
	}

	actual, err := slack.DoPost(Post{
		Mrkdwn:  true,
		Message: "Kill me Please!!",
		Channel: "#random",
		Attachments: []Attachment{
			Attachment{
				ImageURL:   "https://media.giphy.com/media/1vZcE6QIuvRLJr59UG/giphy.gif",
				Color:      "#FFFFFF",
				AuthorIcon: PeopleBaby,
				Footer:     "I am a footer",
			},
		},
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	if expected.OK != actual.OK {
		t.Errorf("\n\tExpected:\t%v\n\tActual:\t\t%v", expected, actual)
	}

}
