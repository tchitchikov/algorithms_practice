
def traverse_linked_list(linked_list, node, sum_so_far):
    sum_so_far = sum_so_far + linked_list[node]['data']
    if not linked_list[node]['next']:
        return sum_so_far, node
    else:
        node = linked_list[node]['next']
        return traverse_linked_list(linked_list, node, sum_so_far)


if __name__ == '__main__':
    linked_list = {
        'a': {'id': 1, 'next': 'b', 'prev': None, 'data': 3},
        'b': {'id': 2, 'next': 'c', 'prev': 'a', 'data': 6},
        'c': {'id': 3, 'next': 'd', 'prev': 'c', 'data': 11},
        'd': {'id': 4, 'next': None, 'prev': 'c', 'data': 100}
    }
    print(traverse_linked_list(linked_list, 'a', 0))