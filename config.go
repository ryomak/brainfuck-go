package bf

import "github.com/BurntSushi/toml"

//tomlファイルをマッピングする構造体
type ConfigData struct {
	//bfのコマンド一覧
	Commands commands
}

//Load is Loading ConfigData
func LoadConfig(filePath string) (*ConfigData, error) {
	config := new(ConfigData)
	_, err := toml.DecodeFile(filePath, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// GetCommands is getting bf commands
func (c *ConfigData) GetCommands() commands {
	return c.Commands
}

// DefaultConfig is getting default config
func DefaultConfig() *ConfigData {
	return &ConfigData{
		commands{
			NEXT:  ">",
			PREV:  "<",
			INC:   "+",
			DEC:   "-",
			READ:  ",",
			WRITE: ".",
			OPEN:  "[",
			CLOSE: "]",
		},
	}
}

// GokuConfig
func GokuConfig() *ConfigData {
	return &ConfigData{
		commands{
			NEXT:  "おめぇ強そうだな！",
			PREV:  "棒よのびろ――っ",
			INC:   "オッス！",
			DEC:   "界王拳!!",
			READ:  "オラのすべてをこの拳にかける",
			WRITE: "バイバイ、みんな",
			OPEN:  "いやだ！クリリンは友だちだ！亀仙人のじいちゃんには世話になったんだ！",
			CLOSE: "クリリンのことかーーーーーーーーっ",
		},
	}
}
