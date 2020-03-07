package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
)

var (
	App     AppConfig
	Etcd    EtcdConfig
	Db      DbConfig
	GServer ServerConfig
)

// Initialize GrantNZ server config
// The config is grant_n_z_server.yaml data structure
func InitGrantNZServerConfig(yamlPath string) {
	yml := readLocalYml(yamlPath)
	App = yml.GetAppConfig()
	GServer = yml.GetServerConfig()
	Db = yml.GetDbConfig()
	Etcd = yml.GetEtcdConfig()
}

// Initialize GrantNZ server config
// The config is grant_n_z.yaml data structure
func InitGrantNZCacheConfig(yamlPath string) {
	yml := readLocalYml(yamlPath)
	App = yml.GetAppConfig()
	Db = yml.GetDbConfig()
	Etcd = yml.GetEtcdConfig()
}

// Read yaml file
func readLocalYml(ymlName string) YmlConfig {
	var yml YmlConfig
	data, err := ioutil.ReadFile(ymlName)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &yml); err != nil {
		panic(err)
	}

	return yml
}
