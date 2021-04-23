package file

import (
  "os"
  "strconv"
  "io"
  "encoding/base64"
)

const fileName string = "keys.pubk"

func CreateKeyFile(num1 int, num2 int) error{
  str1 := strconv.Itoa(num1)
  str2 := strconv.Itoa(num2)

  result := str1 + "," + str2
  result = base64.StdEncoding.EncodeToString([]byte(result))
  file, err := os.Create(fileName)
  if err != nil {
    return err
  }
  defer file.Close()

  file.WriteString(result)
  return nil
}

func ReadFile(path string) ([]byte, error){
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
  for{
    _, err := file.Read(content)
    if err == io.EOF {
      break
    }
  }
  return content, nil
}

func CreateFile(path string, data []int) error{
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()
  arr := make([]string, len(data) * 2)
  var str string
  for i := 0; i < len(data); i++ {
    arr[i] = strconv.Itoa(data[i]) + "\n"
    str += arr[i]
  }
  file.WriteString(base64.StdEncoding.EncodeToString([]byte(str)))
  return nil
}

func CreateFileUncry(path string, data []byte) error{
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()
  file.Write(data)
  return nil
}
