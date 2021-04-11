package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

type Booking struct {
	// bookableDate是自定义验证方法
	CheckIn time.Time `form:"checkIn" binding:"required,bookabledate" time_format:"2006-01-02"`

	// validator支持引用验证，gtfield=CheckIn 当前值要大于CheckIn
	CheckOut time.Time `form:"checkOut" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if date, ok := field.Interface().(time.Time); ok {
		return date.Unix() > time.Now().Unix()
	}
	return false
}

/**
自定义校验方法
http://localhost:8080/test?checkIn=2021-05-01&checkOut=2022-01-01

中文翻译需要用到v9版本
*/
func main() {
	r := gin.Default()

	// 注册自定义校验函数
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/test", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBindWith(&b, binding.Query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "通过校验",
		})
	})
	r.Run()
}
