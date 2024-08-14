//go:build integration

package gin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gopkg.in/stretchr/testify.v1/assert"

	"river0825/cleanarchitecture/module/account/domain/account_error"

	"river0825/cleanarchitecture/core/arch/core_error"
)

func Test_response(t *testing.T) {

	type args struct {
		givenError         error
		expectHttpCode     int
		expectResponseBody *ResponseBody
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "WhenGiveNoError_ShouldReturn200",
			args: args{
				givenError:         nil,
				expectHttpCode:     http.StatusOK,
				expectResponseBody: &ResponseBody{},
			},
		},
		{
			name: "WhenGiveCoreError_ShouldReturnCorrectMessage",
			args: args{
				givenError: core_error.NewCoreError(
					account_error.GetAccountError,
					fmt.Errorf("test error"),
				),
				expectHttpCode: http.StatusBadRequest,
				expectResponseBody: &ResponseBody{
					ErrorCode: account_error.GetAccountError,
					Message:   "test error",
				},
			},
		},
		{
			name: "WhenGiveValidationErrors_ShouldReturnBadRequest",
			args: args{
				givenError:     validator.ValidationErrors{},
				expectHttpCode: http.StatusBadRequest,
				expectResponseBody: &ResponseBody{
					ErrorCode: "validation_error",
				},
			},
		},
		{
			name: "WhenGiveOtherError_ShouldReturnInternalServerError",
			args: args{
				givenError:     fmt.Errorf("test error"),
				expectHttpCode: http.StatusInternalServerError,
				expectResponseBody: &ResponseBody{
					ErrorCode: "internal_server_error",
					Message:   "internal server error",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			// Act
			response(ctx, nil, tt.args.givenError)

			// Assert
			actual := &ResponseBody{}
			_ = json.Unmarshal(w.Body.Bytes(), actual)

			assert.Equal(t, tt.args.expectResponseBody, actual)
			assert.Equal(t, tt.args.expectHttpCode, w.Code)
		})
	}
}
