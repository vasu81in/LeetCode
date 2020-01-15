import math
def sieve(n):
    sieve = [i for i in xrange(2, n+1, 1)]
    for j in range(2, int(math.floor(math.sqrt(n))), 1):
        if j in sieve:
            for i in xrange(2*j, n+1, j):
                if i in sieve:
                    sieve.remove(i)
    print(sieve)


def is_prime(n):
    i = 2
    while i <= math.sqrt(n):
        if n % i == 0:
            return
        i += 1
    return n

def main():
    sieve(100)
    print(map(is_prime, range(2, 100,1)))


if __name__ == '__main__':
    main()
