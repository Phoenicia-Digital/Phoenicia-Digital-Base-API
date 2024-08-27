# Backend API Will be Run on a Ubuntu Server By Default.

FROM ubuntu:latest

# Procedures that will likely Not Be Changed unless neccessary.
# DO NOT CHANGE THIS IF YOU DONT KNOW WHAT YOU ARE DOING!

WORKDIR /Backend
COPY ../dist .

# Do not change this Unless You Have Changed main.go <file name> or the default build output name!
CMD [ "./main" ]