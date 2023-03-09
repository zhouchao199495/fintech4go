package service

type Consumer interface {
	Consume(message string)
}
