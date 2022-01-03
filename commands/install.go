package commands

import (
	// "errors"
	"flag"
	"fmt"
	"github.com/3xxx/engineercms/conf"
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	// "github.com/beego/beego/v2/server/web"
	// "github.com/beego/i18n"
	"os"
	// "time"
)

//系统安装.
func Install() {

	fmt.Println("Initializing...")

	err := orm.RunSyncdb("default", false, true)
	if err == nil {
		// initialization()
	} else {
		panic(err.Error())
	}
	fmt.Println("Install Successfully!")
	os.Exit(0)

}

func Version() {
	if len(os.Args) >= 2 && os.Args[1] == "version" {
		fmt.Println(conf.VERSION)
		os.Exit(0)
	}
}

//修改用户密码
func ModifyPassword() {
	var account, password string

	//账号和密码需要解析参数后才能获取
	if len(os.Args) >= 2 && os.Args[1] == "password" {
		flagSet := flag.NewFlagSet("MinDoc command: ", flag.ExitOnError)

		flagSet.StringVar(&account, "account", "", "用户账号.")
		flagSet.StringVar(&password, "password", "", "用户密码.")

		if err := flagSet.Parse(os.Args[2:]); err != nil {
			logs.Error("解析参数失败 -> ", err)
			os.Exit(1)
		}

		if len(os.Args) < 2 {
			fmt.Println("Parameter error.")
			os.Exit(1)
		}

		if account == "" {
			fmt.Println("Account cannot be empty.")
			os.Exit(1)
		}
		if password == "" {
			fmt.Println("Password cannot be empty.")
			os.Exit(1)
		}
		member, err := models.NewMember().FindByAccount(account)

		if err != nil {
			fmt.Println("Failed to change password:", err)
			os.Exit(1)
		}
		pwd, err := utils.PasswordHash(password)

		if err != nil {
			fmt.Println("Failed to change password:", err)
			os.Exit(1)
		}
		member.Password = pwd

		err = member.Update("password")
		if err != nil {
			fmt.Println("Failed to change password:", err)
			os.Exit(1)
		}
		fmt.Println("Successfully modified.")
		os.Exit(0)
	}

}

//初始化数据——放到main.go中了
// func initialization() {

// 	err := models.NewOption().Init()
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	lang, _ := web.AppConfig.String("default_lang")
// 	err = i18n.SetMessage(lang, "conf/lang/"+lang+".ini")
// 	if err != nil {
// 		panic(fmt.Errorf("initialize locale error: %s", err))
// 	}

// 	member, err := models.NewMember().FindByFieldFirst("account", "admin")
// 	if errors.Is(err, orm.ErrNoRows) {

// 		// create admin user
// 		logs.Info("creating admin user")
// 		member.Account = "admin"
// 		member.Avatar = conf.URLForWithCdnImage("/static/mindoc/images/headimgurl.jpg")
// 		member.Password = "123456"
// 		member.AuthMethod = "local"
// 		member.Role = conf.MemberSuperRole
// 		member.Email = "admin@iminho.me"

// 		if err := member.Add(); err != nil {
// 			panic("Member.Add => " + err.Error())
// 		}

// 		// create demo book
// 		logs.Info("creating demo book")
// 		book := models.NewBook()

// 		book.MemberId = member.MemberId
// 		book.BookName = i18n.Tr(lang, "init.default_proj_name") //"MinDoc演示项目"
// 		book.Status = 0
// 		book.ItemId = 1
// 		book.Description = i18n.Tr(lang, "init.default_proj_desc") //"这是一个MinDoc演示项目，该项目是由系统初始化时自动创建。"
// 		book.CommentCount = 0
// 		book.PrivatelyOwned = 0
// 		book.CommentStatus = "closed"
// 		book.Identify = "mindoc"
// 		book.DocCount = 0
// 		book.CommentCount = 0
// 		book.Version = time.Now().Unix()
// 		book.Cover = conf.GetDefaultCover()
// 		book.Editor = "markdown"
// 		book.Theme = "default"

// 		if err := book.Insert(lang); err != nil {
// 			panic("初始化项目失败 -> " + err.Error())
// 		}
// 	} else if err != nil {
// 		panic(fmt.Errorf("occur errors when initialize: %s", err))
// 	}

// 	if !models.NewItemsets().Exist(1) {
// 		item := models.NewItemsets()
// 		item.ItemName = i18n.Tr(lang, "init.default_proj_space") //"默认项目空间"
// 		item.MemberId = 1
// 		if err := item.Save(); err != nil {
// 			panic("初始化项目空间失败 -> " + err.Error())
// 		}
// 	}
// }
