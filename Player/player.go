package player

import (
	"main/Cartas"
)

type Player struct {
	name string
	Hand []Cartas.Card
}

func (p *Player) AddName(NewName string) {
	p.name = NewName
}

func (p Player) ReturnName() string {
	return p.name
}

func (p *Player) AddCard(typeOfCardNew string, cardNumberNew int) {
	cardToAdd := Cartas.Card{CardType: typeOfCardNew, Numero: cardNumberNew}
	p.Hand = append(p.Hand, cardToAdd)
}

func (p Player) SearchCard(typeToSearch string, numberToSearch int) (Cartas.Card, bool) {
	cartaToReturn := Cartas.Card{CardType: "", Numero: 0}
	foundCard := false

	for _, v := range p.Hand {
		foundCard = Cartas.EqualCards(
			v,
			Cartas.Card{CardType: typeToSearch, Numero: numberToSearch},
		)
		if foundCard {
			cartaToReturn.CardType = typeToSearch
			cartaToReturn.Numero = numberToSearch
		}
	}

	return cartaToReturn, foundCard
}
