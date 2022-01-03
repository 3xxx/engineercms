package models

import (
	"fmt"
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/jinzhu/gorm"
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/gorm"
	"time"
)

// 打赏文章表
type Pay struct {
	// gorm.Model
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UserID    int64   `gorm:"column:user_id;"` // One-To-One (属于 - 本表的BillingAddressID作外键
	User2ID   int64   `gorm:"column:user2_id;"`
	ArticleID int64   `gorm:"column:article_id;"`
	Amount    int     `gorm:"column:amount"`
	User      User    `gorm:"foreignkey:Id;references:UserID;"` //这个外键难道不是错的么？应该是UserID?没错，因为column:user_id
	Article   Article `gorm:"foreignkey:Id;references:ArticleID;"`
}

// 用户付费计算表
type PayMath struct {
	gorm.Model
	// ID         uint `json:"id" gorm:"primary_key"`
	// CreatedAt  time.Time
	// UpdatedAt  time.Time
	// DeletedAt  *time.Time
	UserID       int64 `gorm:"column:user_id;"` // One-To-One (属于 - 本表的BillingAddressID作外键
	User2ID      int64 `gorm:"column:user2_id;"`
	UserTempleID uint  `gorm:"column:user_temple_id;"`
	Amount       int   `gorm:"column:amount"`
	User         User  //`gorm:"foreignkey:UserId"` //对于preload来说，这个后面不影响
	User2        User  //`gorm:"foreignkey:User2Id"`
	UserTemple   UserTemple
}

// 用户付费下载pdf文件表
type PayMathPdf struct {
	gorm.Model
	// ID         uint `json:"id" gorm:"primary_key"`
	// CreatedAt  time.Time
	// UpdatedAt  time.Time
	// DeletedAt  *time.Time
	UserID int64 `gorm:"column:user_id;"` // One-To-One (属于 - 本表的BillingAddressID作外键
	// User2ID int64 `gorm:"column:user2_id;foreignkey:User2Id;"`
	// UserTempleID uint  `gorm:"column:user_temple_id;foreignkey:UserTempleID;"`
	Amount int  `gorm:"column:amount"`
	User   User `gorm:"foreignkey:Id;references:UserID;"` //对于preload来说，这个后面不影响
	// User2    User //`gorm:"foreignkey:User2Id"`
	PdfTitle string
	PdfLink  string
	// UserTemple   UserTemple
}

// 用户付费计算表
type PayExcel struct {
	gorm.Model
	// ID         uint `json:"id" gorm:"primary_key"`
	// CreatedAt  time.Time
	// UpdatedAt  time.Time
	// DeletedAt  *time.Time
	UserID        int64 `gorm:"column:user_id;"` // One-To-One (属于 - 本表的BillingAddressID作外键
	User2ID       int64 `gorm:"column:user2_id;"`
	ExcelTempleID uint  `gorm:"column:excel_temple_id;"`
	Amount        int   `gorm:"column:amount"`
	User          User  `gorm:"foreignkey:Id;references:UserID;"` //对于preload来说，这个后面不影响
	User2         User  //`gorm:"foreignkey:User2Id"`
	ExcelTemple   ExcelTemple
}

// 用户付费下载pdf文件表
type PayExcelPdf struct {
	gorm.Model
	// ID         uint `json:"id" gorm:"primary_key"`
	// CreatedAt  time.Time
	// UpdatedAt  time.Time
	// DeletedAt  *time.Time
	UserID int64 `gorm:"column:user_id;"` // One-To-One (属于 - 本表的BillingAddressID作外键
	// User2ID int64 `gorm:"column:user2_id;foreignkey:User2Id;"`
	// UserTempleID uint  `gorm:"column:user_temple_id;foreignkey:UserTempleID;"`
	Amount int  `gorm:"column:amount"`
	User   User `gorm:"foreignkey:Id;references:UserID;"` //对于preload来说，这个后面不影响
	// User2    User //`gorm:"foreignkey:User2Id"`
	PdfTitle string
	PdfLink  string
	// UserTemple   UserTemple
}

