package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type responseErr struct {
	Field     string `json:"field"`
	Condition string `json:"condition"`
}

func RenderBindingErrors(ctx *gin.Context, validationError validator.ValidationErrors) {
	var responseErrs []responseErr
	for _, fieldError := range validationError {
		field := fieldError.Field()
		responseErrs = append(responseErrs, responseErr{
			Field:     strings.ToLower(field[:1]) + field[1:],
			Condition: fieldError.ActualTag(),
		})
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, responseErrs)
}
