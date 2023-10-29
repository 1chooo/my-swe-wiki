def create_heart_pattern():
    lines = []
    for y in range(30, -30, -1):
        line = ''.join(['Love'[(x - y) % len('Love')] if ((x * 0.05) ** 2 + (y * 0.1) ** 2 - 1) ** 3 - (x * 0.05) ** 2 * (y * 0.1) ** 3 <= 0 else ' ' for x in range(-30, 30)])
        if line != ' ' * len(line):
            lines.append(line)
    return '\n'.join(lines)


if __name__ == '__main__':
    pattern = create_heart_pattern()
    print(pattern)
