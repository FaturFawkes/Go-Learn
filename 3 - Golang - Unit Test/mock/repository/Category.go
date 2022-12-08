package repository

import "mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}