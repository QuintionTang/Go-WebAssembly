package main
import "syscall/js"

func main() {
	// 获取全局对象（网页浏览器为window）
	window := js.Global()
	// window.document.getElementById("helloresult")
	helloresult := window.Get("document").Call("getElementById", "helloresult")
	// HTML
	helloresult.Set("innerHTML", "Hello, WebAssembly")
}
