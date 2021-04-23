package main

import (
  "fmt"
  "rsass/file"
  "errors"
)

func create_keys(num1 int, num2 int) error{
  if num1 == 1 || num2 == 1 {
    fmt.Println("Числа должны быть > 1!")
    return errors.New("Bye")
  }
  var N int
  if num1 >= num2 {
    N = num1
  }else if num2 > num1 {
    N = num2
  }
  res, resNum := check_nums(N, num1, num2)
  if res == false {
    fmt.Printf("Число %d не является простым!\n", resNum)
    return errors.New("Bye")
  }
  n := num1 * num2
  fi := (num1 -1) * (num2 - 1)
  e := get_e(fi)
  d := get_d(e, fi)
  
  if n <= 255 {
    fmt.Println("Выберите другие простые числа. Желательно начинать от 53")
    return errors.New("Bye")
  }
  fmt.Println("Открытый ключ: ", e, n)
  fmt.Println("Закрытый ключ: ", d, n)
  fmt.Println("Обязательно сохраните эти два числа (закрытый ключ)! В противном случае вы НЕ сможете расшифровать данные!")
  err := file.CreateKeyFile(e, n)
  if err != nil {
    return err
  }
  return nil
}

func get_e(fi int) int{
  var index int
  count := fi
  for index = 1; index < count; index++ {
    if (index % 1 == 0) && (fi % 1 == 0) && (fi % index != 0) {
      break
    }
  }
  return index
}

func get_d(e int, fi int) int{
  var d int
  var index int

  for index = 2; index < fi; index++ {
    d = (1 + fi * index) / e
    if d != e && ((e * d) % fi == 1) {
      break;
    }
  }
  return d
}

func check_nums(n int, num1 int, num2 int) (bool, int){//Простое число - должно делиться только на себя и на 1
  for index := 2; index <= n; index++ {
    if num1 % index == 0 && num1 != index {
      return false, num1
    }
    if num2 % index == 0 && num2 != index {
      return false, num2
    }
  }
  return true, 0
}
