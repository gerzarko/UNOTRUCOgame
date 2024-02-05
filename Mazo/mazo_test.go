package Mazo

import (
	"fmt"
	"testing"
)

func TestMazo(t *testing.T) {

	mazoNuevo := new(Mazo)

	t.Run("test good initialization", func(t *testing.T) {
        mazoNuevo.InitializeDeck()
        got := mazoNuevo.FirstCard()

        if got.CardType != "E" && got.Numero != 10{
            t.Errorf("wanted a card but got %v",got)
        }
        

	})
    
    t.Run("test shuffle cards",func(t *testing.T) {
        mazoNuevo.InsertCardInMazo("O",2)
        mazoNuevo.InsertCardInMazo("C",4)
        mazoNuevo.InsertCardInMazo("B",5)
        got := mazoNuevo.FirstCard() 
        mazoNuevo.Shuffle()
        want := mazoNuevo.FirstCard()
       
        if got == want{
            t.Errorf("wanted a different card but got same one %v",got)

        }
        fmt.Printf("%v %v",got,want)


    })

    fmt.Printf("%v ",mazoNuevo)

}
