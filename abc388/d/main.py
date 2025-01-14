import heapq
from dataclasses import dataclass, field


@dataclass(order=True)
class PrioritizedItem:
    priority: int
    index: int = field(compare=False)

    def __init__(self):
        pass


def main() -> None:
    N = int(input())
    A = [int(x) for x in input().split()]

    adults: list[PrioritizedItem] = []

    x = 0
    for i in range(N):
        A[i] += len(adults)
        if i > 0:
            x += 1
        if A[i] > 0:
            adult = PrioritizedItem()
            adult.index = i
            adult.priority = A[i] + x
            heapq.heappush(adults, adult)

        # Delete adults having no stones.
        while adults:
            adult = heapq.heappop(adults)
            if adult.priority - x > 0:
                heapq.heappush(adults, adult)
                break
            else:
                A[adult.index] = 0

    for a in adults:
        A[a.index] = a.priority - x

    print(" ".join([str(b) for b in A]))


if __name__ == "__main__":
    main()
