package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	kk := `{
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
	marshalMapAny(kk)

}

func marshalKV() {

}
func marshalMapAny(data string) {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &result)
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", result)
	for i, v := range result {
		fmt.Printf("%v\n", i)
		fmt.Printf("%v\n", v)
	}
	inbounds, ok := result["inbounds"]
	fmt.Printf("%v\n", inbounds)
	fmt.Printf("%v\n", ok)
	switch inbounds.(type) {
	case []interface{}:
		a := inbounds.([]interface{})
		fmt.Println(a[0]) //
	case map[string]interface{}:
		m := inbounds.(map[string]interface{})
		fmt.Println(m["name"]) // John Doe
		fmt.Println(m["age"])  // 30
	}

}
