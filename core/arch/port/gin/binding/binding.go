package binding

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-errors/errors"
	"github.com/hashicorp/go-multierror"

	"river0825/cleanarchitecture/core/arch/core_error"
)

func ShouldBind(ctx *gin.Context, obj any) error {
	// ignore validator when binding
	tmpValidator := binding.Validator
	binding.Validator = nil
	defer func() {
		binding.Validator = tmpValidator
	}()

	// ignore error when binding
	err := multierror.Append(nil, ctx.ShouldBindUri(obj))  // use 'uri:"param"' in binding tag to bind http://domain/:param
	err = multierror.Append(err, ctx.ShouldBindQuery(obj)) // use 'form:"param"' in binding tag to bind http://domain?param=value
	err = multierror.Append(err, ctx.ShouldBindHeader(obj))

	// bind body FORM / JSON
	var errBindBody error
	switch ctx.ContentType() {
	case binding.MIMEJSON:
		errBindBody = ctx.ShouldBind(&obj)
	default:
		errBindBody = ctx.ShouldBind(obj)
	}

	// we accept empty body of request, so we have to ignore the EOF error when binding form or json
	if !errors.Is(errBindBody, io.EOF) {
		err = multierror.Append(err, errBindBody)
	}

	// bind xrid
	userID, isExist := ctx.Get("nk_user_id")
	v, _ := ctx.Get("nk_is_subscribed")
	orgID, _ := ctx.Get("nk_organization")
	var isSubscribed bool
	switch s := v.(type) {
	case string:
		isSubscribed = s == "true"
	default:
		if s == nil {
			isSubscribed = false
		} else {
			isSubscribed = v.(bool)
		}
	}

	if isExist {
		m := make(map[string][]string)
		s := fmt.Sprintf("%v", userID)
		sb := fmt.Sprintf("%v", isSubscribed)

		m["user_id"] = []string{s}
		m["is_subscribed"] = []string{sb}

		if orgID != nil {
			oid := fmt.Sprintf("%v", orgID)
			m["organization_id"] = []string{oid}
		}

		// ignore error when binding
		_ = binding.MapFormWithTag(obj, m, "token")
	}

	return core_error.NewValidateErrorOrNil(err.ErrorOrNil())
}
