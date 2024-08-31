# 1. Mundo Full Cycle e Estudo de Caso "Cartola FC"

**Teoria**




![](img/1.png)

---

![](img/2.png)

---

![](img/3.png)

---

![](img/4.png)

---

![](img/5.png)

---

## Consolidador de Estatistica 

💢️ Linguagem: **Golang**
💢️ Tipo: **Microserviço**

> **Ferramentas e Conhecimentos**

**Vscode (extension)**
- Go

```
ctrl + shift + p
> Go: Install/Update Tools
Marque todas as opções e OK
```

**Conversor de sql para Golang**

https://sqlc.dev/

---

## Hands-On

```bash
go mod init github.com/sialka/cartola_fc
go mod tidy
```

**Estrutura do Projeto**

* **ms-consolidacao**: Pasta do Microsserviço
* **internal**: Arquivos do banco de dados MySQL
* **pkg**: Pacotes
* **sql**: Queries usadas no banco de dados

**Preparando Banco de Dados**

```bash
$ cd ms-consolidacao
$ docker-composer up-d
```

**Gerando: db.go | models.go | queries.sql.go**

```bash
$ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
$ sudo snap install sqlc
$ sqlc generate
```

**Instalando Migrate**

```bash
$ curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
$ apt-get update
$ apt-get install -y migrate

$ migrate -source file:///home/sidnei/developer/cartola-fc/ms-consolidacao/sql/migrations -database 'mysql://root:root@tcp(mysql:3306)/cartola' up
```

**Gerando as Entity**

Path: internal/domain/entity

> File: ms-consolidacao/internal/domain/entity/player.go

```go
package entity

// Estrutura de Dados
type Player struct {
	ID    string
	Name  string
	Price float64
}

// Construtor
func NewPlayer(id, name string, price float64) *Player {
	return &Player{
		ID:    id,
		Name:  name,
		Price: price,
	}
}

```

> File:  ms-consolidacao/internal/domain/entity/team.go

```go
package entity

type MyTeam struct {
	ID      string
	Name    string
	Players []string
	Score   float64
}

func NewMyTeam(id, name string) *MyTeam {
	return &MyTeam{
		ID:   id,
		Name: name,
	}
}

```

> File:  ms-consolidacao/internal/domain/entity/match.go

```go
package entity

import (
	"strconv"
	"time"
)

// Objeto de Valor
type MatchResult struct {
	teamAScore int
	teamBScore int
}

func NewMatchResult(teamAScore, teamBScore int) *MatchResult {
	return &MatchResult{
		teamAScore: teamAScore,
		teamBScore: teamBScore,
	}
}

func (m *MatchResult) GetResult() string {
	// Exemplo de retorno: 1-0
	return strconv.Itoa(m.teamAScore) + "-" + strconv.Itoa(m.teamBScore)
}

type Match struct {
	ID      string
	TeamA   *Team
	TeamB   *Team
	TeamAID string
	TeamBID string
	Date    time.Time
	Status  string
	Result  MatchResult
	Actions []GameAction
}

func NewMatch(id string, teamA *Team, teamB *Team, date time.Time) *Match {
	return &Match{
		ID:      id,
		TeamA:   teamA,
		TeamB:   teamB,
		TeamAID: teamA.ID,
		TeamBID: teamB.ID,
		Date:    date,
	}
}
```

> File:  ms-consolidacao/internal/domain/entity/my_team.go

```go
package entity

type MyTeam struct {
	ID      string
	Name    string
	Players []string
	Score   float64
}

func NewMyTeam(id, name string) *MyTeam {
	return &MyTeam{
		ID:   id,
		Name: name,
	}
}
```

> File:  ms-consolidacao/internal/domain/entity/match_action.go

