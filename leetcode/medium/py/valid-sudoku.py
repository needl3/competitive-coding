class Solution:
	def isValidSudoku(self, board: list[list[str]]) -> bool:
		for r in range(9):
			for c in range(9):
				if board[r][c] != '.':
					if not self.checkForSpotPossibility(board, board[r][c], (r,c)):
						return 0
		return 1
	def checkForSpotPossibility(self,board, givenVal, index):
		for r in range(9):
			if board[r][index[1]] == givenVal and index != (r,index[1]):
				return 0
		for c in range(9):
			if board[index[0]][c] == givenVal and index != (index[0],c):
				return 0
		for r in range(3):
			for c in range(3):
				if board[index[0]//3*3+r][index[1]//3*3+c] == givenVal and index != (index[0]//3*3+r,index[1]//3*3+c):
					return 0
		return 1
