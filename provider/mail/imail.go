package mail

//IMailService mail service interfaces
type IMailService interface {
	Send() error
}
