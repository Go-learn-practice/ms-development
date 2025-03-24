package login_service_v1

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
	"strconv"
	"strings"
	common "test.com/devCommon"
	"test.com/devCommon/encrypts"
	"test.com/devCommon/errs"
	"test.com/devCommon/jwts"
	"test.com/devCommon/tms"
	"test.com/devGrpc/user/login"
	"test.com/devUser/config"
	"test.com/devUser/internal/dao"
	"test.com/devUser/internal/data/member"
	"test.com/devUser/internal/data/organization"
	"test.com/devUser/internal/database"
	"test.com/devUser/internal/database/tran"
	"test.com/devUser/internal/repo"
	"test.com/devUser/pkg/model"
	"time"
)

// LoginService 定义结构体
type LoginService struct {
	login.UnimplementedLoginServiceServer
	cache            repo.Cache
	memberRepo       repo.MemberRepo
	organizationRepo repo.OrganizationRepo
	transaction      tran.Transaction
}

// NewLoginService 构造函数 注册redis服务
func NewLoginService() *LoginService {
	return &LoginService{
		cache:            dao.RedisCacheInstance,
		memberRepo:       dao.NewMemberDao(),
		organizationRepo: dao.NewOrganizationDao(),
		transaction:      dao.NewTransaction(),
	}
}

// Captcha 获取验证码
func (ls *LoginService) Captcha(ctx context.Context, msg *login.CaptchaRequest) (*login.CaptchaResponse, error) {
	// 1. 获取参数
	mobile := msg.Mobile
	// 2. 校验参数
	if !common.VerifyMobile(mobile) {
		// 错误格式需要转换
		return nil, errs.GrpcError(model.NoLegalMobile)
	}
	// 3. 生成验证码（随机4位1000-9999）或者6位（100000-999999）
	code := "123456"
	// 4. 调用短信平台（放入go协程中执行）
	go func() {
		// 调用短信平台
		time.Sleep(2 * time.Second)
		// zap 日志库
		zap.L().Info("短信平台调用成功，发送短信")
		// redis 后续缓存可能存在 mysql，也可能在存在 mongo 中，也可能在 memCache 中
		// 5. 存储验证码 redis 当中 过期时间 15 分钟
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		// 添加 redis 缓存
		err := ls.cache.Put(c, model.RegisterRedisKey+mobile, code, 15*time.Minute)
		if err != nil {
			log.Printf("save mobile and code to redis fail: %v \n", err)
		}
	}()
	return &login.CaptchaResponse{
		Code: code,
	}, nil
}

