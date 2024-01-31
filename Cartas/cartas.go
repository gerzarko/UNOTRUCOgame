package Cartas


type Card struct{

    CardType string
    Numero int

}

func (c *Card) ReturnValues()(string,int){


    return c.CardType,c.Numero


}
