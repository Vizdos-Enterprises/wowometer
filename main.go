package main

import testing_http "github.com/Vizdos-Enterprises/wowometer/internal"

func main() {
	go testing_http.StartHTTP()
	select {}

}
