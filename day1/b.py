def get_input():
    with open("input.txt", "r") as f:
        return [int(line.strip()) for line in f.readlines()]


def run(coins):
    for a in coins:
        for b in coins:
            for c in coins:
                if a + b + c == 2020:
                    return a, b, c


if __name__ == '__main__':
    nums = run(get_input())
    print("{0} + {1} + {2} = {3}".format(nums[0], nums[1], nums[2],
                                         nums[0] + nums[1] + nums[2]))
    print(nums[0] * nums[1] * nums[2])
