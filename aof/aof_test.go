package aof

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	file, err := os.Open("/home/dev/RunEnv/Redis3.0/redis/6379/appendonly.aof")
	defer file.Close()

	if err != nil {
		t.Errorf("Open file failed. Error:%s", err.Error())
		return
	}
	reader := NewBufioReader(file)

	for {
		op, err := reader.ReadOperation()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Errorf("ReadOperation failed. %s", err.Error())
			break
		}
		fmt.Println(op)
	}
}
