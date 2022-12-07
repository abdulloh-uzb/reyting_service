package postgres

import (
	pbr "reyting_service/genproto/reyting"
	"time"

	"github.com/jmoiron/sqlx"
)

type reytingRepo struct {
	db *sqlx.DB
}

func NewReytingRepo(db *sqlx.DB) *reytingRepo {
	return &reytingRepo{db: db}
}

func (r *reytingRepo) CreateReyting(req *pbr.Ranking) (*pbr.Empty, error) {

	err := r.db.QueryRow(`insert into rankings(name, description, ranking, post_id, customer_id) values($1,$2,$3,$4,$5)`,
		req.Name, req.Description, req.Ranking, req.PostId, req.CustomerId)
	if err.Err() != nil {
		return &pbr.Empty{}, err.Err()
	}
	return &pbr.Empty{}, nil
}

func (r *reytingRepo) GetRankings(id int) (*pbr.Rankings, error) {
	rankings := &pbr.Rankings{}

	rows, err := r.db.Query(`select name,description, ranking, post_id, customer_id from rankings where post_id=$1`, id)

	if err != nil {
		return &pbr.Rankings{}, err
	}
	for rows.Next() {
		ranking := &pbr.Ranking{}
		err := rows.Scan(&ranking.Name, &ranking.Description, &ranking.Ranking, &ranking.PostId, &ranking.CustomerId)
		if err != nil {
			return &pbr.Rankings{}, err
		}
		rankings.Rankings = append(rankings.Rankings, ranking)
	}

	return rankings, nil
}
func (r *reytingRepo) GetRankingsByCustomerId(id int) (*pbr.Rankings, error) {
	rankings := &pbr.Rankings{}

	rows, err := r.db.Query(`select name,description, ranking, post_id, customer_id from rankings where customer_id=$1`, id)

	if err != nil {
		return &pbr.Rankings{}, err
	}
	for rows.Next() {
		ranking := &pbr.Ranking{}
		err := rows.Scan(&ranking.Name, &ranking.Description, &ranking.Ranking, &ranking.PostId, &ranking.CustomerId)
		if err != nil {
			return &pbr.Rankings{}, err
		}
		rankings.Rankings = append(rankings.Rankings, ranking)
	}

	return rankings, nil
}

func (r *reytingRepo) DeleteRankingByPostId(id int) (*pbr.Empty, error) {

	_, err := r.db.Exec(`update rankings set deleted_at = $1 where post_id=$2 and deleted_at is null`, time.Now(), id)
	if err != nil {
		return &pbr.Empty{}, err
	}
	return &pbr.Empty{}, nil
}

func (r *reytingRepo) DeleteRankingByCustomerId(id int) (*pbr.Empty, error) {

	_, err := r.db.Exec(`update posts set deleted_at = $1 where customer_id=$2 and deleted_at is null`, time.Now(), id)
	if err != nil {
		return &pbr.Empty{}, err
	}
	return &pbr.Empty{}, nil
}
