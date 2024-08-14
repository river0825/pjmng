package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	"river0825/cleanarchitecture/core/arch/core_error"
)

type ResponseBody struct {
	Message   string               `json:"message,omitempty"`
	ErrorCode core_error.ErrorCode `json:"error_code,omitempty"`
	Data      any                  `json:"data,omitempty"`
}

// response TODO: ticket: TFB-128 Integrate error code and http code, and then remove switch case section in the gin responser.
func response(ctx *gin.Context, data any, err error) {
	switch v := err.(type) {
	case nil:
		ctx.PureJSON(http.StatusOK, &ResponseBody{
			Data: data,
		})
	case *core_error.CoreError:
		handledCoreError(ctx, *v)
	case validator.ValidationErrors:
		ctx.PureJSON(http.StatusBadRequest, &ResponseBody{
			ErrorCode: core_error.ErrorCodeValidationError,
			Message:   err.Error(),
		})
	default:
		log.Error().Stack().Err(err).Msg("internal server error")
		ctx.PureJSON(http.StatusInternalServerError, &ResponseBody{
			ErrorCode: core_error.ErrorCodeInternalServerError,
			Message:   "internal server error",
		})
	}
}

func handledCoreError(ctx *gin.Context, err core_error.CoreError) {
	log.Error().Err(err.Err).Msgf("error code [%s]", err.Code)
	switch {
	case core_error.ErrorCodeValidationError == err.Code:
		ctx.PureJSON(http.StatusBadRequest, &ResponseBody{
			ErrorCode: err.Code,
			Message:   err.Error(),
		})
	default:
		ctx.PureJSON(http.StatusInternalServerError, &ResponseBody{
			ErrorCode: core_error.ErrorCodeInternalServerError,
			Message:   "internal server error",
		})
		log.Error().Err(err).Msg("internal server error")
	}
}
