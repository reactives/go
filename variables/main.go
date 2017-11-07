/* пример многострочный комментарий*/
// Переменные в golang
package main

func main() {

	InvalFunc();

	// пример объявления  чисeл с плавающей точкой
	var pl float32 = 33.45 // float = float32, float64

	println("float32: ", pl)

	// пример объявления булевыx переменныx
	var varbool = true
	println("bool variable", varbool)

	// пример объявления строки
	var hello string = "Hello\n\t"
	var world = "World"
	println(hello, world)

	// пример объявления бинарныx данныx
	var rawBinaryData byte = '\x27'
	println("rawBinaryData", rawBinaryData)

	// так нельзя
	// var singleQuote string = 'Hello world'
	// missing '
	// syntax error: unexpected ello at end of statement

	/*
		пример короткое объявление
	*/
	meaningOfLive := 42
	println("Meaning of life is ", meaningOfLive)
	// работает только для новых переменных, world объявлен выше, поэтому ошибка
	// world := "Мир"
	// no new variables on left side of :=

	/*
		пример объявления приведение типов
	*/
	println("float to int conversion ", int(pl))
	println("int to string conversion ", string(48))

	//пример объявления комплексныx чисeл
	z := 2 + 3i
	println("complex number: ", z)

	/*
		пример операций со строками
	*/
	str1 := "Reactives"
	str2 := "ru"
	fullName := str1 + str2
	println("site length is: ", fullName, len(fullName))

	escaping := `Reactives\r\n
	Ru`
	println("as-is escaping: ", escaping)

	/*
		пример объявления значение c значением по-умолчанию
	*/
	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println("default values: ", defaultInt, defaultFloat, defaultString, defaultBool)

	/*
		пример объявления нескольких переменных
	*/
	var var1, var2 string = "val1", "val2"
	println(var1, var2)

	var (
		p0 int = 22
		p2     = "string example"
		p3     = 56
	)
	println(p0, p2, p3)

}

func InvalFunc()  {

	// пример объявления целых чисел
	var i int = 10							// платформозависимый тип, 32 или 64 бита
	var autoInt = -10
	var bigInt int64 = 1<<32 - 1			// int8, int16, int32, int64
	var unsignedInt uint = 100500			// платформозависимый тип, 32 или 64 бита
	var unsignedBigInt uint64 = 1<<64 - 1 	// uint8, unit16, uint32, unit64

	println("integers", i, autoInt, bigInt, unsignedInt, unsignedBigInt)
}
