<?php

define('CHR_START', 'A');
define('CHR_END', 'AA');

function isPangram(string $sentence): bool
{
    return isPangramIsSet($sentence);
}

function isPangramIsSet(string $sentence): bool
{
    $sentence = strtoupper($sentence);
    $charArray = str_split($sentence);
    $charArray = array_flip($charArray);

    for ($char = CHR_START; $char != CHR_END; $char++) {
        if (!isset($charArray[$char])) {
            return false;
        }
    }

    return true;
}

function isPangramStrPos(string $sentence): bool
{
    $sentence = strtoupper($sentence);

    for ($char = CHR_START; $char != CHR_END; $char++) {
        if (strpos($sentence, $char) === false) {
            return false;
        }
    }

    return true;
}