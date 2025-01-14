from bisect import bisect_left

if __name__ == "__main__":
    N = int(input())
    A = [int(x) for x in input().split()]

    dp = 0

    for i in range(N):
        a = A[i]
        if a == 0:
            continue
        b = 2 * a
        j = bisect_left(A, b)
        if j < N:
            dp += 1
            A[j] = 0

    print(dp)
