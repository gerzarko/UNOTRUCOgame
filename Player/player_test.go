package player

import (
	"testing"

	"main/Cartas"
)

func TestPlayerInitialization(t *testing.T) {
	player1 := new(Player)

	t.Run("test player name is correcty initialized", func(t *testing.T) {
		player1.AddName("german")
		got := player1.ReturnName()
		want := "german"

		if want != got {
			t.Errorf("got %v when it wanted %v", got, want)
		}
	})

	t.Run("test player can added Card Correctly", func(t *testing.T) {
		player1.AddCard("B", 3)
		got, existsInHand := player1.SearchCard("B", 3)
		want := Cartas.Card{Numero: 3, CardType: "B"}

		if existsInHand == false {
			t.Errorf("card not found in player hand")
		}

		checkedCardEqueals := Cartas.EqualCards(Cartas.Card{Numero: 3, CardType: "B"}, want)
		if checkedCardEqueals == false {
			t.Errorf("I wanted card %v and got card %v", want, got)
		}
	})
}
