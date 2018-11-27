# coding:utf8

# Definition for an interval.
class Interval:
    def __init__(self, s=0, e=0):
        self.start = s
        self.end = e

class Solution:
    def insert(self, intervals, newInterval):
        """
        :type intervals: List[Interval]
        :type newInterval: Interval
        :rtype: List[Interval]
        """
        return_itervals = []
        next_insert = Interval(-1, -1)
        if not intervals:
            return [newInterval]
        for ite in intervals:
            # 融合区间的起始点已经找到的情况
            if (next_insert.start >= 0):
                # 融合区间的终止点也已经找到
                if (next_insert.end >= 0):
                    return_itervals.append(ite)
                    continue
                if (ite.start > newInterval.end):
                    next_insert.end = newInterval.end
                    return_itervals.append(next_insert)
                    return_itervals.append(ite)
                    continue
                if (ite.start <= newInterval.end):
                    if (ite.end >= newInterval.end):
                        next_insert.end = ite.end
                        return_itervals.append(next_insert)
                        continue

            else:
                if (newInterval.start <= ite.start):
                    next_insert.start = newInterval.start
                    if (newInterval.end < ite.start):
                        return_itervals.append(newInterval)
                        return_itervals.append(ite)
                        next_insert.end = newInterval.end
                        continue
                    if (newInterval.end <= ite.end):
                        next_insert.end = ite.end
                        return_itervals.append(next_insert)
                        continue
                if (newInterval.start > ite.start):
                    if (newInterval.start > ite.end):
                        return_itervals.append(ite)
                        continue
                    next_insert.start = ite.start
                    if (newInterval.end <= ite.end):
                        next_insert.end =ite.end
                        return_itervals.append(next_insert)
                        continue
        if next_insert.start < 0:
            next_insert.start = newInterval.start
        if next_insert.end < 0 :
            next_insert.end = newInterval.end
            return_itervals.append(next_insert)
        return return_itervals