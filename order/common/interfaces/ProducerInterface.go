package interfaces

//go:generate mockgen -destination=../../mocks/services/mock_service.go -package=mock_service -source=ProducerInterface.go

type ProducerService interface {
	ProduceMessage(msg interface{}) error
}
