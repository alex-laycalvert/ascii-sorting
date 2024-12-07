package main

import (
	"math/rand"
	"time"

	"github.com/gbin/goncurses"
)

const (
	SELECTED_PAIR = 1
	CORRECT_PAIR  = 2
)

func main() {
	w, err := goncurses.Init()
	panicErr(err)
	w.Clear()
	goncurses.Cursor(0)
	goncurses.StartColor()
	goncurses.UseDefaultColors()
	goncurses.InitPair(SELECTED_PAIR, goncurses.C_RED, -1)
	goncurses.InitPair(CORRECT_PAIR, goncurses.C_GREEN, -1)

	defer w.Clear()
	defer goncurses.Cursor(1)
	defer goncurses.End()

	termHeight, termWidth := w.MaxYX()

	maxWidth := termWidth - 2
	maxHeight := termHeight - 1

	count := min(maxWidth, maxHeight*2)
	start := (termWidth - count) / 2
	// end := (termWidth + count) / 2

	numbers := random(count)

	for i := range numbers {
		for j := range len(numbers) - i - 1 {
			w.Clear()
			drawNumbers(w, start, termHeight, numbers, j, len(numbers)-i)
			w.Refresh()
			time.Sleep(20 * time.Millisecond)
			if numbers[j] <= numbers[j+1] {
				continue
			}
			numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
		}
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func random(length int) []int {
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = i + 1
	}

	for i := len(numbers) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers
}

func drawNumbers(w *goncurses.Window, start int, termHeight int, numbers []int, selected int, firstCorrectIndex int) {
	for i, number := range numbers {
		if selected == i {
			w.ColorOn(SELECTED_PAIR)
		} else if i >= firstCorrectIndex {
			w.ColorOn(CORRECT_PAIR)
		}
		for j := range number / 2 {
			w.MoveAddChar(termHeight-j, start+i, goncurses.ACS_BLOCK)
		}
		w.ColorOff(SELECTED_PAIR)
		w.ColorOff(CORRECT_PAIR)
	}
}
