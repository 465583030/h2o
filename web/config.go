package web

import (
	"fmt"

	"github.com/spf13/viper"
)

// Name get server.name
func Name() string {
	return viper.GetString("server.name")
}

func home(pre string) string {
	name := pre + "." + Name()
	if viper.GetBool("server.ssl") {
		return "https://" + name
	}
	return "http://" + name
}

// Backend backend home
func Backend() string {
	if IsProduction() {
		return home("www")
	}
	return viper.GetString("server.backend")
}

// Frontend frontend home
func Frontend() string {
	if IsProduction() {
		return home("my")
	}
	return viper.GetString("server.frontend")
}

// IsProduction production mode ?
func IsProduction() bool {
	return viper.GetString("env") == "production"
}

// DataSource datasource url
func DataSource() string {
	//"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s"
	args := ""
	for k, v := range viper.GetStringMapString("database.args") {
		args += fmt.Sprintf(" %s=%s ", k, v)
	}
	return args

	//"postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
	// return fmt.Sprintf(
	// 	"%s://%s:%s@%s:%d/%s?sslmode=%s",
	// 	viper.GetString("database.driver"),
	// 	viper.GetString("database.args.user"),
	// 	viper.GetString("database.args.password"),
	// 	viper.GetString("database.args.host"),
	// 	viper.GetInt("database.args.port"),
	// 	viper.GetString("database.args.dbname"),
	// 	viper.GetString("database.args.sslmode"),
	// )
}
