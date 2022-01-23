package main

import "fmt"

type Range struct {
	start, end uint64
	print      bool
}

func main() {
	rgs := []Range{
		{2, 100, true},
	}
	for _, rg := range rgs {
		if rg.start == 2 && rg.end <= 100 {
			fmt.Printf("Following are all Eban numbers upto %d:\n", rg.end)
		}
		count := 0
		for i := rg.start; i <= rg.end; i += 2 {
			b := i / 1000000000
			r := i % 1000000000
			m := r / 1000000
			r = i % 1000000
			t := r / 1000
			r %= 1000
			if m >= 30 && m <= 66 {
				m %= 10
			}
			if t >= 30 && t <= 66 {
				t %= 10
			}
			if r >= 30 && r <= 66 {
				r %= 10
			}
			if b == 0 || b == 2 || b == 4 || b == 6 {
				if m == 0 || m == 2 || m == 4 || m == 6 {
					if t == 0 || t == 2 || t == 4 || t == 6 {
						if r == 0 || r == 2 || r == 4 || r == 6 {
							if rg.print {
								fmt.Printf("%d ", i)
							}
							count++
						}
					}
				}
			}
		}
	}
}