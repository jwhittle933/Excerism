<?php

/*
This is only a SKELETON file for the "Hamming" exercise. It's been provided as a
convenience to get you started writing code faster.

Remove this comment before submitting your exercise.
*/

function distance(string $strandA, string $strandB) : int
{
    if (strlen($strandA) !== strlen($strandB))
    {
        throw new InvalidArgumentException('DNA strands must be of equal length.');
    }

    $arrA = str_split($strandA);
    $arrB  = str_split($strandB);
    $uncommon = 0;

    for ($i = 0; $i < count($arrA); $i++ )
    {
        if ($arrA[$i] === $arrB[$i])
        {
            continue;
        }

        $uncommon++;
    }

    return $uncommon;
}
