module github.com/yourname/hello

go 1.17

require (
	github.com/oneleo/golang-cht/code/c01/c01p01mymath v0.0.0-20210824062953-ef86ef804f9e // indirect
	github.com/yourname/hello/mymath v0.0.0-00010101000000-000000000000 // indirect
	github.com/yourname/hello/mypack v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/yourname/hello/mymath => ./mymath

replace github.com/yourname/hello/mypack => ./mypack
