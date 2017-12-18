def is_isogram(string):
    counts = {}
    for char in string.lower():
        if char in ' -':
            continue
        elif char in counts:
            return False
        else:
            counts[char] = 1
    return True
