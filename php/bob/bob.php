<?php

class Bob
{
    public function respondTo(string $input): string
    {
        if ($this->isSilent($input))
        {
            return "Fine. Be that way!";
        }

        if ($this->isYelling($input))
        {
            if ($this->isAsking($input))
            {
                return "Calm down, I know what I'm doing!";
            }
            return "Whoa, chill out!";
        }

        if ($this->isAsking($input))
        {
            return "Sure.";
        }

        return "Whatever.";
    }

    private function isSilent(string $str): int
    {
        return preg_match('/^\s*$/', $str);
    }

    private function isYelling(string $str): int
    {
        return preg_match('/^[^A-Za-z]*[A-Z]+(?:[^a-z]*)$/', $str);
    }

    private function isAsking(string $str): int
    {
        return preg_match('/\?\s*$/', $str);
    }
}
