package lucky

import (
	"log"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

func (d *Draws) Generate(disableRepeatedDrawCheck bool) *Draw {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var numbers []int
	var stars []int
	draw := &Draw{}
	for {
		for {
			numbers = r.Perm(50)[:5]
			if !duplicates(numbers) {
				break
			}
		}

		for {
			stars = r.Perm(12)[:2]
			if !duplicates(stars) {
				break
			}
		}

		dealWithZeros(numbers)
		dealWithZeros(stars)
		sort.Ints(numbers)
		sort.Ints(stars)

		draw.Numbers = numbers
		draw.Starts = stars

		if disableRepeatedDrawCheck {
			break
		}
		if !d.isARepeatedDraw(draw) {
			break
		}
	}
	return &Draw{
		Numbers: numbers,
		Starts:  stars,
	}
}

func duplicates(slice []int) bool {
	keys := make(map[int]bool)
	for _, entry := range slice {
		if _, value := keys[entry]; value {
			return true
		}
		keys[entry] = true
	}

	return false
}

func dealWithZeros(slice []int) {
	for i := range slice {
		slice[i] += 1
	}
}

func (d *Draws) isARepeatedDraw(draw *Draw) bool {
	for _, dr := range *d {
		if reflect.DeepEqual(draw.Numbers, dr.Numbers) &&
			reflect.DeepEqual(draw.Starts, dr.Starts) {
			log.Println("Found repeated sequence...")
			return true
		}
	}

	return false
}
