package marshalmapany

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	kk := `{
	   "a": 1,
	   "b": {
	       "c": 2,
	       "d": {
	           "e": 3,
	           "f": [4, 5, 6]
	       }},
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
	   ],
	"matrix":[[1,2,3],[4,5,6]]
	}`
	var data interface{}
	err := json.Unmarshal([]byte(kk), &data)
	if err != nil {
		panic(err)
	}
	MarshalMapAny(data, "")
}

func MarshalMapAny(data interface{}, prefix string) {
	switch t := data.(type) {
	case map[string]interface{}:
		for key, value := range t {
			var newKey string
			if prefix != "" {
				newKey = prefix + "." + key
			} else {
				newKey = key
			}

			MarshalMapAny(value, newKey)

		}
	case []interface{}:
		for index, value := range t {
			var newKey string
			if prefix != "" {
				newKey = prefix + ".[" + fmt.Sprintf("%d", index) + "]"
			} else {
				newKey = "[" + fmt.Sprintf("%d", index) + "]"
			}
			MarshalMapAny(value, newKey)
		}
	case string, bool, float64:
		fmt.Printf("%s=%v\n", prefix, t)
	default:
		panic(errors.New("unknown type"))
	}
}
