module example.com

go 1.25.3

replace example.com/counter => ./counter

require (
	example.com/counter v0.0.0-00010101000000-000000000000
	example.com/file v0.0.0-00010101000000-000000000000
)

replace example.com/file => ./file
