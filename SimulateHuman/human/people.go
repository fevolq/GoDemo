/*
@Time: 2023/6/12 22:02
@Description:
*/

package human

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

var people = make([]*Human, 0)

// 初始化人类数量
func initPeople(initPeopleNum int) {
	for i := 0; i < initPeopleNum; i++ {
		human := initHuman()
		people = append(people, human)
	}
}

func Execute(initPeopleNum, years int) {
	initPeople(initPeopleNum)

	for year := 1; year <= years; year++ {
		number := len(people)
		for index := 0; index < number; index++ {
			wg.Add(1)
			go people[index].AddYear()
		}
	}
	wg.Wait()
}

func GetInfo() {
	fmt.Printf("总人数：%s\n", len(people))

	aliveNum := 0
	for _, human := range people {
		if human.Status {
			aliveNum++
		}
	}
	fmt.Printf("存活人数：%s", aliveNum)
}

/*
TODO:
1. 优化human初始化逻辑，使用sync.Pool
	var pool = sync.Pool{
	New: func() interface{} {
		return new(Human)
		},
	}
	human := pool.Get().(*Human)
2. 使用协程
3. 改变people结构，使用map[year][people]，每岁有哪些人——进而可优化GetInfo方法
*/
