<?php

final class RobotNames
{
    protected $issuedNames = [];

    public static function Instance()
    {
        static $inst = null;

        if ($inst === null) {
            $inst = new self();
        }

        return $inst;
    }

    private function __construct() {}
    private function __clone() {}

    public function getIssuedNames(): Array
    {
        return $this->issuedNames;
    }

    public function addIssuedName(String $name)
    {
        $this->issuedNames[] = $name;
        $this->issuedNames = array_unique($this->issuedNames);

        return $this;
    }
}

class Robot
{
    protected $name;
    protected $issuedNames;

    public function __construct()
    {
        $this->issuedNames = RobotNames::Instance();
        $this->reset();
    }

    public function getName(): String
    {
        return $this->name;
    }

    public function setName(String $name)
    {
        $this->name = $name;
        $this->issuedNames->addIssuedName($name);

        return $this;
    }

    public function getIssuedNames(): Array
    {
        return $this->issuedNames->getIssuedNames();
    }

    public function reset()
    {
        $name = $this->generateUniqueName();
        $this->setName($name);

        return $this;
    }

    protected function generateUniqueName()
    {
        do {
            $name = $this->generateRandomName();
        } while (in_array($name, $this->issuedNames->getIssuedNames()));

        return $name;
    }

    protected function generateRandomName()
    {
        $letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
        $lettersCount = strlen($letters);
        $numbers = '0123456789';
        $numbersCount = strlen($numbers);

        $name = '';

        // 2 letters
        for ($i = 0; $i < 2; $i++) {
            $name .= $letters[mt_rand(0, ($lettersCount - 1))];
        }

        // 3 numbers
        for ($i = 0; $i < 3; $i++) {
            $name .= $numbers[mt_rand(0, ($numbersCount - 1))];
        }

        return $name;
    }
}
