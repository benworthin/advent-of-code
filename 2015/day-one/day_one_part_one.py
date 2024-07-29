def main():
    go_up = 1
    go_down = -1
    final_floor = 0

    # Retrieve puzzle input
    with open('input.txt') as file:
        puzzle_input = file.readline().strip()

    for direction in puzzle_input:
        if direction == '(':
            final_floor += go_up;
        else:
            final_floor += go_down

    print(f"Final Floor: {final_floor}")


if __name__ == "__main__":
    main()
