package main

import (
  "fmt"
  "flag"
  "strings"
  "strconv"
)

func main(){
  addKey := flag.String("ak", "0", "23,73 - два простых числа (для создания ключей) без пробела через запятую")
  crypto := flag.String("crypto", "0", "Файл, который нужно зашифровать/расшифровать")
  key    := flag.String("key", "0", "23,73 - два числа (секретный ключ) без пробела через запятую")

  flag.Parse()
  fmt.Printf("%s %s %s\n", *addKey, *crypto, *key)

  if *addKey != "0" && *crypto == "0" && *key == "0" {
    arr := strings.Split(*addKey, ",")
    num1, err := strconv.Atoi(arr[0])
    if err != nil {
      fmt.Println(err)
      return
    }
    num2, err := strconv.Atoi(arr[1])
    if err != nil {
      fmt.Println(err)
      return
    }
    err = create_keys(num1, num2)
    if err != nil {
      fmt.Println(err)
    }
  }else if *addKey == "0" && *crypto != "0" && *key == "0" {
    err := cryp_to(*crypto)
    if err != nil {
      fmt.Println(err)
      return
    }
  }else if *addKey == "0" && *crypto != "0" && *key != "0" {
    arr := strings.Split(*key, ",")
    num1, err := strconv.Atoi(arr[0])
    if err != nil {
      fmt.Println(err)
      return
    }
    num2, err := strconv.Atoi(arr[1])
    if err != nil {
      fmt.Println(err)
      return
    }
    err = uncrypt(*crypto, num1, num2)
    if err != nil {
      fmt.Println(err)
    }
  }
}
