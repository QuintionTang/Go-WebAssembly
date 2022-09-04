package main

import "syscall/js"

func main() {
        window := js.Global()
        // window.document.getElementById("clickresult") 
        message := window.Get("document").Call("getElementById", "clickresult")

		cb := js.FuncOf(func(this js.Value, args []js.Value) any { 
				message.Set("innerHTML", "Go 事件触发")
				return nil
		})
        // message.addEventListener("click", cb)
        message.Call("addEventListener", "click", cb)

        select {}
		
}