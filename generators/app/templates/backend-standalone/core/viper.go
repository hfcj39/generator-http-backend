package core

import (
	"<%= displayName %>/global"
	"<%= displayName %>/utils"
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用<%= displayName %>_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			global.LOG.Error(err.Error())
		}
	})

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		global.LOG.Error(err.Error())
	}
	return v
}
