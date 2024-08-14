package application

import (
	"context"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	"river0825/cleanarchitecture/core/arch/core_error"
)

type IFacade interface {
	Execute(ctx context.Context, cmd any) (any, error)
}

// AbsFacade implicit implement IFacade
type AbsFacade struct {
	useCases map[reflect.Type]IUseCase
}

func NewAbsFacade() *AbsFacade {
	return &AbsFacade{
		useCases: make(map[reflect.Type]IUseCase),
	}
}

func (a *AbsFacade) RegisterUseCase(cmd any, uc IUseCase) {
	if _, ok := a.useCases[reflect.TypeOf(cmd)]; ok {
		panic("use case already registered")
	}

	a.useCases[reflect.TypeOf(cmd)] = uc
}

func (a *AbsFacade) Execute(ctx context.Context, cmd any) (any, error) {
	if c, ok := cmd.(ICommand); ok {
		if err := c.Validate(); err != nil {
			return nil, core_error.NewCoreError(core_error.ErrorCodeValidationError, err)
		}
	} else {
		v := validator.New(validator.WithRequiredStructEnabled())
		if err := v.Struct(cmd); err != nil {
			return nil, core_error.NewCoreError(core_error.ErrorCodeValidationError, err)
		}
	}

	uc := a.useCases[reflect.TypeOf(cmd)]

	if uc == nil {
		log.Fatal().Str("usecase", reflect.TypeOf(uc).String()).Msg("use case not found")
		panic("use case not found")
	}

	resp, err := uc.Execute(ctx, cmd)
	if err != nil {
		log.Error().Err(err).Str("usecase", reflect.TypeOf(uc).String()).Msg("execute use case error")
	}

	return resp, err
}
