class Solution:
	def solveSudoku(self, board: list[list[str]]) -> None:
		for row in range(9):
			for column in range(9):
				if board[row][column] == '.':
					for val in range(1,10):
						if self.checkForSpotPossibility(board, str(val), (row, column)):
							board[row][column] = str(val)
							if self.solveSudoku(board):
								return 1
							board[row][column] = '.'
					return 0
		return 1

	def checkForSpotPossibility(self,board, givenVal, index):
		for r in range(9):
			if board[r][index[1]] == givenVal:
				return 0
		for c in range(9):
			if board[index[0]][c] == givenVal:
				return 0
		for r in range(3):
			for c in range(3):
				if board[index[0]//3*3+r][index[1]//3*3+c] == givenVal:
					return 0
		return 1
