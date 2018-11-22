# coding: utf8

class Solution(object):
    def firstUniqChar(self, s):
        """
        :type s: str
        :rtype: int
        """
        map = {}
        for index, value in enumerate(s):
            if map.has_key(value):
                map[value] = -1
            else:
                map[value] = index
        ans = len(s) + 1
        for v in map.values():
            if v < 0:
                continue
            if ans > v:
                ans = v
        if ans > len(s):
            return -1
        return ans