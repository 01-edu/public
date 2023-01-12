def do_punishment(first_part, second_part, nb_lines):
    first_part = first_part.strip() + " "
    second_part = second_part.strip() + ".\n"

    res = ""
    for _ in range(0, nb_lines):
        res += first_part + " " + second_part + "oiewoiew.\n"
    return res