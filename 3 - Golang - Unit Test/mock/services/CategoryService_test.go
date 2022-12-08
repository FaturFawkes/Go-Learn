package services

import (
	"mock/entity"
	"mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRespository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRespository}

func TestCategoryServices_Get(t *testing.T) {
	// Program Mock
	t.Run("Fail", func(t *testing.T) {
		categoryRespository.Mock.On("FindById", "1").Return(nil)
		category, err := categoryService.Get("1")
		assert.Nil(t, category)
		assert.Error(t, err)
	})
	t.Run("Success", func(t *testing.T) {
		categoryRespository.Mock.On("FindById", "2").Return(entity.Category{"2", "Fatur"})
		category, err := categoryService.Get("2")
		assert.Nil(t, err)
		assert.NotEmpty(t, category)
	})
}