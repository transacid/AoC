def get_input():
    with open("input.txt", "r") as f:
        return [line.strip() for line in f.readlines()]


def run(field):
    x = 0
    y = 0
    line = len(field[0])
    trees = 0
    len_map = len(field)
    while len_map > 0:
        if y >= line:
            y = y - line
        pos = field[x][y]
        len_map -= 1
        if pos == "#":
            trees += 1
        x, y = x + 1, y + 3
    print(trees)


if __name__ == '__main__':
    run(get_input())