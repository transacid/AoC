#!/usr/bin/env python3

from collections import Counter


def get_input():
    with open("input.txt", "r") as f:
        template = f.readline().strip()
        f.readline()
        return template, f.read().strip()


if __name__ == '__main__':
    tl, inputs = get_input()
    d = dict(x.split(" -> ") for x in inputs.split("\n"))
    for i in range(10):
        l = []
        for i, _ in enumerate(tl):
            l.append(tl[i])
            try:
                l.append(d[tl[i:i+2]])
            except KeyError:
                break
        tl = "".join(l)
    c = Counter(tl).most_common()
    max = c[0][1]
    min = c[-1][1]
    print(max - min)
