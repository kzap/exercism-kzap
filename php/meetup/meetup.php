<?php

function meetup_day(int $year, int $month, string $dayOfWeekIteration, string $dayOfWeek): \DateTimeInterface
{
    // validate day of week
    $dayOfWeek = strtolower(trim($dayOfWeek));
    if (!in_array(
        $dayOfWeek,
        [
            'sun',
            'sunday',
            'mon',
            'monday',
            'tue',
            'tuesday',
            'wed',
            'wednesday',
            'thu',
            'thursday',
            'fri',
            'friday',
            'sat',
            'saturday'
        ]
    )) {
        throw new \InvalidArgumentException('Day of Week Invalid. Must be Sun to Sat or Sunday to Saturday');
    }

    $dayOfWeekIteration = strtolower(trim($dayOfWeekIteration));
    switch ($dayOfWeekIteration) {
        case 'first':
        case 'second':
        case 'third':
        case 'fourth':
        case 'fifth':
        case 'last':
            // get month and year as 'July 2015'
            $monthYearString = date('F Y', mktime(0, 0, 0, $month, 1, $year));

            // create string like 'first monday of July 2015'
            $dateString = $dayOfWeekIteration.' '.$dayOfWeek.' of '.$monthYearString;

            $meetupDate = new \DateTime($dateString);
            break;

        // find day of the week from the 13th to 19th of the month
        case 'teenth':
            for ($dayOfMonth = 13; $dayOfMonth <= 19; $dayOfMonth++) {
                $dayOfWeekCheck = [];
                // check Sunday-Saturday
                $dayOfWeekCheck[] = strtolower(date('D', mktime(0, 0, 0, $month, $dayOfMonth, $year)));
                // check Sun-Sat
                $dayOfWeekCheck[] = strtolower(date('l', mktime(0, 0, 0, $month, $dayOfMonth, $year)));

                if (in_array($dayOfWeek, $dayOfWeekCheck)) {
                    // create date string '2015-07-01'
                    $dateString = implode('-', [$year, $month, $dayOfMonth]);
                    
                    $meetupDate = new \DateTime($dateString);
                    break;
                }
            }
            break;

        default:
            throw new \InvalidArgumentException('Day of the Week Iteration Invalid. Allowed values are first, second, third, fourth, firth, last, teenth');
            break;
    }

    return $meetupDate;
}
