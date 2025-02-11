package otp

import (
	"net/textproto"
	"regexp"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/otp", new(Otp))
}

type Otp struct{}

type Email struct {
	Subject string
	Code    string
	Sender  string
	Date    string
	rawDate time.Time
}

func extractOTP(input string) string {
	// Regular expression to match a sequence of 6 digits
	re := regexp.MustCompile(`\b\d{6}\b`)
	otp := re.FindString(input)
	if otp == "" {
		return ""
	}
	return otp
}

func messageToEmail(msg *imap.Message) *Email {

	return &Email{
		Subject: msg.Envelope.Subject,
		Code:    extractOTP(msg.Envelope.Subject),
		Sender:  msg.Envelope.From[0].MailboxName + "@" + msg.Envelope.From[0].HostName,
		Date:    msg.Envelope.Date.Format(time.DateTime),
		rawDate: msg.Envelope.Date,
	}
}

func getLastMessageByDate(messages []*Email) *Email {
	if len(messages) == 0 {
		return nil
	}

	lastMessage := messages[0]
	for _, message := range messages {
		if message.rawDate.After(lastMessage.rawDate) {
			lastMessage = message
		}
	}

	return lastMessage
}

func (*Otp) LastOtpCodeBySender(email, password, senderEmail string) (*Email, string) {

	c, err := client.DialTLS("imap.gmail.com:993", nil)

	if err != nil {
		return nil, err.Error()
	}

	defer c.Logout()

	if err := c.Login(email, password); err != nil {
		return nil, err.Error()
	}

	_, errSelect := c.Select("INBOX", true)

	if errSelect != nil {
		return nil, errSelect.Error()
	}

	// Create a search criteria to look for the sender email in the "From" header
	criteria := &imap.SearchCriteria{
		Header: textproto.MIMEHeader{
			"From": []string{senderEmail},
		},
	}

	// Execute the search using the criteria
	ids, err := c.Search(criteria)

	if err != nil {
		return nil, err.Error()
	}

	// If no messages are found
	if len(ids) == 0 {
		return nil, "No messages found from the specified sender"
	}

	// Fetch the emails
	seqSet := new(imap.SeqSet)
	seqSet.AddNum(ids...)

	items := []imap.FetchItem{imap.FetchItem("FULL")}

	messages := make(chan *imap.Message, len(ids))

	err = c.Fetch(seqSet, items, messages)

	if err != nil {
		return nil, err.Error()
	}

	var resultMessages []*Email
	for msg := range messages {
		resultMessages = append(resultMessages, messageToEmail(msg))
	}
	lastMessage := getLastMessageByDate(resultMessages)

	return lastMessage, ""
}
