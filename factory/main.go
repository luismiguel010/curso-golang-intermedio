package main

import "fmt"

// Interfaces

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// SMS

type SMSNotification struct {

}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {

}

func (SMSNotificationSender) GetSenderMethod () string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel () string {
	return "Twilio"
}

// Email

type EmailNotification struct {

}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {

}

func (EmailNotificationSender) GetSenderMethod () string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel () string {
	return "SES"
}

// Función que retorna las clases concretas que queremos utilizar para enviar notificaciones
func getNotificationFactory (notifactionType string) (INotificationFactory, error) {
	if notifactionType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notifactionType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No notification Type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderChannel())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}