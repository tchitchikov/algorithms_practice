import random

from heaps import max_heaps, min_heaps

def heap_sort(array):
    array = max_heaps.build_max_heap(array)
    i = len(array) - 1
    output = []
    while i >= 0:
        output.insert(0, array[0])
        array = array[1:]
        array = max_heaps.max_heap(array, 0)
        i = i - 1
    return output





if __name__ == '__main__':
    # array = [5, 1, 3, 4, 2]
    array = [random.randrange(0,10) for x in range(15)]
    print(max_heaps.build_max_heap(array))
    print(heap_sort(array))
    print(min_heaps.build_min_heap(array))
