package repository

import (
	"context"
	"database/sql"

	"github.com/sialka/cartola_fc/internal/domain/entity"
	"github.com/sialka/cartola_fc/internal/infra/db"
)

type PlayerRepository struct {
	Repository
}

func NewPlayerRepository(dbConn *sql.DB) *PlayerRepository {
	return &PlayerRepository{
		Repository: Repository{
			dbConn:  dbConn,
			Queries: db.New(dbConn),
		},
	}
}

// ctx context.Context -> Valida se caso demorar cancela a ação
func (r *PlayerRepository) Create(ctx context.Context, player *entity.Player) error {
	// r.Queries.CreatePlayer -> Está no infra/db/queries.sql.go
	err := r.Queries.CreatePlayer(ctx, db.CreatePlayerParams{
		ID:    player.ID,
		Name:  player.Name,
		Price: player.Price,
	})
	return err
}

func (r *PlayerRepository) FindByID(ctx context.Context, id string) (*entity.Player, error) {
	player, err := r.Queries.FindPlayerById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.Player{
		ID:    player.ID,
		Name:  player.Name,
		Price: player.Price,
	}, nil
}

func (r *PlayerRepository) FindByIDForUpdate(ctx context.Context, id string) (*entity.Player, error) {
	player, err := r.Queries.FindPlayerByIdForUpdate(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.Player{
		ID:    player.ID,
		Name:  player.Name,
		Price: player.Price,
	}, nil
}

func (r *PlayerRepository) Update(ctx context.Context, player *entity.Player) error {
	// check if player exists for update
	_, err := r.FindByIDForUpdate(ctx, player.ID)
	if err != nil {
		return err
	}

	err = r.Queries.UpdatePlayer(ctx, db.UpdatePlayerParams{
		ID:    player.ID,
		Name:  player.Name,
		Price: player.Price,
	})
	return err
}

func (r *PlayerRepository) FindAll(ctx context.Context) ([]*entity.Player, error) {
	players, err := r.Queries.FindAllPlayers(ctx)
	if err != nil {
		return nil, err
	}
	var output []*entity.Player
	for _, player := range players {
		output = append(output, &entity.Player{
			ID:    player.ID,
			Name:  player.Name,
			Price: player.Price,
		})
	}
	return output, nil
}

func (r *PlayerRepository) FindAllByIDs(ctx context.Context, ids []string) ([]entity.Player, error) {
	var output []entity.Player
	for _, pID := range ids {
		player, err := r.FindByID(ctx, pID)
		if err != nil {
			return nil, err
		}
		output = append(output, entity.Player{
			ID:    player.ID,
			Name:  player.Name,
			Price: player.Price,
		})
	}
	return output, nil
}
