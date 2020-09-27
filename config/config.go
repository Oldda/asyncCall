package config

import(
	"github.com/spf13/viper"
	"log"
)

func NewConfig(cnfDirPath,cnfName,cnfType string)*viper.Viper{
	viper := viper.New()
	viper.SetConfigName(cnfName)
	viper.SetConfigType(cnfType)
	viper.AddConfigPath(cnfDirPath)
	if err := viper.ReadInConfig();err != nil {
		log.Println(err)
	}
	return viper
}