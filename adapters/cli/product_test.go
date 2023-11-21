package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/landucci/go-exagonal/adapters/cli"
	mock_application "github.com/landucci/go-exagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Hylian Shield"
	productPrice := 3000.0
	productStatus := "enabled"
	productId := "hs"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Produto com ID %s e com nome %s foi criado com preço %f e status %s",
		productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Produto com ID %s e com nome %s foi criado com preço %f foi habilitado",
		productId, productName, productPrice)

	result, err = cli.Run(service, "enable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Produto com ID %s e com nome %s foi criado com preço %f foi desabilitado",
		productId, productName, productPrice)

	result, err = cli.Run(service, "disable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("product: {\n\tID: %s,\n\tName: %s,\n\tPrice: %f,\n\tStatus: %s\n}",
		productId, productName, productPrice, productStatus)

	result, err = cli.Run(service, "asfdasdf", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
