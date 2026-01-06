#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'
FILENAME="student/append-output.sh"
script_dirS=$(cd -P "$(dirname "$BASH_SOURCE")" &>/dev/null && pwd)

if test ! -e append-output; then
	mkdir append-output
	cat <<EOF >append-output/results.txt
"In the End" - Linkin Park
"Crawling" - Linkin Park
"Elevation" - U2
"Get the Party Started" - Pink
"Lady Marmalade" - Christina Aguilera, Lil' Kim, Mya, Pink
EOF
	cat <<EOF >append-output/songs.txt
"Breathe" - Faith Hill
"It Wasn't Me" - Shaggy featuring Ricardo "RikRok" Ducent
"Hanging by a Moment" - Lifehouse
"Shape of My Heart" - Backstreet Boys
"Thank You" - Dido
"I'm Like a Bird" - Nelly Furtado
"Family Affair" - Mary J. Blige
"Fallin'" - Alicia Keys
"All for You" - Janet Jackson
"I Wanna Know" - Joe
"U Remind Me" - Usher
"U Got It Bad" - Usher
"I'm a Believer" - Smash Mouth
"Get the Party Started" - Pink
"Wherever You Will Go" - The Calling
"In the End" - Linkin Park
"Loser" - Beck
"Get Ur Freak On" - Missy Elliott
"I'm Real" - Jennifer Lopez
"Butterfly" - Crazy Town
"Crawling" - Linkin Park
"I'm a Slave 4 U" - Britney Spears
"Elevation" - U2
"Lady Marmalade" - Christina Aguilera, Lil' Kim, Mya, Pink
"Play" - Jennifer Lopez
"I'm Just a Kid" - Simple Plan
"Imitation of Life" - R.E.M.
"Big Pimpin'" - Jay-Z
"Stutter" - Joe featuring Mystikal
"I Wish" - R. Kelly
"This Is the Night" - Clay Aiken
"Hella Good" - No Doubt
"I Know" - Dionne Farris
"I'll Be Missing You" - Puff Daddy and Faith Evans featuring 112
"I Try" - Macy Gray
"Thong Song" - Sisqo
"Survivor" - Destiny's Child
"I Want It That Way" - Backstreet Boys
"Bad Day" - Daniel Powter
"I'm Like a Bird" - Nelly Furtado
"I Need to Know" - Marc Anthony
"Follow Me" - Uncle Kracker
"Hemorrhage (In My Hands)" - Fuel
"Soak Up the Sun" - Sheryl Crow
"I Hope You Dance" - Lee Ann Womack
"Can't Get You Out of My Head" - Kylie Minogue
"I Just Wanna Love U (Give It 2 Me)" - Jay-Z
"My Love Is Your Love" - Whitney Houston
"Bounce with Me" - Lil' Bow Wow
"Where the Party At" - Jagged Edge
"I'm Already There" - Lonestar
"I Don't Want to Miss a Thing" - Aerosmith
"If You Could Read My Mind" - Stars on 54
"My Way" - Usher
"Always on Time" - Ja Rule featuring Ashanti
EOF
fi
challenge() {
	$(cd "$1" && bash "$script_dirS"/$FILENAME)
	submitted=$(cat $1/results.txt)
	rm $1"/results.txt"
	$(cd "$1" && bash "$script_dirS"/solutions/append-output.sh)
	expected=$(cat $1/results.txt)
	diff <(echo "$submitted") <(echo "$expected")
}

if [ -s ${FILENAME} ]; then
	if [[ $(cat $FILENAME | grep "cat <<EOF >results.txt" | wc -l) -ne 0 ]]; then
		echo "cheating is not allowed in this exercise!"
		exit 1
	elif
		[[ $(cat $FILENAME | grep echo | wc -l) -ne 0 ]]
	then
		echo "cheating is not allowed in this exercise!"
		exit 1
	fi
fi

challenge append-output
rm -r append-output
