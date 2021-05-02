package main

import (
	"errors"
	"fmt"
	"rsass/file"
)

// Создание ключей и возврат ошибки, если что-то пошло не так
func createKeys(num1 int, num2 int) error {
	type Keys struct {
		n  int
		fi int
		e  int
		d  int
	}

	if num1 == 1 || num2 == 1 { // Проверка на корректность введённых чисел
		fmt.Println("Числа должны быть > 1!")
		return errors.New("Bye")
	}
	var N int
	if num1 >= num2 {
		N = num1
	} else if num2 > num1 {
		N = num2
	}
	res, resNum := checkNums(N, num1, num2) // Проверка - простые числа были введены или нет.
	if res == false {
		fmt.Printf("Число %d не является простым!\n", resNum)
		return errors.New("Bye")
	}
	keys := Keys{
		n:  num1 * num2,
		fi: (num1 - 1) * (num2 - 1), //Функция Эйлера
	}
	keys.e = getE(keys.fi)
	keys.d = getD(keys.e, keys.fi)

	if keys.n <= 255 {
		fmt.Println("Выберите другие простые числа. Желательно начинать от 53")
		return errors.New("Bye")
	}
	fmt.Println("Открытый ключ: ", keys.e, keys.n)
	fmt.Println("Закрытый ключ: ", keys.d, keys.n)
	fmt.Println("Обязательно сохраните файлы keys.pubk и keys.privk! В противном случае вы НЕ сможете расшифровать/зашифровать данные!")
	err := file.CreateKeyFile(keys.e, keys.n, "keys.pubk") // Сохраняем открытый ключ в файле
	if err != nil {
		return err
	}
	err = file.CreateKeyFile(keys.d, keys.n, "keys.privk")
	if err != nil {
		return err
	}
	return nil
}

func getE(fi int) int { // Получение открытой эскпоненты.
	var index int
	count := fi
	for index = 2; index < count; index++ {
		if (index%1 == 0) && (fi%1 == 0) && (fi%index != 0) {
			break
		}
	}
	return index
}

func getD(e int, fi int) int { // Получение секретной экспоненты
	var d int
	var k int

	for k = 2; k < fi; k++ {
		d = (1 + fi*k) / e
		if d != e && ((e*d)%fi == 1) {
			break
		}
	}
	return d
}

func checkNums(n int, num1 int, num2 int) (bool, int) { //Простое число - должно делиться только на себя и на 1
	for index := 2; index <= n; index++ {
		if num1%index == 0 && num1 != index {
			return false, num1
		}
		if num2%index == 0 && num2 != index {
			return false, num2
		}
	}
	return true, 0
}
