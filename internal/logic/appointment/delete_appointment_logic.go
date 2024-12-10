package appointment

import (
	"context"

    "github.com/suyuan32/simple-admin-member-rpc/ent/appointment"
    "github.com/suyuan32/simple-admin-member-rpc/internal/svc"
    "github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
    "github.com/suyuan32/simple-admin-member-rpc/types/mms"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/suyuan32/simple-admin-common/utils/uuidx"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteAppointmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAppointmentLogic {
	return &DeleteAppointmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAppointmentLogic) DeleteAppointment(in *mms.UUIDsReq) (*mms.BaseResp, error) {
	_, err := l.svcCtx.DB.Appointment.Delete().Where(appointment.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &mms.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
