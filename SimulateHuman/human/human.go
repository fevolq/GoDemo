/*
@Time: 2023/6/12 22:02
@Description:
*/

package human

import "sync"

const addNumber = 2        // 每次生孩子的数量
const minAddHumanYear = 15 // 可生孩子的最小年龄
const maxAddHumanYear = 45 // 可生孩子的最大年龄
const maxYear = 75         // 最大存活年龄

var wg sync.WaitGroup

type Human struct {
	Year   int  `json:"year"`
	Status bool `json:"status"`
}

//var pool = sync.Pool{
//	New: func() interface{} {
//		return new(Human)
//	},
//}

func initHuman() *Human {
	//human := pool.Get().(*Human)
	human := new(Human)
	human.Year = 1
	human.Status = true
	return human
}

func (human Human) AddHuman() []*Human {
	var sonHuman = make([]*Human, 0)
	if !human.conformAddHuman() {
		return sonHuman
	}

	for i := 0; i < addNumber; i++ {
		son := initHuman()
		sonHuman = append(sonHuman, son)
	}
	return sonHuman
}

// 是否符合生孩子的条件
func (human Human) conformAddHuman() bool {
	return human.Status && human.Year >= minAddHumanYear && human.Year <= maxAddHumanYear && human.Year%15 == 0
}

// 是否会死亡
func (human Human) willDie() bool {
	return human.Year >= maxYear
}

func (human *Human) BeDead() {
	human.Status = false
}

func (human *Human) AddYear() {
	defer func() {
		wg.Done()
		//pool.Put(human)
	}()

	if human == nil || !human.Status {
		return
	}

	human.Year += 1

	if human.conformAddHuman() {
		sonHuman := human.AddHuman()

		lock.Lock()
		people = append(people, sonHuman...)
		lock.Unlock()
	}

	if human.willDie() {
		human.BeDead()
	}
}

/*
TODO:
1. 增加随机死亡事件（使用随机数权重）
2. 改变生孩子的逻辑：最大最小区间不变，每两年随机，数量随机，且与上次生孩子的时间和总数有关
3. 增加性别——设计生孩子逻辑（待定）
*/
