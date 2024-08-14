package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.host-h.net/skunkworks/domain-api-dashboard/models"
)

const (
	DomainsGetSQL = `select d.ID,
       d.Domain_Number,
       d.Domain_Name,
       d.Account_Number,
       d.Client_Number,
       if(d.Bill_For_Registration = 'Yes', true, false) Bill_For_Registration,
       d.Status
from Domain d
where %s like '%s';`
)

type MysqlDBRepo struct {
	db *sqlx.DB
}

func (r *MysqlDBRepo) GetDomains(ctx context.Context, searchBy models.DomainSearchByType, searchStr string) ([]models.Domain, error) {
	ctxInner, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	result := []models.Domain{}
	sql := fmt.Sprintf(DomainsGetSQL, models.DomainSearchByColumn(searchBy), "%"+searchStr+"%")

	err := r.db.SelectContext(ctxInner, &result, sql)
	if err != nil {
		return nil, err
	}

	return result, nil
}
