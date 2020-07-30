# Mattermost Incoming Webhook

![tutorial](tutorial.png)

## Example

```go
w := Payload{
		URL:      "http://example.com/hooks/",
		Text:    "Hallo :tada:",
	}
	err := m.Do()
	if err != nil {
		panic(err)
	}
```