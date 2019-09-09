<?php

function distance(string $strandA, string $strandB) : int
{
    if (strlen($strandA) !== strlen($strandB))
    {
        throw new InvalidArgumentException('DNA strands must be of equal length.');
    }

    $arrA = str_split($strandA);
    $arrB  = str_split($strandB);
    $result = array_diff_assoc($arrA, $arrB);

    return count($result);
}
