class Solution:
    def countGoodRectangles(self, rectangles: List[List[int]]) -> int:
        m = [0,0]
        #count, side
        for i in rectangles:
            side = min(i)
            if side > m[1]:
                m = [1,side]
            elif side == m[1]:
                m[0] += 1
        return m[0]
