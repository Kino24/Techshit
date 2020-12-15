package main

import (
	"fmt"
	"math"
)

//distance of two points
type distance struct {
	X float64
	Y float64
}

func main() {
	fmt.Println("Choose a calculator:")
	fmt.Println("Stoichiometry[1] or Distance Between 2 points?[2]")
	fmt.Println("Enter a number")
	var ch int
	fmt.Scanln(&ch)

	if ch == 2 {
		fmt.Println("Enter value for x1")
		var x1 float64
		fmt.Scanln(&x1)
		fmt.Println("Enter value for y1")
		var y1 float64
		fmt.Scanln(&y1)

		v := distance{0, 0}
		v.X = x1
		v.Y = y1

		fmt.Println("Enter value for x2")
		var x2 float64
		fmt.Scanln(&x2)
		fmt.Println("Enter value for y2")
		var y2 float64
		fmt.Scanln(&y2)

		b := distance{0, 0}
		b.X = x2
		b.Y = y2

		var xans = (b.X - v.X)
		var yans = (b.Y - v.Y)
		var a = math.Pow(xans, 2)
		var d = math.Pow(yans, 2)
		var c = a + d
		var ans = math.Sqrt(c)
		fmt.Println("The distance between point", v, " and point", b, "is :")
		fmt.Println(ans)
	} else {
		if ch == 1 {
			var choice = 1
			var choice2 = 2

			fmt.Println("Please enter a number.")
			fmt.Println("moles[1] grams[2]")
			var find int
			fmt.Scanln(&find)

			if find == choice {
				fmt.Println("Enter grams of first element: ")
				var g1 float64
				fmt.Scanln(&g1)

				fmt.Println("Enter enter molar mass of first element: ")
				var mol1 float64
				fmt.Scanln(&mol1)

				var moles1 = g1 * 1 / mol1

				fmt.Println("Enter Mole Ratio of second element: ")
				var mr1 float64
				fmt.Scanln(&mr1)

				fmt.Println("Enter mole ratio of first element: ")
				var mr2 float64
				fmt.Scanln(&mr2)

				var mr = mr1 / mr2
				var moles2 = moles1 * mr
				var finale = math.Round(moles2)

				fmt.Println("The element contains ", finale, " moles.")
			} else {
				if find == choice2 {
					fmt.Println("Enter grams of first element: ")
					var g1 float32
					fmt.Scanln(&g1)

					fmt.Println("Enter enter molar mass of first element: ")
					var mol1 float32
					fmt.Scanln(&mol1)

					var moles1 = g1 * 1 / mol1

					fmt.Println("Enter Mole Ratio of second element: ")
					var mr1 float32
					fmt.Scanln(&mr1)

					fmt.Println("Enter mole ratio of first element: ")
					var mr2 float32
					fmt.Scanln(&mr2)

					var mr = mr1 / mr2
					var moles2 = moles1 * mr

					fmt.Println("Enter molar mass of second element")
					var mol2 float32
					fmt.Scanln(&mol2)

					var mm = moles2 * mol2

					fmt.Println("The element weighs ", mm, " grams.")
				} else {
					fmt.Println(" ")
					fmt.Println(find, "is not part of the choices bitch!")
				}
			}
		} else {
			fmt.Println(ch, "is not part of the choices mother fu-!")
		}

	}
}
