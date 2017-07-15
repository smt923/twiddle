# twiddle
Shift characters and fiddle with words to generate common variations on them, designed for usernames, logins, etc

## Usage
Twiddle will read from stdin so either run the program and type stuff in, or more commonly pipe stuff into stdin for it

Basic usage:
```
cat usernames.txt | ./twiddle > moreusernames.txt
```
adding -l will run only a few common variatios, -e will add a lot more

## Variations/Output
Given an example input of 'aliceh', this is the output:

No flags:
```
alicehaliceh
aliceh0
aliceh1
aliceh
aaliceh
haliceh
```

Using -l (light variations)
```
aliceh
aaliceh
haliceh
```

Using -e (every variation)
```
alicehaliceh
aliceh0
aliceh1
ALICEH
Aliceh
Aaliceh
halice
hhaliceh
hhalice
haliceh0
haliceh1
alaliceh
aliceh
aaliceh
haliceh
```
