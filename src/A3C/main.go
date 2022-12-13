package main

import (
	"os"
	"strconv"
)

func main() {
	s := makeString()
	saveToFile(s, "program.c")
	s = makeString2()
	saveToFile(s, "program2.c")
	s = makeString3()
	saveToFile(s, "program3.c")
}
func makeString() string {
	var s string
	s = "#include stdio.h\n"
	s = s + "\nvoid p(int x) {\n  switch (x) {\n"
	for i := 0; i < 1000; i++ {
		s = s + "    case" + strconv.Itoa(i) + ": q" + strconv.Itoa(i) + "(); break;\n"
	}
	s = s + "  }\n}"
	return s
}
func saveToFile(s string, name string) {
	os.Create(name)
	os.WriteFile(name, []byte(s), 0644)
}
func makeString2() string {
	var s string
	s = "#include stdio.h\n"
	s = s + "\nvoid p(int x) {\n  switch (x) {\n"
	for i := 0; i < 100000; i = i + 100 {
		s = s + "    case" + strconv.Itoa(i) + ": q" + strconv.Itoa(i) + "(); break;\n"
	}
	s = s + "  }\n}"
	return s
}
func makeString3() string {
	var s string
	s = "#include stdio.h\n"
	s = s + "\nvoid p(int x) {\n  switch (x) {\n"
	for i := 0; i < 10000000; i = i + 10000 {
		s = s + "    case" + strconv.Itoa(i) + ": q" + strconv.Itoa(i) + "(); break;\n"
	}
	s = s + "  }\n}"
	return s
}
