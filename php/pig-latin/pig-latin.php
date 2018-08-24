<?php

define('VOWELS', ['a','e','i','o','u']);

define('CONSONANT_DIGRAPHS', ['bl', 'br', 'ch', 'cl', 'cr', 'dr', 'fl', 'fr', 'gl', 'gr', 'qu', 'pl', 'pr', 'sc', 'sh', 'sk', 'sl', 'sm', 'sn', 'sp', 'st', 'sw', 'sw', 'th', 'tr', 'tw', 'wh', 'wr']);

define('CONSONANT_TRIGRAPHS', ['sch', 'scr', 'shr', 'sph', 'spl', 'spr', 'squ', 'str', 'thr']);

function translate(String $sentence): String
{
    $sentence = trim($sentence);
    $words = explode(' ', $sentence);

    $pigLatin = [];
    
    foreach ($words as $word) {
        $wordToCheck = strtolower($word);

        if (in_array($wordToCheck[0], VOWELS)) {
            $word = $word.'ay';

        } elseif (in_array(substr($wordToCheck, 0, 3), CONSONANT_TRIGRAPHS)) {
            $word = substr($word, 3).substr($word, 0, 3).'ay';

        } elseif (in_array(substr($wordToCheck, 0, 2), CONSONANT_DIGRAPHS)) {
            $word = substr($word, 2).substr($word, 0, 2).'ay';

        } elseif (!in_array($wordToCheck[0], VOWELS) && !in_array($wordToCheck[1], VOWELS)) {
            $word = $word.'ay';

        } else {
            $word = substr($word, 1).$word[0].'ay';
        }

        $pigLatin[] = $word;
    }

    $pigLatinSentence = implode(' ', $pigLatin);

    return $pigLatinSentence;
}
