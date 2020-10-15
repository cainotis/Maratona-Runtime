package compiler

import (
	"fmt"
	"os/exec"
	"testing"
)

func setup(lang string, file string) {
	exec.Command("cp", "tests/"+file, "./").Output()
}

func teardown(file string) {
	exec.Command("rm", "programa.out", file).Output()
}

func TestCompilation(t *testing.T) {
	tests := []struct {
		lang string
		file string
	}{
		{"C", "programa.c"},
		{"Python", "programa.py"},
		{"C++11", "programa.cpp"},
		{"Go", "programa.go"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s/%s", test.lang, test.file), func(t *testing.T) {
			setup(test.lang, test.file)
			if Compile(test.lang) != 0 {
				t.Errorf("Compilation failed!")
			}
			teardown(test.file)
		})
	}
}
