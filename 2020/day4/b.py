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


YEAR = re.compile('^[0-9]{4}$')
HEIGHT = re.compile('^[0-9]{2,3}(cm|in)$')
HEX_COLOR = re.compile('^#[0-9a-f]{6}$')
PID = re.compile('^[0-9]{9}$')

checks = {'byr': lambda v: YEAR.fullmatch(v) and 1920 <= int(v) <= 2002,
          'iyr': lambda v: YEAR.fullmatch(v) and 2010 <= int(v) <= 2020,
          'eyr': lambda v: YEAR.fullmatch(v) and 2020 <= int(v) <= 2030,
          'hgt': lambda v: HEIGHT.fullmatch(v) and (
                  (v[-2:] == 'cm' and 150 <= int(v[:-2]) <= 193) or
                  (v[-2:] == 'in' and 59 <= int(v[:-2]) <= 76)),
          'hcl': lambda v: HEX_COLOR.fullmatch(v),
          'ecl': lambda v: v in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl',
                                 'oth'],
          'pid': lambda v: PID.fullmatch(v)}


def run(batch):
    invalid = 0
    valid = []
    for item in batch:
        item.pop("cid", None)
        if len(item) < 7:
            invalid += 1
        else:
            valid.append(item)
    return valid


def is_valid(passport):
    for field, validation in checks.items():
        if field not in passport:
            print('Passport {0} is missing {1}.'.format(passport, field))
            return False
        if not validation(passport[field]):
            print('Passport {0} field {1} failed validation.'.format(passport,
                                                                     field))
            return False
    return True


def count_valid_passports(passports):
    count = 0
    for passport in passports:
        count += 1 if is_valid(passport) else 0
    return count


if __name__ == '__main__':
    print('{0} valid passports.'.format(count_valid_passports(get_input())))
