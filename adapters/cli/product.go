package cli

import (
	"fmt"

	"github.com/landucci/go-exagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Produto com ID %s e com nome %s foi criado com preço %f e status %s",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, _ := service.Enable(product)
		result = fmt.Sprintf("Produto com ID %s e com nome %s foi criado com preço %f foi habilitado",
			res.GetId(), res.GetName(), res.GetPrice())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, _ := service.Disable(product)
		result = fmt.Sprintf("Produto com ID %s e com nome %s foi criado com preço %f foi desabilitado",
			res.GetId(), res.GetName(), res.GetPrice())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("product: {\n\tID: %s,\n\tName: %s,\n\tPrice: %f,\n\tStatus: %s\n}",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil
}
