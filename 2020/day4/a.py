import re


def get_input():
    with open("input.txt", "r") as f:
        batch = f.read()
        sl = re.split('\n\n', batch)
        dict_list = []
        for passport in sl:
            passports = re.sub(r"\n", " ", passport).rstrip().split(" ")
            dict_list.append(dict(s.split(':') for s in passports))
        return dict_list


def run(batch):
    all_b = len(batch)
    invalid = 0
    for item in batch:
        item.pop("cid", None)
        if len(item) < 7:
            invalid += 1
    return all_b - invalid


if __name__ == '__main__':
    print(run(get_input()))
