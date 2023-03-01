package userinfra

import (
	"dataflow/db/mysql/model"
	userentity "dataflow/pkg/domain/entity/user"
	"github.com/volatiletech/null"
)

func convertToDto(entity *userentity.User) *model.User {
	return &model.User{
		ID:        entity.ID,
		UID:       entity.UID,
		Name:      entity.Name,
		Thumbnail: null.StringFrom(entity.Thumbnail),
	}
}

func convertToUserEntity(dto *model.User) *userentity.User {
	return &userentity.User{
		ID:        dto.ID,
		Name:      dto.Name,
		Thumbnail: dto.Thumbnail.String,
	}
}

func convertToUserSliceEntity(userSlice model.UserSlice) userentity.UserSlice {
	res := make(userentity.UserSlice, 0, len(userSlice))
	for _, userData := range userSlice {
		res = append(res, convertToUserEntity(userData))
	}
	return res
}
