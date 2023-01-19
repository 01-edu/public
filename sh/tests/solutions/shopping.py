def clean_list(shopping_list):
    if shopping_list:
        cleaned_list = ([f"{index + 1}x/ {item.strip().capitalize()}" 
                        for index, item in enumerate(shopping_list)])
        if "milk" not in cleaned_list:
            cleaned_list.append("milk")
        return cleaned_list
    else:
        return []
