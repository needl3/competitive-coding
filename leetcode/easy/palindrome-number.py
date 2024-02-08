class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        else:
            reversedNumber = self.reverseTheNumber(x)
            if reversedNumber == x:
                return True
    def reverseTheNumber(self,number):
        reversedOne = 0
        while number:
            temp = number % 10
            number = number // 10
            reversedOne = reversedOne*10+temp
        return reversedOne
