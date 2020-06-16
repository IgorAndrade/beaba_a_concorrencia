package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s string
	s = "oi"
	fmt.Println(s)

	ss := "Texto qualquer"
	fmt.Println(ss)

	var i int
	i = 6
	fmt.Println(i)
	ii := 8
	fmt.Println(ii)

	var f float64 = 8.5
	fmt.Println(f)

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)
	ponteiro()
}

func ponteiro() {
	i := 6
	var p *int
	p = &i // pegando o endereço da variável
	fmt.Println(p, i, &i)
	fmt.Println(p, *p, i, &i)
	// Go não tem aritmética de ponteiros
	// p++
	*p++ // desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)

}
