module example.com/wc

go 1.25.3

replace example.com/counter => ../counter

replace example.com/file => ../file

require (
	example.com/counter v0.0.0-00010101000000-000000000000
	example.com/file v0.0.0-00010101000000-000000000000
)
