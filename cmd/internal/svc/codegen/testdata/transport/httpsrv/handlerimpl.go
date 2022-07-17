package httpsrv

import (
	"context"
	"github.com/goccy/go-json"
	"mime/multipart"
	"net/http"
	service "testdata"
	"testdata/vo"

	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/toolkit/cast"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type UsersvcHandlerImpl struct {
	usersvc service.Usersvc
}

func (receiver *UsersvcHandlerImpl) PageUsers(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx   context.Context
		query vo.PageQuery
		code  int
		data  vo.PageRet
		msg   error
	)
	ctx = _req.Context()
	if _req.Body == nil {
		http.Error(_writer, "missing request body", http.StatusBadRequest)
		return
	} else {
		if _err := json.NewDecoder(_req.Body).Decode(&query); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		} else {
			if _err := ddhttp.ValidateVar(query, "", ""); _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	code, data, msg = receiver.usersvc.PageUsers(
		ctx,
		query,
	)
	if msg != nil {
		if errors.Is(msg, context.Canceled) {
			http.Error(_writer, msg.Error(), http.StatusBadRequest)
		} else if _err, ok := msg.(*ddhttp.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, msg.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Code int        `json:"code,omitempty"`
		Data vo.PageRet `json:"data,omitempty"`
	}{
		Code: code,
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *UsersvcHandlerImpl) GetUser(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		userId string
		photo  string
		code   int
		data   string
		msg    error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["userId"]; exists {
		userId = _req.FormValue("userId")
		if _err := ddhttp.ValidateVar(userId, "", "userId"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter userId", http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["photo"]; exists {
		photo = _req.FormValue("photo")
		if _err := ddhttp.ValidateVar(photo, "", "photo"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter photo", http.StatusBadRequest)
		return
	}
	code, data, msg = receiver.usersvc.GetUser(
		ctx,
		userId,
		photo,
	)
	if msg != nil {
		if errors.Is(msg, context.Canceled) {
			http.Error(_writer, msg.Error(), http.StatusBadRequest)
		} else if _err, ok := msg.(*ddhttp.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, msg.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Code int    `json:"code,omitempty"`
		Data string `json:"data,omitempty"`
	}{
		Code: code,
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *UsersvcHandlerImpl) SignUp(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx      context.Context
		username string
		password int
		actived  bool
		score    []int
		code     int
		data     string
		msg      error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["username"]; exists {
		username = _req.FormValue("username")
		if _err := ddhttp.ValidateVar(username, "", "username"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter username", http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["password"]; exists {
		if casted, _err := cast.ToIntE(_req.FormValue("password")); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		} else {
			password = casted
		}
		if _err := ddhttp.ValidateVar(password, "", "password"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter password", http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["actived"]; exists {
		if casted, _err := cast.ToBoolE(_req.FormValue("actived")); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		} else {
			actived = casted
		}
		if _err := ddhttp.ValidateVar(actived, "", "actived"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter actived", http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["score"]; exists {
		if casted, _err := cast.ToIntSliceE(_req.Form["score"]); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		} else {
			score = casted
		}
		if _err := ddhttp.ValidateVar(score, "", "score"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		if _, exists := _req.Form["score[]"]; exists {
			if casted, _err := cast.ToIntSliceE(_req.Form["score[]"]); _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			} else {
				score = casted
			}
			if _err := ddhttp.ValidateVar(score, "", "score"); _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			http.Error(_writer, "missing parameter score", http.StatusBadRequest)
			return
		}
	}
	code, data, msg = receiver.usersvc.SignUp(
		ctx,
		username,
		password,
		actived,
		score,
	)
	if msg != nil {
		if errors.Is(msg, context.Canceled) {
			http.Error(_writer, msg.Error(), http.StatusBadRequest)
		} else if _err, ok := msg.(*ddhttp.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, msg.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Code int    `json:"code,omitempty"`
		Data string `json:"data,omitempty"`
	}{
		Code: code,
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *UsersvcHandlerImpl) UploadAvatar(_writer http.ResponseWriter, _req *http.Request) {
	var (
		pc  context.Context
		pf  []v3.FileModel
		ps  string
		pf2 v3.FileModel
		pf3 *multipart.FileHeader
		pf4 []*multipart.FileHeader
		ri  int
		ri2 interface{}
		re  error
	)
	pc = _req.Context()
	if _err := _req.ParseMultipartForm(32 << 20); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	pfFileHeaders, exists := _req.MultipartForm.File["pf"]
	if exists {
		if len(pfFileHeaders) == 0 {
			http.Error(_writer, "no file uploaded for parameter pf", http.StatusBadRequest)
			return
		}
		for _, _fh := range pfFileHeaders {
			_f, _err := _fh.Open()
			if _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
			pf = append(pf, v3.FileModel{
				Filename: _fh.Filename,
				Reader:   _f,
			})
		}
	} else {
		http.Error(_writer, "missing parameter pf", http.StatusBadRequest)
		return
	}
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["ps"]; exists {
		ps = _req.FormValue("ps")
		if _err := ddhttp.ValidateVar(ps, "", "ps"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter ps", http.StatusBadRequest)
		return
	}
	pf2FileHeaders, exists := _req.MultipartForm.File["pf2"]
	if exists {
		if len(pf2FileHeaders) == 0 {
			http.Error(_writer, "no file uploaded for parameter pf2", http.StatusBadRequest)
			return
		}
		if len(pf2FileHeaders) > 0 {
			_fh := pf2FileHeaders[0]
			_f, _err := _fh.Open()
			if _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
			pf2 = v3.FileModel{
				Filename: _fh.Filename,
				Reader:   _f,
			}
		}
	} else {
		http.Error(_writer, "missing parameter pf2", http.StatusBadRequest)
		return
	}
	pf3Files := _req.MultipartForm.File["pf3"]
	if len(pf3Files) > 0 {
		pf3 = pf3Files[0]
	}
	pf4 = _req.MultipartForm.File["pf4"]
	ri, ri2, re = receiver.usersvc.UploadAvatar(
		pc,
		pf,
		ps,
		pf2,
		pf3,
		pf4,
	)
	if re != nil {
		if errors.Is(re, context.Canceled) {
			http.Error(_writer, re.Error(), http.StatusBadRequest)
		} else if _err, ok := re.(*ddhttp.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, re.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Ri  int         `json:"ri,omitempty"`
		Ri2 interface{} `json:"ri2,omitempty"`
	}{
		Ri:  ri,
		Ri2: ri2,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewUsersvcHandler(usersvc service.Usersvc) UsersvcHandler {
	return &UsersvcHandlerImpl{
		usersvc,
	}
}
