package service

import "github.com/unionj-cloud/go-doudou/framework/http/onlinedoc"

func init() {
	onlinedoc.Oas = `{"openapi":"3.0.2","info":{"title":"Testsvc","version":"v20210909"},"servers":[{}],"paths":{"/page/users":{"post":{"description":"You can define your service methods as your need. Below is an example.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/PageQuery"}}},"required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/PageUsersResp"}}}}}}}},"components":{"schemas":{"Order":{"title":"Order","type":"object","properties":{"Col":{"type":"string"},"Sort":{"type":"string"}}},"Page":{"title":"Page","type":"object","properties":{"Orders":{"type":"array","items":{"$ref":"#/components/schemas/Order"},"description":"排序规则"},"PageNo":{"type":"integer","format":"int32","description":"页码"},"Size":{"type":"integer","format":"int32","description":"每页行数"}}},"PageFilter":{"title":"PageFilter","type":"object","properties":{"Dept":{"type":"integer","format":"int32","description":"所属部门ID"},"Name":{"type":"string","description":"真实姓名，前缀匹配"}}},"PageQuery":{"title":"PageQuery","type":"object","properties":{"Filter":{"$ref":"#/components/schemas/PageFilter"},"Page":{"$ref":"#/components/schemas/Page"}},"description":"分页筛选条件"},"PageRet":{"title":"PageRet","type":"object","properties":{"HasNext":{"type":"boolean"},"Items":{"type":"object"},"PageNo":{"type":"integer","format":"int32"},"PageSize":{"type":"integer","format":"int32"},"Total":{"type":"integer","format":"int32"}}},"PageUsersResp":{"title":"PageUsersResp","type":"object","properties":{"code":{"type":"integer","format":"int32"},"data":{"$ref":"#/components/schemas/PageRet"},"err":{"type":"string"}}},"UserVo":{"title":"UserVo","type":"object","properties":{"Dept":{"type":"string"},"Id":{"type":"integer","format":"int32"},"Name":{"type":"string"},"Phone":{"type":"string"}}}}}}`
}
