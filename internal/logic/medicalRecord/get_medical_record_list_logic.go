package medicalRecord

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/ent/medicalrecord"
	"github.com/suyuan32/simple-admin-member-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicalRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMedicalRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicalRecordListLogic {
	return &GetMedicalRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMedicalRecordListLogic) GetMedicalRecordList(in *mms.MedicalRecordListReq) (*mms.MedicalRecordListResp, error) {
	var predicates []predicate.MedicalRecord
	if in.PatientName != nil {
		predicates = append(predicates, medicalrecord.PatientNameContains(*in.PatientName))
	}
	if in.PhoneNumber != nil {
		predicates = append(predicates, medicalrecord.PhoneNumberContains(*in.PhoneNumber))
	}
	result, err := l.svcCtx.DB.MedicalRecord.Query().Where(medicalrecord.UserIDContains(in.UserId)).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mms.MedicalRecordListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mms.MedicalRecordInfo{
			Id:                 pointy.GetPointer(v.ID.String()),
			CreatedAt:          pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:          pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			PatientName:        &v.PatientName,
			PhoneNumber:        &v.PhoneNumber,
			Gender:             &v.Gender,
			Age:                &v.Age,
			VisitTime:          &v.VisitTime,
			Diagnosis:          &v.Diagnosis,
			TreatmentPlan:      &v.TreatmentPlan,
			Prescription:       &v.Prescription,
			ExaminationResults: &v.ExaminationResults,
			DoctorAdvice:       &v.DoctorAdvice,
			DoctorId:           &v.DoctorID,
			Department:         &v.Department,
			AppointmentId:      &v.AppointmentID,
			Remarks:            &v.Remarks,
			UserId:             &v.UserID,
		})
	}

	return resp, nil
}