// Register 注册
func (ls *LoginService) Register(ctx context.Context, msg *login.RegisterRequest) (*login.RegisterResponse, error) {
	c := context.Background()
	// 1. 校验参数
	// 2. 校验验证码（从redis获取）
	redisCode, err := ls.cache.Get(c, model.RegisterRedisKey+msg.Mobile)
	if errors.Is(err, redis.Nil) {
		// 验证码过期
		return nil, errs.GrpcError(model.CaptchaNotExist)
	}
	if err != nil {
		//log.Fatalf("get redis code fail: %v \n", err)
		zap.L().Error("Register redis get error", zap.Error(err))
		return nil, errs.GrpcError(model.RedisError)
	}
	if redisCode != msg.Captcha {
		return nil, errs.GrpcError(model.CaptchaError)
	}
	// 3. 校验业务逻辑（邮箱是否被注册 账号是否被注册 手机号是否被注册）
	exist, err := ls.memberRepo.GetMemberByEmail(c, msg.Email)
	if err != nil {
		zap.L().Error("Register db get error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.EmailExist)
	}
	exist, err = ls.memberRepo.GetMemberByAccount(c, msg.Name)
	if err != nil {
		zap.L().Error("Register db get error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.AccountExist)
	}
	exist, err = ls.memberRepo.GetMemberByMobile(c, msg.Mobile)
	if err != nil {
		zap.L().Error("Register db get error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if exist {
		return nil, errs.GrpcError(model.MobileExist)
	}
	// 4. 执行业务 将数据存入member表 生成一个数据 存入组织表 organization
	pwd := encrypts.Md5(msg.Password)
	// member表数据
	mem := &member.Member{
		Account:       msg.Name,
		Password:      pwd,
		Name:          msg.Name,
		Mobile:        msg.Mobile,
		Email:         msg.Email,
		CreateTime:    time.Now().UnixMilli(),
		LastLoginTime: time.Now().UnixMilli(),
		Status:        model.Normal,
	}
	// 执行事务操作
	err = ls.transaction.Action(func(conn database.DbConn) error {
		err = ls.memberRepo.SaveMember(c, conn, mem)
		if err != nil {
			zap.L().Error("Register db SaveMember error", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}
		// 存入组织
		org := &organization.Organization{
			Name:       mem.Name + "个人组织",
			MemberId:   mem.Id,
			CreateTime: time.Now().UnixMilli(),
			Personal:   model.Personal,
			Avatar:     "",
		}
		err = ls.organizationRepo.SaveOrganization(c, conn, org)
		if err != nil {
			zap.L().Error("Register db SaveOrganization error", zap.Error(err))
			return errs.GrpcError(model.DBError)
		}
		return nil
	})
	// 执行事务过程的错误
	if err != nil {
		return nil, err
	}
	// 5. 返回结果
	return &login.RegisterResponse{}, nil
}

// Login 登录
func (ls *LoginService) Login(ctx context.Context, msg *login.LoginRequest) (*login.LoginResponse, error) {
	c := context.Background()
	// 1. 数据库查询 账号密码 是否正确
	pwd := encrypts.Md5(msg.Password)
	mem, err := ls.memberRepo.FindMember(c, msg.Account, pwd)
	if err != nil {
		zap.L().Error("Login db FindMember error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if mem == nil {
		return nil, errs.GrpcError(model.AccountAndPwdError)
	}
	memMsg := &login.MemberMessage{}
	err = copier.Copy(memMsg, mem)
	memMsg.Code, _ = encrypts.EncryptInt64(mem.Id, model.AESKey)
	memMsg.LastLoginTime = tms.FormatByMilli(mem.LastLoginTime)
	memMsg.CreateTime = tms.FormatByMilli(mem.CreateTime)
	// 2. 根据用户id查 organization 信息
	orgs, err := ls.organizationRepo.FindOrganizationByMemId(c, mem.Id)
	if err != nil {
		zap.L().Error("Login db FindOrganizationByMemId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	var orgMessages []*login.OrganizationMessage
	err = copier.Copy(&orgMessages, orgs)
	for _, v := range orgMessages {
		v.Code, _ = encrypts.EncryptInt64(v.Id, model.AESKey)
		v.OwnerCode = memMsg.Code
		o := organization.ToMap(orgs)[v.Id]
		v.CreateTime = tms.FormatByMilli(o.CreateTime)
	}
	if len(orgs) > 0 {
		memMsg.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)
	}
	// 3. 用 jwt 生成token
	memIdStr := strconv.FormatInt(mem.Id, 10)
	// 7days
	exp := time.Duration(config.Conf.Jwt.AccessExp*24) * time.Hour
	// 14days
	rExp := time.Duration(config.Conf.Jwt.RefreshExp*24) * time.Hour
	token := jwts.CreateToken(memIdStr, exp, config.Conf.Jwt.AccessSecret, rExp, config.Conf.Jwt.RefreshSecret)
	tokenList := &login.TokenMessage{
		AccessToken:    token.AccessToken,
		AccessTokenExp: token.AccessExp,
		RefreshToken:   token.RefreshToken,
		TokenType:      "bearer",
	}
	// 4. 返回结果
	return &login.LoginResponse{
		Member:           memMsg,
		OrganizationList: orgMessages,
		TokenList:        tokenList,
	}, nil
}

func (ls *LoginService) TokenVerify(ctx context.Context, msg *login.LoginRequest) (*login.LoginResponse, error) {
	token := msg.Token
	if strings.Contains(token, "bearer") {
		token = strings.ReplaceAll(token, "bearer ", "")
	}
	parseToken, err := jwts.ParseToken(token, config.Conf.Jwt.AccessSecret)
	if err != nil {
		zap.L().Error("Login TokenVerify error", zap.Error(err))
		return nil, errs.GrpcError(model.NoLogin)
	}
	//TODO: 数据库查询 优化点 登录之后 应该把用户信息缓存起来
	id, _ := strconv.ParseInt(parseToken, 10, 64)
	memberById, err := ls.memberRepo.FindMemberById(context.Background(), id)
	if err != nil {
		zap.L().Error("TokenVerify db FindMemberById error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	memMsg := &login.MemberMessage{}
	_ = copier.Copy(memMsg, memberById)
	memMsg.Code, _ = encrypts.EncryptInt64(memberById.Id, model.AESKey)
	orgs, err := ls.organizationRepo.FindOrganizationByMemId(context.Background(), memberById.Id)
	if err != nil {
		zap.L().Error("TokenVerify db FindOrganizationByMemId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if len(orgs) > 0 {
		memMsg.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)
	}
	// 格式化时间戳
	memMsg.CreateTime = tms.FormatByMilli(memberById.CreateTime)
	return &login.LoginResponse{Member: memMsg}, nil
}

func (ls *LoginService) MyOrgList(ctx context.Context, msg *login.UserRequest) (*login.OrgListResponse, error) {
	memId := msg.MemId
	// TODO: 这里是否可以使用 MyOrgList 方法参数中的 ctx
	orgs, err := ls.organizationRepo.FindOrganizationByMemId(context.Background(), memId)
	if err != nil {
		zap.L().Error("MyOrgList db FindOrganizationByMemId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	var orgMessages []*login.OrganizationMessage
	err = copier.Copy(&orgMessages, orgs)
	for _, org := range orgMessages {
		org.Code, _ = encrypts.EncryptInt64(org.Id, model.AESKey)
	}
	return &login.OrgListResponse{OrganizationList: orgMessages}, nil
}

func (ls *LoginService) FindMemInfoById(ctx context.Context, msg *login.UserRequest) (*login.MemberMessage, error) {
	memberById, err := ls.memberRepo.FindMemberById(context.Background(), msg.MemId)
	if err != nil {
		zap.L().Error("FindMemInfoById db FindMemberById error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	memMsg := &login.MemberMessage{}
	_ = copier.Copy(memMsg, memberById)
	memMsg.Code, _ = encrypts.EncryptInt64(memberById.Id, model.AESKey)
	orgs, err := ls.organizationRepo.FindOrganizationByMemId(context.Background(), memberById.Id)
	if err != nil {
		zap.L().Error("FindMemInfoById db FindOrganizationByMemId error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if len(orgs) > 0 {
		memMsg.OrganizationCode, _ = encrypts.EncryptInt64(orgs[0].Id, model.AESKey)
	}
	memMsg.CreateTime = tms.FormatByMilli(memberById.CreateTime)
	return memMsg, nil
}
