package inits

import (
	"api-gatware/cmd/config"
	"encoding/json"
	"log"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func Nacos() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "676630af-2c50-46c6-ac66-1e5f2587f970", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            "nacos",
		Password:            "Lmy031129",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "14.103.198.194",
			Port:   8848,
		},
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		panic("nacos 连接失败" + err.Error())
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-srv",
		Group:  "DEFAULT_GROUP"})

	if err != nil {
		panic("nacos 序列化失败" + err.Error())
	}

	json.Unmarshal([]byte(content), &config.ConfigAppData)
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "user-srv",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			log.Println("进入监测")
			json.Unmarshal([]byte(content), &config.ConfigAppData)
		},
	})
	log.Println("nacos  连接成功", config.ConfigAppData)

}
