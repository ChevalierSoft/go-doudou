package service

import "github.com/unionj-cloud/go-doudou/v2/framework/rest"

func init() {
	rest.Oas = `{"openapi":"3.0.2","info":{"title":"Usersvc","description":"用户服务接口\nv1版本","version":"v20221106"},"servers":[{"url":"http://localhost:6060"}],"paths":{"/usersvc/downloadavatar":{"post":{"description":"comment5","parameters":[{"name":"data","in":"query","required":true,"schema":{"type":"string"}},{"name":"userAttrs","in":"query","schema":{"type":"array","items":{"type":"string"}}}],"requestBody":{"content":{"application/json":{"schema":{"type":"object"}}},"required":true},"responses":{"200":{"description":"","content":{"application/octet-stream":{"schema":{"type":"string","format":"binary"}}}}}}},"/usersvc/pageusers":{"post":{"description":"You can define your service methods as your need. Below is an example.@role(user)","requestBody":{"content":{"application/json":{"schema":{"type":"object"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/PageUsersResp"}}}}}}},"/usersvc/signup":{"post":{"description":"comment3\n@permission(create,update)@role(admin)","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"$ref":"#/components/schemas/SignUpReq"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/SignUpResp"}}}}}}},"/usersvc/uploadavatar":{"post":{"description":"comment4\n@role(user)","requestBody":{"content":{"multipart/form-data":{"schema":{"$ref":"#/components/schemas/UploadAvatarReq"}}},"required":true},"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/UploadAvatarResp"}}}}}}},"/usersvc/user":{"get":{"description":"comment1\ncomment2\n@role(admin)","parameters":[{"name":"userId","in":"query","description":"用户ID","required":true,"schema":{"type":"string","description":"用户ID"}},{"name":"photo","in":"query","description":"图片地址","required":true,"schema":{"type":"string","description":"图片地址"}}],"responses":{"200":{"description":"","content":{"application/json":{"schema":{"$ref":"#/components/schemas/GetUserResp"}}}}}}}},"components":{"schemas":{"GetUserResp":{"title":"GetUserResp","type":"object","properties":{"code":{"type":"integer","format":"int32"},"data":{"type":"string"}},"required":["code","data"]},"PageUsersResp":{"title":"PageUsersResp","type":"object","properties":{"code":{"type":"integer","format":"int32"},"data":{"type":"object"}},"required":["code","data"]},"SignUpReq":{"title":"SignUpReq","type":"object","properties":{"actived":{"type":"boolean"},"password":{"type":"integer","format":"int32"},"score":{"type":"array","items":{"type":"integer","format":"int32"}},"username":{"type":"string"}},"required":["username","password","actived","score"]},"SignUpResp":{"title":"SignUpResp","type":"object","properties":{"code":{"type":"integer","format":"int32"},"data":{"type":"string"}},"required":["code","data"]},"UploadAvatarReq":{"title":"UploadAvatarReq","type":"object","properties":{"pf":{"type":"array","items":{"type":"string","format":"binary"}},"pf2":{"type":"string","format":"binary"},"pf3":{"type":"string","format":"binary"},"pf4":{"type":"array","items":{"type":"string","format":"binary"}},"ps":{"type":"string"}},"required":["pf","ps","pf2","pf4"]},"UploadAvatarResp":{"title":"UploadAvatarResp","type":"object","properties":{"ri":{"type":"integer","format":"int32"},"ri2":{"type":"object"}},"required":["ri","ri2"]}}}}`
}
