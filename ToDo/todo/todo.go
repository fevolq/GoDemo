/*
@Time: 2023/6/12 21:55
@Description:
*/

package todo

import (
	"fmt"
	"reflect"
)

type ToDo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	Remark    string `json:"remark"`
}

var todoId = 0

// Init 构造对象
func Init(name, remark string) ToDo {
	// TODO: 随机生成id
	todoId += 1
	id := todoId
	return ToDo{ID: id, Name: name, Completed: false, Remark: remark}
}

func (todo *ToDo) updateToDo(attrs map[string]interface{}) error {
	for key, value := range attrs {
		switch key {
		case "name":
			todo.Name = value.(string)
		case "completed":
			todo.Completed = value.(bool)
		case "remark":
			todo.Remark = value.(string)
		default:
			return fmt.Errorf("unknown attribute: %s", key)
		}
	}
	return nil
}

func (todo ToDo) Info() map[interface{}]interface{} {
	result := make(map[interface{}]interface{})

	// 获取结构体类型的元数据
	typeFields := reflect.TypeOf(todo)
	// 获取结构体值的元数据
	valueFields := reflect.ValueOf(todo)

	for i := 0; i < typeFields.NumField(); i++ {
		field := typeFields.Field(i)
		result[field.Name] = valueFields.Field(i).Interface()
	}
	return result
}
