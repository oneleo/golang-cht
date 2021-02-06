module github.com/yourname/hello

go 1.15

require (
	github.com/oneleo/golang-cht/code/c01/c01p01mymath v0.0.0-20210205112407-768dcaee0f79
	github.com/yourname/hello/mymath v0.0.0-00010101000000-000000000000
)

replace github.com/yourname/hello/mymath => ./mymath
