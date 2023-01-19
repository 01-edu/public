import sys

sys.path.append('student')

from shopping import clean_list

def test_normal_input():
    shopping_list = ['book', 'pen', 'milk', 'bread']
    expected_output = ['1/ Book', '2/ Pen', '3/ Milk', '4/ Bread']
    assert clean_list(shopping_list) == expected_output

def test_empty_input():
    shopping_list = []
    expected_output = []
    assert clean_list(shopping_list) == expected_output

def test_multiple_spaces():
    shopping_list = ['    shirt    ', '  pants  ', 'milk', 'shoes']
    expected_output = ['1/ Shirt', '2/ Pants', '3/ Milk', '4/ Shoes']
    assert clean_list(shopping_list) == expected_output
  
def test_missing_milk():
    shopping_list = ['computer', 'chair', 'desk']
    expected_output = ['1/ Computer', '2/ Chair', '3/ Desk', '4/ Milk']
    assert clean_list(shopping_list) == expected_output

def test_word_with_space():
    shopping_list = ['tooth brush', 'hand sanitizer', 'milk', 'toilet paper']
    expected_output = ['1/ Tooth brush', '2/ Hand sanitizer', '3/ Milk', '4/ Toilet Paper']
    assert clean_list(shopping_list) == expected_output