```go
package entity

import "github.com/google/uuid"

type GameAction struct {
	ID         string
	PlayerID   string
	PlayerName string
	TeamID     string
	Minute     int
	Action     string
	Score      int
}

func NewGameAction(playerID string, minute int, action string, score int, teamID string) *GameAction {
	return &GameAction{
		ID:       uuid.New().String(),
		PlayerID: playerID,
		Minute:   minute,
		//TeamID:   teamID,
		Action:   action,
		Score:    score,
	}
}
```
**Criando o Domain Services**

Toda regra de negocio ficaram aqui ms-consolidacao/internal/domain

> File: ms-consolidacao/internal/domain/service/my_team_choose_players.go

```go
package service

import (
	"errors"

	"github.com/sialka/cartola_fc/internal/domain/entity"
)

var errNotEnoughMoney = errors.New("not enough money")

/*
* Escolha de Jogador
* 1 - Qdo retiro um jogador do meu time estou vendendo
* 2 - Qdo adiciono estou comprando e preciso ter dinheiro para isso
 */
func ChoosePlayers(myTeam entity.MyTeam, myPlayers []entity.Player, players []entity.Player) error {
	totalCost := 0.0
	totalEarned := 0.0 //calculateTotalEarned(myPlayers, players) // Total de ganho

	// Levantando o Custo do novo jogador
	for _, player := range players {

		if playerInMyTeam(player, myTeam) && !playerInPlayersList(player, players) {
			totalEarned += player.Price
		}

		if !playerInMyTeam(player, myTeam) && playerInPlayersList(player, players) {
			totalCost += player.Price
		}
	}

	// Retorna Erro caso não tenha dinheiro para comprar o novo jogador
	if totalCost > myTeam.Score+totalEarned {
		return errNotEnoughMoney
	}

	// Meu dinheiro = Total - Custo
	myTeam.Score += totalEarned - totalCost
	myTeam.Players = []string{}

	// Adicionando os novos jogadores
	for _, player := range players {
		myTeam.Players = append(myTeam.Players, player.ID)
	}
	return nil
}

// Verifica se o Jogador está no meu Time
func playerInMyTeam(player entity.Player, myTeam entity.MyTeam) bool {
	for _, p := range myTeam.Players {
		if p == player.ID {
			return true
		}
	}
	return false
}

// Verifica se o Jogador está na Lista
func playerInPlayersList(player entity.Player, players []entity.Player) bool {
	for _, p := range players {
		if p.ID == player.ID {
			return true
		}
	}
	return false
}

// get the difference between two slices
/*
func calculateTotalEarned(myPlayers []entity.Player, players []entity.Player) float64 {
	var totalEarned float64
	for _, myPlayer := range myPlayers {
		if !playerInPlayersList(myPlayer, players) {
			totalEarned += myPlayer.Price
		}
	}
	return totalEarned
}
*/

```

Repositorio tambénm é um serviço de dominio, vamos criar uma interface para se comunicar com o banco de dados.

> File: ms-consolidacao/internal/domain/repository/interface.go

```go
package repository

import (
	"context"

	"github.com/sialka/cartola_fc/internal/domain/entity"
)

type MyTeamRepositoryInterface interface {
	FindByID(ctx context.Context, id string) (*entity.MyTeam, error)
	FindByIDForUpdate(ctx context.Context, id string) (*entity.MyTeam, error)
	AddScore(ctx context.Context, team *entity.MyTeam, score float64) error
	Create(ctx context.Context, team *entity.MyTeam) error
	FindAllPlayers(ctx context.Context, teamID string) ([]entity.Player, error)
	SavePlayers(ctx context.Context, myTeam *entity.MyTeam) error
}

type MatchRepositoryInterface interface {
	Create(ctx context.Context, match *entity.Match) error
	SaveActions(ctx context.Context, match *entity.Match, score float64) error
	FindByID(ctx context.Context, id string) (*entity.Match, error)
	Update(ctx context.Context, match *entity.Match) error
	FindAll(ctx context.Context) ([]*entity.Match, error)
}

type PlayerRepositoryInterface interface {
	FindByID(ctx context.Context, id string) (*entity.Player, error)
	FindByIDForUpdate(ctx context.Context, id string) (*entity.Player, error)
	Create(ctx context.Context, player *entity.Player) error
	Update(ctx context.Context, player *entity.Player) error
	FindAll(ctx context.Context) ([]*entity.Player, error)
	FindAllByIDs(ctx context.Context, ids []string) ([]entity.Player, error)
}

type TeamRepositoryInterface interface {
	FindByID(ctx context.Context, id string) (*entity.Team, error)
}

```

