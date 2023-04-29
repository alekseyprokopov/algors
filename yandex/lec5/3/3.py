def get_prefix(array):
    prefix = [0] * (len(array) + 1)
    for i in range(1, len(array)):
        h = array[i] - array[i - 1]
        if h > 0:
            prefix[i + 1] = prefix[i] + h
        else:
            prefix[i + 1] = prefix[i]
    return prefix


if __name__ == '__main__':
    n = int(input())
    coords = [int(input().split()[1]) for i in range(n)]
    m = int(input())
    ran = [list(map(int, input().split())) for i in range(m)]

    prefix_count = get_prefix(coords)
    reverse_prefix = get_prefix(coords[::-1])

    for item in ran:
        l, r = item
        print(prefix_count)
        print(reverse_prefix)
        if l < r:
            count = prefix_count[r] - prefix_count[l]
        else:
            count = reverse_prefix[l] - reverse_prefix[r + 1]
        print(count)