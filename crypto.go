package main

import (
  "rsass/file"
  "strings"
  "strconv"
  "bytes"
  "encoding/base64"
)

func cryp_to(path string) error{
  data, err := file.ReadFile(path)
  if err != nil {
    return err
  }
  result, err := file.ReadFile("keys.pubk")
  if err != nil {
    return err
  }
  utf8, err := base64.StdEncoding.DecodeString(string(result)) //Декодируем
  if err != nil {
    return err
  }
  arr := strings.Split(string(utf8), ",")
  e, err := strconv.Atoi(arr[0])
  if err != nil {
    return nil
  }
  n, err := strconv.Atoi(arr[1])
  if err != nil {
    return nil
  }

  size := len(data)
  cont := make([]int, size)
  newData := make([]int, size)
  for index := 0; index < size; index++ {
    c := 1
    b := int(data[index])
    if index > 0 {
      newData[index] = b + newData[index - 1]
    }else{
      newData[index] = b
    }
    for ad := 0; ad < e; ad++ {
      c = (c * newData[index]) % n
    }
    cont[index] = c
  }
  err = file.CreateFile(path + ".cry", cont)
  if err != nil {
    return err
  }
  return nil
}

func uncrypt(path string, d int, n int) error{
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
  for index := size - 1; index >= 0; index-- {
    num, err := strconv.Atoi(newData[index])
    if err != nil {
      return err
    }
    c := 1
    b := num
    for ad := 0; ad < d; ad++ {
      c = (c * b) % n
    }
    res[index] = byte(c)
  }
  for index := size - 1; index >= 0; index-- {
    if index != 0 {
      res[index] = byte(int(res[index]) - int(res[index - 1]))
    }else{
      res[index] = res[index]
    }
  }
  path = strings.Replace(path, ".cry", "", 1)
  err = file.CreateFileUncry(path, res)
  if err != nil {
    return err
  }
  return nil
}
