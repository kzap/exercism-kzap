<?php

//
// This is only a SKELETON file for the "Roman Numerals" exercise. It's been provided as a
// convenience to get you started writing code faster.
//

$romanNumeralsDict = [
    1 => 'I',
    5 => 'V',
    10 => 'X',
    50 => 'L',
    100 => 'C',
    500 => 'D',
    1000 => 'M',
];

function toRoman(int $arabicNumeral)
{
    global $romanNumeralsDict;

    $arabicRanges = array_keys($romanNumeralsDict);

    if ($arabicNumeral > 3000) {
        throw new \InvalidArgumentException('Max number to convert is 3000');
    }

    if ($arabicNumeral < 1) {
        throw new \InvalidArgumentException('Minimum number to convert is 1');
    }

    for ($i = 0; $i < strlen($arabicNumeral); $i++) {
        $numToConvert = substr($arabicNumeral, $i);

        // ignore 0
        if ((int) $numToConvert == 0) {
            continue;
        }

        $numBase10 = convertToBase10($numToConvert);
var_dump($numBase10);
        foreach ($arabicRanges as $key => $val) {
            $arabicRangeMin = $arabicRanges[$key];
            $arabicRangeMax = null;
            if (isset($arabicRanges[$key+1])) {
                $arabicRangeMax = $arabicRanges[$key+1];
            }

            if (!isset($arabicRangeMax)) {
                $romanNumeralChar1 = $romanNumeralsDict[$arabicRangeMin];
                $romanNumeralChar2 = null;

            } elseif ($numBase10 >= $arabicRangeMin && $numBase10 < $arabicRangeMax) {
                $romanNumeralChar1 = $romanNumeralsDict[$arabicRangeMin];
                $romanNumeralChar2 = $romanNumeralsDict[$arabicRangeMax];
            }

            $repetitions = $numBase10 / $arabicRangeMin;
            if ($repetitions <= 3) {
                $romanNumeral .= str_repeat($romanNumeralChar1, $repetitions);
            } elseif ($repetitions == 4) {
                $romanNumeral .= $romanNumeralChar1 . $romanNumeralChar2;
            }
        }
    }

    return $romanNumeral;
}

function convertToBase10(int $number)
{
    $firstDigit = substr($number, 0, 1);

    $base10 = str_pad($firstDigit, strlen($number), '0');

    return (int) $base10;
}
