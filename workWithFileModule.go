package workWithFileModule

import (
	"bufio"
	"bytes"
	"os"
)

func GetCountLines(fileName string) (int, error) {
	count := 0

	file, err := os.Open(fileName)
	if err != nil {
		return count, err
	}
	defer file.Close()

	buf := make([]byte, 32*1024)
	lineSep := []byte{'\n'}

	c, err := file.Read(buf)
	if err != nil {
		return count, err
	}

	count += bytes.Count(buf[:c], lineSep)
	return count, nil
}

func GetCountWords(fileName string) (int, error) {
	count := 0

	file, err := os.Open(fileName)
	if err != nil {
		return count, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}

	return count, nil
}
