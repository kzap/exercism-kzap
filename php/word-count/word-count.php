<?php

function wordCount(string $sentence): array
{
    $sentence = preg_replace('/[^\p{Latin}\d\s?]/u', '', $sentence);
    $sentence = preg_replace('/\s+/u', ' ', $sentence);
    $sentence = mb_convert_case($sentence, MB_CASE_LOWER, 'UTF-8');

    $words = explode(' ', $sentence);
    $words = array_filter($words);
    
    $wordCount = array_count_values($words);

    return $wordCount;
}
