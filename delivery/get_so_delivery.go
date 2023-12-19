package delivery

import (
	"errors"
	"io"
	"net/http"
	"sbs-be/model/constant"
	"sbs-be/model/dto"
	"sbs-be/model/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (delivery *sbsDelivery) GetSo(c *gin.Context) {
	var filter dto.RequestSo
	errBind := c.ShouldBindJSON(&filter)

	validate := validator.New()
	if errBind != nil && errors.Is(errBind, io.EOF) { // checking if body req is empty
		errResp := response.BuildEmptyBodyReqResponse(constant.RESPONSE_MESSAGE_BODY_REQ_EMPTY, errBind.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	} else if errBind != nil {
		errResp := response.BuildInvalidTypeResponse(constant.RESPONSE_MESSAGE_INVALID_DATA_TYPE, errBind.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return

	} else if errBind := validate.Struct(filter); errBind != nil {
		errResp := response.BuildInvalidTypeResponse(constant.RESPONSE_MESSAGE_INVALID_BODY_REQ, errBind.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	data := delivery.SbsUsecase.GetSo(c, filter)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
