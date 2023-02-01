## Read file

### Instructions

Create a file `read_file.py` which will have a function `get_recipes(file_name)` that takes the `file name` as an argument.

Create another file `recipes_data.json` and paste the following content on it:

```json
[
  {
    "title": "Vegetable Lasagna",
    "ingredients": [
      "1 lb. lasagna noodles",
      "1 onion, diced",
      "2 cloves of garlic, minced",
      "2 bell peppers, diced",
      "1 zucchini, diced",
      "1 eggplant, diced",
      "1 (28 oz) can crushed tomatoes",
      "1/4 cup chopped fresh basil",
      "1/4 cup chopped fresh parsley",
      "1/4 cup grated Parmesan cheese",
      "1 cup ricotta cheese",
      "2 cups shredded mozzarella cheese"
    ],
    "people": 6
  },
  {
    "title": "Stir-Fried Shrimp with Vegetables",
    "ingredients": [
      "1 lb. large shrimp, peeled and deveined",
      "2 tbsp. vegetable oil",
      "1 red bell pepper, sliced",
      "1 onion, sliced",
      "1 cup sliced mushrooms",
      "2 cloves of garlic, minced",
      "1/4 cup soy sauce",
      "1/4 cup hoisin sauce",
      "2 tbsp. rice vinegar",
      "1 tsp. sesame oil",
      "1 tsp. sugar",
      "1/4 tsp. red pepper flakes"
    ],
    "people": 4
  },
  {
    "title": "Spicy Chicken Tacos",
    "ingredients": [
      "1 lb. boneless, skinless chicken breasts",
      "1 tbsp. chili powder",
      "1 tsp. cumin",
      "1/2 tsp. paprika",
      "1/4 tsp. salt",
      "1/4 tsp. black pepper",
      "1 tbsp. vegetable oil",
      "8 small flour or corn tortillas",
      "1 cup shredded lettuce",
      "1/2 cup diced tomatoes",
      "1/4 cup sour cream",
      "1/4 cup chopped fresh cilantro"
    ],
    "people": 4
  }
]
```

Your function has to read the file `recipes_data.json` and load its content into a list of dictionaries.

Expected output format:

```python
[{...}, {...}, {...}]
```

### Usage

Here is a possible `test.py` to test your functions:

```python
import read_file

print(read_file.get_recipes('recipes_data.json'))
```

```bash
$ python test.py
[{'title': 'Vegetable Lasagna', 'ingredients': ['1 lb. lasagna noodles', '1 onion, diced', '2 cloves of garlic, minced', '2 bell peppers, diced', '1 zucchini, diced', '1 eggplant, diced', '1 (28 oz) can crushed tomatoes', '1/4 cup chopped fresh basil', '1/4 cup chopped fresh parsley', '1/4 cup grated Parmesan cheese', '1 cup ricotta cheese', '2 cups shredded mozzarella cheese'], 'people': 6}, {'title': 'Stir-Fried Shrimp with Vegetables', 'ingredients': ['1 lb. large shrimp, peeled and deveined', '2 tbsp. vegetable oil', '1 red bell pepper, sliced', '1 onion, sliced', '1 cup sliced mushrooms', '2 cloves of garlic, minced', '1/4 cup soy sauce', '1/4 cup hoisin sauce', '2 tbsp. rice vinegar', '1 tsp. sesame oil', '1 tsp. sugar', '1/4 tsp. red pepper flakes'], 'people': 4}, {'title': 'Spicy Chicken Tacos', 'ingredients': ['1 lb. boneless, skinless chicken breasts', '1 tbsp. chili powder', '1 tsp. cumin', '1/2 tsp. paprika', '1/4 tsp. salt', '1/4 tsp. black pepper', '1 tbsp. vegetable oil', '8 small flour or corn tortillas', '1 cup shredded lettuce', '1/2 cup diced tomatoes', '1/4 cup sour cream', '1/4 cup chopped fresh cilantro'], 'people': 4}]
```

### References

- [Python and json basics](https://www.w3schools.com/python/python_json.asp)