Agora o implementação dos repositorio ficaram na infra

Na pasta de infra já temos a pasta db com 

repository
+ match_repository.go
+ my_team_respository.go
+ player_repository.go
+ team_repository.go

> File: ms-consolidacao/internal/infra/repository/repository.go

```go
package repository

import (
	"database/sql"

	"github.com/sialka/cartola_fc/internal/infra/db"
)

// Informa que precisamos do conteudo de infra/db/queries.sql.go
type Repository struct {
	dbConn *sql.DB
	*db.Queries
}

```

> File: ms-consolidacao/internal/infra/repository/match_repository.go

```go
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sialka/cartola_fc/internal/domain/entity"
	"github.com/sialka/cartola_fc/internal/infra/db"
)

type MatchRepository struct {
	dbConn *sql.DB
	*db.Queries
}

func NewMatchRepository(dbConn *sql.DB) *MatchRepository {
	return &MatchRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *MatchRepository) Create(ctx context.Context, match *entity.Match) error {
	return r.Queries.CreateMatch(ctx, db.CreateMatchParams{
		ID:        match.ID,
		MatchDate: sql.NullTime{Time: time.Now(), Valid: true},
		TeamAID:   sql.NullString{String: match.TeamA.ID, Valid: true},
		TeamBID:   sql.NullString{String: match.TeamB.ID, Valid: true},
		TeamAName: sql.NullString{String: match.TeamA.Name, Valid: true},
		TeamBName: sql.NullString{String: match.TeamB.Name, Valid: true},
	})
}

func (r *MatchRepository) SaveActions(ctx context.Context, match *entity.Match, score float64) error {
	err := r.deleteMatchActions(ctx, match.ID)
	if err != nil {
		return err
	}
	for _, action := range match.Actions {
		err := r.Queries.CreateAction(ctx, db.CreateActionParams{
			ID:       action.ID,
			MatchID:  match.ID,
			TeamID:   action.TeamID,
			PlayerID: action.PlayerID,
			Minute:   int32(action.Minute),
			Action:   action.Action,
			Score:    score,
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (r *MatchRepository) deleteMatchActions(ctx context.Context, matchID string) error {
	err := r.Queries.RemoveActionFromMatch(ctx, matchID)
	return err
}

func (r *MatchRepository) FindByID(ctx context.Context, id string) (*entity.Match, error) {
	match, err := r.Queries.FindMatchById(ctx, id)
	if err != nil {
		return nil, err
	}

	teamAResult := 0
	teamBResult := 0
	if match.Result.String != "" {
		matchResult := strings.Split(match.Result.String, "-")
		teamAResult, err = strconv.Atoi(matchResult[0])
		if err != nil {
			return nil, err
		}
		teamBResult, err = strconv.Atoi(matchResult[1])
		if err != nil {
			return nil, err
		}
	}
	// get actions
	actions, err := r.Queries.GetMatchActions(ctx, id)
	if err != nil {
		return nil, err
	}

	// convert actions to entity
	var gameActions []entity.GameAction
	for _, action := range actions {
		playerInfo, err := r.Queries.FindPlayerById(ctx, action.PlayerID)
		if err != nil {
			return nil, err
		}
		gameActions = append(gameActions, entity.GameAction{
			ID:         action.ID,
			PlayerID:   action.PlayerID,
			PlayerName: playerInfo.Name,
			Minute:     int(action.Minute),
			Action:     action.Action,
		})
	}

	return &entity.Match{
		ID:      match.ID,
		TeamA:   &entity.Team{ID: match.TeamAID.String, Name: match.TeamAName.String},
		TeamB:   &entity.Team{ID: match.TeamBID.String, Name: match.TeamBName.String},
		Result:  *entity.NewMatchResult(teamAResult, teamBResult),
		Date:    match.MatchDate.Time,
		Actions: gameActions,
	}, nil
}

func (r *MatchRepository) Update(ctx context.Context, match *entity.Match) error {
	matchResult := match.Result.GetResult()
	return r.Queries.UpdateMatch(ctx, db.UpdateMatchParams{
		ID:        match.ID,
		Result:    sql.NullString{String: matchResult, Valid: true},
		MatchDate: sql.NullTime{Time: match.Date, Valid: true},
		TeamAID:   sql.NullString{String: match.TeamA.ID, Valid: true},
		TeamAName: sql.NullString{String: match.TeamA.Name, Valid: true},
		TeamBID:   sql.NullString{String: match.TeamB.ID, Valid: true},
		TeamBName: sql.NullString{String: match.TeamB.Name, Valid: true},
	})
}

func (r *MatchRepository) FindAll(ctx context.Context) ([]*entity.Match, error) {
	matches, err := r.Queries.FindAllMatches(ctx)
	if err != nil {
		return nil, err
	}
	var gameActions []entity.GameAction
	var matchList []*entity.Match

	for _, match := range matches {
		teamAResult := 0
		teamBResult := 0
		if match.Result.String != "" {
			matchResult := strings.Split(match.Result.String, "-")
			teamAResult, err = strconv.Atoi(matchResult[0])
			if err != nil {
				return nil, err
			}
			teamBResult, err = strconv.Atoi(matchResult[1])
			if err != nil {
				return nil, err
			}
		}
		// get actions
		actions, err := r.Queries.GetMatchActions(ctx, match.ID)
		if err != nil {
			return nil, err
		}

		// convert actions to entity
		for _, action := range actions {
			playerInfo, err := r.Queries.FindPlayerById(ctx, action.PlayerID)
			if err != nil {
				return nil, err
			}
			gameActions = append(gameActions, entity.GameAction{
				ID:         action.ID,
				PlayerID:   action.PlayerID,
				PlayerName: playerInfo.Name,
				Minute:     int(action.Minute),
				Action:     action.Action,
				Score:      int(action.Score),
			})
		}

		matchList = append(matchList, &entity.Match{
			ID:      match.ID,
			TeamA:   &entity.Team{ID: match.TeamAID.String, Name: match.TeamAName.String},
			TeamB:   &entity.Team{ID: match.TeamBID.String, Name: match.TeamBName.String},
			TeamAID: match.TeamAID.String,
			TeamBID: match.TeamBID.String,
			Result:  *entity.NewMatchResult(teamAResult, teamBResult),
			Date:    match.MatchDate.Time,
			Actions: gameActions,
		})
	}
	return matchList, nil
}

```

