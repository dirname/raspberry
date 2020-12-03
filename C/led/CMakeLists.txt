cmake_minimum_required(VERSION 3.13)
project(led C)

set(CMAKE_C_STANDARD 90)

add_executable(led main.c)

find_library(WIRINGPI_LIBRARIES NAMES wiringPi)
target_link_libraries(led ${WIRINGPI_LIBRARIES})