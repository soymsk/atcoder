import sys

sys.set_int_max_str_digits(10**6)


def main() -> None:
    N, M, A, B = (int(x) for x in input().split())
    L, R = [], []

    for i in range(M):
        l, r = (int(x) for x in input().split())
        L.append(l)
        R.append(r)

    dp = [False] * N
    for x in range(N):
        if x == 0:
            dp[0] = True
            continue

        for i in range(A, B + 1):
            for m in range(M):
                # if dp[x - i] and not (L[m] <= x - i + 1 and x - i + 1 <= R[m]):
                if x - i >= 0 and dp[x - i]:
                    dp[x] = True

                if L[m] <= x + 1 and x + 1 <= R[m]:
                    # current is bad.
                    dp[x] = False

    # print(dp)
    print("Yes" if dp[-1] else "No")


if __name__ == "__main__":
    main()
