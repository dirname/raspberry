cmake_minimum_required(VERSION 3.13)
project(touch C)

set(CMAKE_C_STANDARD 90)

add_executable(touch main.c)

find_library(WIRINGPI_LIBRARIES NAMES wiringPi)
target_link_libraries(touch ${WIRINGPI_LIBRARIES})