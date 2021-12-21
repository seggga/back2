package model

import "github.com/google/uuid"

type Storage interface {
	AddUser(User) error
	AddUnion(Union) error
	GetUser(uuid.UUID) (User, error)
	GetUnion(uuid.UUID) (Union, error)
	AddUserToUnion(userID, unionID uuid.UUID) error
	ExcludeUserFromUnion(userID, unionID uuid.UUID) error
	SearchUserByName(string) []User
	SearchUserByUnion(Union) []User
	SearchUnionByName(string) []Union
	SearchUnionByUsers([]User) []Union
}
