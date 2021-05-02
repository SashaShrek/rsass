package file

import (
	"encoding/base64"
	"io"
	"os"
	"strconv"
	"strings"
)

// CreateKeyFile Создаёт файл с открытым ключом
func CreateKeyFile(num1 int, num2 int, fileName string) error {
	str1 := strconv.Itoa(num1) // int to ascii
	str2 := strconv.Itoa(num2)

	result := str1 + "," + str2
	result = base64.StdEncoding.EncodeToString([]byte(result)) // utf8 -> base64
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(result)
	return nil
}

// ReadFile чтение файла с ключом
func ReadFile(path string) (int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	stat, err := file.Stat() //Получить размер файла
	if err != nil {
		return 0, 0, err
	}

	content := make([]byte, stat.Size())
	for {
		_, err := file.Read(content)
		if err == io.EOF {
			break
		}
	}
	utf8, err := base64.StdEncoding.DecodeString(string(content)) //Декодируем из base64 -> utf8
	if err != nil {
		return 0, 0, err
	}
	arr := strings.Split(string(utf8), ",") // Делим байты на массив
	ed, err := strconv.Atoi(arr[0])         // ascii to int
	if err != nil {
		return 0, 0, err
	}
	n, err := strconv.Atoi(arr[1])
	if err != nil {
		return 0, 0, err
	}
	return ed, n, nil
}
