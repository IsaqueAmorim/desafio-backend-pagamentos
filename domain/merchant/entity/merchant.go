package entity

import (
	"back-end-challenge/domain/customer/value-object/cnpj"
	user "back-end-challenge/domain/shared/user/entity"
)

type Merchant struct {
	User   user.User
	CNPJ   cnpj.CNPJ
	Amount float64
}
