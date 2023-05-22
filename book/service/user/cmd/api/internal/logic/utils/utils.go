package utils

import (
	"go-zero-demo/book/service/user/cmd/api/internal/types"

	"go-zero-demo/book/common/errorx"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

func CheckRegisterParams(req *types.RegisterReq) error {

	err := validation.Validate(req.Number,
		validation.Required,
		validation.Length(1, 9),
	)
	if err != nil {
		return errorx.NewDefaultError("register params number invalid")
	}

	err = validation.Validate(req.Name,
		validation.Required,      // not empty
		validation.Length(2, 10), // length between 5 and 100
	)
	if err != nil {
		return errorx.NewDefaultError("register params name invalid")
	}

	err = validation.Validate(req.Gender,
		validation.Required, // not empty
		validation.In("男", "女"),
	)
	if err != nil {
		return errorx.NewDefaultError("register params gender invalid")
	}

	err = validation.Validate(req.Password,
		validation.Required, // not empty
		validation.Length(10, 15),
	)
	if err != nil {
		return errorx.NewDefaultError("register params passwd invalid")
	}

	return nil
}
