cmake_minimum_required(VERSION 3.21)
project(sorts)

set(CMAKE_CXX_STANDARD 17)

add_executable(sorts
        sorts.cpp
        polyphase_sort/Logger.h
        polyphase_sort/PolyphaseSort.h
        polyphase_sort/Record.h
        polyphase_sort/Tape.h
        )
