package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"rsass/file"
	"strconv"
	"strings"
)

func crypTo(path string) error { // Шифрует файл
	// Чтение указанного файла в -crypto
	data, err := file.ReadFile(path)
	if err != nil {
		return err
	}
	result, err := file.ReadFile("keys.pubk") // Чтение открытого ключа
	if err != nil {
		return err
	}
	utf8, err := base64.StdEncoding.DecodeString(string(result)) //Декодируем из base64 -> utf8
	if err != nil {
		return err
	}
	arr := strings.Split(string(utf8), ",") // Делим байты на массив
	e, err := strconv.Atoi(arr[0])          // ascii to int
	if err != nil {
		return nil
	}
	n, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil
	}

	fmt.Println("Wait...")
	size := len(data)
	cont := make([]int, size)
	var c int
	var b int
	for index := 0; index < size; index++ { // Находим остаток от деления
		c = 1
		b = int(data[index])
		for ad := 0; ad < e; ad++ {
			c = (c * b) % n
		}
		cont[index] = c << 3
		fmt.Printf("%d ", c)
	}
	err = file.CreateFile(path+".cry", cont) // Сохраняем получившиеся значения
	if err != nil {
		return err
	}
	fmt.Println("\nReady")
	return nil
}

func uncrypt(path string, d int, n int) error { // Дешифратор
	data, err := file.ReadFile(path)
	if err != nil {
		return err
	}
	utf8, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return err
	}
	data = []byte(utf8)
	size := bytes.Count(data, []byte("\n"))
	arr := bytes.Split(data, []byte("\n"))
	newData := make([]string, size)
	for i := 0; i < size; i++ {
		newData[i] += string(arr[i])
	}
	res := make([]byte, len(newData))
	var c int
	var b int
	for index := size - 1; index >= 0; index-- { // Находим первоначальные значения
		num, err := strconv.Atoi(newData[index])
		if err != nil {
			return err
		}
		c = 1
		b = num >> 3
		for ad := 0; ad < d; ad++ {
			c = (c * b) % n
		}
		res[index] = byte(c)
		fmt.Printf("%d", c)
	}
	path = strings.Replace(path, ".cry", "", 1)
	err = file.CreateFileUncry(path, res) // Записываем получившиеся данные
	if err != nil {
		return err
	}
	return nil
}
