package models

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/jeevatkm/go-model.v1"
	"gorm.io/gorm"
	"strconv"
)

func BindAndValidate(context echo.Context, source interface{}, destination interface{}) error {
	if err := ValidateRequest(context, source); err != nil {
		return err
	}

	if err := MapStruct(source, destination); err != nil {
		return err
	}
	return nil
}

func MapStruct(source interface{}, destination interface{}) error {
	model.Copy(destination, source)

	return nil
}

func ValidateRequest(context echo.Context, source interface{}) error {
	if err := context.Bind(source); err != nil {
		return err
	}

	if errs := context.Validate(source); errs != nil {
		return errs
	}

	return nil
}

func Paginate(context echo.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(context.QueryParam("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(context.QueryParam("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
