package notification

type Notification struct {
	Text   string
	Button *Button
}

type Button struct {
	Text string
	Type string
	URL  string
}

func New(text string, buttonText, buttonType, buttonURL string) Notification {
	var button *Button
	if buttonText != "" {
		button = &Button{
			Text: buttonText,
			Type: buttonType,
			URL:  buttonURL,
		}
	}
	return Notification{
		Text:   text,
		Button: button,
	}
}
