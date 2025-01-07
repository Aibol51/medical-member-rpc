package appointment

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/ent/appointment"
	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppointmentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppointmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppointmentListLogic {
	return &GetAppointmentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppointmentListLogic) GetAppointmentList(in *mms.AppointmentListReq) (*mms.AppointmentListResp, error) {
	userId := l.ctx.Value("userId").(string)
	result, err := l.svcCtx.DB.Appointment.Query().Where(appointment.UserIDContains(userId)).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mms.AppointmentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mms.AppointmentInfo{
			Id:              pointy.GetPointer(v.ID.String()),
			CreatedAt:       pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:       pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			PatientName:     &v.PatientName,
			PhoneNumber:     &v.PhoneNumber,
			IdCard:          &v.IDCard,
			Gender:          &v.Gender,
			Age:             &v.Age,
			AppointmentTime: &v.AppointmentTime,
			Symptoms:        &v.Symptoms,
			Status:          &v.Status,
			Remarks:         &v.Remarks,
			UserId:          &v.UserID,
		})
	}

	return resp, nil
}
