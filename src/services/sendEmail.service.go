package services

import (
	"easy-wallet-be/src/utils"
	"log"

	"github.com/mailjet/mailjet-apiv3-go/v4"
)

func SendEmail(email string, name string, subject string, htmlPart string) {
	mailjetClient := mailjet.NewMailjetClient(
		utils.GetEnv("MAILJET_PUBLIC_KEY"),
		utils.GetEnv("MAILJET_SECRET_KEY"),
	)

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "tiggy.ribeiro@gmail.com",
				Name:  "Easy Wallet Team",
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
		log.Fatal(err)
	}
}
