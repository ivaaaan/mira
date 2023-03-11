package config

import _ "github.com/BurntSushi/toml"

type Config struct {
	Jira struct {
		URL         string `toml:"url"`
		Project_Key string `toml:"project_key"`
		Username    string `toml:"username"`
		// For some reason viper does not want to umarshal this if field name is APIToken
		Api_Token string `toml:"api_token"`
	} `toml:"jira"`
}
