# coding:utf8


class Solution:
    def checkRecord(self, s):
        """
        :type s: str
        :rtype: bool
        """
        if len(s.split("A")) > 2:
            return False
        if len(s.split("LLL")) > 1:
            return False
        return True
