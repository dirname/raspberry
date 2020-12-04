#!/bin/bash
# Run as user pi
session=$(xauth list :0.0)
echo "sudo -S xauth add ${session}"
