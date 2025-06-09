# Serial Randomness test

This repository contains an experiment to verify parameters for a certificate
lint.

Serial numbers are required to have 64 bits of randomness.

How can we write a lint to verify that's true?

We'd expect on average a serial with 64 bits of randomness to have some number
of zero bits and one bits.

Naively, you might say 32 of each, but of course due to randomness it could be
less. As well, leading zero bits aren't encoded, so they're "invisible".

As part of choosing the parameters for the lint, this repository contains a
test which generates a given number of random 64-bit integers, and calculates
the number of zeros and ones in each.

