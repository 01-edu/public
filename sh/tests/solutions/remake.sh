#!/usr/bin/env bash

IFS='
'

if [[ $# -eq 1 && -d "$1" ]]; then
	cd $1
	touch -t 01010001 ciao
	chmod 442 ciao
	mkdir mamma
	touch -t 01020001 mamma
	chmod 777 mamma
	touch -t 01030001 guarda
	chmod 400 guarda
	touch -t 01040001 come
	chmod 642 come
	mkdir mi
	touch -t 01050001 mi
	chmod 452 mi
	touch -t 01060001 diverto
	chmod 421 diverto
else
	echo Error
	exit 1
fi
