package main

func createHugeString(int) string {
	return ""
}

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  if len(v) > 100 {
  	justString = v[:100]
  }
}

func main() {
  someFunc()
}
