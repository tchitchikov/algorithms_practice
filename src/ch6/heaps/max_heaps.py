import math


def build_max_heap(array):
    i = math.floor(len(array)/2)
    while i >= 0:
        array = max_heap(array, i)
        i = i - 1
    return array


def max_heap(array, loc):
    if loc == 0:
        array = largest_value_first(array)
        max_heap(array, 1)
    left = 2 * loc
    right = 2 * loc + 1
    if left <= len(array)-1 and array[left] > array[loc]:
        largest = left
    else:
        largest = loc
    if right <= len(array)-1 and array[right] > array[largest]:
        largest = right
    if largest != loc:
        array[loc], array[largest] = array[largest], array[loc]
        max_heap(array, largest)
    return array


def largest_value_first(array):
    max_val_loc = 0
    for i in range(1, len(array)-1):
        if array[i] > array[max_val_loc]:
            max_val_loc = i
    if max_val_loc != 0:
        array[0], array[max_val_loc] = array[max_val_loc], array[0]
    return array