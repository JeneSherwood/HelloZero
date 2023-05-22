package logic

import (
	"context"
	"strconv"
	"time"

	"go-zero-demo/book/common/errorx"
	"go-zero-demo/book/service/user/cmd/api/internal/logic/utils"
	"go-zero-demo/book/service/user/cmd/api/internal/svc"
	"go-zero-demo/book/service/user/cmd/api/internal/types"
	"go-zero-demo/book/service/user/model"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Exist(req *types.RegisterReq) error {
	_, err := l.svcCtx.UserModel.FindOneByNumber(req.Number)
	switch err {
	case nil:
		return errorx.NewDefaultError("用户已存在")
	case model.ErrNotFound:
		return nil
	default:
		return err
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterReply, err error) {

	err = utils.CheckRegisterParams(req)
	if err != nil {
		return &types.RegisterReply{
			Number:         "",
			RegisterResult: false,
		}, err
	}

	err = l.Exist(req)
	if err != nil {
		return &types.RegisterReply{
			Number:         "",
			RegisterResult: false,
		}, errorx.NewDefaultError("学号已注册")
	}

	newID, err := uuid.NewUUID()
	if err != nil {
		return &types.RegisterReply{
			Number:         "",
			RegisterResult: false,
		}, errorx.NewDefaultError("UUID generation failed")
	}

	insertArgs := model.User{
		Id:         int64(newID.ID()),
		Number:     req.Number,
		Name:       req.Name,
		Password:   req.Password,
		Gender:     req.Gender,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	_, err = l.svcCtx.UserModel.Insert(insertArgs)
	if err != nil {
		return &types.RegisterReply{
			Number:         "",
			RegisterResult: false,
		}, errorx.NewDefaultError("插入数据失败")
	}

	respNum := strconv.Itoa(int(insertArgs.Id))

	return &types.RegisterReply{
		Number:         respNum,
		RegisterResult: true,
	}, nil
}
