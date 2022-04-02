package external

import (
	"ecommerce-microservice/order/common"
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/exception"
	"ecommerce-microservice/order/domain"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func NewGormDB() *gorm.DB {
	dsn := command.PostgreConfig()

	if dbInstance == nil {
		var err error

		DBInstance, err := gorm.Open(postgres.New(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,
		})
		if err != nil {
			exception.PanicIfNeeded(err)
		}

		if common.IsMigrate && dbInstance == nil {
			err = DBInstance.AutoMigrate(&domain.Order{}, &domain.OrderProduct{}, &domain.OrderHistory{}, &domain.SellerRating{})
			exception.PanicIfNeeded(err)

			if err == nil {
				fmt.Println("Migrasi berhasil!")
			}
		}

		dbInstance = DBInstance

		return dbInstance
	} else {
		return dbInstance
	}
}
