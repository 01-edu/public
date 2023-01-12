import sys

sys.path.append('/jail/app/student')

from punishment import do_punishment

def test_empty_lines():
    assert do_punishment("", "", 3) == " .\n .\n .\n"

def test_empty():
    assert do_punishment("", "", 0) == ""

def test_truncation():
    assert do_punishment("   Test for   ", "   truncating strings ", 1) == "Test for truncating strings.\n"