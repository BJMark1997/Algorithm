package main

//input: posX = 15 posY = 20
//arr = [[35,20],[30,10],[20,25],[15,30],[20,15],[25,20]]
//output: 3

import "math"

func GetMatchPoint(posX int, posY int, arr [][]int) int {
	minIndex := math.MaxInt64
	minDis := math.MaxInt64
	for i, pos := range arr {
		if !patch(posX, posY, pos[0], pos[1]) {
			continue
		}
		dis := getDis(posX, posY, pos[0], pos[1])
		if dis < minDis {
			minDis = dis
			minIndex = i
		}
		if minIndex == math.MaxInt64 {
			return -1
		}

	}
	return minIndex
}

func getDis(posX1 int, posX2 int, posY1 int, posY2 int) int {
	dis := math.Abs(float64(posX1-posX2)) + math.Abs(float64(posY1-posY2))
	return int(dis)
}

func patch(posX1 int, posX2 int, posY1 int, posY2 int) bool {
	return posX1 == posX2 || posY1 == posY2
}
