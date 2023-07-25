package states

import (
	"fmt"
)

type Select struct {
	state *state
}

func NewSelect(state *state) Interface {
	fmt.Println("Selecione o(s) item(s) desejado(s)")
	return &Select{
		state: state,
	}
}

func (i *Select) InsertCoin(_ int) {
	fmt.Println("Retorne à tela anterior para inserir mais dinheiro")
}

func (i *Select) Select(itemID int) {
	item, ok := i.state.ItemList[itemID]
	if !ok {
		fmt.Println("Item não encontrado")
		return
	}
	if item.Price > i.state.InsertedMoney {
		fmt.Println("Dinheiro insuficiente para o item selecionado")
		fmt.Printf("Item: %s, Preço: %.2f\n", item.Name, float64(item.Price)/100)
		fmt.Printf("Dinheiro restante: %.2f\n", float64(i.state.InsertedMoney)/100)
		return
	}
	i.state.SelectedItems = append(i.state.SelectedItems, item)
	i.state.InsertedMoney -= item.Price
	fmt.Printf("- Item %s selecionado\n", item.Name)
	fmt.Printf("Dinheiro restante: %.2f\n", float64(i.state.InsertedMoney)/100)
}

func (i *Select) Remove(index int) {
	if index >= len(i.state.SelectedItems) {
		fmt.Println("Item não encontrado")
		return
	}
	item := i.state.SelectedItems[index]
	i.state.SelectedItems = append(i.state.SelectedItems[:index], i.state.SelectedItems[index+1:]...)
	i.state.InsertedMoney += item.Price
	fmt.Printf("Item %s removido\n", item.Name)
	fmt.Printf("Dinheiro restante: %.2f\n", float64(i.state.InsertedMoney)/100)
}

func (i *Select) Cancel() {
	for _, item := range i.state.SelectedItems {
		i.state.InsertedMoney += item.Price
	}
	fmt.Printf("Insira o valor desejado ou cancele o pedido. Valor já inserido: %.2f\n", float64(i.state.InsertedMoney)/100)
	i.state.SelectedItems = nil
	i.state.SetState(NewReady(i.state))
}

func (i *Select) Next() {
	if len(i.state.SelectedItems) == 0 {
		fmt.Println("Selecione pelo menos um item antes de prosseguir")
		return
	}
	i.state.SetState(NewDispense(i.state))
}
