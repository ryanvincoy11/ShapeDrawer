// Header comment here!

package main

import (
	"fmt"
)

var display Display
var validColors = map[string]Color{
	"red":     {255, 0, 0},
	"green":   {0, 255, 0},
	"blue":    {0, 0, 255},
	"yellow":  {255, 255, 0},
	"orange":  {255, 164, 0},
	"purple":  {128, 0, 128},
	"brown":   {165, 42, 42},
	"black":   {0, 0, 0},
	"white":   {255, 255, 255},
	"invalid": {-1, -1, -1},
}

func checkColor(s string) (c Color) {
	if s == "red" || s == "green" || s == "blue" || s == "yellow" || s == "orange" || s == "purple" || s == "brown" || s == "black" || s == "white" {
		return validColors[s]
	}
	return validColors["invalid"]
}

func main() {
	// Print intro message

	fmt.Println("Homework 5: Geometry Using Go Interfaces")
	fmt.Println("CS 341, Fall 2025")
	fmt.Println("")
	fmt.Println("This application allows you to draw various shapes")
	fmt.Println("of different colors via interfaces in Go.")
	fmt.Println("")

	//
	// SOME PRINT STATEMENTS YOU WILL NEED CAN BE FOUND BELOW
	//
	// Ask user for dimensions for display
	var x, y int
	fmt.Print("Enter the number of rows (x) that you would like the display to have: ")
	fmt.Scan(&x)
	fmt.Print("Enter the number of columns (y) that you would like the display to have: ")
	fmt.Scan(&y)

	display.initialize(x, y)
	//
	// Menu options
	for {
		fmt.Println("")
		fmt.Println("Select a shape to draw: ")
		fmt.Println("	 R for a rectangle")
		fmt.Println("	 T for a triangle")
		fmt.Println("	 C for a circle")
		fmt.Println(" or X to stop drawing shapes.")
		fmt.Print("Your choice --> ")
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "R":
			//
			// Drawing a rectangle
			var x1, y1, x2, y2 int
			var color string
			fmt.Print("Enter the X and Y values of the lower left corner of the rectangle: ")
			fmt.Scan(&x1, &y1)
			fmt.Print("Enter the X and Y values of the upper right corner of the rectangle: ")
			fmt.Scan(&x2, &y2)
			fmt.Print("Enter the color of the rectangle: ")
			fmt.Scan(&color)
			fmt.Printf("Rectangle: (%d,%d) to (%d,%d)\n", x1, y1, x2, y2)

			var c = checkColor(color)
			rect := Rectangle{Point{x1, y1}, Point{x2, y2}, c}
			if err := rect.draw(&display); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Rectangle drawn successfully.")
			}

			//
			// Drawing a triangle
		case "T":
			var x0, y0, x1, y1, x2, y2 int
			var color string
			fmt.Print("Enter the X and Y values of the first point of the triangle: ")
			fmt.Scan(&x0, &y0)
			fmt.Print("Enter the X and Y values of the second point of the triangle: ")
			fmt.Scan(&x1, &y1)
			fmt.Print("Enter the X and Y values of the third point of the triangle: ")
			fmt.Scan(&x2, &y2)
			fmt.Print("Enter the color of the triangle: ")
			fmt.Scan(&color)
			fmt.Printf("Triangle: (%d,%d), (%d,%d), (%d,%d)\n", x0, y0, x1, y1, x2, y2)

			var c = checkColor(color)
			tri := Triangle{Point{x0, y0}, Point{x1, y1}, Point{x2, y2}, c}
			if err := tri.draw(&display); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Triangle drawn successfully.")
			}

			//
			// Drawing a circle
		case "C":
			var cx, cy, rad int
			var color string
			fmt.Print("Enter the X and Y values of the center of the circle: ")
			fmt.Scan(&cx, &cy)
			fmt.Print("Enter the value of the radius of the circle: ")
			fmt.Scan(&rad)
			fmt.Print("Enter the color of the circle: ")
			fmt.Scan(&color)
			fmt.Printf("Circle: centered around (%d,%d) with radius %d\n", cx, cy, rad)
			var c = checkColor(color)
			circ := Circle{Point{cx, cy}, rad, c}
			if err := circ.draw(&display); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Circle drawn successfully.")
			}

		case "X":
			goto SAVE

		default:
			fmt.Println("**Error, unknown command. Try again.")
		}
	}
	//
	// Saving the results in a file
SAVE:
	var filename string
	fmt.Print("Enter the name of the .ppm file in which the results should be saved: ")
	fmt.Scan(&filename)
	//
	// Exiting program
	if err := display.screenShot(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done. Exiting program...")
	}
}
