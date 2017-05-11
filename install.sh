#!/usr/bash

echo "Build garden..."
go build
echo "End build garden..."
echo "Start write environment variables..."
pwd
ls garden.source 
`source env.gd`
echo "Start garden..."
./garden &

runstats=`ps -ef | grep garden`
echo "Garden is running... "$runstats
