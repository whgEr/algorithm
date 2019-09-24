#!/usr/bin/env python3
# -*- coding: utf-8 -*-

def insert_sort(data):
    """ 插入排序
        原理：首元素认为有序序列，从下标为i=1开始处理，提出当前元素，与前i-1个元素比较
        value < data[i-1]  则元素后移 data[i]=data[i-1]
        value >= data[i-1] 至此必是有序列，则break

        Important:data[i]=value
    """

    l = len(data)
    if l <= 1:
        return data
    for i in range(1, l):
        value, k = data[i], i-1
        while k >= 0:
            if data[k] > value:
                data[k+1] = data[k]
            else:
                break
            # 此句可以解决第0个元素不能处理的情况
            data[k] = value
            k -= 1

    return data
