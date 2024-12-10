// Code generated by goctl. DO NOT EDIT.
// Source: mms.proto

package server

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/internal/logic/appointment"
	"github.com/suyuan32/simple-admin-member-rpc/internal/logic/base"
	"github.com/suyuan32/simple-admin-member-rpc/internal/logic/medicine"
	"github.com/suyuan32/simple-admin-member-rpc/internal/logic/member"
	"github.com/suyuan32/simple-admin-member-rpc/internal/logic/memberrank"
	"github.com/suyuan32/simple-admin-member-rpc/internal/logic/news"
	"github.com/suyuan32/simple-admin-member-rpc/internal/logic/oauthprovider"
	"github.com/suyuan32/simple-admin-member-rpc/internal/logic/swiper"
	"github.com/suyuan32/simple-admin-member-rpc/internal/logic/token"
	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"
)

type MmsServer struct {
	svcCtx *svc.ServiceContext
	mms.UnimplementedMmsServer
}

func NewMmsServer(svcCtx *svc.ServiceContext) *MmsServer {
	return &MmsServer{
		svcCtx: svcCtx,
	}
}

// Appointment management
func (s *MmsServer) CreateAppointment(ctx context.Context, in *mms.AppointmentInfo) (*mms.BaseUUIDResp, error) {
	l := appointment.NewCreateAppointmentLogic(ctx, s.svcCtx)
	return l.CreateAppointment(in)
}

func (s *MmsServer) UpdateAppointment(ctx context.Context, in *mms.AppointmentInfo) (*mms.BaseResp, error) {
	l := appointment.NewUpdateAppointmentLogic(ctx, s.svcCtx)
	return l.UpdateAppointment(in)
}

func (s *MmsServer) GetAppointmentList(ctx context.Context, in *mms.AppointmentListReq) (*mms.AppointmentListResp, error) {
	l := appointment.NewGetAppointmentListLogic(ctx, s.svcCtx)
	return l.GetAppointmentList(in)
}

func (s *MmsServer) GetAppointmentById(ctx context.Context, in *mms.UUIDReq) (*mms.AppointmentInfo, error) {
	l := appointment.NewGetAppointmentByIdLogic(ctx, s.svcCtx)
	return l.GetAppointmentById(in)
}

func (s *MmsServer) DeleteAppointment(ctx context.Context, in *mms.UUIDsReq) (*mms.BaseResp, error) {
	l := appointment.NewDeleteAppointmentLogic(ctx, s.svcCtx)
	return l.DeleteAppointment(in)
}

