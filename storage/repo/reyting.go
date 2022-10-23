package repo

import (
	pbr "reyting-service/genproto/reyting"
)

type ReytingStorageI interface {
	CreateReyting(*pbr.Ranking) (*pbr.Empty, error)
	GetRankings(id int) (*pbr.Rankings, error)
	GetRankingsByCustomerId(id int) (*pbr.Rankings, error)
}
