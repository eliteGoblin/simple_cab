package db



import (
	"fmt"
	"log"
	"simple_cab/config"
	"io"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var MyDB *gorm.DB

func SetLoggerFile(out io.Writer) {
	if out != nil {
		MyDB.SetLogger(log.New(out, "	\r\n", 0))
	}
}
func InitConn(dbName string) {
	var err error
	mysqlCnf := config.GetInstance().MySQL
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlCnf.User,
		mysqlCnf.Password,
		mysqlCnf.Host,
		mysqlCnf.Port,
		dbName,
	)

	MyDB, err = gorm.Open("mysql", dsn)
	// 连接池
	if err == nil {
		MyDB.DB().SetMaxIdleConns(mysqlCnf.MaxIdleConns)
		MyDB.DB().SetMaxOpenConns(mysqlCnf.MaxOpenConns)
		MyDB.DB().Ping()
		MyDB.LogMode(mysqlCnf.EnableLog)
	} else {
		MyDB = nil
		log.Panic("Gorm Open Error", err)
	}
}

