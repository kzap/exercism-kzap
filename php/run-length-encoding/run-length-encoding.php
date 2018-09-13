<?php

function encode(string $input): string
{
	$regex = '!([\w\s])\1*!';

	$encodedString = preg_replace_callback(
    	$regex, 
        function ($matches) {
            return encodeChar($matches[0]);
        },
        $input
    );

    return $encodedString;
}

function decode(string $input): string
{
	$decodedString = '';

    $regex = '!\d+[\w|\s]|[\w|\s]!';

    $decodedString = preg_replace_callback(
    	$regex, 
        function ($matches) {
            return decodeChar($matches[0]);
        },
        $input
    );
    
    return $decodedString;
}

function encodeChar(string $charString): string
{
	$charRepetitions = strlen($charString);

	if ($charRepetitions === 1) {
		return $charString;
	}

	return (string) $charRepetitions . $charString[0];
}

function decodeChar(string $encodedChar): string
{
	if (strlen($encodedChar) === 1) {
		return $encodedChar;
	}

	$decodedChar = $encodedChar[-1];
	$decodedCharCount = (int) substr($encodedChar, 0, -1);

	return str_repeat($decodedChar, $decodedCharCount);
}
