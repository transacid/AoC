import re


def get_input():
    with open("input.txt", "r") as f:
        batch = f.read()
        sl = re.split('\n\n', batch)
        dict_list = []
        for passport in sl:
            passports = re.sub(r"\n", " ", passport).strip().split(" ")
            dict_list.append(passports)
        return dict_list


def run(passes):
    yeses = 0
    for group in passes:
        indi_yes = len("".join(set("".join(group))))
        yeses = yeses + indi_yes
    return yeses


if __name__ == '__main__':
    print(run(get_input()))
