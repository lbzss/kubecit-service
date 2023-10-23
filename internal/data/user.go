package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit-service/ent"
	"kubecit-service/ent/account"
	"kubecit-service/ent/user"
	"kubecit-service/ent/viporder"
	"kubecit-service/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo 用户数据仓库构造方法
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (repo *userRepo) FindById(ctx context.Context, id uint64) (po *biz.UserPO, err error) {
	user, err := repo.data.db.User.Query().Where(user.IDEQ(int(id))).First(ctx)
	if err != nil {
		return nil, err
	}
	if user.RoleID == 3 {
		teacher, err := user.QueryTeacher().Only(ctx)
		if err != nil {
			return nil, err
		}
		return &biz.UserPO{
			Id:        uint64(user.ID),
			Username:  user.Username,
			Channel:   user.Channel,
			RoleId:    user.RoleID,
			UserId:    user.ID,
			TeacherId: teacher.ID,
		}, nil
	}

	return &biz.UserPO{
		Id:       uint64(user.ID),
		Username: user.Username,
		Channel:  user.Channel,
		RoleId:   user.RoleID,
		UserId:   user.ID,
	}, nil

}

func (repo *userRepo) Save(ctx context.Context, userPO *biz.UserPO) error {
	if userPO.Id == 0 {
		nUser, err := repo.data.db.User.Create().
			SetUsername(userPO.Username).
			SetChannel(userPO.Channel).
			SetRoleID(userPO.RoleId).
			Save(ctx)
		if err != nil {
			return err
		}
		userPO.Id = uint64(nUser.ID)
		return nil
	} else {
		_, err := repo.data.db.User.Update().
			SetUsername(userPO.Username).
			SetChannel(userPO.Channel).
			SetRoleID(userPO.RoleId).
			Where(user.IDEQ(int(userPO.Id))).
			Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}

}

func (repo *userRepo) SaveAccountAndUserTx(ctx context.Context, accountPO *biz.AccountPO, userPO *biz.UserPO) error {
	if err := repo.data.WithTx(ctx, func(tx *ent.Tx) error {
		return repo.saveAccountAndUser(ctx, accountPO, userPO)
	}); err != nil {
		return err
	}
	return nil
}

func (repo *userRepo) saveAccountAndUser(ctx context.Context, accountPO *biz.AccountPO, userPO *biz.UserPO) error {
	if userPO.Id == 0 {
		nUser, err := repo.data.db.User.Create().
			SetUsername(userPO.Username).
			SetChannel(userPO.Channel).
			SetRoleID(userPO.RoleId).
			Save(ctx)
		if err != nil {
			return err
		}
		userPO.Id = uint64(nUser.ID)
	} else {
		_, err := repo.data.db.User.Update().
			SetUsername(userPO.Username).
			SetChannel(userPO.Channel).
			SetRoleID(userPO.RoleId).
			Where(user.IDEQ(int(userPO.Id))).
			Save(ctx)
		if err != nil {
			return err
		}

	}
	if accountPO.Id == 0 {
		_, err := repo.data.db.Account.Create().
			SetOpenid(accountPO.Openid).
			SetPassword(accountPO.Password).
			SetUserID(userPO.Id).
			SetMethod(accountPO.Method).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err := repo.data.db.Account.Update().
			SetOpenid(accountPO.Openid).
			SetPassword(accountPO.Password).
			SetUserID(accountPO.UserId).
			SetMethod(accountPO.Method).
			Where(account.IDEQ(int(accountPO.Id))).
			Save(ctx)
		if err != nil {
			return err
		}

	}
	return nil
}

func (repo *userRepo) Become(ctx context.Context, req *biz.BecomeOrderInfo) error {
	_, err := repo.data.db.VipOrder.Create().SetBizID(req.BizId).SetPrice(req.Price).SetVipType(int8(req.VipType)).
		SetPayType(int8(req.PayType)).SetUserID(req.UserId).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepo) Callback(ctx context.Context, req *biz.VipInfo, payStatus int) (*biz.VipInfo, error) {
	result, err := repo.data.WithResultTx(ctx, func(tx *ent.Tx) (interface{}, error) {
		return repo.callback(ctx, req, payStatus)
	})
	if err != nil {
		return nil, err
	}
	if vipInfo, ok := result.(*biz.VipInfo); ok {
		return vipInfo, nil
	}
	return nil, err
}

func (repo *userRepo) callback(ctx context.Context, req *biz.VipInfo, payStatus int) (*biz.VipInfo, error) {
	order, err := repo.data.db.VipOrder.Query().Where(viporder.BizIDEQ(req.VipOrderId)).Only(ctx)
	if err != nil {
		repo.log.Errorf("query vip order error: %v", err)
		return nil, err
	}
	order, err = repo.data.db.VipOrder.UpdateOne(order).SetPayStatus(int8(payStatus)).Save(ctx)
	if err != nil {
		repo.log.Errorf("update vip order error: %v", err)
		return nil, err
	}
	info, err := repo.data.db.VipInfo.Create().SetUserID(order.UserID).SetVipType(order.VipType).SetStartAt(req.StartAt).
		SetExpireAt(req.ExpireAt).Save(ctx)
	if err != nil {
		repo.log.Errorf("create  vip info error: %v", err)
		return nil, err
	}
	return &biz.VipInfo{
		Id:       info.ID,
		VipType:  int(info.VipType),
		StartAt:  info.StartAt,
		ExpireAt: info.ExpireAt,
		UserId:   info.UserID,
	}, nil
}

func (repo *userRepo) GetOrderInfo(ctx context.Context, bizId int64) (*biz.BecomeOrderInfo, error) {
	record, err := repo.data.db.Debug().VipOrder.Query().Where(viporder.BizIDEQ(bizId)).Only(ctx)
	if err != nil {
		repo.log.Errorf("query order info error: %v", err)
		return nil, err
	}
	return &biz.BecomeOrderInfo{
		UserId:    record.UserID,
		PayType:   int(record.PayType),
		VipType:   int(record.VipType),
		PayStatus: int(record.PayStatus),
		BizId:     record.BizID,
		Price:     record.Price,
	}, nil
}
