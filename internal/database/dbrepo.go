package database

import (
	"context"

	"gitlab.host-h.net/skunkworks/domain-api-dashboard/internal/models"
)

type DBRepo interface {
	GetDomains(ctx context.Context, searchBy models.DomainSearchByType, searchStr string) ([]models.Domain, error)
}
