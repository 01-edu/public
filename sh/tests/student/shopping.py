def clean_list(shopping_list):
    if shopping_list:
        cleaned_list = ([f"{index + 1}/ {item.strip().capitalize()}" 
                        for index, item in enumerate(shopping_list)])
        if "Milk" not in shopping_list:
            cleaned_list.append(f"{len(cleaned_list)}/ Milk")
        return cleaned_list
    else:
        return []
