#!/bin/bash

docker run -d \
--name=MARL111_DAQ \
-p 4222:4222 \
-v /data:/data \
-v /dev/ttyS0:/dev/ttyS0 \
lemi011b_daq --site=Marion --number=1

docker run -d \
--name=MARL11_INGRESS \
-p 4222:4222 \
-v /data:/data \
lemi011b_ingress --site=Marion
