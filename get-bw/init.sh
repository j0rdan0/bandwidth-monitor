#!/bin/bash

cd ../web-ui

ng serve >/dev/null 2>&1 & 

cd -

./get-bw
