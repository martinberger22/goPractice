package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Probably should have upped your sneak skill. [Heist Failed]")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("The vault is open, grab as much as you can and try to make it out!")
	} else if isHeistOn && openedVault <= 70 {
		isHeistOn = false
		fmt.Println("Couldn't open the vault, better up your lock picking skills next time. [Heist Failed]")
	}

	leftSafely := rand.Intn(6)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("RIP, guards snuck up on you while you were loading the cash. [Heist Failed]")
		case 1:
			isHeistOn = false
			fmt.Println("Next time make sure the door doesn't lock you in. [Heist Failed]")
		case 2:
			isHeistOn = false
			fmt.Println("Maybe don't hire an undercover FBI agent next time. [Heist Failed]")
		case 3:
			isHeistOn = false
			fmt.Println("Probably shouldn't have tried to carry so much, too slow. [Heist Failed]")
		default:
			fmt.Println("Start the getaway car!")
		}

		if isHeistOn {
			amtStolen := 10000 + rand.Intn(1000000)
			fmt.Println(amtStolen)
		}

	}

	fmt.Println(isHeistOn)
}
