package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql struct {
		Host     string
		Password string
		User     string
		Data     string
		Port     int
	}
	Redis struct {
		Host     string
		Password string
		Port     int
	}
}

var ConfigAppData Config

func InitViper() {
	viper.SetConfigFile("../../cmd/config/dev.yaml")
	viper.ReadInConfig()
	viper.Unmarshal(&ConfigAppData)
	log.Println("viper is success", ConfigAppData)
}

//
//func Nacos() {
//	clientConfig := constant.ClientConfig{
//		NamespaceId:         "676630af-2c50-46c6-ac66-1e5f2587f970", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		LogDir:              "/tmp/nacos/log",
//		CacheDir:            "/tmp/nacos/cache",
//		LogLevel:            "debug",
//		Username:            "nacos",
//		Password:            "Lmy031129",
//	}
//	serverConfigs := []constant.ServerConfig{
//		{
//			IpAddr: "14.103.198.194",
//			Port:   8848,
//		},
//	}
//	configClient, err := clients.NewConfigClient(
//		vo.NacosClientParam{
//			ClientConfig:  &clientConfig,
//			ServerConfigs: serverConfigs,
//		},
//	)
//	if err != nil {
//		panic("Nacos连接失败" + err.Error())
//	}
//	content, err := configClient.GetConfig(vo.ConfigParam{
//		DataId: "user-srv",
//		Group:  "DEFAULT_GROUP"})
//	if err != nil {
//		panic("Nacos序列化失败" + err.Error())
//	}
//	fmt.Println(content)
//	fmt.Println("nacos", config.ConfigAppData)
//	err = json.Unmarshal([]byte(content), &config.ConfigAppData)
//	if err != nil {
//		return
//	}
//	err = configClient.ListenConfig(vo.ConfigParam{
//		DataId: "user-srv",
//		Group:  "DEFAULT_GROUP",
//		OnChange: func(namespace, group, dataId, data string) {
//			log.Println("进入监测")
//			_ = json.Unmarshal([]byte(data), &config.ConfigAppData)
//		},
//	})
//	log.Println("Nacos连接成功", config.ConfigAppData)
//}
