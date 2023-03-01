package userinfra

import (
	userentity "dataflow/pkg/domain/entity/user"
)

func convertToDto(entity *userentity.User) *User {
	return &User{
		UID:       entity.UID,
		Name:      entity.Name,
		Thumbnail: entity.Thumbnail,
	}
}

func convertToUserEntity(userDto *User) *userentity.User {
	return &userentity.User{
		Name:      userDto.Name,
		Thumbnail: userDto.Thumbnail,
	}
}
