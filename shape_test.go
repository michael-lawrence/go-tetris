package main

import (
	"fmt"
	"log"
	"testing"
)

func printShape(s *Shape) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			v := s[i][j]
			fmt.Printf("%v, ", v)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func compareShapes(s1 *Shape, s2 *Shape) error {
	len1 := len(s1)
	len2 := len(s1)

	if len1 != len2 {
		printShape(s1)
		fmt.Printf("expected\n")
		printShape(s2)
		return fmt.Errorf("shapes are not the same length: %v, %v", len1, len2)
	}

	for i := 0; i < len(s1); i++ {
		row1 := s1[i]
		row2 := s2[i]
		row1Len := len(row1)
		row2Len := len(row2)

		if row1Len != row2Len {
			printShape(s1)
			fmt.Printf("expected\n")
			printShape(s2)
			return fmt.Errorf("shapes rows (%v) are not the same length: %v, %v", i, row1Len, row2Len)
		}

		for j := 0; j < row1Len; j++ {
			val1 := row1[j]
			val2 := row2[j]

			if val1 != val2 {
				printShape(s1)
				fmt.Printf("expected\n")
				printShape(s2)
				return fmt.Errorf("shapes are not equal at position: [%v][%v], values are %v, %v", i, j, val1, val2)
			}
		}
	}

	return nil
}

// TestRotate Tests rotating shapes
func TestRotate(t *testing.T) {
	shape := Shape{
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expectations := [4]Shape{
		{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
		},
		{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
		},
		{
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
		},
		{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	}

	for i := 0; i < len(expectations); i++ {
		shape.Rotate()
		err := compareShapes(&shape, &expectations[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}
