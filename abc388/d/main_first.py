def main() -> None:
    N = int(input())
    A = [int(x) for x in input().split()]
    non_zeros = {i for i, a in enumerate(A) if a > 0}

    for i, a in enumerate(A):
        if i > 0:
            non_zeros2 = {*non_zeros}
            for left in non_zeros2:
                if left >= i:
                    break

                A[left] -= 1
                A[i] += 1
                if A[left] == 0:
                    non_zeros.remove(left)

                if A[i] > 0:
                    non_zeros.add(i)

    print(" ".join([str(b) for b in A]))


if __name__ == "__main__":
    main()
