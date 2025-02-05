package expert

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/ent/expert"
	"github.com/suyuan32/simple-admin-member-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetExpertListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExpertListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExpertListLogic {
	return &GetExpertListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExpertListLogic) GetExpertList(in *mms.ExpertListReq) (*mms.ExpertListResp, error) {
	var predicates []predicate.Expert
	if in.NameZh != nil {
		predicates = append(predicates, expert.NameZhContains(*in.NameZh))
	}
	if in.NameEn != nil {
		predicates = append(predicates, expert.NameEnContains(*in.NameEn))
	}
	if in.NameRu != nil {
		predicates = append(predicates, expert.NameRuContains(*in.NameRu))
	}
	result, err := l.svcCtx.DB.Expert.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mms.ExpertListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mms.ExpertInfo{
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
