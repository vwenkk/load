package group

import (
	"context"
	"github.com/vwenkk/load/pkg/ent"
	"github.com/vwenkk/load/pkg/ent/group"
	"github.com/vwenkk/load/pkg/ent/predicate"
	"github.com/vwenkk/load/pkg/i18n"
	"github.com/vwenkk/load/pkg/statuserr"

	"github.com/vwenkk/load/rpc/internal/svc"
	"github.com/vwenkk/load/rpc/types/load"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupListLogic {
	return &GetGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupListLogic) GetGroupList(in *load.GroupListReq) (*load.GroupListResp, error) {
	var predicates []predicate.Group
	if in.GetName() != "" {
		predicates = append(predicates, group.NameContains(in.GetName()))
	}
	if in.GetRemark() != "" {
		predicates = append(predicates, group.RemarkContains(in.GetRemark()))
	}
	result, err := l.svcCtx.DB.Group.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize, func(pager *ent.GroupPager) {
		pager.Order = ent.Asc(group.FieldUpdatedAt)
	})
	if err != nil {
		logx.Errorf(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &load.GroupListResp{
		Total: result.PageDetails.Total,
	}
	for _, v := range result.List {
		resp.Data = append(resp.Data, &load.GroupInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.UpdatedAt.Unix(),
			Status:    uint32(v.Status),
			Name:      v.Name,
			Remark:    v.Remark,
		})
	}

	return resp, nil
}
