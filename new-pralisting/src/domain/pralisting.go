package domain

import "context"

// Pralisting ...
type Pralisting struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PralistingEntity interface {
	Fetch(ctx context.Context) ([]Pralisting, error)
	GetByID(ctx context.Context, id int) (Pralisting, error)
}

type PralistingRepo interface {
	Fetch(ctx context.Context) (res []Pralisting, err error)
	GetByID(ctx context.Context, id int) (Pralisting, error)
}
