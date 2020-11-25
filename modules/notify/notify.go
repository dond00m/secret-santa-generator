package notify

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"

	"github.com/go-gomail/gomail"
	"github.com/matcornic/hermes/v2"
)

// SantaPair captures santa and their receipent
type SantaPair struct {
	SantaName         string
	SantaEmail        string
	ReceipentName     string
	ReceipentEmail    string
	ReceipentWishlist string
	ReceipentAddress  string
}

// smtpServer data to smtp server
type smtpAuthentication struct {
	Server         string
	Port           int
	SenderEmail    string
	SenderIdentity string
	SMTPUser       string
	SMTPPassword   string
}

// sendOptions are options for sending an email
type sendOptions struct {
	To      string
	Subject string
}

// SendEmail will send an email to the secret santa with reciepient information
func SendEmail(pair SantaPair) {
	log.SetPrefix("Notify-SendEmail: ")
	log.Printf("Send email to %s at %s", pair.ReceipentName, pair.ReceipentEmail)

	// // smtp server configuration.
	smtpConfig := smtpAuthentication{
		Server:         "smtp.gmail.com",
		Port:           465,
		SenderIdentity: "Secret Cow Santa",
		SenderEmail:    os.Getenv("HERMES_SENDER_EMAIL"),
		SMTPPassword:   os.Getenv("HERMES_SMTP_PASSWORD"),
		SMTPUser:       os.Getenv("HERMES_SENDER_EMAIL"),
	}
	options := sendOptions{
		To:      pair.SantaEmail,
		Subject: "Secret Santa 2020",
	}

	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Secret Cow Santa",
			// Optional product logo
			Logo: "https://i.imgur.com/iM0h7Cd.jpeg",
		},
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: pair.SantaName,
			Intros: []string{
				"It's family secret santa time, here's your match!",
				"Rules/Info:",
				" * Gift limit is $50",
				" * Plan on mailing your gift in case we can't get together",
				"You are the secret santa for:",
			},
			Table: hermes.Table{
				Data: [][]hermes.Entry{
					// List of rows
					{
						// Key is the column name, Value is the cell value
						// First object defines what columns will be displayed
						{Key: "Name", Value: pair.ReceipentName},
						{Key: "Wishlist", Value: pair.ReceipentWishlist},
						{Key: "Address", Value: pair.ReceipentAddress},
					},
				},
				Columns: hermes.Columns{
					// Custom style for each rows
					CustomWidth: map[string]string{
						"Item":  "20%",
						"Price": "15%",
					},
					CustomAlignment: map[string]string{
						"Price": "right",
					},
				},
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Optionally, preview the generated HTML e-mail by writing it to a local file
	err = ioutil.WriteFile("preview.html", []byte(emailBody), 0644)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}
	htmlBytes, err := ioutil.ReadFile(fmt.Sprintf("preview.html"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sending email to %s...", pair.SantaEmail)
	err = send(smtpConfig, options, string(htmlBytes))
	if err != nil {
		panic(err)
	}

}

// send sends the email
func send(smtpConfig smtpAuthentication, options sendOptions, htmlBody string) error {

	if smtpConfig.Server == "" {
		return errors.New("SMTP server config is empty")
	}
	if smtpConfig.Port == 0 {
		return errors.New("SMTP port config is empty")
	}

	if smtpConfig.SMTPUser == "" {
		return errors.New("SMTP user is empty")
	}

	if smtpConfig.SenderIdentity == "" {
		return errors.New("SMTP sender identity is empty")
	}

	if smtpConfig.SenderEmail == "" {
		return errors.New("SMTP sender email is empty")
	}

	if options.To == "" {
		return errors.New("no receiver emails configured")
	}

	from := mail.Address{
		Name:    smtpConfig.SenderIdentity,
		Address: smtpConfig.SenderEmail,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from.String())
	m.SetHeader("To", options.To)
	m.SetHeader("Subject", options.Subject)
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer(smtpConfig.Server, smtpConfig.Port, smtpConfig.SMTPUser, smtpConfig.SMTPPassword)

	return d.DialAndSend(m)
}
