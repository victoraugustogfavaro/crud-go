package view

import (
	"github.com/victoraugustogfavaro/crud-go/src/controller/model/response"
	"github.com/victoraugustogfavaro/crud-go/src/model"
)

// função para converter e passar as informações
// para a 'classe' user structure
func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}