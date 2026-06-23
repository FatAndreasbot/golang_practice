package t3_test

import (
	"io"
	"os"
	"strings"
	t3 "t_3"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	out, _ := io.ReadAll(r)
	return string(out)
}

func TestProcessError(t *testing.T) {
	err := t3.SimulateRequest()
	msg := captureStdout(func() {
		t3.ProcessError(err)
	})
	if err == nil {
		return
	}

	if strings.Contains(err.Error(), "запрос не выполнен\n") {
		if msg != "Требуется повторная попытка" {
			t.Error()
		}
	}
	if strings.Contains(err.Error(), "ошибка:") {
		if msg != "Требуется повторная попытка\n" {
			t.Error()
		}
	}
}
