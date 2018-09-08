<?php

function toDecimal(string $trinary): int
{
    $decimalValue = 0;

    $trinary = trim($trinary);
    $trinaryNumbers = str_split($trinary);
    $trinaryCount = count($trinaryNumbers);
    
    foreach ($trinaryNumbers as $key => $trinaryNumber) {
        $trinaryNumber = (int) $trinaryNumber;

        if (!in_array($trinaryNumber, [0,1,2])) {
            return 0;
        }

        $exponent = $trinaryCount - $key - 1;

        $decimalValue += $trinaryNumber * (3**$exponent);
    }

    return $decimalValue;
}
