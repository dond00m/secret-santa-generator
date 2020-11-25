module github.com/dond00m/secret-santa-generator

go 1.12

replace local/notify => ./modules/notify

replace local/randomize => ./modules/randomize

require (
	local/notify v0.0.0-00010101000000-000000000000
	local/randomize v0.0.0-00010101000000-000000000000
)
