def get_input():
    with open("input.txt", "r") as f:
        return [line.strip() for line in f.readlines()]


def run(passes):
    rows = [r for r in range(128)]
    cols = [r for r in range(8)]
    for action in passes:
        rlength = len(rows)
        clength = len(cols)
        if action == "F":
            del rows[int(rlength / 2):]
        elif action == "B":
            del rows[:int(rlength / 2)]
        elif action == "R":
            del cols[:int(clength / 2)]
        elif action == "L":
            del cols[int(clength / 2):]
    final_row = rows[0]
    final_col = cols[0]
    return final_row * 8 + final_col


if __name__ == '__main__':
    bpasses = []
    for boardinpass in get_input():
        bpass = run(boardinpass)
        bpasses.append(bpass)
    print(sorted(bpasses)[-1:])
