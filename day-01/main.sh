#!/bin/sh
cat input.txt | paste -sd '+' - | sed 's/++/\n/g' | xargs -I {} echo {} | bc | sort -n | tail -n 1
cat input.txt | paste -sd '+' - | sed 's/++/\n/g' | xargs -I {} echo {} | bc | sort -n | tail -n 3 | paste -sd '+' - | bc
