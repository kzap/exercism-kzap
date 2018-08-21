<?php

define('GIGASECOND', 1000000000);

function from(\DateTime $fromDate)
{
    // clone object so we do not modify original object
    $newDate = clone $fromDate;

    // add GIGASECOND to $fromDate
    $newDate = $newDate->add(new \DateInterval('PT'.GIGASECOND.'S'));

    return $newDate;
}
