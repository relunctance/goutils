
## 使用示例:

* 用于删除json中字段

```go
package main

import (
	"fmt"

	"github.com/relunctance/goutils/jsondel"
	sj "github.com/guyannanfei25/go-simplejson"
)

func main() {

	const json = `{
				 "ipinfo": {
				   	"1.0.0.1001":  {
							 "name":{
								"china_admin_code": "330001",
								"city": "",
								"city_name": "",
								"continent_code": "AP",
								"country": "",
								"country_code": "CN",
								"country_code2": "",
								"country_name": "中国",
								"idd_code": "86",
								"isp": "",
								"isp_domain": "电信",
								"latitude": "30.287459",
								"longitude": "120.153576",
								"owner_domain": "",
								"proxy_type": "",
								"region": "",
								"region_name": "浙江",
								"timezone": "Asia/Shanghai",
								"ip": "36.17.88.60",
								"utc_offset": "UTC+8"
							}
					},
				   	"val":  {
							 "name":{
								"china_admin_code": "330001",
								"city": "",
								"city_name": "",
								"continent_code": "AP",
								"country": "",
								"country_code": "CN",
								"country_code2": "",
								"country_name": "中国",
								"idd_code": "86",
								"isp": "",
								"isp_domain": "电信",
								"latitude": "30.287459",
								"longitude": "120.153576",
								"owner_domain": "",
								"proxy_type": "",
								"region": "",
								"region_name": "浙江",
								"timezone": "Asia/Shanghai",
								"ip": "36.17.88.60",
								"utc_offset": "UTC+8"
							}
					},
					"slice": [ 
						 {
							"china_admin_code": "330001",
							"city": "",
							"city_name": "",
							"continent_code": "AP",
							"country": "",
							"country_code": "CN",
							"country_code2": "",
							"country_name": "中国",
							"idd_code": "86",
							"isp": "",
							"isp_domain": "电信",
							"latitude": "30.287459",
							"longitude": "120.153576",
							"owner_domain": "",
							"proxy_type": "",
							"region": "",
							"region_name": "浙江",
							"timezone": "Asia/Shanghai",
							"ip": "36.17.88.60",
							"utc_offset": "UTC+8"
						},
						{
							"china_admin_code": "330002",
							"city": "",
							"city_name": "",
							"continent_code": "AS",
							"country": "",
							"country_code": "CN",
							"country_code2": "",
							"country_name": "中国",
							"idd_code": "86",
							"isp": "",
							"isp_domain": "电信",
							"latitude": "30.287459",
							"longitude": "120.153576",
							"owner_domain": "",
							"proxy_type": "",
							"region": "",
							"region_name": "浙江",
							"timezone": "Asia/Shanghai",
							"ip": "36.17.88.60",
							"utc_offset": "UTC+8"
						}
					]
			}
    }`

	paths := []string{
        // 单个path
		"ipinfo.val.name.china_admin_code",
		"ipinfo.val.name.city",

        // 通配path
		"ipinfo.*.*.isp",
		"ipinfo.*.name.idd_code",

        // 特殊支持 
		"ipinfo.'1.0.0.1001'.name.city",

        // slice
		"ipinfo.slice.#.city",  
		"ipinfo.slice.#.city_name",
	}
	newjson, _ := jsondel.JsonDeleteBytes([]byte(json), paths) // 用法1
	fmt.Println(string(newjson))
	fmt.Println("----\n\n\n")
	newjson2, _ := jsondel.JsonDeleteString(json, paths) // 用法2
	fmt.Println(newjson2)
	fmt.Println("----\n\n\n")
	j, _ := sj.NewJson([]byte(json))
	j = jsondel.JsonDelete(j, paths) // 用法3
	fmt.Println(j)
	fmt.Println("----\n\n\n")
}
```
