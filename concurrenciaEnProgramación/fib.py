# función fibonacci sin concurrencia en python
import random
import time
import sys


def fibonacci(n):  # función fibonacci
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)


def random_num():  # generador de números
    return random.randint(38, 42)


# cantidad de números y arreglo
nums = 10
if len(sys.argv) > 1:
    nums = int(sys.argv[1])
else:
    nums = 10
print("Cantidad de números:", nums)
arr = []

# llenar arreglo
for i in range(nums):
    arr.append(random_num())

# mostar arreglo
print(arr)
# calcular fibonacci y mostar
print("número  fibonacci  tiempo")
for num in arr:
    inicio = time.time()
    num_fib = fibonacci(num)
    print(num, "    ", num_fib, " ", time.time() - inicio)  # final

print("Tomó:", time.time() - inicio)
