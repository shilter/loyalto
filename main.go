    package main

	import (
		"log"
		"os"	
		"github.com/jcheng31/diceroller/dice github.com/jcheng31/diceroller/roller"
	)

	func main() {
		predetermined := roller.WithSequence([]int{1, 2, 3, 4, 5})

		src := rand.NewSource(time.Now().UnixNano())
		random := roller.WithRandomSource(src)
	
		seededSrc := rand.NewSource(42)
		seededRandom := roller.WithRandomSource(seededSrc)
	
		d6 := dice.Regular(seededRandom, 6)
	
		result := d6.RollN(4)
		fmt.Println(result.Total) 
		
		fmt.Println(result.Rolls) 
	
		seededSrc = rand.NewSource(42)
		seededRandom = roller.WithRandomSource(seededSrc)
	
		exploD6 := dice.Exploding(seededRandom, 6, 10)
	
		result = exploD6.RollN(1)
		fmt.Println(result.Total)
		fmt.Println(result.Rolls)
	}