package application

import (
	"river0825/cleanarchitecture/core/arch/application"
	"river0825/cleanarchitecture/core/arch/example/domain/repository"
)

type AFacade struct {
	*application.AbsFacade
	roomRepo repository.IExampleRoomRepository
}

var _ application.IFacade = (*AFacade)(nil)

func NewAFacade(roomRepo repository.IExampleRoomRepository) *AFacade {
	f := &AFacade{
		roomRepo:  roomRepo,
		AbsFacade: application.NewAbsFacade(),
	}

	f.RegisterUseCase(&LeaveCommand{}, newLeaveUseCase(roomRepo))
	f.RegisterUseCase(&JoinCommand{}, newJoinUseCase(roomRepo))
	f.RegisterUseCase(&ErrorCommand{}, newErrorUseCase())

	return f
}
