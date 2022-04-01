package api

import (
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := configSetENV()

	if err != nil {
		panic(err)
	}

	systemEnv := os.Getenv("APP_ENV")

	err = godotenv.Load(systemEnv + ".env")
	if err != nil {
		err = godotenv.Load("../" + systemEnv + ".env")
		if err != nil {
			panic(err)
		}
	}

}

func configSetENV() (err error) {
	if len(os.Args) > 1 {
		err = os.Setenv("APP_ENV", os.Args[1])
	} else {
		err = os.Setenv("APP_ENV", "local")
	}
	return
}
