package service

import "github.com/unionj-cloud/go-doudou/svc/http/onlinedoc"

func init() {
	onlinedoc.Oas = `{"openapi":"3.0.2","info":{"title":"Usersvc","description":"用户服务接口\nv1版本","version":"v20220202"},"paths":{"/usersvc/downloadavatar":{"post":{"description":"comment5","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"$ref":"#/components/schemas/DownloadAvatarReq"}}},"required":true},"responses":{"200":{"content":{"application/octet-stream":{"schema":{"type":"string","format":"binary"}}}}}}},"/usersvc/pageusers":{"post":{"description":"You can define your service methods as your need. Below is an example.","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/PageQuery"}}},"required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/PageUsersResp"}}}}}}},"/usersvc/signup":{"post":{"description":"comment3","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"$ref":"#/components/schemas/SignUpReq"}}},"required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/SignUpResp"}}}}}}},"/usersvc/uploadavatar":{"post":{"description":"comment4","requestBody":{"content":{"multipart/form-data":{"schema":{"$ref":"#/components/schemas/UploadAvatarReq"}}},"required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/UploadAvatarResp"}}}}}}},"/usersvc/user":{"get":{"description":"comment1\ncomment2","parameters":[{"name":"userId","in":"query","description":"用户ID","required":true,"schema":{"type":"string","description":"用户ID"}},{"name":"photo","in":"query","description":"图片地址","required":true,"schema":{"type":"string","description":"图片地址"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetUserResp"}}}}}}}},"components":{"schemas":{"DownloadAvatarReq":{"title":"DownloadAvatarReq","type":"object","properties":{"userId":{"type":"string"}},"required":["userId"]},"Event":{"title":"Event","type":"object","properties":{"EventType":{"type":"integer","format":"int32"},"Name":{"type":"string"}},"required":["Name","EventType"]},"GetUserResp":{"title":"GetUserResp","type":"object","properties":{"code":{"type":"integer","format":"int32"},"data":{"type":"string"}},"required":["code","data"]},"Order":{"title":"Order","type":"object","properties":{"Col":{"type":"string"},"Sort":{"type":"string"}},"description":"排序条件","required":["Col","Sort"]},"Page":{"title":"Page","type":"object","properties":{"Orders":{"type":"array","items":{"$ref":"#/components/schemas/Order"},"description":"排序规则"},"PageNo":{"type":"integer","format":"int32","description":"页码"},"Size":{"type":"integer","format":"int32","description":"每页行数"},"User":{"$ref":"#/components/schemas/UserVo"}},"required":["Orders","PageNo","Size","User"]},"PageFilter":{"title":"PageFilter","type":"object","properties":{"Dept":{"type":"integer","format":"int32","description":"所属部门ID"},"Name":{"type":"string","description":"真实姓名，前缀匹配"}},"description":"筛选条件","required":["Name","Dept"]},"PageQuery":{"title":"PageQuery","type":"object","properties":{"Filter":{"$ref":"#/components/schemas/PageFilter"},"Page":{"$ref":"#/components/schemas/Page"}},"description":"分页筛选条件","required":["Filter","Page"]},"PageRet":{"title":"PageRet","type":"object","properties":{"HasNext":{"type":"boolean"},"Items":{"type":"object"},"PageNo":{"type":"integer","format":"int32"},"PageSize":{"type":"integer","format":"int32"},"Total":{"type":"integer","format":"int32"}},"required":["Items","PageNo","PageSize","Total","HasNext"]},"PageUsersResp":{"title":"PageUsersResp","type":"object","properties":{"code":{"type":"integer","format":"int32"},"data":{"$ref":"#/components/schemas/PageRet"}},"required":["code","data"]},"SignUpReq":{"title":"SignUpReq","type":"object","properties":{"actived":{"type":"boolean"},"password":{"type":"integer","format":"int32"},"score":{"type":"array","items":{"type":"integer","format":"int32"}},"username":{"type":"string"}},"required":["username","password","actived","score"]},"SignUpResp":{"title":"SignUpResp","type":"object","properties":{"code":{"type":"integer","format":"int32"},"data":{"type":"string"}},"required":["code","data"]},"TestAlias":{"title":"TestAlias","type":"object","properties":{"Age":{"type":"object"},"School":{"type":"array","items":{"type":"object","properties":{"Addr":{"type":"object","properties":{"Block":{"type":"string"},"Full":{"type":"string"},"Zip":{"type":"string"}},"required":["Zip","Block","Full"]},"Name":{"type":"string"}},"required":["Name","Addr"]}}},"required":["Age","School"]},"UploadAvatarReq":{"title":"UploadAvatarReq","type":"object","properties":{"pf":{"type":"array","items":{"type":"string","format":"binary"}},"pf2":{"type":"string","format":"binary"},"pf3":{"type":"string","format":"binary"},"pf4":{"type":"array","items":{"type":"string","format":"binary"}},"ps":{"type":"string"}},"required":["pf","ps","pf2","pf4"]},"UploadAvatarResp":{"title":"UploadAvatarResp","type":"object","properties":{"ri":{"type":"integer","format":"int32"},"rs":{"type":"string"}},"required":["ri","rs"]},"UserVo":{"title":"UserVo","type":"object","properties":{"Dept":{"type":"string"},"Id":{"type":"integer","format":"int32"},"Name":{"type":"string"},"Phone":{"type":"string"}},"required":["Id","Name","Phone","Dept"]}}}}`
}
