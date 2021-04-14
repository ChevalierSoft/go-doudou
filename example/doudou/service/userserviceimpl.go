package service

import (
	"context"
	"example/doudou/dao"
	"example/doudou/db"
	"example/doudou/domain"
	"example/doudou/vo"
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou/ddl/query"
)

type UserServiceImpl struct{}

func (u UserServiceImpl) PostSignUp(ctx context.Context, form vo.SignUpForm) (int, error) {
	panic("implement me")
}

func (u UserServiceImpl) PostLogIn(ctx context.Context, form vo.LogInForm) (vo.Auth, error) {
	panic("implement me")
}

func (u UserServiceImpl) GetUser(ctx context.Context, id int) (vo.UserVo, error) {
	userdao := dao.NewUserDao(db.Db())
	var user domain.User
	var err error
	if user, err = userdao.GetUser(ctx, id); err != nil {
		return vo.UserVo{}, errors.Wrap(err, "Error from calling userdao.GetUser")
	}
	return vo.UserVo{
		Id:    user.Id,
		Name:  user.Name,
		Phone: user.Phone,
		Dept:  user.Dept,
	}, nil
}

func (u UserServiceImpl) PostPageUsers(ctx context.Context, query vo.PageQuery) (query.PageRet, error) {
	panic("implement me")
}

func NewUserService() UserService {
	return UserServiceImpl{}
}
