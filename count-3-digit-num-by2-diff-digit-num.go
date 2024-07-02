package main

func main() {

	// v := 0
	// for i := 100; i < 1000; i++ {
	// 	if i%10 == 0 {
	// 		continue
	// 	}
	// 	if i%111 == 0 || i%101 == 0 {
	// 		continue
	// 	}
	// 	if (i/10)%11 == 0 || (i%100)%11 == 0 || i%10 == i/100 {
	// 		v++
	// 	}
	// }
	// println(v)

	v := 0

	for i := 12; i < 99; i++ {
		if i%11 == 0 || i%10 == 0 {
			continue
		}
		v += 3
	}
	println(v)
}
