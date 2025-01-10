package medicalRecord

import (
	"context"

    "github.com/suyuan32/simple-admin-member-rpc/ent/medicalrecord"
    "github.com/suyuan32/simple-admin-member-rpc/internal/svc"
    "github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
    "github.com/suyuan32/simple-admin-member-rpc/types/mms"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/suyuan32/simple-admin-common/utils/uuidx"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteMedicalRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMedicalRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMedicalRecordLogic {
	return &DeleteMedicalRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMedicalRecordLogic) DeleteMedicalRecord(in *mms.UUIDsReq) (*mms.BaseResp, error) {
	_, err := l.svcCtx.DB.MedicalRecord.Delete().Where(medicalrecord.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &mms.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
