def get_input():
    with open("input.txt", "r") as f:
        return [line.strip() for line in f.readlines()]


def run(passwords):
    valid = 0
    for password in passwords:
        s = password.split()
        o, l, p = s[0].split("-"), s[1].strip(":"), s[2]
        first = int(o[0]) - 1
        second = int(o[1]) - 1
        if (l is p[first]) ^ (l is p[second]):
            valid += 1
    print(valid)


if __name__ == '__main__':
    run(get_input())