> File: ms-consolidacao/internal/infra/repository/my_team_repository.go

```go
package repository

import (
	"context"
	"database/sql"

	"github.com/sialka/cartola_fc/internal/domain/entity"
	"github.com/sialka/cartola_fc/internal/infra/db"
)

type MyTeamRepository struct {
	dbConn *sql.DB
	*db.Queries
}

func NewMyTeamRepository(dbConn *sql.DB) *MyTeamRepository {
	return &MyTeamRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (m *MyTeamRepository) AddScore(ctx context.Context, myTeam *entity.MyTeam, score float64) error {
	// check if my team exists for update
	_, err := m.FindByIDForUpdate(ctx, myTeam.ID)
	if err != nil {
		return err
	}
	myTeam.Score = score + myTeam.Score
	err = m.Queries.UpdateMyTeamScore(ctx, db.UpdateMyTeamScoreParams{
		ID:    myTeam.ID,
		Score: myTeam.Score,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MyTeamRepository) FindByID(ctx context.Context, id string) (*entity.MyTeam, error) {
	team, err := m.Queries.FindMyTeamById(ctx, id)
	if err != nil {
		return nil, err
	}

	players, err := m.GetPlayersByMyTeamID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.MyTeam{
		ID:    team.ID,
		Name:  team.Name,
		Score: team.Score,
		Players: func() []string {
			var output []string
			for _, player := range players {
				output = append(output, player.ID)
			}
			return output
		}(),
	}, nil
}

func (m *MyTeamRepository) FindByIDForUpdate(ctx context.Context, id string) (*entity.MyTeam, error) {
	team, err := m.Queries.FindMyTeamByIdForUpdate(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.MyTeam{
		ID:    team.ID,
		Name:  team.Name,
		Score: team.Score,
	}, nil
}

func (m *MyTeamRepository) Create(ctx context.Context, myTeam *entity.MyTeam) error {
	err := m.Queries.CreateMyTeam(ctx, db.CreateMyTeamParams{
		ID:    myTeam.ID,
		Name:  myTeam.Name,
		Score: myTeam.Score,
	})
	return err
}

func (m *MyTeamRepository) FindAllPlayers(ctx context.Context, teamID string) ([]entity.Player, error) {
	players, err := m.Queries.GetPlayersByMyTeamID(ctx, teamID)
	if err != nil {
		return nil, err
	}

	var output []entity.Player
	for _, player := range players {
		output = append(output, entity.Player{
			ID:    player.ID,
			Name:  player.Name,
			Price: player.Price,
		})
	}

	return output, nil
}

func (m *MyTeamRepository) SavePlayers(ctx context.Context, myTeam *entity.MyTeam) error {
	// delete all players
	err := m.Queries.DeleteAllPlayersFromMyTeam(ctx, myTeam.ID)
	if err != nil {
		return err
	}

	// add all players
	for _, playerId := range myTeam.Players {
		err = m.addPlayer(ctx, myTeam.ID, playerId)
		if err != nil {
			return err
		}
	}

	err = m.Queries.UpdateMyTeamScore(ctx, db.UpdateMyTeamScoreParams{
		ID:    myTeam.ID,
		Score: myTeam.Score,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MyTeamRepository) addPlayer(ctx context.Context, teamID string, playerID string) error {
	err := m.Queries.AddPlayerToMyTeam(ctx, db.AddPlayerToMyTeamParams{
		PlayerID: playerID,
		MyTeamID: teamID,
	})
	return err
}

```

