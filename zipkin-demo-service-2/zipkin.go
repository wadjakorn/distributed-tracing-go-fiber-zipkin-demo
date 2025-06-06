package main

func Zipkin() {
	port := uint16(5000)
	endpoint := "http://localhost:9411/api/v2/spans"
	rate := 1.0
	name := AppName
	err := NewTracer(name, port, endpoint, rate)
	if err != nil {
		panic(err)
	}

}
