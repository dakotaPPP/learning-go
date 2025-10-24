module example.com

go 1.25.3

replace example.com/counter => ./counter

replace example.com/file => ./file

replace example.com/wc => ./wc

require example.com/wc v0.0.0-00010101000000-000000000000

require (
	example.com/counter v0.0.0-00010101000000-000000000000 // indirect
	example.com/file v0.0.0-00010101000000-000000000000 // indirect
)
