package main

import (
	"fmt"
)

type Board struct {
	NailsDriven int
	NailsNeeded int
}

//NailDriver drives a nail into a board
type NailDriver interface {
	DriveNail(nailInventory *int, b *Board)
}

//NailPuller pulls a nail from a board
type NailPuller interface {
	PullNail(nailInventory *int, b *Board)
}

//NailDrivePuller drives a nail into a board and pulls a nail from board
type NailDrivePuller interface {
	NailDriver
	NailPuller
}

//Mallet has no state
type Mallet struct{}

//DriveNail implements NailDriver
func (m Mallet) DriveNail(nailInventory *int, b *Board) {
	*nailInventory--
	b.NailsDriven++
}

//CrowBar has no state
type CrowBar struct{}

//PullNail implements NailPuller
func (c CrowBar) PullNail(nailInventory *int, b *Board) {
	*nailInventory++
	b.NailsNeeded++
}

//Toolbox has a NailDriver, a NailPuller and a nails
type Toolbox struct {
	NailDriver
	NailPuller
	nails int
}

//Contractor has no state
type Contractor struct{}

//Fasten a board with a NailDriver and nails from an inventory
func (c Contractor) Fasten(nd NailDriver, nailInventory *int, b *Board) {
	for b.NailsDriven < b.NailsNeeded {
		nd.DriveNail(nailInventory, b)
		b.NailsDriven++
	}
	fmt.Println("Board has been fastened")
}

//UnFasten a board with a NailDriver and return nails back to inventory
func (c Contractor) UnFasten(np NailPuller, nailInventory *int, b *Board) {
	for b.NailsDriven > 0 {
		np.PullNail(nailInventory, b)
		b.NailsDriven--
	}
	fmt.Println("Board has been unfastened")
}

//ProcessBoards fasten and unfasten multiple boards with a NailDriverPuller and an inventory of nails
func (c Contractor) ProcessBoards(ndp NailDrivePuller, nailInventory *int, boards []Board) {
	for _, v := range boards {
		switch {
		case v.NailsDriven > v.NailsNeeded:
			{
				c.UnFasten(ndp, nailInventory, &v)
			}
		case v.NailsDriven < v.NailsNeeded:
			{
				c.Fasten(ndp, nailInventory, &v)
			}
		}
	}
}

func main() {
	tb := Toolbox{
		NailDriver: Mallet{},
		NailPuller: CrowBar{},
		nails:      10,
	}

	boards := []Board{
		{NailsDriven: 6},
		{NailsDriven: 4},
		{NailsNeeded: 5},
		{NailsNeeded: 7},
	}

	c := Contractor{}
	c.ProcessBoards(tb, &tb.nails, boards)

}
