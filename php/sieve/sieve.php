<?php

// Algorithm from https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes#Pseudocode

function sieve(int $maxInt): Array
{
    $primeIntegers = [];

    if ($maxInt >= 2) {
        // Let A be an array of Boolean values, indexed by integers 2 to n, initially all set to true.
        for ($i = 2; $i <= $maxInt; $i++) {
            $primeIntegers[$i] = true;
        }

        // for i = 2, 3, 4, ..., not exceeding √n:
        $sqrtMaxInt = (int) floor(sqrt($maxInt));

        for ($i = 2; $i <= $sqrtMaxInt; $i++) {
            // if A[i] is true:
            if (isset($primeIntegers[$i]) && $primeIntegers[$i] === true) {
                // for j = i2, i2+i, i2+2i, i2+3i, ..., not exceeding n:
                $a = 0;
                do {
                    $j = ($i ** 2) + ($i * $a);

                    if (isset($primeIntegers[$j])) {
                        // A[j] := false.
                        $primeIntegers[$j] = false;
                    }

                    $a++;
                } while ($j <= $maxInt);
            }
        }

        $primeIntegers = array_filter($primeIntegers);
        $primeIntegers = array_keys($primeIntegers);
    }

    return $primeIntegers;
}
