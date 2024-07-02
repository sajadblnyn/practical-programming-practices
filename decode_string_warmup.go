package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "T4 l16 _36 510 _27 s26 _11 320 414 {6 }39 C2 T0 m28 317 y35 d31 F1 m22 g19 d38 z34 423 l15 329 c12 ;37 19 h13 _30 F5 t7 C3 325 z33 _21 h8 n18 132 k24"

	sl := strings.Split(s, " ")

	ls := make(map[int]string)
	i := 0
	for _, v := range sl {
		i, _ = strconv.Atoi(strings.Join(strings.Split(v, "")[1:], ""))

		ls[i] = strings.Split(v, "")[0]
	}

	v := ""
	for i := 0; i < len(ls); i++ {
		v += ls[i]
	}
	fmt.Println(v)
}
