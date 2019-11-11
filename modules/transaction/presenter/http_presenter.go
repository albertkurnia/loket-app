package presenter

import (
	"loket-app/modules/transaction/usecase"

	"github.com/labstack/echo"
)

type transactionServiceHTTPHandler struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewTransactionServiceHandler(transactionUseCase usecase.TransactionUseCase) *transactionServiceHTTPHandler {
	return &transactionServiceHTTPHandler{
		transactionUseCase: transactionUseCase,
	}
}

func (h *transactionServiceHTTPHandler) Mount(group *echo.Group) {
	group.POST("/purchase", h.PurchaseTicket)
	group.GET("/get_info", h.GetTransactionDetail)
}

func (h *transactionServiceHTTPHandler) PurchaseTicket(c echo.Context) error {
	return nil
}

func (h *transactionServiceHTTPHandler) GetTransactionDetail(c echo.Context) error {
	return nil
}
