package service

import (
	"context"
	pbr "exam/reyting_service/genproto/reyting"
	l "exam/reyting_service/pkg/logger"
	"exam/reyting_service/storage"

	"github.com/jmoiron/sqlx"
)

type ReytingService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewReytingService(db *sqlx.DB, log l.Logger) *ReytingService {
	return &ReytingService{
		storage: storage.NewStorage(db),
		logger:  log,
	}
}

func (r *ReytingService) CreateRanking(ctx context.Context, req *pbr.Ranking) (*pbr.Empty, error) {
	a, err := r.storage.Reyting().CreateReyting(req)
	if err != nil {
		return &pbr.Empty{}, err
	}
	return a, nil
}

func (r *ReytingService) GetRankings(ctx context.Context, req *pbr.Id) (*pbr.Rankings, error) {
	rankings, err := r.storage.Reyting().GetRankings(int(req.Id))
	if err != nil {
		r.logger.Error("error while get rankings by post id", l.Error(err))
		return &pbr.Rankings{}, err
	}
	return rankings, err
}
func (r *ReytingService) GetRankingsByCustomerId(ctx context.Context, req *pbr.Id) (*pbr.Rankings, error) {
	rankings, err := r.storage.Reyting().GetRankingsByCustomerId(int(req.Id))
	if err != nil {
		r.logger.Error("error while get rankings by customer id", l.Error(err))
		return &pbr.Rankings{}, err
	}
	return rankings, err
}
