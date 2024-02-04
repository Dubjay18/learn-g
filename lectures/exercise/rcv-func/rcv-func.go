//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Health uint
type Energy uint
type Name string

type Player struct {
	health    Health
	maxHealth Health
	energy    Energy
	maxEnergy Energy
	name      Name
}

func (p *Player) modifyHealth(health Health) {
	if p.health+health > p.maxHealth {
		health = p.maxHealth - p.health
	}
	p.health += health
	fmt.Println(p.name, "health is now", p.health, "added", health)
}

func (p *Player) modifyEnergy(energy Energy) {
	if p.energy+energy > p.maxEnergy {
		energy = p.maxEnergy - p.energy
	}
	p.energy += energy
	fmt.Println(p.name, "energy is now", p.energy, "added", energy)
}

func main() {
	master := Player{maxHealth: 100, maxEnergy: 100, name: "Master"}
	master.modifyHealth(50)
	master.modifyEnergy(50)

	fmt.Println(master)
}
