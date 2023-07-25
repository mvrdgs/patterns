package states

import (
	"fmt"
)

type Dispense struct {
	state *state
}

func NewDispense(state *state) Interface {
	fmt.Println("Confirme seu pedido, itens selecionados:")
	for _, item := range state.SelectedItems {
		fmt.Printf(" - %s\n", item.Name)
	}
	return &Dispense{
		state: state,
	}
}

func (i *Dispense) InsertCoin(_ int) {
	fmt.Println("Finalize o pedido ou retorne à tela inicial para inserir mais dinheiro")
}

func (i *Dispense) Select(_ int) {
	fmt.Println("Finalize o pedido ou retorne à tela anterior para selecionar mais itens")
}

func (i *Dispense) Remove(_ int) {
	fmt.Println("Retorne à tela anterior para remover itens")
}

func (i *Dispense) Cancel() {
	fmt.Println("Retornar à tela de seleção, altere os itens selecionados ou retorne à tela anterior")
	i.state.SetState(NewSelect(i.state))
}

func (i *Dispense) Next() {
	defer fmt.Println("Pedido finalizado!")
	if i.state.InsertedMoney > 0 {
		fmt.Printf("Troco: %.2f\n", float64(i.state.InsertedMoney)/100)
	}
	i.state.InsertedMoney = 0
	i.state.SelectedItems = nil
	i.state.SetState(NewReady(i.state))
}
