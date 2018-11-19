# coding:utf8


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:

    def rotateRight(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        # 如果传入的链表为空或者只有一个节点或是k的值为0时，直接返回原链表
        if (not head) or k == 0:
            return head
        # 获取链表长度同时获取到原链表的最后一个节点
        end = head
        Length = 1
        while end.next is not None:
            Length += 1
            end = end.next

        # 如果k的值是链表长度的整数倍，直接返回原链表
        if not (k % Length):
            return head

        end.next = head
        for i in range(0, (Length - k % Length)):
            end = end.next
        head = end.next
        end.next = None
        return head



