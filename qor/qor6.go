package main

import (
	"github.com/qor/auth"
	"github.com/qor/auth/providers/password"
	"strings"

	//"github.com/qor/auth/providers/facebook"
	//"github.com/qor/auth/providers/twitter"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/admin"
	"github.com/qor/auth/claims"
	"github.com/qor/qor"
	"github.com/qor/roles"
	"github.com/qor/session/manager"
	"net/http"
	"time"
)

//##########################  定义数据库对象  ##########################
type AdminUser struct {
	gorm.Model
	UserName    string
	Password    string
	ConfirmedAt *time.Time
	DeletedAt   *time.Time
}

// ToClaims convert to auth Claims
func (admin AdminUser) ToClaims() *claims.Claims {
	claims := claims.Claims{}
	claims.Provider = "password"
	claims.Id = admin.UserName
	claims.UserID = admin.UserName
	return &claims
}
func (admin AdminUser) DisplayName() string {
	return admin.UserName
}

// //########################## 定义admin 权限 ##########################
type AdminAuth struct {
}

// 使用写死的用户名密码登陆。
var DefaultLoginHandler = func(context *auth.Context) (*claims.Claims, error) {
	var (
		req       = context.Request
		w         = context.Writer
		adminUser AdminUser
	)
	req.ParseForm()
	loginName := strings.TrimSpace(req.Form.Get("login"))
	password := strings.TrimSpace(req.Form.Get("password"))
	adminUser.UserName = loginName
	logs.Info("get DefaultAuthorizeHandler , loginName: %v , password : %v  ", loginName, password)

	if loginName == "admin" && password == "admin" {
		claims := adminUser.ToClaims()
		//登陆成功，注册session。
		manager.SessionManager.Add(w, req, "AdminUser", adminUser.UserName)
		logs.Info("##################### login success #####################")
		return claims, nil
	} else {
		return nil, auth.ErrInvalidPassword
	}
}

var DefaultRegisterHandler = func(context *auth.Context) (*claims.Claims, error) {
	return nil, auth.ErrInvalidAccount
}

func (AdminAuth) LoginURL(c *admin.Context) string {
	logs.Info(" user not login ")
	return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
	logs.Info(" user  logout ")
	return "/auth/logout"
}

//从session中获得当前用户。
func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	adminUserName := manager.SessionManager.Get(c.Request, "AdminUser")
	logs.Info("########## adminUser %v", adminUserName)
	if adminUserName != "" {
		adminUser := AdminUser{}
		adminUser.UserName = adminUserName
		return &adminUser
	}
	return nil
}

//声明全局变量
var (
	// 注册数据库，可以是sqlite3 也可以是 mysql 换下驱动就可以了。
	DB, _ = gorm.Open("sqlite3", "demo.db")

	// 初始化admin 还有其他的，比如API
	Auth  = auth.New(&auth.Config{DB: DB, UserModel: AdminUser{}})
	Admin = admin.New(&admin.AdminConfig{DB: DB, Auth: AdminAuth{}})
)

//noinspection GoInvalidCompositeLiteral
func init() {

	DB.AutoMigrate(AdminUser{}) //自动创建表。
	// Register Auth providers
	// Allow use username/password
	Auth.RegisterProvider(password.New(&password.Config{
		AuthorizeHandler: DefaultLoginHandler,
		ResetPasswordHandler: func(context *auth.Context) error {
			return nil
		},
		RecoverPasswordHandler: func(context *auth.Context) error {
			return nil
		},
		RegisterHandler: func(context *auth.Context) (*claims.Claims, error) {
			return nil, nil
		},
	}))

	// 创建admin后台对象资源。
	Admin.AddResource(&AdminUser{})

	roles.Register("admin", func(req *http.Request, currentUser interface{}) bool {
		return currentUser != nil && currentUser.(*AdminUser).UserName == "admin"
	})
}

func main() {
	// 启动服务
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	mux.Handle("/auth/", Auth.NewServeMux())
	http.ListenAndServe(":9000", manager.SessionManager.Middleware(mux))
	//访问 http://localhost:9000/auth/register 注册账号
	// admin ： http://localhost:9000/admin/auth_identities
}
