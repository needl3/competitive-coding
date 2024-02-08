class Solution:
    def reverse(self, x: int) -> int:
        num=0
        small = 1
        if x<0:
            x=x*-1
            small = -1
        while x:
            temp = x%10
            x=x//10
            num = num*10+temp
        if num > pow(2,31)-1 or num<pow(-2,31):
            return 0
        return num*small
