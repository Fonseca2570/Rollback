package transaction

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"transaction/models/transaction"
	"transaction/validator"
)

func Update(c *gin.Context) {
	updateRequest := UpdateRequest{}

	err := c.BindJSON(&updateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = validator.Validate.Struct(updateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	updateModel := transaction.UpdateModel{
		IdUser:    updateRequest.IdUser,
		FirstName: updateRequest.FirstName,
		LastName:  updateRequest.LastName,
		Email:     updateRequest.Email,
	}

	err = updateModel.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, updateModel)
	return
}