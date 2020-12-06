def get_input():
    with open("input.txt", "r") as f:
        return [line.strip() for line in f.readlines()]


def run(field, r, d):
    x = 0
    y = 0
    line = len(field[0])
    trees = 0
    len_map = len(field)
    len_field = len(field)
    while len_map > 0:
        if x > len_field:
            break
        if y >= line:
            y = y - line
        pos = field[x][y]
        len_map -= 1
        if pos == "#":
            trees += 1
        x, y = x + d, y + r
    return trees


if __name__ == '__main__':
    print(run(get_input(), 3, 1))
