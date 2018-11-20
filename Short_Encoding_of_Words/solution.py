# coding:utf8

class Solution:

    def minimumLengthEncoding(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        # 先去重并排序
        words = sorted([word[::-1] for word in set(words)])
        print words
        ans = 0
        last = ""
        words.append("")
        for word in words:
            if not word.startswith(last):
                ans += len(last)+1
            last = word
        return ans









