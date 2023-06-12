/*
@Time: 2023/6/12 21:55
@Description:
*/

package main

import (
	"fmt"
	"todoPro/todo"
)

var todos = make(todo.Todos, 0, 10)

func main() {
	todo.AddToDo(&todos, "Test", "")
	todo.AddToDo(&todos, "test", "2")
	todo.AddToDo(&todos, "测试", "3")
	infos := todo.GetAll(todos, nil)
	fmt.Println(infos)

	if err := todo.UpdateToDo(&todos, 2, map[string]interface{}{"completed": true, "name": "Demo"}); err == nil {
		// completed 为 nil，表示未分配指针
		var completed *bool

		// 进行赋值
		//completed = new(bool)
		//*completed = true
		fmt.Println(todo.GetAll(todos, completed))
	} else {
		fmt.Print(err)
	}
}
