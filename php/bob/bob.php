<?php

class Bob
{
    public function respondTo(String $input): String
    {
        $input = preg_replace('/[^\p{Latin}\d?]/u', '', $input);

        $hasLetters = (preg_match('/[\p{Latin}]/u', $input) === 1) ?? false;

        // He says 'Fine. Be that way!' if you address him without actually saying anything.
        if (!strlen($input)) {
            return 'Fine. Be that way!';

        // Bob answers 'Sure.' if you ask him a question.
        } elseif (substr($input, -1) === '?') {
            // He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
            if ($hasLetters && $input === mb_convert_case($input, MB_CASE_UPPER, 'UTF-8')) {
                return 'Calm down, I know what I\'m doing!';

            } else {
                return 'Sure.';
            }

        // He answers 'Whoa, chill out!' if you yell at him.
        } elseif ($hasLetters && $input === mb_convert_case($input, MB_CASE_UPPER, 'UTF-8')) {
            return 'Whoa, chill out!';

        // He answers 'Whatever.' to anything else.
        } else {
            return 'Whatever.';
        }
    }
}
