package file

import (
	"encoding/base64"
	"io"
	"os"
	"strconv"
)

const fileName string = "keys.pubk"

// CreateKeyFile Создаёт файл с открытым ключом
func CreateKeyFile(num1 int, num2 int) error {
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

// ReadFile чтение файла, который надо зашифровать
func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat() //Получить размер файла
	if err != nil {
		return nil, err
	}

	content := make([]byte, stat.Size())
	for {
		_, err := file.Read(content)
		if err == io.EOF {
			break
		}
	}
	return content, nil
}
