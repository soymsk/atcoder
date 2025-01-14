if __name__ == "__main__":
    N, D = (int(x) for x in input().split())
    T = []
    L = []
    W = []
    for i in range(N):
        tl = [int(x) for x in input().split()]
        T.append(tl[0])
        L.append(tl[1])

    for k in range(D + 1):
        if k == 0:
            continue
        max_w = 0
        for i in range(N):
            w = T[i] * (L[i] + k)
            if w > max_w:
                max_w = w
                # max_i = i

        print(max_w)
