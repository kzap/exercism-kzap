<?php

//
// This is only a SKELETON file for the "Roman Numerals" exercise. It's been provided as a
// convenience to get you started writing code faster.
//

function toRoman(int $arabicNumber)
{
    global $romanNumeralsDict;

    if ($arabicNumber > 3999) {
        throw new \InvalidArgumentException('Max number to convert is 3999');
    }

    if ($arabicNumber < 1) {
        throw new \InvalidArgumentException('Minimum number to convert is 1');
    }

    $romanNumeral = '';

    $placeValues = strlen($arabicNumber);
    for ($placeValue = $placeValues; $placeValue >= 1; $placeValue--) {
        $stringPos = ($placeValues - $placeValue);
        $numberToConvert = (int) substr($arabicNumber, $stringPos, 1);

        // ignore 0
        if ($numberToConvert === 0) {
            continue;
        }

        // convert 1 into 1000, 2 into 20 etc depending on $placeValue
        $arabricNumeralToConvert = $numberToConvert * (10 ** ($placeValue - 1));

        $romanNumeral .= calculateRomanNumeral($arabricNumeralToConvert);
    }

    return $romanNumeral;
}

function calculateRomanNumeral(int $arabicNumber)
{
    $romanNumeralsDict = [
        1 => 'I',
        5 => 'V',
        10 => 'X',
        50 => 'L',
        100 => 'C',
        500 => 'D',
        1000 => 'M',
    ];

    // lets get ranges of the arabic numerals
    $arabicNumerals = array_keys($romanNumeralsDict);

    foreach ($arabicNumerals as $key => $val) {
        // if the number we are trying to convert higher than the next arabic index
        if (isset($arabicNumerals[($key + 1)]) && $arabicNumber >= $arabicNumerals[($key + 1)]) {
            continue;
        }

        $currentNumeral = $arabicNumerals[$key];
        $prevNumeral = $arabicNumerals[($key - 1)];

        if (isset($arabicNumerals[$key + 1])) {
            $nextNumeral = $arabicNumerals[($key + 1)];
        } else {
            $nextNumeral = null;
        }

        // For 5-9, 50-90, 500-900
        if (substr($currentNumeral, 0, 1) === '5') {
            $modulus = $arabicNumber % $currentNumeral;
            $quotient = $modulus / $prevNumeral;

            // for 9 return IX, 90 return XC, 900 return CM
            if ($quotient === 4) {
                return $romanNumeralsDict[$prevNumeral].$romanNumeralsDict[$nextNumeral];

            // for 8 return VIII, 70 return LXX, 600 return LC etc
            } else {
                return $romanNumeralsDict[$currentNumeral].str_repeat($romanNumeralsDict[$prevNumeral], $quotient);
            }

        // For 1-4, 10-40, 100-400, 1000-3000
        } else {
            $quotient = $arabicNumber / $currentNumeral;

            // for 4 return IV, 40 return XL, 400 return CD
            if ($quotient === 4 && isset($nextNumeral)) {
                return $romanNumeralsDict[$currentNumeral].$romanNumeralsDict[$nextNumeral];

            // for 3 return III, 20 return XX, 100 return C, 2000 return MM
            } else {
                return str_repeat($romanNumeralsDict[$currentNumeral], $quotient);
            }
        }

        break;
    }
}
