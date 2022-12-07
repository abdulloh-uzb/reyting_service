package service

import (
	"context"
	pbr "reyting_service/genproto/reyting"
	l "reyting_service/pkg/logger"
	"reyting_service/storage"

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
	return rankings, nil
}
func (r *ReytingService) GetRankingsByCustomerId(ctx context.Context, req *pbr.Id) (*pbr.Rankings, error) {
	rankings, err := r.storage.Reyting().GetRankingsByCustomerId(int(req.Id))
	if err != nil {
		r.logger.Error("error while get rankings by customer id", l.Error(err))
		return &pbr.Rankings{}, err
	}
	return rankings, nil
}

func (r *ReytingService) DeleteRankingByPostId(ctx context.Context, req *pbr.Id) (*pbr.Empty, error) {
	_, err := r.storage.Reyting().DeleteRankingByPostId(int(req.Id))
	if err != nil {
		r.logger.Error("error while delete rankings by post id", l.Error(err))
		return &pbr.Empty{}, err
	}
	return &pbr.Empty{}, nil
}

func (r *ReytingService) DeleteRankingByCustomerId(ctx context.Context, req *pbr.Id) (*pbr.Empty, error) {
	_, err := r.storage.Reyting().DeleteRankingByCustomerId(int(req.Id))
	if err != nil {
		r.logger.Error("error while delete rankings by customer id", l.Error(err))
		return &pbr.Empty{}, err
	}
	return &pbr.Empty{}, nil
}
