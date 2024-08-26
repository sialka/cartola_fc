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
func ChoosePlayers(myTeam *entity.MyTeam, myPlayers []entity.Player, players []entity.Player) error {
	totalCost := 0.0
	totalEarned := 0.0 //calculateTotalEarned(myPlayers, players) // Total de ganho

	// Levantando o Custo do novo jogador
	for _, player := range players {

		if playerInMyTeam(player, *myTeam) && !playerInPlayersList(player, players) {
			totalEarned += player.Price
		}

		if !playerInMyTeam(player, *myTeam) && playerInPlayersList(player, players) {
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
