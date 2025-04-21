// Starting code for the exercise
package main

import "fmt"

// Item struct
type Item struct {
	Name string
	Type string
}

// Player struct
type Player struct {
	Name      string
	Inventory []Item
}


func main() {
	// Create a new player
	player := Player{
		Name:      "Hero",
		Inventory: []Item{},
	}
	// Create a new item
	item := Item{
		Name: "Health Potion",
		Type: "Potion",
	}
	// Add item to player's inventory
	player.PickUpItem(item)
	// Print player's inventory
	for _, i := range player.Inventory {
		println("Item in inventory:", i.Name)
	}
}

// Pick up an item (modifies the player's inventory)
func (p *Player) PickUpItem(item Item) {
	// TODO: Implement function to add an item to inventory
	p.Inventory = append(p.Inventory, item)
	fmt.Println("Picked up item:", item.Name)

}

// Drop an item (removes from inventory)
func (p *Player) DropItem(itemName string) {
	// TODO: Implement function to remove an item from inventory
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Println("Dropped item:", item.Name)
			return
		}
	}
	fmt.Println("Item not found in inventory:", itemName)
}

// Use an item (if potion, remove it after use)
func (p *Player) UseItem(itemName string) {
	// TODO: Implement function to use an item
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "Potion" {
				fmt.Println("Used potion:", item.Name)
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			} else {
				fmt.Println("Used item:", item.Name)
			}
			return
		}
	}
	fmt.Println("Item not found in inventory:", itemName)
}
