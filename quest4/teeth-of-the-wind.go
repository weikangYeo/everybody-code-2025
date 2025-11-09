package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// part 1
	part1Total := 2025 * 1000 / 130
	fmt.Printf("Part1: %d\n", part1Total)

	// part 2
	// part2Total := 10000000000000 * 13 / 102
	part2Total := 10000000000000*289/993 + 1
	fmt.Printf("Part2: %d\n", part2Total)

	// part 3
	input := strings.TrimSpace(`
	619
608|608
592|1184
582|582
576|2304
573|573
568|2272
567|567
563|1689
562|562
551|1102
545|545
527|1054
513|513
511|2044
508|508
490|1470
479|479
467|1868
459|459
458|458
445|445
432|864
428|428
421|842
420|420
415|1660
402|402
383|766
375|375
374|1496
369|369
356|356
353|353
341|1023
338|338
324|324
307|307
302|604
294|294
293|1172
283|283
282|282
263|263
253|506
234|234
225|225
209|209
191|573
171`)

	// sanitizing input, so it will be seperated by ","
	input = strings.ReplaceAll(input, "\t", "")
	input = strings.ReplaceAll(input, "\n", ",")
	input = strings.ReplaceAll(input, "|", ",")
	sanitizedInput := strings.Split(input, ",")

	// a 2 way flag, true if divide, false if multiple
	isDivide := false

	total := 100.0
	for _, value := range sanitizedInput {
		parsedFloat, err := strconv.ParseFloat(value, 64)

		if err != nil {
			fmt.Println("something went wrong ", err)
		}

		if isDivide {
			total /= parsedFloat
		} else {
			total *= parsedFloat
		}
		isDivide = !isDivide
	}
	fmt.Printf("Total: %f\n", total)
}
