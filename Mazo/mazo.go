package Mazo

import (
	"main/Cartas"
	"math/rand"
)


type Mazo struct{

    AllCards []Cartas.Card

}


func (m *Mazo)InsertCardInMazo(cardType string,cardNumber int){

    cardToInsert := Cartas.Card{CardType: cardType,Numero: cardNumber} 
    m.AllCards = append(m.AllCards, cardToInsert)

}


func (m *Mazo)InitializeDeck(){





}


func (m Mazo) FirstCard()Cartas.Card{
    
    return m.AllCards[0]

}

func (m *Mazo)Shuffle(){

    rand.Shuffle(len(m.AllCards), func(i, j int) {
        m.AllCards[i], m.AllCards[j] = m.AllCards[j], m.AllCards[i]
    })

}
