package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	kk := `{
"k":"v",
    "inbounds": [ 
        {
            "listen": "0.0.0.0",
            "port": 443,
            "protocol": "vless",
            "settings": {
                "clients": [
                    {
                        "id": "",
                        "flow": "xtls-rprx-vision" 
                    }
                ],
                "decryption": "none"
            },
            "streamSettings": {
                "network": "tcp",
                "security": "reality",
                "realitySettings": {
                    "show": false, 
                    "dest": "example.com:443", 
                    "xver": 0,
                    "serverNames": [
                        "example.com",
                        "www.example.com"
                    ],
                    "privateKey": "",
                    "minClientVer": "", 
                    "maxClientVer": "", 
                    "maxTimeDiff": 0, 
                    "shortIds": [
                        "", 
                        "0123456789abcdef" 
                    ]
                }
            }
        }
    ]
}`
	var data interface{}
	err := json.Unmarshal([]byte(kk), &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", err)

	marshalMapAny(data, "")

}

func marshalKV() {

}
func marshalMapAny(data interface{}, prefix string) {
	switch t := data.(type) {
	case map[string]interface{}:
		for key, value := range t {
			//fmt.Printf("%v\n", key)
			//fmt.Printf("%v\n", value)
			var newKey string
			if prefix != "" {
				newKey = prefix + "." + key + "."
			} else {
				newKey = key + "."
			}

			switch value.(type) {
			case string:
				fmt.Printf("%s:%s\n", newKey, value)
			case bool:
				fmt.Printf("%s:%b\n", newKey, value)
			case float32:
				fmt.Printf("%s:%f\n", newKey, value)
			case float64:
				fmt.Printf("%s:%f\n", newKey, value)
			case int:
				fmt.Printf("%s:%d\n", newKey, value)
			case int32:
				fmt.Printf("%s:%d\n", newKey, value)
			case int64:
				fmt.Printf("%s:%d\n", newKey, value)
			default:
				marshalMapAny(value, newKey)
			}
		}
	case []interface{}:
		for index, value := range t {

			var newKey string
			if prefix != "" {
				newKey = prefix + "[" + fmt.Sprintf("%d", index) + "]."
			} else {
				newKey = "[" + fmt.Sprintf("%d", index) + "]."
			}

			marshalMapAny(value, newKey)
		}
	default:
		panic(errors.New("unknown type"))
	}
}
