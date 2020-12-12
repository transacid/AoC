def get_input():
    with open("input.txt", "r") as f:
        return [line.strip().split(" ") for line in f.readlines()]


def nop():
    pass


def acc(accumulator, quantity):
    sign = quantity[0]
    count = int(quantity[1:])
    if sign == "+":
        accumulator += count
    else:
        accumulator -= count
    return accumulator


def jmp(stack, quantity):
    sign = quantity[0]
    count = int(quantity[1:])
    if sign == "+":
        stack += count
    else:
        stack -= count
    return stack


def run(ops):
    accumulator = 0
    stack = 0
    stack_list = []
    while True:
        if ops[stack][0] == "nop":
            nop()
            stack += 1
        elif ops[stack][0] == "acc":
            accumulator = acc(accumulator, ops[stack][1])
            stack += 1
        elif ops[stack][0] == "jmp":
            stack = jmp(stack, ops[stack][1])
        if stack in stack_list:
            return accumulator
        stack_list.append(stack)


if __name__ == '__main__':
    print(run(get_input()))
