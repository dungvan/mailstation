package common

type PubSubTopic string

func (t PubSubTopic) String() string {
	return string(t)
}

const (
	// Magic prefix string for all pubsub topics
	pubsubPrefix = "mailstation:"
	// INCOMING_EMAIL_TOPIC is the topic for incoming email messages.
	INCOMING_EMAIL_TOPIC PubSubTopic = pubsubPrefix + "INCOMING_EMAIL"
	// NEW_OTP_TOPIC is the topic for new OTP parsed messages.
	NEW_OTP_TOPIC PubSubTopic = pubsubPrefix + "NEW_OTP_TOPIC"
)

// IncomingEmail is the struct for incoming email message.
type IncomingEmail struct {
	From         string   `json:"from"`
	To           []string `json:"to"`
	Subject      string   `json:"subject"`
	TextBody     string   `json:"text_body"`
	MailFilePath string   `json:"-"`
}

// MailOTP is the struct for one time password message.
type MailOTP struct {
	Email   string `json:"email"`
	Service string `json:"service"`
	OTP     string `json:"otp"`
}
