<?php

function detectAnagrams(string $word, array $wordsToCheck): array
{
    $anagrams = [];

    $word = mb_convert_case($word, MB_CASE_LOWER, 'UTF-8');
    $wordArray = str_split($word);
    sort($wordArray);

    foreach ($wordsToCheck as $wordToCheck) {
        $wordToCheckTest = mb_convert_case($wordToCheck, MB_CASE_LOWER, 'UTF-8');

        if ($word === $wordToCheckTest) {
            continue;
        }

        $wordToCheckArray = str_split($wordToCheckTest);
        sort($wordToCheckArray);

        if ($wordArray === $wordToCheckArray) {
            $anagrams[] = $wordToCheck;
        }
    }

    return $anagrams;
}
