package app_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	app "github.com/lucasf5/app/Domain/Service"
	mock_repositories "github.com/lucasf5/app/mocks"
)

func TestProductServiceGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_repositories.NewMockProductInterface(ctrl)
	persistense := mock_repositories.NewMockProductPersistenceInterface(ctrl)

	persistense.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	productService := app.ProductService{Persistence: persistense}

	result, err := productService.Get("1")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductServiceSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_repositories.NewMockProductInterface(ctrl)
	persistense := mock_repositories.NewMockProductPersistenceInterface(ctrl)

	persistense.EXPECT().Save(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(product, nil).AnyTimes()

	productService := app.ProductService{Persistence: persistense}

	result, err := productService.Save("name", 1.0, "description", "enabled")

	require.Nil(t, err)
	require.Equal(t, product, result)
}
