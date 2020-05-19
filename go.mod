module github.com/lokalise/lokalise-cli-2-go

require (
	github.com/lokalise/go-lokalise-api/v2 v2.0.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0
	github.com/urfave/cli v1.22.1
)

//replace github.com/lokalise/go-lokalise-api => ../go-lokalise-api

go 1.13
