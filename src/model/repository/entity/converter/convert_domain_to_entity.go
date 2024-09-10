package converter

import (
	"github.com/victoraugustogfavaro/crud-go/src/model"
	"github.com/victoraugustogfavaro/crud-go/src/model/repository/entity"
)

// converter de domain para entity
func ConvertDomainToEntity(
	domain model.UserDomainInterface, 
) *entity.UserEntity {
	return &entity.UserEntity{
		Email: domain.GetEmail(),
		Password: domain.GetPassword(),
		Name: domain.GetName(),
		Age: domain.GetAge(),
	}
}
