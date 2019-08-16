package interactors

import (
	"fmt"

	"github.com/CESARBR/knot-gateway-webui/backend/pkg/entities"
)

// StateRepository represents the state repository interface
type StateRepository interface {
	Update(state entities.State)
	GetCurrent() entities.State
}

// Logger represents the logger interface
type Logger interface {
	Log(message string) error
}

// UpdateStateInteractor represents the update state use case
type UpdateStateInteractor struct {
	stateRepository StateRepository
}

// NewUpdateStateInteractor creates a new UpdateStateInteractor intance
func NewUpdateStateInteractor(stateRepository StateRepository) *UpdateStateInteractor {
	return &UpdateStateInteractor{stateRepository}
}

// Execute runs the use case logic
func (i *UpdateStateInteractor) Execute(t string) error {
	fmt.Println(t)
	// state := i.stateRepository.GetCurrent()
	// if err := state.Update(t); err != nil {
	// 	return err
	// }
	// i.stateRepository.Update(state)
	return nil
}
