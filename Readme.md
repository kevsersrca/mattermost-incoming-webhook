# Mattermost Incoming Webhook

![tutorial](tutorial.png)

## Installation

```
go get github.com/kevsersrca/mattermost-incoming-webhook
```
## Example

```go
function main() {
	w := Payload{
		URL:      "http://example.com/hooks/",
		Text:    "Hallo :tada:",
	}
	err := m.Do()
	if err != nil {
		panic(err)
	}
}
	
```