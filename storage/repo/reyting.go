package repo

import (
	pbr "reyting_service/genproto/reyting"
)

type ReytingStorageI interface {
	CreateReyting(*pbr.Ranking) (*pbr.Empty, error)
	GetRankings(id int) (*pbr.Rankings, error)
	GetRankingsByCustomerId(id int) (*pbr.Rankings, error)
	DeleteRankingByPostId(id int) (*pbr.Empty, error)
	DeleteRankingByCustomerId(id int) (*pbr.Empty, error)
}
