package main

/*
	Блок импорта внешних зависимостей или стандартных библиотек
*/

import (
	"fmt"
	"math"
)

/*
	Блок создания НЕизменяемых переменных
*/

const (
	Pi           = math.Pi // реимпорт из другого пакета
	Greet string = "Привет"
)

/*
	Блок создания изменяемых переменных
*/

var (
	globalVar = 0 // переменная в глобальной области

	a int     = 1    // псевдоним для int32
	b bool    = true // истина иили лож
	c uint32  = 3    // только положительные целые числа
	d int32   = 4    // любые целые числа
	e rune    = 'f'  // один символ из строки
	f float32 = 3.14 // число с плавающей точкой одинарной точности (float64 двойной точности)
	g byte    = 255  // минимально возможная адресуемая единица памяти 1 байт (8 бит)
)

func myPrintln(greet string, a, b int) (int, string) {
	fmt.Println(greet, a, b)
	return a + b, greet + " друг"
}

/*
	main входная точка в программу
	ничего не принимает и не возвращает
*/

func main() {
	var localVar = 0 // переменная в локальной области
	localVar, message := myPrintln(Greet, localVar, int(d))
	fmt.Println(localVar, message)
}
