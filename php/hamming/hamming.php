<?php

//
// This is only a SKELETON file for the "Hamming" exercise. It's been provided as a
// convenience to get you started writing code faster.
//

function distance($a, $b)
{
    $stringA = strtoupper(trim($a));
    $stringB = strtoupper(trim($b));

    if (strlen($stringA) !== strlen($stringB)) {
        throw new \InvalidArgumentException('DNA strands must be of equal length.');
    }

    if (!strlen($stringA) || !strlen($stringA)) {
        throw new \InvalidArgumentException('DNA strands have atleast 1 letter.');
    }

    $hammingDistance = 0;
    for ($i = 0; $i < strlen($stringA); $i++) {
        if ($stringA[$i] !== $stringB[$i]) {
            $hammingDistance++;
        }
    }

    return $hammingDistance;
}
