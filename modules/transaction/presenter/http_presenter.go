package presenter

import (
	"errors"
	"fmt"
	"loket-app/helper"
	"loket-app/modules/transaction/model"
	"loket-app/modules/transaction/sanitizer"
	"loket-app/modules/transaction/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// transactionServiceHTTPHandler - http handler struct for transaction service.
type transactionServiceHTTPHandler struct {
	transactionUseCase usecase.TransactionUseCase
}

// NewTransactionServiceHandler - function for initiating new transaction service handler.
func NewTransactionServiceHandler(transactionUseCase usecase.TransactionUseCase) *transactionServiceHTTPHandler {
	return &transactionServiceHTTPHandler{
		transactionUseCase: transactionUseCase,
	}
}

// Mount - mounting endpoint(s) by echo framework grouping.
func (h *transactionServiceHTTPHandler) Mount(group *echo.Group) {
	group.POST("/purchase", h.PurchaseTicket)
	group.GET("/get_info", h.GetTransactionDetail)
}

// PurchaseTicket - http handler function for purchasing ticket.
func (h *transactionServiceHTTPHandler) PurchaseTicket(c echo.Context) error {
	logCtx := fmt.Sprintf("%T.PurchaseTicket", *h)

	params := model.PurchaseTicketReq{}
	if err := c.Bind(&params); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_bind_params")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	if err := sanitizer.ValidatePurchaseTicketRequest(&params); err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_validate_purchase_ticket")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	resp, err := h.transactionUseCase.PurchaseTicket(&params)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_purchase_ticket")
		return helper.NewResponse(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), nil).WriteResponse(c)
	}

	data := make(map[string]interface{})
	data["transactionId"] = resp
	return helper.NewResponse(http.StatusOK, http.StatusOK, "Success", data).WriteResponse(c)
}

// GetTransactionDetail - http handler function for getting transaction detail.
func (h *transactionServiceHTTPHandler) GetTransactionDetail(c echo.Context) error {
	logCtx := fmt.Sprintf("%T.GetTransactionDetail", *h)

	txID, err := strconv.ParseUint(c.QueryParam("id"), 10, 64)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_parse_uint")
		return helper.NewResponse(http.StatusInternalServerError, http.StatusInternalServerError, err.Error(), nil).WriteResponse(c)
	}

	if txID <= 0 {
		err := errors.New("invalid transaction id")
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_tx_id")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	resp, err := h.transactionUseCase.GetTransactionDetail(txID)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), logCtx, "error_get_tx_information")
		return helper.NewResponse(http.StatusBadRequest, http.StatusBadRequest, err.Error(), nil).WriteResponse(c)
	}

	data := make(map[string]interface{})
	data["transaction"] = resp
	return helper.NewResponse(http.StatusOK, http.StatusOK, "Success", data).WriteResponse(c)
}
