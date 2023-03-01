package reserveinfra

import (
	"dataflow/pkg/domain/entity/reserve"
	"time"
)

func convertToDto(entity *reserveentity.Reserve) *Reserve {
	return &Reserve{
		ID:        entity.ID,
		StartedAt: entity.StartedAt.Format(time.RFC3339),
	}
}

func convertToReserveEntity(dto *Reserve) (*reserveentity.Reserve, error) {
	parsedStartedAt, err := time.Parse(dto.StartedAt, time.RFC3339)
	if err != nil {
		return nil, err
	}

	return &reserveentity.Reserve{
		ID:        dto.ID,
		StartedAt: parsedStartedAt,
	}, nil
}
