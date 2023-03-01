package planinfra

import (
	"dataflow/pkg/domain/entity/plan"
)

func convertToDto(entity *planentity.Plan) *Plan {
	return &Plan{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func convertToPlanEntity(dto *Plan) *planentity.Plan {
	return &planentity.Plan{
		ID:   dto.ID,
		Name: dto.Name,
	}
}
