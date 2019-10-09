#!/bin/sh
echo Hello $(echo $REPOSITORY | cut -d'/' -f4)!
