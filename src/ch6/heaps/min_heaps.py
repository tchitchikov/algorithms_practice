import math


def build_min_heap(array):
    i = 0
    while i <= math.floor(len(array)/2):
        array = min_heap(array, i)
        i = i + 1
    return array


def min_heap(array, loc):
    if loc == 0:
        array = smallest_value_first(array)
        min_heap(array, 1)
    left = 2 * loc
    right = 2 * loc + 1
    if left <= len(array)-1 and array[left] < array[loc]:
        smallest = left
    else:
        smallest = loc
    if right <= len(array)-1 and array[right] < array[smallest]:
        smallest = right
    if smallest != loc:
        array[loc], array[smallest] = array[smallest], array[loc]
        min_heap(array, smallest)
    return array


def smallest_value_first(array):
    min_val_loc = 0
    for i in range(1, len(array)-1):
        if array[i] < array[min_val_loc]:
            min_val_loc = i
    if min_val_loc != 0:
        array[0], array[min_val_loc] = array[min_val_loc], array[0]
    return array
