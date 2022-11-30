#!/usr/bin/env bash
echo "compiling ..."
clang++ `pkg-config --cflags gtk4` main.cpp `pkg-config --libs gtk4`