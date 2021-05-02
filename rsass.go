package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	addKey := flag.String("ak", "0", "23,73 - два простых числа (для создания ключей) без пробела через запятую")
	crypto := flag.String("crypto", "0", "Файл, который нужно зашифровать")
	uncry := flag.String("uncry", "0", "Файл, который нужно расшифровать")

	flag.Parse()
	fmt.Printf("%s %s %s\n", *addKey, *crypto, *uncry)

	if *addKey != "0" && *crypto == "0" {
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
		err = createKeys(num1, num2)
		if err != nil {
			fmt.Println(err)
		}
	} else if *addKey == "0" && *crypto != "0" {
		err := crypTo(*crypto, "keys.pubk")
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if *addKey == "0" && *crypto == "0" && *uncry != "0" {
		err := unCrypt(*uncry, "keys.privk")
		if err != nil {
			fmt.Println(err)
		}
	}
}
