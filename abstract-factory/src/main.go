package main

import "fmt"

// SMS | EMAIL
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// SMS Implementations
type SMSNotification struct {
}

func (smsn SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}
func (smsn SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// Estructura que implementa ISender Methods
type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// EMAIL Implementations
type EmailNotification struct {
}

func (en EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}
func (en EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

// Estructura que implementa ISender Methods
type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		// Instancia un nuevo SMSNotification y devuelve la referencia a memoria
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		// Instancia un nuevo EmailNotification y devuelve la referencia a memoria
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No Notification type")
}

// Metodos genéricos para cualquier instancia del tipo INotificationFactory
func sendNotification(f INotificationFactory) {
	f.SendNotification()
}
func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}
func getChannel(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderChannel())
}

func main() {
	/*
		Se tiene un sofware que envía notificaciones:
		- SMS | Push notification | Email

		El reto es crear un programa capaz de manejar los
		diferentes tipos de notificaciones con un proceso polimórfico

	*/

	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	getMethod(smsFactory)
	getChannel(smsFactory)

	sendNotification(emailFactory)
	getMethod(emailFactory)
	getChannel(emailFactory)

}