// 余额表
type Money struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID int64 `gorm:"column:user_id;"` // 外键 (属于), tag `index`是为该列创建索引
	Amount int   `json:"amount" gorm:"column:amount"`
	User   User  `gorm:"foreignkey:Id;references:UserID;"`
}

// 用户充值记录
type Recharge struct {
	gorm.Model
	// ID       int    `gorm:"primary_key"`
	UserID int64 `gorm:"column:user_id;"` // 外键 (属于), tag `index`是为该列创建索引
	Amount int   `gorm:"column:amount"`
	User   User  `gorm:"foreignkey:Id;references:UserID;"`
}

/*20201008这些都应该放到models文件夹里第一个文件里，这样，程序初始化的时候先定义这些全局变量……*/
//定义全局的db对象，我们执行数据库操作主要通过他实现。
// var _db *gorm.DB

func init() {
	// var err error
	// var dns string
	// db_type := beego.AppConfig.String("db_type")
	// db_name := beego.AppConfig.String("db_name")
	// db_path := beego.AppConfig.String("db_path")
	// if db_path == "" {
	// 	db_path = "./"
	// }
	// dns = fmt.Sprintf("%s%s.db", db_path, db_name)
	// _db, err = gorm.Open(db_type, dns)

	// defer _db.Close()//20200803这个不能打开。
	// _db.LogMode(true)

	// if err != nil {
	// 	panic("连接数据库失败, error=" + err.Error())
	// }

	// defer gdb.Close()
	//禁止表名复数形式

	// _db.SingularTable(true)

	// 开发的时候需要打开调试日志
	// _db.LogMode(true)
	//设置数据库连接池参数

	// _db.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
	// _db.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	// _db.AutoMigrate(&Pay{}, &Money{}, &Recharge{}, &PayMath{}, &PayMathPdf{}, &PayExcel{}, &PayExcelPdf{})
	// _db.CreateTable(&Pay{}, &Money{}, &Recharge{}) //当第一个存在时，后面的不再建立，bug！！！
	// _db.CreateTable(&Money{})
	// _db.CreateTable(&Recharge{})
	// _db.CreateTable(&PayMath{})
	// _db.CreateTable(&PayMathPdf{})
	// _db.CreateTable(&PayExcel{})
	// _db.CreateTable(&PayExcelPdf{})
	// if !gdb.HasTable(&Pay1{}) {
	// 	if err = gdb.CreateTable(&Pay1{}).Error; err != nil {
	// 		panic(err)
	// 	}
	// }
}

//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools._db//GetDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，
// db对象在调用他的方法的时候会从数据库连接池中获取新的连接
// 注意：使用连接池技术后，千万不要使用完db后调用db.Close关闭数据库连接，
// 这样会导致整个数据库连接池关闭，导致连接池没有可用的连接
// func _db//GetDB() *gorm.DB {
// 	return _db
// }

//查询某个打赏记录
func GetPay(id int64) (pay Pay, err error) {
	//获取DB
	db := _db //GetDB()
	db.Find(&pay, 1)
	db.Model(&pay).Find(&pay)
	// err = db.Model(&pay).Related(&pay.User, "User").Error
	err = db.Model(&pay).Association("User").Error
	return pay, err
	// db.Find(&pays, 1)
	// var user []User
	// db.Model(&pays).Related(&User, "Users")
	//db.Where("id=?", id).First(&topic)
	// 关联的关键代码
	//db.Model(&topic).Related(&topic.Category, "CategoryId")
	//result := db.Model(&topic).Related(&topic.User, "UserId")
}

//查询某个用户账户余额
func GetUserMoney(uid int64) (money Money, err error) {
	//获取DB
	db := _db //GetDB()
	err = db.Where("user_id = ?", uid).First(&money).Error
	return money, err
}