> File: ms-consolidacao/internal/infra/repository/player_repository.go

```go
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

```

> File: ms-consolidacao/internal/infra/repository/team_repository.go

```go
package repository

import (
	"context"
	"database/sql"

	"github.com/sialka/cartola_fc/internal/domain/entity"
	"github.com/sialka/cartola_fc/internal/infra/db"
)

type TeamRepository struct {
	dbConn *sql.DB
	*db.Queries
	Repository
}

func NewTeamRepository(dbConn *sql.DB) *TeamRepository {
	return &TeamRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *TeamRepository) AddScore(ctx context.Context, player *entity.Player, score float64) error {
	err := r.Queries.AddScoreToTeam(ctx, db.AddScoreToTeamParams{
		ID:    player.ID,
		Score: score,
	})
	return err
}

func (r *TeamRepository) FindByID(ctx context.Context, id string) (*entity.Team, error) {
	team, err := r.Queries.FindTeamById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Team{
		ID:   team.ID,
		Name: team.Name,
	}, nil
}

```

**Use Cases**

Trata-se de toda a ação/intensão de um usuário final para realizar uma transação no sistema.

Pacote uow

> File: ms-consolidacao/pkg/uow/uow.go

```go
package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

type RepositoryFactory func(tx *sql.Tx) interface{}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, db *sql.DB) (*Uow, error) {
	return &Uow{
		Repositories: make(map[string]RepositoryFactory),
		Db:           db,
	}, nil
}

func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}
	repo := u.Repositories[name](u.Tx)
	return repo, nil
}

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {
	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}
	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.Tx = tx
	err = fn(u)
	if err != nil {
		return u.Rollback()
	}
	return u.CommitOrRollback()
}

func (u *Uow) CommitOrRollback() error {
	err := u.Tx.Commit()
	if err != nil {
		if errRb := u.Rollback(); errRb != nil {
			return errors.New(fmt.Sprintf("commit error: %s, rollback error: %s", err, errRb))
		}
		return err
	}
	u.Tx = nil
	return nil
}

func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return errors.New("no transaction to rollback")
	}
	err := u.Tx.Rollback()
	if err != nil {
		return err
	}
	u.Tx = nil
	return nil
}

func (u *Uow) UnRegister(name string) {
	delete(u.Repositories, name)
}

```

