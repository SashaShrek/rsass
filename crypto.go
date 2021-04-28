package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"rsass/file"
	"strconv"
	"strings"
)

// crypTo шифрует файл
func crypTo(path string, pathKeys string) error { // Шифрует файл
	type Percent struct {
		percent int64
		step    int64
		size    int64
		index   int64
	}

	type Module struct {
		c int
		b int
	}

	result, err := file.ReadFile(pathKeys) // Чтение открытого ключа
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

	fmt.Println("Шифрую данные и начинаю запись в буфер...")

	file, err := os.Create(path + ".cry")
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	open, err := os.Open(path)
	if err != nil {
		return err
	}
	defer open.Close()

	stat, _ := open.Stat()

	var per Percent
	per.size = stat.Size()
	per.percent = per.size / 100
	per.step = 1

	var mod Module

	reader := bufio.NewReader(open)

	for per.index = 0; per.index < per.size; per.index++ {
		num, _ := reader.ReadByte()
		mod.c = 1
		mod.b = int(num)
		for ad := 0; ad < e; ad++ {
			mod.c = (mod.c * mod.b) % n
		}
		writer.WriteString(strconv.Itoa(mod.c<<3) + "\n")
		if per.percent > 0 && per.percent*per.step == per.index {
			fmt.Printf("%d%c ", per.step, '%')
			per.step++
		}
	}
	fmt.Println("\nСброс буфера в файл")
	err = writer.Flush()
	if err != nil {
		return err
	}
	fmt.Println("Готово")
	return nil
}

// unCrypt - это дешифратор
func unCrypt(path string, d int, n int) error {
	type Percent struct {
		percent int
		step    int
		size    int
		index   int
	}

	type Module struct {
		c int
		b int
	}

	open, err := os.Open(path)
	if err != nil {
		return err
	}
	defer open.Close()

	save, err := os.Create(strings.Replace(path, ".cry", "", 1))
	if err != nil {
		return err
	}
	defer save.Close()

	fmt.Println("Подсчёт данных...")
	per := Percent{
		step:  1,
		index: 1,
	}
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		per.size++
	}
	file.Close()
	fmt.Println("Готово")

	fmt.Println("Дешифрую данные и начинаю запись в буфер...\nЭто может занять какое-то время")
	per.percent = per.size / 100

	var mod Module
	scanner = bufio.NewScanner(open)
	writer := bufio.NewWriter(save)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		mod.c = 1
		mod.b = num >> 3
		for ad := 0; ad < d; ad++ {
			mod.c = (mod.c * mod.b) % n
		}
		writer.WriteByte(byte(mod.c))
		if per.percent > 0 && per.percent*per.step == per.index {
			fmt.Printf("%d%c ", per.step, '%')
			per.step++
		}
		per.index++
	}
	fmt.Println("\nСброс буфера в файл")
	err = writer.Flush()
	if err != nil {
		return err
	}
	fmt.Println("Готово")
	return nil
}
