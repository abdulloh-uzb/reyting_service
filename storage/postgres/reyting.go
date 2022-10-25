package postgres

import (
	pbr "reyting-service/genproto/reyting"

	"github.com/jmoiron/sqlx"
)

type reytingRepo struct {
	db *sqlx.DB
}

func NewReytingRepo(db *sqlx.DB) *reytingRepo {
	return &reytingRepo{db: db}
}

func (r *reytingRepo) CreateReyting(req *pbr.Ranking) (*pbr.Empty, error) {

	err := r.db.QueryRow(`insert into rankings(name, description, ranking, post_id, customer_id) values($1,$2,$3,$4,$5) select * from WHERE
	ranking BETWEEN 0 AND 5`,
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
