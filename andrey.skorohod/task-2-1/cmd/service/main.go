package main

import (
	"fmt"
)

const (
	highBord = 30
	lowBord  = 15
)

func main() {
	var departments, employees int
	if _, err := fmt.Scan(&departments); err != nil || departments < 1 || departments > 1000 {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departments {
		if _, err := fmt.Scan(&employees); err != nil || employees < 1 || employees > 1000 {
			fmt.Println("Invalid number of employees")

			return
		}

		low, high := lowBord, highBord

		for range employees {
			var (
				sign string
				temp int
			)

			if _, err := fmt.Scan(&sign, &temp); err != nil {
				fmt.Println("Invalid temperature value")

				return
			}

			switch sign {
			case ">=":
				if temp > low {
					low = temp
				}
			case "<=":
				if temp < high {
					high = temp
				}
			default:
				fmt.Println("Invalid sign")

				return
			}

			if low > high {
				fmt.Println("-1")
			} else {
				fmt.Println(low)
			}
		}
	}
}
