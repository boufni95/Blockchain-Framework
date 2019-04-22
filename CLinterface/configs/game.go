package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

//ExtractGameConfig : extracts the confogiguration from a json slice of bytes
func ExtractGameConfig(b []byte) (GameServerConfig, error) {
	var config GameServerConfig
	var conf map[string]interface{}
	json.Unmarshal(b, &conf)
	//spew.Dump(conf)
	if err := gameCheckUnmarshalMap(conf); err != nil {
		fmt.Println("error found!")
		return config, err
	}

	return config, nil

}

//GameObjectConfig : config of a game object in the game
type GameObjectConfig struct {
	NAME   string
	CODE   byte
	LIFES  byte
	DAMAGE byte
}

//GameServerConfig : config of the game server
type GameServerConfig struct {
	VERSION string
	IMPORTS []string
	PORT    string
	GAME    GameConfig
	USER    GameUserConfig
}

//GameConfig : config of the game on the server
type GameConfig struct {
	NAME     string
	MAXROOMS int
	CLANS    bool
	GAMEMODS []GameModeConfig
}

//GameModeConfig : config of a gamemode
type GameModeConfig struct {
	NAME       string
	CODE       byte
	MAXPLAYERS int
	LIFES      int
	OBJECTS    []GameObjectConfig
	RULES      []RuleConfig
}

//GameUserConfig : Configuration ogf a user on this instance
type GameUserConfig struct {
	CLAN  bool
	TYPES []GUserTypeConfig
}

//GUserTypeConfig types of user on this instance
type GUserTypeConfig struct {
	NAME    string
	CODE    byte
	LIFES   byte
	OBJECTS []GameObjectConfig
	RULES   []RuleConfig
}

func gameCheckUnmarshalMap(m map[string]interface{}) error {
	if m["VERSION"] != nil {
		fmt.Println("Version: ", m["VERSION"])
	} else {
		return errors.New("version not found")
	}

	if m["IMPORTS"] != nil {
		imp := m["IMPORTS"].([]interface{})
		fmt.Println("found:", len(imp), "imports")
		if len(imp) > 0 {
			var s strings.Builder
			for i := 0; i < len(imp); i++ {
				subs := imp[i].(string)
				s.WriteString(subs)
				s.WriteString("; ")
			}
			fmt.Println("importing:", s.String())
		}
	} else {
		return errors.New("imports not found")
	}
	if m["PORT"] != nil {
		fmt.Println("Port: ", m["PORT"])
	} else {
		return errors.New("port not found")
	}
	if m["GAME"] != nil {
		fmt.Println("-----Game configuration found")
	} else {
		return errors.New("Game configuration not found")
	}

	GAME := m["GAME"].(map[string]interface{})

	if GAME["NAME"] != nil {
		fmt.Println("Game name:", GAME["NAME"])
	} else {
		return errors.New("Game name not found")
	}

	if GAME["MAXROOMS"] != nil {
		fmt.Println("Max rooms:", GAME["MAXROOMS"])
	} else {
		return errors.New("Max rooms not found")
	}

	if GAME["CLANS"] != nil {
		fmt.Println("Clans active:", GAME["CLANS"])
	} else {
		return errors.New("Clans flag not found")
	}

	if GAME["GAMEMODS"] == nil {
		return errors.New("gamemods not found")
	}

	GMODS := GAME["GAMEMODS"].([]interface{})
	fmt.Println("-----found", len(GMODS), "gamemods")
	for i := 0; i < len(GMODS); i++ {
		fmt.Println("-----Gamemode#", i)
		GMODE := GMODS[i].(map[string]interface{})
		if GMODE["NAME"] != nil {
			fmt.Println("Gamemode name #", i, ":", GMODE["NAME"])
		} else {
			return errors.New("Gamemode name not fund")
		}
		if GMODE["CODE"] != nil {
			fmt.Println("Gamemode #", i, "code:", GMODE["CODE"])
		} else {
			return errors.New("Gamemode code not found")
		}
		if GMODE["MAXPLAYERS"] != nil {
			fmt.Println("Gamemode #", i, "max players:", GMODE["MAXPLAYERS"])
		} else {
			return errors.New("Gamemode max players not found")
		}
		if GMODE["LIFES"] != nil {
			fmt.Println("Gamemode #", i, "lifes:", GMODE["LIFES"])
		} else {
			return errors.New("Gamemode lifes not found")
		}
		if GMODE["OBJECTS"] != nil {
			fmt.Println("-----Gamemode #", i, "objects found")
		} else {
			return errors.New("Gamemode objects not found")
		}
		OBJS := GMODE["OBJECTS"].([]interface{})
		fmt.Println("-----found", len(OBJS), "objects in gamemode#", i)
		for j := 0; j < len(OBJS); j++ {
			err := checkObjectConfig(i, j, OBJS[j])
			if err != nil {
				return err
			}
		}

	}
	if m["USER"] != nil {
		fmt.Println("-----User config found")
	} else {
		return errors.New("User config not found")
	}
	USER := m["USER"].(map[string]interface{})

	if USER["CLAN"] != nil {
		fmt.Println("user clan:", USER["CLAN"])
	} else {
		return errors.New("user clan not found")
	}

	if USER["TYPES"] == nil {
		return errors.New("user types not found")
	}

	UTYPES := USER["TYPES"].([]interface{})
	fmt.Println("-----found", len(UTYPES), "user types")
	for i := 0; i < len(UTYPES); i++ {
		fmt.Println("User type #", i)
		UTYPE := UTYPES[i].(map[string]interface{})
		if UTYPE["NAME"] != nil {
			fmt.Println("user type name #", i, ":", UTYPE["NAME"])
		} else {
			return errors.New("user type name not fund")
		}
		if UTYPE["CODE"] != nil {
			fmt.Println("user type code #", i, ":", UTYPE["CODE"])
		} else {
			return errors.New("user type code not fund")
		}
		if UTYPE["LIFES"] != nil {
			fmt.Println("user type lifes #", i, ":", UTYPE["LIFES"])
		} else {
			return errors.New("user type lifes not fund")
		}

		if UTYPE["OBJECTS"] != nil {
			fmt.Println("-----user type #", i, "objects found")
		} else {
			return errors.New("user type objects not found")
		}
		OBJS := UTYPE["OBJECTS"].([]interface{})
		fmt.Println("-----found", len(OBJS), "objects in user type #", i)
		for j := 0; j < len(OBJS); j++ {
			err := checkObjectConfig(i, j, OBJS[j])
			if err != nil {
				return err
			}
		}

	}

	return nil
}
func checkObjectConfig(i int, j int, o interface{}) error {
	OBJ := o.(map[string]interface{})
	if OBJ["NAME"] != nil {
		fmt.Println("Gamemode #", i, ",object #", j, "name:", OBJ["NAME"])
	} else {
		return errors.New("Gamemode object name not found")
	}
	if OBJ["CODE"] != nil {
		fmt.Println("Gamemode #", i, ",object #", j, "code:", OBJ["CODE"])
	} else {
		return errors.New("Gamemode object code not found")
	}
	if OBJ["LIFES"] != nil {
		fmt.Println("Gamemode #", i, ",object #", j, "lifes:", OBJ["LIFES"])
	} else {
		return errors.New("Gamemode object code not found")
	}
	if OBJ["DAMAGE"] != nil {
		fmt.Println("Gamemode #", i, ",object #", j, "damage:", OBJ["DAMAGE"])
	} else {
		return errors.New("Gamemode object damage not found")
	}
	return nil
}
