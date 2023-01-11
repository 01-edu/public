#!/usr/bin/env bash

IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {

    submitted=$(echo -e "0\n3\n2\n5\n7\n1\n4\n9\n8\n6\n10" | bash -c ""$script_dirS"/student/greatest-of-all.sh")
    expected=$(echo -e "0\n3\n2\n5\n7\n1\n4\n9\n8\n6\n10" | bash -c ""$script_dirS"/solutions/greatest-of-all.sh")
    diff <(echo "$submitted") <(echo "$expected")
    submitted=$(echo -e "26\n85\n21\n94\n68\n60\n99\n31\n10\n98\n" | bash -c ""$script_dirS"/student/greatest-of-all.sh")
    expected=$(echo -e "26\n85\n21\n94\n68\n60\n99\n31\n10\n98\n" | bash -c ""$script_dirS"/solutions/greatest-of-all.sh")
    diff <(echo "$submitted") <(echo "$expected")
    submitted=$(echo -e "152\n485\n569\n611\n871\n551\n984\n895\n285\n989\n" | bash -c ""$script_dirS"/student/greatest-of-all.sh")
    expected=$(echo -e "152\n485\n569\n611\n871\n551\n984\n895\n285\n989\n" | bash -c ""$script_dirS"/solutions/greatest-of-all.sh")
    diff <(echo "$submitted") <(echo "$expected")
    submitted=$(echo -e "152\n10001\n569" | bash -c ""$script_dirS"/student/greatest-of-all.sh")
    expected=$(echo -e "152\n10001\n569" | bash -c ""$script_dirS"/solutions/greatest-of-all.sh")
    diff <(echo "$submitted") <(echo "$expected")
    submitted=$(echo -e "152\n485\nalpha\n" | bash -c ""$script_dirS"/student/greatest-of-all.sh")
    expected=$(echo -e "152\n485\nalpha\n" | bash -c ""$script_dirS"/solutions/greatest-of-all.sh")
    diff <(echo "$submitted") <(echo "$expected")
    submitted=$(echo -e "152\n485\n-45\n45\n" | bash -c ""$script_dirS"/student/greatest-of-all.sh")
    expected=$(echo -e "152\n485\n-45\n45" | bash -c ""$script_dirS"/solutions/greatest-of-all.sh")
    diff <(echo "$submitted") <(echo "$expected")
}

challenge
