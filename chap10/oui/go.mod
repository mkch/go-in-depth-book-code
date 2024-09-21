module oui

go 1.22.4

replace ouidb => .\ouidb

replace ouisql => .\ouidb\ouisql

require (
	ouidb v0.0.0-00010101000000-000000000000
	ouisql v0.0.0-00010101000000-000000000000
)

require github.com/mattn/go-sqlite3 v1.14.22 // indirect
