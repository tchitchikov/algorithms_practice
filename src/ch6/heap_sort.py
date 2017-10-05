import copy
import math


def heap_sort(array):
    array = build_max_heap(array)
    i = len(array) - 1
    output = []
    while i >= 0:
        output.insert(0, array[0])
        array = array[1:]
        array = max_heap(array, 0)
        i = i - 1
    return output


def build_max_heap(array):
    i = math.floor(len(array)/2)
    while i > 0:
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


if __name__ == '__main__':
    array = [5, 1, 3, 4, 2]
    print(build_max_heap(array))
    print(heap_sort(array))
