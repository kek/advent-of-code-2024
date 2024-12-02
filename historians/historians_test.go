package historians

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("Test Advent of Code 2024", func(t *testing.T) {
		example := `3   4
4   3
2   5
1   3
3   9
3   3`

		gotLeft, gotRight := ParseInput(example)
		wantLeft, wantRight := []int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}

		if !reflect.DeepEqual(gotLeft, wantLeft) {
			t.Errorf("got %q want %q", gotLeft, wantLeft)
		}
		if !reflect.DeepEqual(gotRight, wantRight) {
			t.Errorf("got %q want %q", gotRight, wantRight)
		}
	})
}

func TestSolveDay1(t *testing.T) {
	t.Run("Test Solve Day 1", func(t *testing.T) {
		example := `3   4
4   3
2   5
1   3
3   9
3   3`
		left, right := ParseInput(example)
		got := SolveDay1(left, right)
		want := 11

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestDay1(t *testing.T) {
	t.Run("Test Day 1", func(t *testing.T) {
		// read day1-input.txt
		input := ReadFile("../day1-input.txt")

		left, right := ParseInput(input)
		got := SolveDay1(left, right)
		want := 2031679

		fmt.Println("Day 1: ", got)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	return string(data)
}
