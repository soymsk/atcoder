from bisect import bisect

if __name__ == "__main__":
    N = int(input())
    A = [int(x) for x in input().split()]

    dp = 0

    for i in range(N):
        b = A[i]
        j = bisect(A, b / 2)
        dp += j

    print(dp)
