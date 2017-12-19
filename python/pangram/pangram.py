def is_pangram(sentence):
    alphabet = set(range(ord('a'), ord('z') + 1))
    chars = set([ord(c) for c in list(sentence.lower()) if ord(c) in alphabet])
    return True if alphabet == chars else False
