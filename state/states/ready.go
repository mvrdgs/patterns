package states

import (
	"fmt"
)

type Ready struct {
	state *state
}

func NewReady(state *state) Interface {
	fmt.Println("Olá! Insira o valor desejado")
	return &Ready{
		state: state,
	}
}

func (i *Ready) InsertCoin(money int) {
	i.state.InsertedMoney += money
	fmt.Printf("Dinheiro inserido: %.2f. Total: %.2f\n", float64(money)/100, float64(i.state.InsertedMoney)/100)
}

func (i *Ready) Select(_ int) {
	fmt.Println("Adicione o valor e clique em prosseguir")
}

func (i *Ready) Remove(_ int) {
	return
}

func (i *Ready) Next() {
	i.state.SetState(NewSelect(i.state))
}

func (i *Ready) Cancel() {
	fmt.Printf("Pedido cancelado, valor reembolsado: %.2f\n", float64(i.state.InsertedMoney)/100)
	i.state.InsertedMoney = 0
}
