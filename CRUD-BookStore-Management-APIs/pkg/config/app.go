package config
import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// the connect function helps us to connect with our database
func Connect(){
	d,err := gorm.Open("mysql","username:password/databaseName?charset=utf8parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d

}
// to just return the database
func GetDB() *gorm.DB{
	return db
}