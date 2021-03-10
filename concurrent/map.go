package main


// 并发读写 map 会造成问题，编译器会检测出来
// fatal error: concurrent map read and map write
func main() {
	m := make(map[string]string)

	go func() {
			for {
				m["a"] = "c"
			}
	}()

	go func() {
		for {
			_ = m["a"]
		}
	}()

	for {}
}