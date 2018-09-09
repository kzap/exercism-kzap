<?php

class Bob
{
    public function respondTo(String $input): String
    {
        // trim extra chars
        $input = preg_replace('/[^\p{Latin}\d?]/u', '', $input);

        $saysNothing = (strlen($input) === 0);

        $isShouting = ($input === mb_convert_case($input, MB_CASE_UPPER, 'UTF-8') && $input !== mb_convert_case($input, MB_CASE_LOWER, 'UTF-8'));

        $isQuestion = (substr($input, -1) === '?');

        // He says 'Fine. Be that way!' if you address him without actually saying anything.
        if ($saysNothing === true) {
            return 'Fine. Be that way!';
        }
        
        if ($isQuestion === true && $isShouting === true) {
            // He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
            return 'Calm down, I know what I\'m doing!';
        }

        // Bob answers 'Sure.' if you ask him a question.
        if ($isQuestion === true) {
            return 'Sure.';
        }

        // He answers 'Whoa, chill out!' if you yell at him.
        if ($isShouting === true) {
            return 'Whoa, chill out!';
        }

        // He answers 'Whatever.' to anything else.
        return 'Whatever.';
    }
}
