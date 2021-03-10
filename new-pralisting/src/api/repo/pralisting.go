package repo

import (
	"context"
	"database/sql"
	"new-pralisting/src/api/domain"
)

// PralistingRepo ...
type PralistingRepo struct {
	Conn *sql.DB
}

// NewPralistingRepository ...
func NewPralistingRepository(Conn *sql.DB) domain.PralistingRepo {
	return &PralistingRepo{Conn}
}

// Fetch ...
func (m *PralistingRepo) Fetch(ctx context.Context) (res []domain.Pralisting, err error) {
	query := "SELECT id, company_name FROM emitens"
	rows, err := m.Conn.Query(query)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		d := domain.Pralisting{}
		err = rows.Scan(&d.ID, &d.Name)
		if err != nil {
			return res, err
		}
		res = append(res, d)
	}
	return res, err
}

// GetByID ...
func (m *PralistingRepo) GetByID(ctx context.Context, id int) (res domain.Pralisting, err error) {
	query := "SELECT id,company_name FROM emitens WHERE id = ?"
	err = m.Conn.QueryRow(query, id).Scan(&res.ID, &res.Name)
	if err != nil {
		return res, err
	}
	return res, err
}
