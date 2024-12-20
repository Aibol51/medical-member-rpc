package medicine

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMedicineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMedicineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMedicineLogic {
	return &CreateMedicineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMedicineLogic) CreateMedicine(in *mms.MedicineInfo) (*mms.BaseIDResp, error) {
    query := l.svcCtx.DB.Medicine.Create().
			SetNotNilSort(in.Sort).
			SetNotNilNameZh(in.NameZh).
			SetNotNilNameEn(in.NameEn).
			SetNotNilNameRu(in.NameRu).
			SetNotNilNameKk(in.NameKk).
			SetNotNilQuantity(in.Quantity).
			SetNotNilDescriptionZh(in.DescriptionZh).
			SetNotNilDescriptionEn(in.DescriptionEn).
			SetNotNilDescriptionRu(in.DescriptionRu).
			SetNotNilDescriptionKk(in.DescriptionKk).
			SetNotNilRemarks(in.Remarks).
			SetNotNilImages(in.Images)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	result, err := query.Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &mms.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
