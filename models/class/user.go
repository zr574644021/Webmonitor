package class

import (
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
)

type User struct {
	UserName string `orm:"pk"` //用户名
	Password string            //密码
}

type LoginRecord struct {
	Id        int
	LoginTime int64                  //登陆时间
	LoginIP   string                 //登陆IP
	Password  string `orm:"null"`    //使用的错误的密码(正确时为空)
	User      *User  `orm:"rel(fk)"` //User表外键
}

func (c *User) Login_matching() bool { //明文形式密码比对
	o := orm.NewOrm()
	var user User
	if err := o.QueryTable("user").Filter("user_name", c.UserName).One(&user); err != nil {
		beego.Error("username error is ",err)
		return false
	} else {
		cryptograph := Encrypt(c.Password)
		if cryptograph == user.Password {
		} else {
			return false
		}
	}
	return true
}

func (c *User) Login_matching_crypte() bool { //密文形式密码比对
	o := orm.NewOrm()
	var user User
	if err := o.QueryTable("user").Filter("user_name", c.UserName).One(&user); err != nil {
		beego.Error("username error")
		return false
	} else {
		if c.Password == user.Password {
		}
	}
	return true
}

func (c *User) Login_record(password, loginIp string) bool {
	var record LoginRecord
	record.LoginTime = time.Now().Unix()
	record.LoginIP = loginIp
	record.Password = password
	record.User = c
	if _, err := orm.NewOrm().Insert(&record); err != nil {
		return false
	}
	return true
}

func Encrypt(password string) (cryptograph string) {
	salt1 := "fi22.ij5.,2432!i"
	salt2 := "fo2.43o5h2f(juaz"
	md5_obj := md5.New()
	md5_obj.Write([]byte(salt1 + password + salt2))
	md5_encode := md5_obj.Sum(nil)
	md5_string := hex.EncodeToString(md5_encode)

	salt3 := "easfcvadwa"
	salt4 := "ofkafjdisa"
	sha1_obj := sha1.New()
	sha1_obj.Write([]byte(salt3 + md5_string + salt4))
	sha1_encode := sha1_obj.Sum(nil)
	cryptograph = hex.EncodeToString(sha1_encode)
	//fmt.Println(cryptograph)

	return cryptograph
}
