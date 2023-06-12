/*
@Time: 2023/6/12 21:55
@Description:
*/

package todo

import (
	"fmt"
	"reflect"
)

type Todos []*ToDo

// AddToDo 增加一个todo
func AddToDo(todos *Todos, name, remark string) {
	todo := Init(name, remark)
	*todos = append(*todos, &todo)
}

func GetType(args interface{}) {
	fmt.Printf("%s 的类型是 %s", args, reflect.TypeOf(args))
	fmt.Println()
}

// UpdateToDo 更新指定的todo
func UpdateToDo(todos *Todos, id int, attrs map[string]interface{}) error {
	for _, todo := range *todos {
		// 此时的todo是指针类型
		if id == (*todo).ID {
			if err := todo.updateToDo(attrs); err != nil {
				return err
			} else {
				return nil
			}
		}
	}
	return fmt.Errorf("no todo found with ID %d", id)
}

// GetAll 获取所有todo的信息
func GetAll(todos Todos, complete *bool) []map[interface{}]interface{} {
	result := make([]map[interface{}]interface{}, 0, len(todos))
	for _, todo := range todos {
		if complete == nil || (*todo).Completed == *complete {
			info := (*todo).Info()
			result = append(result, info)
		}
	}
	return result
}

// 值类型与指针类型
