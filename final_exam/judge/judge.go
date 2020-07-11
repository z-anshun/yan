package judge

//判断获胜的

//黑棋是1，白棋是-1
func Winter(board *[15][15]int) int {
	//我们是从左->右，再从上 -> 下 左边，右上，上都不用判断
	for k1, v1 := range board {
		for k2, v2 := range v1 {
			//下面连成5个
			if v2 != 0 {
				if k1+4 <= 15 && board[k1+1][k2] == v2 && board[k1+2][k2] == v2 && board[k1+3][k2] == v2 && board[k1+4][k2] == v2 {
					return v2
				} else if k2+4 <= 15 && board[k1][k2+1] == v2 && board[k1][k2+2] == v2 && board[k1][k2+3] == v2 && board[k1][k2+4] == v2 {
					return v2
				} else if k1+4 <= 15 && k2-4 <= 15 && board[k1+1][k2-1] == v2 && board[k1+2][k2-2] == v2 && board[k1+3][k2-3] == v2 && board[k1+4][k2-4] == v2 {
					return v2
				}
			}

		}
	}
	return 0
}

//禁手规则  只针对黑棋（1）
//不难发现，只针对该该棋四周的，并且大于3的，大于4的，5，6
func Forbid(board *[15][15]int, x int, y int) string {

	Co_three := judge(board, x, y, 3)
	Co_four := judge(board, x, y, 4)
	Co_five := judge(board, x, y, 5)
	Co_six := judge(board, x, y, 6)

	if Co_three >= 2 {
		return "三三禁手"
	} else if Co_four >= 2 && Co_three == 1 {
		return "三四禁手"
	} else if Co_four >= 2 || Co_five > 0 {
		return "四四禁手"
	} else if Co_six > 0 {
		return "多连禁手"
	}
	//**************
	return ""

}
func judge(board *[15][15]int, y int, x int, k int) int {
	sum := 0
	count := 0

	//水平
	for i := x - 4; i < x+4 && i <= 15; i++ {
		if i < 0 {
			continue
		}
		sum += board[i][y]
	}
	if sum == k {
		count++
	}
	sum = 0
	//竖直
	for i := y - 4; i < y+4 && i <= 15; i++ {
		if i < 0 {
			continue
		}
		sum += board[x][i]
	}
	if sum == k {
		count++
	} else if sum == 2*k-1 {
		count += 2
	}
	//右下
	sum = 0
	n := x - 4
	for i := y - 4; i < y+4 && i <= 15 && n < x+4 && n <= 15; i++ {
		if i < 0 || n < 0 {
			n++
			continue
		}
		sum += board[n][i]
		n++
	}
	if sum == k {
		count++
	}
	//左上
	sum = 0
	n = x + 4
	for i := y - 4; i < y+4 && i <= 15 && n > x-4 && n <= 15; i++ {
		if i < 0 || n > 15 {
			n--
			continue
		}
		sum += board[n][i]
		n--
	}
	if sum == k {
		count++
	}
	return count
}
