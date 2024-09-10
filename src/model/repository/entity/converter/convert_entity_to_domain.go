package converter

import (
	"github.com/victoraugustogfavaro/crud-go/src/model"
	"github.com/victoraugustogfavaro/crud-go/src/model/repository/entity"
)

// entity é o que vai ser salvo no banco
// domain é o que vai usado na aplicação como negócio

// converter de entity para domain
func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	// setando id e retornando o domain
	domain.SetID(entity.ID.Hex())
	return domain
}