Importando o Pacote no go.mod

```bash
$ go mod edit -replace github.com/sialka/cartola_fc/pkg/uow=/home/sidnei/developer/cartola-fc/ms-consolidacao/pkg/uow;
```

Criando UseCase Player_add.go

> File: ms-consolidacao/internal/usecase/player_add.go

```go
package usecase

import (
	"context"

	"github.com/sialka/cartola_fc/internal/domain/entity"
	"github.com/sialka/cartola_fc/internal/domain/repository"
	"github.com/sialka/cartola_fc/pkg/uow"
)

// {"id":"1A","name":"Cristiano Ronaldo","initialPrice":10.0}
type AddPlayerInput struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	InitialPrice float64 `json:"initial_price"`
}

type AddPlayerUseCase struct {
	Uow uow.UowInterface
}

func NewAddPlayerUseCase(uow uow.UowInterface) *AddPlayerUseCase {
	return &AddPlayerUseCase{
		Uow: uow,
	}
}

func (a *AddPlayerUseCase) Execute(ctx context.Context, input AddPlayerInput) error {
	playerRepository := a.getPlayerRepository(ctx)
	player := entity.NewPlayer(input.ID, input.Name, input.InitialPrice)
	err := playerRepository.Create(ctx, player)
	if err != nil {
		return err
	}
	a.Uow.CommitOrRollback()
	return nil
}

func (a *AddPlayerUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := a.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}

```

**Casos de Uso**

Sempre que executarmos uma ação estamos chamando um caso de uso.

**Player Add**

> File: internal/usecase/player_add.go

```go
package usecase

import (
	"context"

	"github.com/sialka/cartola_fc/internal/domain/entity"
	"github.com/sialka/cartola_fc/internal/domain/repository"
	"github.com/sialka/cartola_fc/pkg/uow"
)

// {"id":"1A","name":"Cristiano Ronaldo","initialPrice":10.0}
type AddPlayerInput struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	InitialPrice float64 `json:"initial_price"`
}

type AddPlayerUseCase struct {
	Uow uow.UowInterface
}

func (a *AddPlayerUseCase) Execute(ctx context.Context, input AddPlayerInput) error {
	playerRepository := a.getPlayerRepository(ctx)
	player := entity.NewPlayer(input.ID, input.Name, input.InitialPrice)
	err := playerRepository.Create(ctx, player)
	if err != nil {
		return err
	}
	return a.Uow.CommitOrRollback()
}

func (a *AddPlayerUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := a.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}
```
> File: internal/usecase/my_team_add.go

