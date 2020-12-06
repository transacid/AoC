def get_input():
    with open("input.txt", "r") as f:
        return [line.strip() for line in f.readlines()]


def run(passwords):
    valid = 0
    for password in passwords:
        s = password.split()
        o, l, p = s[0].split("-"), s[1].strip(":"), s[2]
        occurrence = p.count(l)
        if l in p:
            if int(o[0]) <= occurrence <= int(o[1]):
                valid += 1
    print(valid)


if __name__ == '__main__':
    run(get_input())
