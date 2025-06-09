# Serial Randomness test

This repository contains an experiment to verify parameters for a certificate
lint.

Serial numbers are required to have 64 bits of randomness.

How can we write a lint to verify that's true?

We'd expect on average a serial with 64 bits of randomness to have some number
of zero bits and one bits.

As part of choosing the parameters for the lint, this repository contains a
test which generates a given number of random 64-bit integers, and calculates
the number of zeros and ones in each.

## Results

Generating 100 billion unsigned integers is recorded in 100b.csv.
I believe that is a scale that appropriately represents the WebPKI.

As expected:

Half of these numbers had no leading zeros.
The most common number of 1-bits is 32, and 31 for zero bits.

Up to 35 leading zeros were detected.
All serials had at least 4 zeros and 5 one bits set.

The data, graphed:

![leading, zeros and ones](https://github.com/user-attachments/assets/a07fc4aa-8093-4b1c-978a-1f1e4431e058)

## Linting

From this, we can tell that making a lint requiring at least a certain number of zeros and ones is challening
to be useful without any false positives.
