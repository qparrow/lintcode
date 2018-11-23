# coding:utf8


class Solution(object):
    def isToeplitzMatrix(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: bool
        """
        row = len(matrix)
        col = len(matrix[0])

        for i in range(0, row-1):
            for j in range(0, col-1):
                if matrix[i][j] != matrix[i+1][j+1]:
                    return False
        return True


if __name__ == "__main__":
    s =Solution()
    # m = [
    #     [44, 35, 39],
    #     [15, 44, 35],
    #     [17, 15, 44],
    #     [80, 17, 15],
    #     [43, 80,  0],
    #     [77, 43, 80]
    # ]
    m = [[18],[66]]
    s.isToeplitzMatrix(m)