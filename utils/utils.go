package utils

import (
	"fmt"
	"github.com/kataras/iris"
	"log"
	"reflect"
	"time"
)

func SetObjByJson(obj interface{}, data map[string]interface{}) error {
	for key, value := range data {
		if err := setField(obj, key, value); err != nil {
			log.Println(err.Error())
			return err
		}
	}

	return nil
}

func setField(obj interface{}, name string, value interface{}) error {
	structData := reflect.TypeOf(obj).Elem()
	fieldValue, result := structData.FieldByName(name)

	if !result {
		log.Print("SetObjByJson set field fail")
		return fmt.Errorf("No such field %s", name)
	}

	fieldType := fieldValue.Type
	val := reflect.ValueOf(value)
	valTypeStr := fieldType.String()
	fieldTypeStr := fieldType.String()

	if valTypeStr == "float64" && fieldTypeStr == "int" {
		val = val.Convert(fieldType)
	}

	if fieldType != val.Type() {
		return fmt.Errorf("value type %s didn't match obj field type %s ", valTypeStr, fieldTypeStr)
	}

	return nil
}

func LogInfo(app *iris.Application, v ...interface{}) {
	app.Logger().Info(v)
}

func LogError(app *iris.Application, v ...interface{}) {
	app.Logger().Error(v)
}

func LogDebug(app *iris.Application, v ...interface{}) {
	app.Logger().Debug(v)
}

/**
 * 格式化数据
 */
func FormatDatetime(time time.Time) string {
	return time.Format("2006-01-02 03:04:05")
}
