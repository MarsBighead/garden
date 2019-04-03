package garden

// Input items for garden server
type Input interface {
	GetOperationType() string
	Description() string
	Router() string
	Transform()
}
