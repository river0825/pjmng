package inmem

import (
	"river0825/cleanarchitecture/core/arch/example/domain/entity"
	"river0825/cleanarchitecture/core/arch/example/domain/repository"
	"river0825/cleanarchitecture/core/arch/port/repository/inmem"
)

func NewRoomRepository() repository.IExampleRoomRepository {
	return inmem.NewInMemGenericRepository[*entity.ExampleRoom]()
}
