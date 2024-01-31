package Cartas

import (
	"fmt"
	"testing"
)

func TestCardValue(t *testing.T) {

	card1 := Card{
		CardType: "B",
		Numero:   2,
	}

	t.Run("return card content", func(t *testing.T) {

		cardType, number := card1.ReturnValues()
		wantType, wantNumber := "B", 2

		if cardType == wantType && number == wantNumber {

			fmt.Printf("paso el test")

		} else {
			t.Errorf("test fail,got type:%s and wanted %s", cardType, wantType)
			t.Errorf("test fail,got type:%d and wanted %d", number, wantNumber)

		}

	})

}
