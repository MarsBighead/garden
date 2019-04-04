######################
# Author: Toger      #
# Date  : 2017.02.10 #
#####################

1. Encode protobuf binary format from text via message schema xiaodu.BidRequest.
   cat  request_xiaodu.txt | protoc --encode=xiaodu.BidRequest xiaodu.proto >request_xiaodu.bin
2. Decode protobuf binary to key-value text via the same message schema.
   protoc --decode=xiaodu.BidRequest xiaodu.proto < request_xiaodu.bin 
3. Simple decode from protobuf binary 
   protoc --decode_raw < request_xiaodu.bin 

