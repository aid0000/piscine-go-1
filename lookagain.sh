#!/bin/sh
find . -type f -name "*.sh" | sed 's/.sh//' | sed 's/\(.\)//' | sed 's/\(\/\)//'