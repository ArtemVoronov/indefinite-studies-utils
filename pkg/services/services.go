package services

type Service interface {
	Shutdown() error
}
