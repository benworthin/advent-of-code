def main():
    go_up = 1
    go_down = -1
    final_floor = 0

    # Retrieve puzzle input
    with open('input.txt') as file:
        puzzle_input = file.readline().strip()

    for position, direction in enumerate(puzzle_input):
        if final_floor >= 0:
            if direction == '(':
                final_floor += go_up;
            else:
                final_floor += go_down
        else:
            print(f"Santa hit the basement at position: {position}")
            break


if __name__ == "__main__":
    main()
