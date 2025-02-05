package service

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-member-rpc/ent/service"
	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServiceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceListLogic {
	return &GetServiceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServiceListLogic) GetServiceList(in *mms.ServiceListReq) (*mms.ServiceListResp, error) {
	var predicates []predicate.Service
	if in.NameZh != nil {
		predicates = append(predicates, service.NameZhContains(*in.NameZh))
	}
	if in.NameEn != nil {
		predicates = append(predicates, service.NameEnContains(*in.NameEn))
	}
	if in.NameRu != nil {
		predicates = append(predicates, service.NameRuContains(*in.NameRu))
	}
	result, err := l.svcCtx.DB.Service.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mms.ServiceListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mms.ServiceInfo{
			Id:        &v.ID,
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:    pointy.GetPointer(uint32(v.Status)),
			Sort:      &v.Sort,
			NameZh:    &v.NameZh,
			NameEn:    &v.NameEn,
			NameRu:    &v.NameRu,
			NameKk:    &v.NameKk,
			CoverUrl:  &v.CoverURL,
		})
	}

	return resp, nil
}
