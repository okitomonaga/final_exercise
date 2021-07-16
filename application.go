package main

import (
	//"reflect"
	"fmt"
	"strconv"
	"strings"
)

func conrvetIntSlice(str string) []int {

	var result []int
	slice := strings.Split(str, ",")
	if len(slice) != 2 {
		return []int{-1, -1}
	}
	r1, error1 := strconv.Atoi(slice[0])
	r2, error2 := strconv.Atoi(slice[1])
	if error1 != nil || error2 != nil {
		return []int{-1, -1}
	}
	result = append(result, r1)
	result = append(result, r2)
	return result

}

func storeBan(input []int, player int, ban [][]int) [][]int {

	result := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	//result := make([][]int, 9)

	copy(result, ban)

	result[input[0]][input[1]] = player //

	return result
}

func canPut(input []int, ban [][]int) bool {
	if input[0] > 2 || input[1] > 2 || input[0] < 0 || input[1] < 0 {
		return false
	}
	return ban[input[0]][input[1]] == 0
}

func generateBanString(ban [][]int) []string {
	result := make([]string, 3)
	for h := 0; h < 3; h++ {
		for w := 0; w < 3; w++ {

			switch ban[h][w] {
			case 0:
				result[h] += "."
			case 1:
				result[h] += "o"
			case 2:
				result[h] += "x"
			default:
			}

		}
	}

	return result
}

func isWin(ban [][]int) bool {
	for i := 0; i < 3; i++ {
		if ban[i][0] == ban[i][1] && ban[i][0] == ban[i][2] && ban[i][0] != 0 {
			return true
		}
		if ban[0][i] == ban[1][i] && ban[0][i] == ban[2][i] && ban[0][i] != 0 {
			return true
		}
	}

	if ban[0][0] == ban[1][1] && ban[0][0] == ban[2][2] && ban[0][0] != 0 {
		return true
	}

	if ban[0][2] == ban[1][1] && ban[0][2] == ban[2][0] && ban[0][2] != 0 {
		return true
	}

	return false
}

func isDraw(ban [][]int) bool {

	if isWin(ban) {
		return false
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if ban[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	inputPut := make([]int, 2)
	//ban := make([][]int, 9)
	ban := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	outputString := make([]string, 3)

	var input string

	playernum := 1          ///先手から開始
	for i := 0; true; i++ { //無限ループ
		fmt.Printf("Player %d: Input (x,y) : ", playernum)
		fmt.Scan(&input)
		copy(inputPut, conrvetIntSlice(input))

		//おけた時の処理
		if canPut(inputPut, ban) {
			copy(ban, storeBan(inputPut, playernum, ban))
			if isWin(ban) {
				fmt.Printf("Player %d won.\n", playernum)
				break
			}
			if isDraw(ban) {
				fmt.Println("Draw.")
				break
			}
			if playernum == 1 { //ここで先手・後手交代
				playernum = 2
			} else {
				playernum = 1
			}
		} else {
			fmt.Println("input is invalid.")
		}

		copy(outputString, generateBanString(ban))
		for j := 0; j < len(outputString); j++ {
			fmt.Println(outputString[j])
		}
		//勝利条件を満たしたときbreak
	}

	//fmt.Print("Player 1: Input (x,y) : ")

	//fmt.Scan(&input)

	//copy(inputPut, conrvetIntSlice(input))
	//copy(ban, storeBan(inputPut, 1))
	copy(outputString, generateBanString(ban))
	for i := 0; i < len(outputString); i++ {
		fmt.Println(outputString[i])
	}
}
