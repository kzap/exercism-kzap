<?php

function isLeap(int $year): bool
{
    $isLeapYear = false;

    // on every year that is evenly divisible by 4
    if ($year % 4 === 0) {
        $isLeapYear = true;

        // except every year that is evenly divisible by 100
        if ($year % 100 === 0) {
            $isLeapYear = false;

            // unless the year is also evenly divisible by 400
            if ($year % 400 === 0) {
                $isLeapYear = true;
            }
        }
    }

    return $isLeapYear;
}
