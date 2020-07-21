# Mattermost Incoming Webhook


## Example

```go
w := Webhook{
		URL:      "http://example.com/hooks/",
		Username: "bot",
		Channel:  "youtube",
	}

	m := Message{
		Webhook: w,
		Text:    "Hallo",
	}

	err := m.Do()
	if err != nil {
		panic(err)
	}
```