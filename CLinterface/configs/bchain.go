package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func ExtractBChainConfig(path string, debug bool) (BChainConfig, error) {
	var config BChainConfig
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error :(")
		return config, err
	}

	if debug {
		var conf map[string]interface{}
		json.Unmarshal(b, &conf)
		if err := bchainCheckUnmarshalMap(conf); err != nil {
			return config, err
		}
	}
	json.Unmarshal(b, &config)
	return config, nil
}

type BChainConfig struct {
	VERSION       string
	IMPORTS       []string
	PORT          string
	SOURCEIPS     []string
	NODE          NodeConfig
	BLOCK         BlockConfig
	MAXUTXO       int
	PROPOSEWINDOW int
}

func (bcc BChainConfig) Port() string {
	return bcc.PORT
}

type NodeConfig struct {
	MAXCONNECTED int
	MINCONNECTED int
	RETRYTIMEOUT int
}
type BlockConfig struct {
	MAXPERMINUTE float64
	MAXBYTES     int
}

func bchainCheckUnmarshalMap(m map[string]interface{}) error {
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
	if m["SOURCEIPS"] != nil {
		imp := m["SOURCEIPS"].([]interface{})
		fmt.Println("found:", len(imp), "source ips")
		if len(imp) > 0 {
			var s strings.Builder
			for i := 0; i < len(imp); i++ {
				subs := imp[i].(string)
				s.WriteString(subs)
				s.WriteString("; ")
			}
			fmt.Println("ips:", s.String())
		}
	} else {
		return errors.New("source ips not found")
	}

	if m["NODE"] != nil {
		fmt.Println("node config found")
		c := m["NODE"].(map[string]interface{})

		if c["MAXCONNECTED"] != nil {
			fmt.Println("max nodes connected", c["MAXCONNECTED"])
		} else {
			return errors.New("max nodes connected not found")
		}

		if c["MINCONNECTED"] != nil {
			fmt.Println("min nodes connected", c["MINCONNECTED"])
		} else {
			return errors.New("min nodes connected not found")
		}

		if c["RETRYTIMEOUT"] != nil {
			fmt.Println("nodes retry timeout", c["RETRYTIMEOUT"])
		} else {
			return errors.New("nodes retry timeout not found")
		}

	} else {
		return errors.New("node config not found")
	}

	if m["BLOCK"] != nil {
		fmt.Println("block config found")
		c := m["BLOCK"].(map[string]interface{})

		if c["MAXPERMINUTE"] != nil {
			fmt.Println("max blocks per minute", c["MAXPERMINUTE"])
		} else {
			return errors.New("max blocks per minute not found")
		}

		if c["MAXBYTES"] != nil {
			fmt.Println("max blocks bytes", c["MAXBYTES"])
		} else {
			return errors.New("max blocks bytes not found")
		}

	} else {
		return errors.New("block config not found")
	}
	if m["MAXUTXO"] != nil {
		fmt.Println("max UTXO", m["MAXUTXO"])
	} else {
		return errors.New("max UTXO not found")
	}

	if m["PROPOSEWINDOW"] != nil {
		fmt.Println("propose window", m["PROPOSEWINDOW"])
	} else {
		return errors.New("propose window not found")
	}

	return nil
}
