package registry

import (
	"awesomeProject1/external/persistence/mysql"
	"awesomeProject1/usecase"
)

func InjectedPhoneUseCase() usecase.PhoneUsecaseI {
	phoneRepo := mysql.NewPhoneRepository()
	manuRepo := mysql.NewManufacturerRepository()
	return usecase.NewPhoneUseCase(phoneRepo, manuRepo)
}