```go
package usecase

import (
	"context"

	"github.com/sialka/cartola_fc/internal/domain/entity"
	"github.com/sialka/cartola_fc/internal/domain/repository"
	"github.com/sialka/cartola_fc/pkg/uow"
)

type AddMyTeamInput struct {
	ID    string
	Name  string
	Score int
}

type AddMyTeamUseCase struct {
	Uow uow.UowInterface
}

func (a *AddMyTeamUseCase) Execute(ctx context.Context, input AddMyTeamInput) error {
	myTeamRepository := a.getMyTeamRepository(ctx)
	myTeam := entity.NewMyTeam(input.ID, input.Name)
	err := myTeamRepository.Create(ctx, myTeam)
	if err != nil {
		return err
	}
	return a.Uow.CommitOrRollback()
}

func (a *AddMyTeamUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := a.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}
	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}
```

> File: internal/usecase/match_add.go

```go
package usecase

import (
	"context"
	"time"

	"github.com/sialka/cartola_fc/internal/domain/entity"
	"github.com/sialka/cartola_fc/internal/domain/repository"
	"github.com/sialka/cartola_fc/pkg/uow"
)

type MatchInput struct {
	ID      string    `json:"id"`
	Date    time.Time `json:"match_date"`
	TeamAID string    `json:"team_a_id"`
	TeamBID string    `json:"team_b_id"`
}

type MatchUseCase struct {
	Uow uow.UowInterface
}

func (u *MatchUseCase) Execute(ctx context.Context, input MatchInput) error {
	return u.Uow.Do(ctx, func(_ *uow.Uow) error {
		matchRepository := u.getMatchRepository(ctx)
		teamRepository := u.getTeamRepository(ctx)

		teamA, err := teamRepository.FindByID(ctx, input.TeamAID)
		if err != nil {
			return err
		}
		teamB, err := teamRepository.FindByID(ctx, input.TeamBID)
		if err != nil {
			return err
		}

		match := entity.NewMatch(input.ID, teamA, teamB, input.Date)

		// Create match
		err = matchRepository.Create(ctx, match)
		if err != nil {
			return err
		}
		return nil
	})
}

func (u *MatchUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := u.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}

func (u *MatchUseCase) getTeamRepository(ctx context.Context) repository.TeamRepositoryInterface {
	teamRepository, err := u.Uow.GetRepository(ctx, "TeamRepository")
	if err != nil {
		panic(err)
	}
	return teamRepository.(repository.TeamRepositoryInterface)
}

```
> File: internal/usecase/action_add.go

```go
package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/sialka/cartola_fc/internal/domain/entity"
	"github.com/sialka/cartola_fc/internal/domain/repository"
	"github.com/sialka/cartola_fc/pkg/uow"
)

var errActionNotFound = errors.New("action not found")

type ActionAddInput struct {
	MatchID  string `json:"match_id"`
	TeamID   string `json:"team_id"`
	PlayerID string `json:"player_id"`
	Minute   int    `json:"minutes"`
	Action   string `json:"action"`
}

type ActionAddUseCase struct {
	Uow         uow.UowInterface
	ActionTable entity.ActionTableInterface
}

/*
func NewActionAddUseCase(uow uow.UowInterface, actionTable entity.ActionTableInterface) *ActionAddUseCase {
	return &ActionAddUseCase{
		Uow:         uow,
		ActionTable: actionTable,
	}
}*/

// execute
func (a *ActionAddUseCase) Execute(ctx context.Context, input ActionAddInput) error {
	return a.Uow.Do(ctx, func(_ *uow.Uow) error {
		matchRepo := a.getMatchRepository(ctx)
		playerRepo := a.getPlayerRepository(ctx)
		myTeamRepo := a.getMyTeamRepository(ctx)

		match, err := matchRepo.FindByID(ctx, input.MatchID)
		if err != nil {
			return err
		}
		// fmt.Printf("match: %v", match)

		score, err := a.ActionTable.GetScore(input.Action)
		if err != nil {
			return errActionNotFound
		}
		theAction := entity.NewGameAction(input.PlayerID, input.Minute, input.Action, score)
		match.Actions = append(match.Actions, *theAction)
		fmt.Println("match.Actions: ", theAction)

		err = matchRepo.SaveActions(ctx, match, float64(score))
		if err != nil {
			return err
		}

		player, err := playerRepo.FindByID(ctx, input.PlayerID)
		if err != nil {
			return err
		}

		player.Price += float64(score)
		err = playerRepo.Update(ctx, player)
		if err != nil {
			return err
		}

		myTeam, err := myTeamRepo.FindByID(ctx, input.TeamID)
		if err != nil {
			return err
		}
		err = myTeamRepo.AddScore(ctx, myTeam, float64(score))
		if err != nil {
			return err
		}
		return nil
	})

}

func (a *ActionAddUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := a.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}

func (a *ActionAddUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := a.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}
	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}

func (a *ActionAddUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := a.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}
```
> File: internal/usecase/my_team_choose_players.go