//添加某个文章某个用户打赏记录
func AddUserPay(articleid, uid int64, amount int) error {
	//获取DB
	db := _db //GetDB()
	// 注意，当你在一个事务中应使用 tx 作为数据库句柄
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	//保证文章id正确
	var article Article
	err := db.Where("id = ?", articleid).First(&article).Error
	if err != nil {
		logs.Info(err)
		tx.Rollback()
		return err
	}
	//根据article prodid查出作者userid
	var product Product
	err = db.Where("id = ?", article.ProductId).First(&product).Error
	if err != nil {
		logs.Info(err)
		tx.Rollback()
		return err
	}
	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	// db.Create(&user)
	newamount := 0 - amount
	if err = tx.Create(&Pay{UserID: uid, User2ID: product.Uid, ArticleID: articleid, Amount: newamount}).Error; err != nil {
		logs.Info(err)
		tx.Rollback()
		return err
	}

	//给作者账户增加赞赏
	// if err := tx.Create(&Pay{UserID: product.Uid, ArticleID: articleid, Amount: amount}).Error; err != nil {
	// 	logs.Error(err)
	// 	tx.Rollback()
	// 	return err
	// }

	//更新用户账户余额money
	//1.首先保证账户存在
	// money := Money{UserID: uid}

	//赞赏者账户
	var moneyA Money
	err = db.Where("user_id = ?", uid).First(&moneyA).Error
	if err != nil {
		logs.Info(err)
		tx.Rollback()
		return err
	}
	//2.赞赏者账户余额保证大于0
	newamount = moneyA.Amount - amount
	if newamount < 0 {
		tx.Rollback()
		return err
		logs.Info(err)
	}
	//3.赞赏者账户修改余额
	rowsAffected := tx.Model(&moneyA).Update("amount", newamount).RowsAffected
	if rowsAffected == 0 {
		logs.Info(err)
		tx.Rollback()
		return err
	}

	//作者账户
	var moneyB Money
	//没有查到则新增一条
	err = db.Where("user_id = ?", product.Uid).First(&moneyB).Error
	if err != nil {
		logs.Error(err)
		if err = tx.Create(&Money{UserID: product.Uid, Amount: 100}).Error; err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
		//事务新增后要查出来moneyB，供下面的update用
		//记得下面这个查询用tx，不能用db，因为数据还没真正写入！！！
		err = tx.Where("user_id = ?", product.Uid).First(&moneyB).Error
		if err != nil {
			tx.Rollback()
			logs.Error(err)
			return err
		}
		//账户充值——这句重复？上面是账户余额表，这个是充值表
		err := tx.Create(&Recharge{UserID: product.Uid, Amount: 100}).Error
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
	}
	// err = tx.FirstOrCreate(&moneyB, Money{UserID: product.Uid}).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// 	logs.Error(err)
	// }

	newamount = moneyB.Amount + amount
	rowsAffected = tx.Model(&moneyB).Update("amount", newamount).RowsAffected
	if rowsAffected == 0 {
		logs.Error(err)
		tx.Rollback()
		return err
	}
	//5.作者账户修改余额
	// newamount = moneyB.Amount + amount
	// rowsAffected = tx.Model(&moneyB).Where(query, args).Update("amount", newamount).RowsAffected
	// if rowsAffected == 0 {
	// 	tx.Rollback()
	// 	return err
	// 	logs.Error(err)
	// }
	//6.不能自己给自己赞赏

	return tx.Commit().Error
	// 更新单个属性，如果它有变化
	// db.Model(&user).Update("name", "hello")
	// if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
	//    tx.Rollback()
	//    return err
	// }

	// 开启事务
	// tx := db.Begin()
	//在事务中执行数据库操作，使用的是tx变量，不是db。
	//库存减一
	//等价于: UPDATE `foods` SET `stock` = stock - 1  WHERE `foods`.`id` = '2' and stock > 0
	//RowsAffected用于返回sql执行后影响的行数
	// rowsAffected := tx.Model(&food).Where("stock > 0").Update("stock", gorm.Expr("stock - 1")).RowsAffected
	// if rowsAffected == 0 {
	//如果更新库存操作，返回影响行数为0，说明没有库存了，结束下单流程
	//这里回滚作用不大，因为前面没成功执行什么数据库更新操作，也没什么数据需要回滚。
	//这里就是举个例子，事务中可以执行多个sql语句，错误了可以回滚事务
	// tx.Rollback()
	// return
	//}
	// err := tx.Create(保存订单).Error

	//保存订单失败，则回滚事务
	// if err != nil {
	//     tx.Rollback()
	// } else {
	//     tx.Commit()
	// }
}

