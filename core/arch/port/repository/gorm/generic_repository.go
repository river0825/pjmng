package gorm

import (
	"context"
	"errors"
	"reflect"

	"river0825/cleanarchitecture/core/arch/domain/repository"
	"river0825/cleanarchitecture/core/dependency/storage/gorm"
)

type GenericGormRepository[T any] struct {
	db *gorm.DB
}

var _ repository.IDomainRepository[string] = (*GenericGormRepository[string])(nil)

func NewGormGenericRepository[T any](db *gorm.DB) *GenericGormRepository[T] {
	return &GenericGormRepository[T]{
		db: db,
	}
}

func (g *GenericGormRepository[T]) Save(_ context.Context, item T) error {
	g.db.DB.Save(item)
	return nil
}

func (g *GenericGormRepository[T]) Get(_ context.Context, id string) (result T, err error) {
	// 當 T 不是 pointer type 的時後 `var result T` 會建立一個 instance
	// 在 FindByID 的設計，在 no document 的時後，會希望回傳 nil, nil
	// 這個時後如果不是 pointer type 的時後，會無法達成
	if reflect.TypeOf(result).Kind() != reflect.Ptr {
		return result, errors.New("type must be a pointer")
	}

	g.db.DB.First(&result, id)

	return result, nil
}
