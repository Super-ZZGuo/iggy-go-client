import (
	binaryserialization "github.com/iggy-rs/iggy-go-client/binary_serialization"
	. "github.com/iggy-rs/iggy-go-client/contracts"
)

func (tms *IggyHttpClient) GetUser() (*UserResponse, error) {

}

func (tms *IggyHttpClient) GetUsers() ([]*UserResponse, error) {

}

func (tms *IggyHttpClient) CreateUser(request CreateUserRequest) error {

}

func (tms *IggyHttpClient) UpdateUser(request UpdateUserRequest) error {

}

func (tms *IggyHttpClient) DeleteUser(identifier Identifier) error {

}

func (tms *IggyHttpClient) UpdateUserPermissions(request UpdateUserPermissionsRequest) error {

}

func (tms *IggyHttpClient) ChangePassword(request ChangePasswordRequest) error {

}
