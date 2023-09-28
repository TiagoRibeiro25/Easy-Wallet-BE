package services

import (
	"easy-wallet-be/src/configs"
	"easy-wallet-be/src/utils"

	"github.com/mailjet/mailjet-apiv3-go/v4"
)

// SendEmail sends an email to the specified recipient with the given subject and HTML content.
// It uses the Mailjet API to send the email, and retrieves the necessary API keys
// and sender information from the configuration file.
func SendEmail(email string, name string, subject string, htmlPart string) {
	sendEmailData := configs.GetSendEmailData()

	mailjetClient := mailjet.NewMailjetClient(
		sendEmailData.APIPublicKey,
		sendEmailData.APISecretKey,
	)

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: sendEmailData.FromEmail,
				Name:  sendEmailData.FromName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  name,
				},
			},
			Subject:  subject,
			HTMLPart: htmlPart,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		utils.HandleError(err, "Error sending email", false)
	}
}
