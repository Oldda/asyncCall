package db

import(
	"asyncCall/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

/*
数据库连接池的原理
只用sql.Open函数创建连接池，可是此时只是初始化了连接池，并没有创建任何连接。
连接创建都是惰性的，只有当你真正使用到连接的时候，连接池才会创建连接。连接池很重要，它直接影响着你的程序行为。
连接池的工作原来却相当简单。当你的函数(例如Exec，Query)调用需要访问底层数据库的时候，函数首先会向连接池请求一个连接。
如果连接池有空闲的连接，则返回给函数。否则连接池将会创建一个新的连接给函数。
一旦连接给了函数，连接则归属于函数。函数执行完毕后，要不把连接所属权归还给连接池，要么传递给下一个需要连接的（Rows）对象，
最后使用完连接的对象也会把连接释放回到连接池。
总计：数据库链接在sql.Open调用的时候初始化。需要使用到链接的时候惰性分配闲置链接或者新生成一个链接。
故而，每个数据库一般只调用一次sql.Open函数来生成一个连接池。每次需要操作数据库的时候就去调用open返回的对象指针。
func Open(driverName, dataSourceName string) (*DB, error)

配置连接池有两个的方法：
db.SetMaxOpenConns(n int) 设置打开数据库的最大连接数。包含正在使用的连接和连接池的连接。如果你的函数调用需要申请一个连接，并且连接池已经没有了连接或者连接数达到了最大连接数。此时的函数调用将会被block，直到有可用的连接才会返回。设置这个值可以避免并发太高导致连接mysql出现too many connections的错误。该函数的默认设置是0，表示无限制。
db.SetMaxIdleConns(n int) 设置连接池中的保持连接的最大连接数。默认也是0，表示连接池不会保持释放会连接池中的连接的连接状态：即当连接释放回到连接池的时候，连接将会被关闭。这会导致连接再连接池中频繁的关闭和创建。
*/

var MysqlEngine *gorm.DB

func init(){
	var err error
	config := config.NewConfig("./config","dev","json")
	db_user := config.GetString("database.db_user")
	db_host := config.GetString("database.db_host")
	db_port := config.GetString("database.db_port")
	db_password := config.GetString("database.db_password")
	db_name := config.GetString("database.db_name")

	dns := db_user + ":"+ db_password + "@tcp("+ db_host +":" + db_port + ")/" + db_name + "?charset=utf8&parseTime=True&loc=Local"
	MysqlEngine, err = gorm.Open(mysql.New(mysql.Config{
		DSN:dns,
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	 	DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	  	DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	  	SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}),&gorm.Config{})
  	if err != nil{
  		log.Println(err)
  	}
}