```go
package usecase

import (
	"context"

	"github.com/sialka/cartola_fc/internal/domain/repository"
	"github.com/sialka/cartola_fc/internal/domain/service"
	"github.com/sialka/cartola_fc/pkg/uow"
)

type MyTeamChoosePlayersInput struct {
	ID        string   `json:"my_team_id"`
	PlayersID []string `json:"players"`
}

type MyTeamChoosePlayersUseCase struct {
	Uow uow.UowInterface
}

func NewMyTeamChoosePlayersUseCase(uow uow.UowInterface) *MyTeamChoosePlayersUseCase {
	return &MyTeamChoosePlayersUseCase{
		Uow: uow,
	}
}

func (u *MyTeamChoosePlayersUseCase) Execute(ctx context.Context, input MyTeamChoosePlayersInput) error {
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		playerRepo := u.getPlayerRepository(ctx)
		myTeamRepo := u.getMyTeamRepository(ctx)
		myTeam, err := myTeamRepo.FindByID(ctx, input.ID)
		if err != nil {
			return err
		}
		myPlayers, err := playerRepo.FindAllByIDs(ctx, myTeam.Players)
		if err != nil {
			return err
		}

		players, err := playerRepo.FindAllByIDs(ctx, input.PlayersID)
		if err != nil {
			return err
		}
		service.ChoosePlayers(myTeam, myPlayers, players)
		err = myTeamRepo.SavePlayers(ctx, myTeam)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *MyTeamChoosePlayersUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := u.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}
	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}

func (u *MyTeamChoosePlayersUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := u.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}
```

> File: internal/usecase/match_update_result.go

```go
package usecase

import (
	"context"
	"strconv"
	"strings"

	"github.com/sialka/cartola_fc/internal/domain/entity"
	"github.com/sialka/cartola_fc/internal/domain/repository"
	"github.com/sialka/cartola_fc/pkg/uow"
)

type MatchUpdateResultInput struct {
	ID     string `json:"match_id"`
	Result string `json:"result"`
}

type MatchUpdateResultUseCase struct {
	Uow uow.UowInterface
}

func NewMatchUpdateResultUseCase(uow uow.UowInterface) *MatchUpdateResultUseCase {
	return &MatchUpdateResultUseCase{
		Uow: uow,
	}
}

func (u *MatchUpdateResultUseCase) Execute(ctx context.Context, input MatchUpdateResultInput) error {
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		matchRepo := u.getMatchRepository(ctx)
		match, err := matchRepo.FindByID(ctx, input.ID)
		if err != nil {
			return err
		}
		matchResult := strings.Split(input.Result, "-")
		// convert results to int
		teamAResult, _ := strconv.Atoi(matchResult[0])
		teamBResult, _ := strconv.Atoi(matchResult[1])
		match.Result = *entity.NewMatchResult(teamAResult, teamBResult)
		err = matchRepo.Update(ctx, match)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *MatchUpdateResultUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := u.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}

```
