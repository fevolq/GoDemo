/*
@Time: 2023/6/12 22:01
@Description: 人类数量模拟
*/

package main

import (
	"SimulateHuman/human"
)

const testYears = 100
const initPeopleNum = 10000

func main() {
	human.Execute(initPeopleNum, testYears)

	human.GetInfo()
}
