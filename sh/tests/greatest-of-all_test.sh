#!/usr/bin/env bash

IFS='
'
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

challenge() {
	submitted=$(echo -e $1 | bash "./student/greatest-of-all.sh")
	expected=$(echo -e $1 | bash "./solutions/greatest-of-all.sh")
	diff <(echo "$submitted") <(echo "$expected")
	if [ $? != 0 ]; then
		exit 1
	fi
}

challenge "0\n3\n2\n5\n7\n1\n4\n9\n8\n6\n"

challenge "26\n85\n21\n94\n68\n60\n99\n31\n10\n98\n"

challenge "152\n485\n569\n611\n871\n551\n984\n895\n285\n989\n"

challenge "152\n10001\n569"

challenge "152\n485\nalpha\n"

challenge "152\n485\n-45\n45\n"
