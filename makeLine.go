package utils

// 추세선을 그리기 위해 만들었음.

func MakeLine(p1, p2 float64, di int) []float64 {
	var rtv []float64
	if p1 > p2 {
		for i := 1; i < 100; i++ {
			v := (p1 - p2) / float64(di)
			rtv = append(rtv, p2-(v*float64(i)))
		}
	} else if p2 > p1 {
		for i := 1; i < 100; i++ {
			v := (p2 - p1) / float64(di)
			rtv = append(rtv, p2+(v*float64(i)))
		}
	}

	return rtv
}

/*
example:
func main() {
	p2 := 42200.0
	p1 := 42162.2
	di := 14

	fmt.Println("p2 이후의 추세선의 값")
	var res = makeLine_(p2, p1, di)
	fmt.Print(res)
}
*/
