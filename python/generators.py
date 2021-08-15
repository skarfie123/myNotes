from typing import Generator


def natural_numbers(start: int):
    yield start
    yield from natural_numbers(start + 1)  # joins another "sequence"
    # generator equivalent of recursion gives an infinite sequence


def prime_numbers(nats: Generator = natural_numbers(2)):
    """https://www.youtube.com/watch?v=5jwV3zxXc8E"""
    n = next(nats)
    yield n
    yield from prime_numbers(
        i for i in nats if i % n != 0
    )  # pass in a new generator that excludes the next set of multiples


def fibonacci_numbers(a, b: int):
    yield a
    yield from fibonacci_numbers(b, a + b)


def print_ten(label: str, g: Generator):
    print(label, *(next(g) for _ in range(10)))


def main():
    print_ten("Natural numbers:", natural_numbers(1))
    print_ten("Prime numbers:", prime_numbers())
    print_ten("Fibonacci numbers:", fibonacci_numbers(1, 1))
    print_ten("Lucas numbers:", fibonacci_numbers(2, 1))


if __name__ == "__main__":
    main()
