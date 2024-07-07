module petrolistan.com/main

go 1.22.2

require (
	github.com/gorilla/mux v1.8.1
	petrolistan.com/opet v0.0.0-00010101000000-000000000000
	petrolistan.com/utils v1.2.3
)

replace petrolistan.com/opet => ../opet

replace petrolistan.com/utils => ../utils
