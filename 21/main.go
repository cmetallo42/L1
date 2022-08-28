package main

import "fmt"

type (
	Human struct {
		Food
	}
	Robot struct {
		Energy
	}
	Adapter struct {
		*Human
		*Robot
	}
	Food interface {
		GiveFood()
	}
	Energy interface {
		GiveEnergy()
	}
)

func (h *Human) GiveFood() {
	fmt.Println("Human is eating Food")
}

func (r *Robot) GiveEnergy() {
	fmt.Println("Robot is getting Energy")
}

func (r *Robot) SomeFood() {
	fmt.Println("Robot has no idea what to do with food")
}

func (h *Human) SomeEnergy() {
	fmt.Println("Human having a good time with 220V cable")
}

func (a *Adapter) GiveFood() {
	a.SomeEnergy()
}

func (a *Adapter) GiveEnergy() {
	a.SomeFood()
}

func HumanToRobot(h *Human) Energy {
	return &Adapter{}
}

func RobotToHuman(r *Robot) Food {
	return &Adapter{}
}

func main() {
	human := Human{}
	robot := Robot{}

	human.GiveFood()
	robot.GiveEnergy()

	humanToRobotAdapter := HumanToRobot(&human)
	humanToRobotAdapter.GiveEnergy()

	robotToHumanAdapter := RobotToHuman(&robot)
	robotToHumanAdapter.GiveFood()
}