//添加某个math某个用户计算扣费记录
func AddUserPayMath(templeid uint, uid int64, amount int) error {
	//获取DB
	db := _db //GetDB()
	// 注意，当你在一个事务中应使用 tx 作为数据库句柄
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	//保证模板id正确
	var usertemple UserTemple
	err := db.Where("id = ?", templeid).First(&usertemple).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}
	//根据模板id查出作者userid
	// var product Product
	// err = db.Where("id = ?", article.ProductId).First(&product).Error
	// if err != nil {
	// 	logs.Error(err)
	// 	tx.Rollback()
	// 	return err
	// }

	newamount := 0 - amount
	if err = tx.Create(&PayMath{UserID: uid, User2ID: usertemple.UserID, UserTempleID: templeid, Amount: newamount}).Error; err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	//赞赏者账户
	var moneyA Money
	err = db.Where("user_id = ?", uid).First(&moneyA).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}
	//2.赞赏者账户余额保证大于0
	newamount = moneyA.Amount - amount
	if newamount < 0 {
		tx.Rollback()
		return err
		logs.Error(err)
	}
	//3.赞赏者账户修改余额
	rowsAffected := tx.Model(&moneyA).Update("amount", newamount).RowsAffected
	if rowsAffected == 0 {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	// 模板作者账户充值记录
	var moneyB Money
	//没有查到则新增一条
	err = db.Where("user_id = ?", usertemple.UserID).First(&moneyB).Error
	if err != nil {
		logs.Error(err)
		if err = tx.Create(&Money{UserID: usertemple.UserID, Amount: 100}).Error; err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
		//事务新增后要查出来moneyB，供下面的update用
		// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
		// result := db.Create(&user) // 通过数据的指针来创建
		// tx也能返回id啊result :=  result.Error
		// user.ID             // 返回插入数据的主键
		// result.Error        // 返回 error
		// result.RowsAffected // 返回插入记录的条数

		//记得下面这个查询用tx，不能用db，因为数据还没真正写入！！！
		// 用money:=Money{}来创建，然后取到money.ID供后续使用
		err = tx.Where("user_id = ?", usertemple.UserID).First(&moneyB).Error
		if err != nil {
			tx.Rollback()
			logs.Error(err)
			return err
		}
		//账户充值_这句干嘛？重复？——这个是充值数据表，上面是账户余额表，不重复
		err := tx.Create(&Recharge{UserID: usertemple.UserID, Amount: 100}).Error
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
	}

	newamount = moneyB.Amount + amount
	rowsAffected = tx.Model(&moneyB).Update("amount", newamount).RowsAffected
	if rowsAffected == 0 {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//添加某个mathpdf某个用户计算扣费记录,归平台所有
func AddUserPayMathPdf(pdftitle, pdflink string, uid int64, amount int) error {
	//获取DB
	db := _db //GetDB()
	// 注意，当你在一个事务中应使用 tx 作为数据库句柄
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	//保证pdf正确_目录下有这个文件存在，controllers里判断吧

	newamount := 0 - amount
	if err := tx.Create(&PayMathPdf{UserID: uid, PdfTitle: pdftitle, PdfLink: pdflink, Amount: newamount}).Error; err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	//赞赏者账户
	var moneyA Money
	err := db.Where("user_id = ?", uid).First(&moneyA).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}
	//2.赞赏者账户余额保证大于0
	newamount = moneyA.Amount - amount
	if newamount < 0 {
		tx.Rollback()
		return err
		logs.Error(err)
	}
	//3.赞赏者账户修改余额
	rowsAffected := tx.Model(&moneyA).Update("amount", newamount).RowsAffected
	if rowsAffected == 0 {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//查询某个用户下载过的pdf记录
func GetUserPayMathPdf(uid int64, pdftitle string) (paymathpdf PayMathPdf, err error) {
	//获取DB
	db := _db //GetDB()
	err = db.Where("user_id = ? AND pdf_title=?", uid, pdftitle).First(&paymathpdf).Error
	return paymathpdf, err
}

//添加某个math某个用户计算扣费记录
func AddUserPayExcel(templeid uint, uid int64, amount int) error {
	//获取DB
	db := _db //GetDB()
	// 注意，当你在一个事务中应使用 tx 作为数据库句柄
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	//保证模板id正确
	var exceltemple ExcelTemple
	err := db.Where("id = ?", templeid).First(&exceltemple).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}
	//根据模板id查出作者userid
	// var product Product
	// err = db.Where("id = ?", article.ProductId).First(&product).Error
	// if err != nil {
	// 	logs.Error(err)
	// 	tx.Rollback()
	// 	return err
	// }

	newamount := 0 - amount
	if err = tx.Create(&PayExcel{UserID: uid, User2ID: exceltemple.UserID, ExcelTempleID: templeid, Amount: newamount}).Error; err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	//赞赏者账户
	var moneyA Money
	err = db.Where("user_id = ?", uid).First(&moneyA).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}
	//2.赞赏者账户余额保证大于0
	newamount = moneyA.Amount - amount
	if newamount < 0 {
		tx.Rollback()
		return err
		logs.Error(err)
	}
	//3.赞赏者账户修改余额
	rowsAffected := tx.Model(&moneyA).Update("amount", newamount).RowsAffected
	if rowsAffected == 0 {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	// 模板作者账户充值记录
	var moneyB Money
	//没有查到则新增一条
	err = db.Where("user_id = ?", exceltemple.UserID).First(&moneyB).Error
	if err != nil {
		logs.Error(err)
		if err = tx.Create(&Money{UserID: exceltemple.UserID, Amount: 100}).Error; err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
		//事务新增后要查出来moneyB，供下面的update用
		// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
		// result := db.Create(&user) // 通过数据的指针来创建
		// tx也能返回id啊result :=  result.Error
		// user.ID             // 返回插入数据的主键
		// result.Error        // 返回 error
		// result.RowsAffected // 返回插入记录的条数

		//记得下面这个查询用tx，不能用db，因为数据还没真正写入！！！
		// 用money:=Money{}来创建，然后取到money.ID供后续使用
		err = tx.Where("user_id = ?", exceltemple.UserID).First(&moneyB).Error
		if err != nil {
			tx.Rollback()
			logs.Error(err)
			return err
		}
		//账户充值_这句干嘛？重复？——这个是充值数据表，上面是账户余额表，不重复
		err := tx.Create(&Recharge{UserID: exceltemple.UserID, Amount: 100}).Error
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
	}

	newamount = moneyB.Amount + amount
	rowsAffected = tx.Model(&moneyB).Update("amount", newamount).RowsAffected
	if rowsAffected == 0 {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//添加某个mathpdf某个用户计算扣费记录,归平台所有
func AddUserPayExcelPdf(pdftitle, pdflink string, uid int64, amount int) error {
	//获取DB
	db := _db //GetDB()
	// 注意，当你在一个事务中应使用 tx 作为数据库句柄
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	//保证pdf正确_目录下有这个文件存在，controllers里判断吧

	newamount := 0 - amount
	if err := tx.Create(&PayExcelPdf{UserID: uid, PdfTitle: pdftitle, PdfLink: pdflink, Amount: newamount}).Error; err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	//赞赏者账户
	var moneyA Money
	err := db.Where("user_id = ?", uid).First(&moneyA).Error
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}
	//2.赞赏者账户余额保证大于0
	newamount = moneyA.Amount - amount
	if newamount < 0 {
		tx.Rollback()
		return err
		logs.Error(err)
	}
	//3.赞赏者账户修改余额
	rowsAffected := tx.Model(&moneyA).Update("amount", newamount).RowsAffected
	if rowsAffected == 0 {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//查询某个用户下载过的pdf记录
func GetUserPayExcelPdf(uid int64, pdftitle string) (payexcelpdf PayExcelPdf, err error) {
	//获取DB
	db := _db //GetDB()
	err = db.Where("user_id = ? AND pdf_title=?", uid, pdftitle).First(&payexcelpdf).Error
	return payexcelpdf, err
}

//用户充值和提现
func AddUserRecharge(uid int64, amount int) error {
	//获取DB
	db := _db //GetDB()
	// 注意，当你在一个事务中应使用 tx 作为数据库句柄
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	//账户添加一条充值记录
	// var recharge Recharge
	//没有查到则新增一条
	// err := tx.FirstOrCreate(&recharge, Recharge{UserID: uid}).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// rowsAffected := tx.Model(&recharge).Update("amount", amount).RowsAffected
	// if rowsAffected == 0 {
	// 	tx.Rollback()
	// 	return err
	// }
	if err := tx.Create(&Recharge{UserID: uid, Amount: amount}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//账户修改余额
	var money Money
	//没有查到则新增一条
	err := tx.FirstOrCreate(&money, Money{UserID: uid}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	newamount := money.Amount + amount
	rowsAffected := tx.Model(&money).Update("amount", newamount).RowsAffected
	if rowsAffected == 0 {
		tx.Rollback()
		return err
	}
	// if err := tx.Create(&Money{UserID: uid, Amount: newamount}).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	return tx.Commit().Error
	// 更新单个属性，如果它有变化
	// db.Model(&user).Update("name", "hello")
	// if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
	//    tx.Rollback()
	//    return err
	// }

	// 开启事务
	// tx := db.Begin()
	//在事务中执行数据库操作，使用的是tx变量，不是db。
	//库存减一
	//等价于: UPDATE `foods` SET `stock` = stock - 1  WHERE `foods`.`id` = '2' and stock > 0
	//RowsAffected用于返回sql执行后影响的行数
	// rowsAffected := tx.Model(&food).Where("stock > 0").Update("stock", gorm.Expr("stock - 1")).RowsAffected
	// if rowsAffected == 0 {
	//如果更新库存操作，返回影响行数为0，说明没有库存了，结束下单流程
	//这里回滚作用不大，因为前面没成功执行什么数据库更新操作，也没什么数据需要回滚。
	//这里就是举个例子，事务中可以执行多个sql语句，错误了可以回滚事务
	// tx.Rollback()
	// return
	//}
	// err := tx.Create(保存订单).Error

	//保存订单失败，则回滚事务
	// if err != nil {
	//     tx.Rollback()
	// } else {
	//     tx.Commit()
	// }
}

// 用户添加充值申请——软删除db.Delete(&User{}, 10)，采用事务，如果软删除出错，则回滚
func AddApplyRecharge(uid int64, amount int) error {
	//获取DB
	db := _db //GetDB()
	// 注意，当你在一个事务中应使用 tx 作为数据库句柄
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	//账户添加一条充值记录
	recharge := Recharge{UserID: uid, Amount: amount}
	if err := tx.Create(&recharge).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 进行软删除
	// db.Delete(&User{}, 10)
	// db.Delete(&user)
	if err := tx.Delete(&recharge).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 管理员查询到所有申请
func GetApplyRecharge(limit, offset int) (recharge []Recharge, err error) {
	//获取DB
	db := _db //GetDB()
	if err := db.Unscoped().Preload("User").Where("deleted_at <> ?", "").Limit(limit).Offset(offset).Find(&recharge).Error; err != nil {
		return recharge, err
	}
	return recharge, err
}

// 管理员查询软删除的充值申请，同意则取消软删除，修改申请数额db.Unscoped().Where("age = 20").Find(&users)
// db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
func UpdateRecharge(rid uint, amount int) error {
	//获取DB
	db := _db //GetDB()
	// 注意，当你在一个事务中应使用 tx 作为数据库句柄
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 查询申请人id和数额
	var recharge Recharge
	if err := tx.Unscoped().Where("id = ?", rid).Find(&recharge).Error; err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(&recharge)
	// db.Model(&user).Update("name", "hello")
	if err := tx.Unscoped().Model(&recharge).Update("deleted_at", nil).Error; err != nil {
		tx.Rollback()
		return err
	}

	//账户修改余额
	var money Money
	//没有查到则新增一条
	err := tx.FirstOrCreate(&money, Money{UserID: recharge.UserID}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 修改余额记录
	newamount := money.Amount + amount
	rowsAffected := tx.Model(&money).Update("amount", newamount).RowsAffected
	if rowsAffected == 0 {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//查询某个用户花费赞赏 和 获得 文章 赞赏记录
func GetUserPay(uid int64, limit, offset int) (pays []*Pay, err error) {
	//获取DB
	db := _db //GetDB()
	// err = db.Where("user_id", uid).Find(&pays).Error
	err = db.Order("updated_at desc").Model(&pays).Preload("User").Preload("Article").Where("user_id = ?", uid).Or("user2_id = ?", uid).Limit(limit).Offset(offset).Find(&pays).Error //查询所有device记录
	// err = db.Model(&pays).Related(&pays.User, "Users").Error
	return pays, err
	// 多连接及参数
	// db.Joins("JOIN pays ON pays.user_id = users.id", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("user_id = ?", uid).Find(&pays)
}

//查询某个用户花费赞赏 和 获得 math 赞赏记录
func GetUserPayMath(uid int64, limit, offset int) (paymaths []*PayMath, err error) {
	//获取DB
	db := _db //GetDB()
	// err = db.Where("user_id", uid).Find(&pays).Error
	err = db.Order("updated_at desc").Preload("User").Preload("User2").Preload("UserTemple").Where("user_id = ?", uid).Or("user2_id = ?", uid).Limit(limit).Offset(offset).Find(&paymaths).Error //查询所有device记录
	// err = db.Model(&pays).Related(&pays.User, "Users").Error
	return paymaths, err
	// 多连接及参数
	// db.Joins("JOIN pays ON pays.user_id = users.id", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("user_id = ?", uid).Find(&pays)
}

//查询某个用户 花费 赞赏（赞赏别人）文章 的记录
func GetUserPayAppreciation(uid int64, limit, offset int) (pays []*Pay, err error) {
	//获取DB
	db := _db                                                                                                                                               //GetDB()
	err = db.Order("updated_at desc").Model(&pays).Preload("User").Preload("Article").Where("user_id=?", uid).Limit(limit).Offset(offset).Find(&pays).Error //查询所有device记录
	return pays, err
	// 多连接及参数
	// db.Joins("JOIN pays ON pays.user_id = users.id", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("user_id = ?", uid).Find(&pays)
}

//查询某个用户 花费 赞赏（赞赏别人）math 的记录
func GetUserPayMathAppreciation(uid int64, limit, offset int) (paymaths []*PayMath, err error) {
	//获取DB
	db := _db                                                                                                                                                          //GetDB()
	err = db.Order("updated_at desc").Model(&paymaths).Preload("User").Preload("UserTemple").Where("user_id=?", uid).Limit(limit).Offset(offset).Find(&paymaths).Error //查询所有device记录
	return paymaths, err
	// 多连接及参数
	// db.Joins("JOIN pays ON pays.user_id = users.id", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("user_id = ?", uid).Find(&pays)
}

//查询某个用户 获得 赞赏（别人赞赏自己）的记录
func GetUserGetAppreciation(uid int64, limit, offset int) (pays []*Pay, err error) {
	//获取DB
	db := _db                                                                                                                                                //GetDB()
	err = db.Order("updated_at desc").Model(&pays).Preload("User").Preload("Article").Where("user2_id=?", uid).Limit(limit).Offset(offset).Find(&pays).Error //查询所有device记录
	return pays, err
	// 多连接及参数
	// db.Joins("JOIN pays ON pays.user_id = users.id", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("user_id = ?", uid).Find(&pays)
}

//查询某个用户 获得 赞赏（别人赞赏自己）math 的记录
func GetUserGetMathAppreciation(uid int64, limit, offset int) (paymaths []*PayMath, err error) {
	//获取DB
	db := _db                                                                                                                                                           //GetDB()
	err = db.Order("updated_at desc").Model(&paymaths).Preload("User").Preload("UserTemple").Where("user2_id=?", uid).Limit(limit).Offset(offset).Find(&paymaths).Error //查询所有device记录
	return paymaths, err
	// 多连接及参数
	// db.Joins("JOIN pays ON pays.user_id = users.id", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("user_id = ?", uid).Find(&pays)
}

// url := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "localhost:3306", "demo")
// db, err := gorm.Open("mysql", url)
// db.LogMode(true)
// if err != nil {
// 	panic("failed to connect database")
// }
