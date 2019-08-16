package services

// UpdateStateInteractor represents the UpdateState use case interface
type UpdateStateInteractor interface {
	Execute(stateType string) error
}

// StateService represents the capabilities associated with configuration state
type StateService struct {
	UpdateStateInteractor
}

// NewStateService creates a new StateService instance
func NewStateService(updateStateInteractor UpdateStateInteractor) *StateService {
	return &StateService{updateStateInteractor}
}
