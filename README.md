# script-notify
Notifies user when desired output is read in their script. Useful for bruteforcing in ctfs or long running programs.<br>
Usage Example: ./script-notify scripts/test.sh test3

*Only works on linux as it requires notify-send to be installed.

*For python scripts add *#!/usr/bin/env python3* to the top to make them executable and also add *flush = True* to your print statements as python buffers its output.