func (s *MmsServer) InitDatabase(ctx context.Context, in *mms.Empty) (*mms.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

// Medicine management
func (s *MmsServer) GetMedicineList(ctx context.Context, in *mms.MedicineListReq) (*mms.MedicineListResp, error) {
	l := medicine.NewGetMedicineListLogic(ctx, s.svcCtx)
	return l.GetMedicineList(in)
}

func (s *MmsServer) GetMedicineById(ctx context.Context, in *mms.IDReq) (*mms.MedicineInfo, error) {
	l := medicine.NewGetMedicineByIdLogic(ctx, s.svcCtx)
	return l.GetMedicineById(in)
}

// Member management
func (s *MmsServer) CreateMember(ctx context.Context, in *mms.MemberInfo) (*mms.BaseUUIDResp, error) {
	l := member.NewCreateMemberLogic(ctx, s.svcCtx)
	return l.CreateMember(in)
}

func (s *MmsServer) UpdateMember(ctx context.Context, in *mms.MemberInfo) (*mms.BaseResp, error) {
	l := member.NewUpdateMemberLogic(ctx, s.svcCtx)
	return l.UpdateMember(in)
}

func (s *MmsServer) GetMemberList(ctx context.Context, in *mms.MemberListReq) (*mms.MemberListResp, error) {
	l := member.NewGetMemberListLogic(ctx, s.svcCtx)
	return l.GetMemberList(in)
}

func (s *MmsServer) DeleteMember(ctx context.Context, in *mms.UUIDsReq) (*mms.BaseResp, error) {
	l := member.NewDeleteMemberLogic(ctx, s.svcCtx)
	return l.DeleteMember(in)
}

func (s *MmsServer) GetMemberById(ctx context.Context, in *mms.UUIDReq) (*mms.MemberInfo, error) {
	l := member.NewGetMemberByIdLogic(ctx, s.svcCtx)
	return l.GetMemberById(in)
}

func (s *MmsServer) GetMemberByUsername(ctx context.Context, in *mms.UsernameReq) (*mms.MemberInfo, error) {
	l := member.NewGetMemberByUsernameLogic(ctx, s.svcCtx)
	return l.GetMemberByUsername(in)
}

func (s *MmsServer) GetMemberByMobile(ctx context.Context, in *mms.MobileReq) (*mms.MemberInfo, error) {
	l := member.NewGetMemberByMobileLogic(ctx, s.svcCtx)
	return l.GetMemberByMobile(in)
}

// MemberRank management
func (s *MmsServer) CreateMemberRank(ctx context.Context, in *mms.MemberRankInfo) (*mms.BaseIDResp, error) {
	l := memberrank.NewCreateMemberRankLogic(ctx, s.svcCtx)
	return l.CreateMemberRank(in)
}

func (s *MmsServer) UpdateMemberRank(ctx context.Context, in *mms.MemberRankInfo) (*mms.BaseResp, error) {
	l := memberrank.NewUpdateMemberRankLogic(ctx, s.svcCtx)
	return l.UpdateMemberRank(in)
}

func (s *MmsServer) GetMemberRankList(ctx context.Context, in *mms.MemberRankListReq) (*mms.MemberRankListResp, error) {
	l := memberrank.NewGetMemberRankListLogic(ctx, s.svcCtx)
	return l.GetMemberRankList(in)
}

func (s *MmsServer) GetMemberRankById(ctx context.Context, in *mms.IDReq) (*mms.MemberRankInfo, error) {
	l := memberrank.NewGetMemberRankByIdLogic(ctx, s.svcCtx)
	return l.GetMemberRankById(in)
}

func (s *MmsServer) DeleteMemberRank(ctx context.Context, in *mms.IDsReq) (*mms.BaseResp, error) {
	l := memberrank.NewDeleteMemberRankLogic(ctx, s.svcCtx)
	return l.DeleteMemberRank(in)
}

// News management
func (s *MmsServer) GetNewsList(ctx context.Context, in *mms.NewsListReq) (*mms.NewsListResp, error) {
	l := news.NewGetNewsListLogic(ctx, s.svcCtx)
	return l.GetNewsList(in)
}

func (s *MmsServer) GetNewsById(ctx context.Context, in *mms.IDReq) (*mms.NewsInfo, error) {
	l := news.NewGetNewsByIdLogic(ctx, s.svcCtx)
	return l.GetNewsById(in)
}

// OauthProvider management
func (s *MmsServer) CreateOauthProvider(ctx context.Context, in *mms.OauthProviderInfo) (*mms.BaseIDResp, error) {
	l := oauthprovider.NewCreateOauthProviderLogic(ctx, s.svcCtx)
	return l.CreateOauthProvider(in)
}

func (s *MmsServer) UpdateOauthProvider(ctx context.Context, in *mms.OauthProviderInfo) (*mms.BaseResp, error) {
	l := oauthprovider.NewUpdateOauthProviderLogic(ctx, s.svcCtx)
	return l.UpdateOauthProvider(in)
}

func (s *MmsServer) GetOauthProviderList(ctx context.Context, in *mms.OauthProviderListReq) (*mms.OauthProviderListResp, error) {
	l := oauthprovider.NewGetOauthProviderListLogic(ctx, s.svcCtx)
	return l.GetOauthProviderList(in)
}

func (s *MmsServer) GetOauthProviderById(ctx context.Context, in *mms.IDReq) (*mms.OauthProviderInfo, error) {
	l := oauthprovider.NewGetOauthProviderByIdLogic(ctx, s.svcCtx)
	return l.GetOauthProviderById(in)
}

func (s *MmsServer) DeleteOauthProvider(ctx context.Context, in *mms.IDsReq) (*mms.BaseResp, error) {
	l := oauthprovider.NewDeleteOauthProviderLogic(ctx, s.svcCtx)
	return l.DeleteOauthProvider(in)
}

func (s *MmsServer) OauthLogin(ctx context.Context, in *mms.OauthLoginReq) (*mms.OauthRedirectResp, error) {
	l := oauthprovider.NewOauthLoginLogic(ctx, s.svcCtx)
	return l.OauthLogin(in)
}

func (s *MmsServer) OauthCallback(ctx context.Context, in *mms.CallbackReq) (*mms.MemberInfo, error) {
	l := oauthprovider.NewOauthCallbackLogic(ctx, s.svcCtx)
	return l.OauthCallback(in)
}

func (s *MmsServer) WechatMiniProgramLogin(ctx context.Context, in *mms.OauthLoginReq) (*mms.BaseResp, error) {
	l := oauthprovider.NewWechatMiniProgramLoginLogic(ctx, s.svcCtx)
	return l.WechatMiniProgramLogin(in)
}

// Swiper management
func (s *MmsServer) GetSwiperList(ctx context.Context, in *mms.SwiperListReq) (*mms.SwiperListResp, error) {
	l := swiper.NewGetSwiperListLogic(ctx, s.svcCtx)
	return l.GetSwiperList(in)
}

func (s *MmsServer) GetSwiperById(ctx context.Context, in *mms.IDReq) (*mms.SwiperInfo, error) {
	l := swiper.NewGetSwiperByIdLogic(ctx, s.svcCtx)
	return l.GetSwiperById(in)
}

// Token management
func (s *MmsServer) CreateToken(ctx context.Context, in *mms.TokenInfo) (*mms.BaseUUIDResp, error) {
	l := token.NewCreateTokenLogic(ctx, s.svcCtx)
	return l.CreateToken(in)
}

func (s *MmsServer) DeleteToken(ctx context.Context, in *mms.UUIDsReq) (*mms.BaseResp, error) {
	l := token.NewDeleteTokenLogic(ctx, s.svcCtx)
	return l.DeleteToken(in)
}

func (s *MmsServer) GetTokenList(ctx context.Context, in *mms.TokenListReq) (*mms.TokenListResp, error) {
	l := token.NewGetTokenListLogic(ctx, s.svcCtx)
	return l.GetTokenList(in)
}

func (s *MmsServer) GetTokenById(ctx context.Context, in *mms.UUIDReq) (*mms.TokenInfo, error) {
	l := token.NewGetTokenByIdLogic(ctx, s.svcCtx)
	return l.GetTokenById(in)
}

func (s *MmsServer) BlockUserAllToken(ctx context.Context, in *mms.UUIDReq) (*mms.BaseResp, error) {
	l := token.NewBlockUserAllTokenLogic(ctx, s.svcCtx)
	return l.BlockUserAllToken(in)
}

func (s *MmsServer) UpdateToken(ctx context.Context, in *mms.TokenInfo) (*mms.BaseResp, error) {
	l := token.NewUpdateTokenLogic(ctx, s.svcCtx)
	return l.UpdateToken(in)
}
