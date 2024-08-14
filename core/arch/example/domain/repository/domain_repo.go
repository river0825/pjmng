package repository

import (
	"river0825/cleanarchitecture/core/arch/domain/repository"
	"river0825/cleanarchitecture/core/arch/example/domain/entity"
)

type IExampleRoomRepository repository.IDomainRepository[*entity.ExampleRoom]
