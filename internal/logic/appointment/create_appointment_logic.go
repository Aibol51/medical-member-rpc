package appointment

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAppointmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAppointmentLogic {
	return &CreateAppointmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAppointmentLogic) CreateAppointment(in *mms.AppointmentInfo) (*mms.BaseUUIDResp, error) {
	result, err := l.svcCtx.DB.Appointment.Create().
		SetNotNilPatientName(in.PatientName).
		SetNotNilPhoneNumber(in.PhoneNumber).
		SetNotNilIDCard(in.IdCard).
		SetNotNilGender(in.Gender).
		SetNotNilAge(in.Age).
		SetNotNilAppointmentTime(in.AppointmentTime).
		SetNotNilSymptoms(in.Symptoms).
		SetNotNilStatus(in.Status).
		SetNotNilRemarks(in.Remarks).
		SetNotNilUserID(in.UserId).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
