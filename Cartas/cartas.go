package Cartas

type Card struct {
	CardType string
	Numero   int
}

func (c *Card) ReturnValues() (string, int) {
	return c.CardType, c.Numero
}

func EqualCards(carta1 Card, carta2 Card) bool {
	if carta1.Numero == carta2.Numero && carta1.CardType == carta2.CardType {
		return true
	}

	return false
}
