# coding:utf8

class Solution(object):
    def selfDividingNumbers(self, left, right):
        """
        :type left: int
        :type right: int
        :rtype: List[int]
        """
        returnData = []
        for k in range(left, right+1):
            # 所有能被10整除的数都跳过
            if k % 10 == 0:
                continue
            needAdd = True
            for i in str(k):
                if not int(i) or k % int(i) != 0:
                    needAdd = False
                    break
            if needAdd:
                returnData.append(k)
        return returnData