package main

import testing_http "github.com/vizdosenterprises/wowometer/internal"

func main() {
	go testing_http.StartHTTP()
	select {}

}
