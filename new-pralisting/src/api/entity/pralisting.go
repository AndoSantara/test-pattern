package entity

import (
	"context"
	"new-pralisting/src/api/domain"
)

// PralistingEntity ...
type PralistingEntity struct {
	entityRepo domain.PralistingRepo
}

// NewPralistingEntity ...
func NewPralistingEntity(a domain.PralistingRepo) domain.PralistingEntity {
	return &PralistingEntity{
		entityRepo: a,
	}
}

// Fetch ...
func (a *PralistingEntity) Fetch(c context.Context) (res []domain.Pralisting, err error) {
	res, err = a.entityRepo.Fetch(c)
	if err != nil {
		return nil, err
	}
	return
}

// GetByID ...
func (a *PralistingEntity) GetByID(c context.Context, id int) (res domain.Pralisting, err error) {
	res, err = a.entityRepo.GetByID(c, id)
	if err != nil {
		return res, err
	}
	return
}
