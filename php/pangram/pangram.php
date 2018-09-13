<?php

define('CHR_START', 97);
define('CHR_END', 122);

function isPangram(string $sentence): bool
{
    return isPangramIsSet($sentence);
}

function isPangramIsSet(string $sentence): bool
{
    $sentence = strtolower($sentence);
    $charArray = str_split($sentence);
    $charArray = array_flip($charArray);

    for ($i = CHR_START; $i <= CHR_END; $i++) {
        $char = chr($i);
        if (!isset($charArray[$char])) {
            return false;
        }
    }

    return true;
}

function isPangramStrPos(string $sentence): bool
{
    $sentence = strtolower($sentence);

    for ($i = CHR_START; $i <= CHR_END; $i++) {
        $char = chr($i);
        if (strpos($sentence, $char) === false) {
            return false;
        }
    }

    return true;
}