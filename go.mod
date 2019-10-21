module github.com/lokalise/lokalise-cli-2-go

require (
	github.com/go-resty/resty/v2 v2.1.0 // indirect
	github.com/lokalise/go-lokalise-api v0.0.0-20191021085541-b39b015ca091
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0
	github.com/urfave/cli v1.22.1
	golang.org/x/net v0.0.0-20191014212845-da9a3fd4c582 // indirect
)

//replace github.com/lokalise/go-lokalise-api => ../go-lokalise-api

go 1.13
