package validation

import (
	"errors"
	"net/http"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string
	Msg   string
}

func message(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Wrong email format"
	case "alphaunicode":
		return "This field should contain unicode alpha characters only"
	case "alpha":
		return "This field should contain ASCII alpha characters only"
	case "numeric":
		return "This field should contain a basic numeric value only"
	case "alphanum":
		return "This field should contain ASCII alphanumeric characters only"
	}
	return ""
}

func SendError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), message(fe.Tag())}
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		return
	}

	c.JSON(http.StatusBadRequest, api.ERROR_MESSAGE_PARSING_BODY_JSON)
}
