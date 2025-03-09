package main

import (
	"fmt"
	"time"
)

func switch_main() {
	i := 1
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	fmt.Println( time.Now().Date())
	fmt.Println( time.Now().Clock())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default: fmt.Println("It's a week day")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default: 
	fmt.Println("It's after noon")
	}

	whatAmI := func (i interface{})  {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int: 
			fmt.Println("I'm an int")
		default: 
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}