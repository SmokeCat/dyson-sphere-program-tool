module github.com/smokecat/dyson-sphere-program-tool

go 1.15

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
)

replace (
github.com/smokecat/dyson-sphere-program-tool/cmd => ./cmd
github.com/smokecat/dyson-sphere-program-tool/internal/gamedata => ./internal/gamedata
)
