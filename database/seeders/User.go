package seeders

import (
	"github.com/dyfun/memorization-app/app/Helper"
	"github.com/dyfun/memorization-app/app/Models"
	"github.com/dyfun/memorization-app/config"
	"github.com/go-faker/faker/v4"
)

func UserSeeder() {
	for i := 0; i < 10; i++ {
		result := config.Db.Create(&Models.User{
			Email:     faker.Email(),
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Password:  Helper.EncryptPassword("784512"),
		})

		if result.Error != nil {
			panic(result.Error)
		}
	}
}
