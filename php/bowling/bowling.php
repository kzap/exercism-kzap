<?php

class Game
{
    const STARTING_FRAME = 1;
    const TOTAL_FRAMES = 10;
    const MAX_PINS = 10;

    protected $pinsStanding = 0;
    protected $currentFrame = 1;
    protected $currentDelivery = 0;
    protected $bowlingScoreCard = [];
    protected $gameEnded = false;

    public function __construct()
    {
        $this->resetGame();
    }

    public function roll(int $pinsKnockedDown = null)
    {
        if (!isset($pinsKnockedDown)) {
            $pinsKnockedDown = mt_rand(0, $this->pinsStanding);
        }

        $this->recordScore($pinsKnockedDown);

        return $this;
    }

    public function score(): int
    {
        if (!$this->gameEnded) {
            throw new \LogicException('Incomplete Game Can Not Be Scored.');
        }

        return $this->calculateFinalScore();
    }

    public function recordScore(int $pinsKnockedDown)
    {
        if ($this->gameEnded) {
            throw new \LogicException('Game has ended already. Stop throwing the ball.');
        }

        if ($pinsKnockedDown > $this->pinsStanding || $pinsKnockedDown > self::MAX_PINS) {
            throw new \InvalidArgumentException(sprintf('I can count. Invalid $pinsKnockedDown: %s', $pinsKnockedDown));
        }

        if (!in_array($pinsKnockedDown, range(0, self::MAX_PINS))) {
            throw new \InvalidArgumentException(sprintf('Nope. Invalid $pinsKnockedDown: %s', $pinsKnockedDown));
        }

        $this->currentDelivery++;
        $this->bowlingScoreCard[] = $pinsKnockedDown;
        $this->pinsStanding -= $pinsKnockedDown;

        // figure out if we should go to the next frame
        if ($this->currentFrame < self::TOTAL_FRAMES) {
            // if there are no more pins standing
            if ($this->pinsStanding === 0) {
                $this->startNextFrame();

            // if we have bowled 2 times
            } elseif ($this->currentDelivery === 2) {
                $this->startNextFrame();
            }

        // figure out if we should end the game on the 10th frame
        } elseif ($this->currentFrame === self::TOTAL_FRAMES) {
            // if we have bowled 3 times
            if ($this->currentDelivery === 3) {
                $this->endGame();

            // if we knocked down all pins, lets reset them
            } elseif ($this->pinsStanding === 0) {
                $this->resetPins();
                $this->bonusDelivery = true;

            // if we have bowled 2 times but there are still pins remaining
            } elseif ($this->currentDelivery === 2 && !$this->bonusDelivery) {
                $this->endGame();
            }
        }

        return $this;
    }

    public function startNextFrame()
    {
        $this->resetPins();
        $this->currentFrame++;
        $this->currentDelivery = 0;

        return $this;
    }

    public function resetGame()
    {
        $this->resetPins();
        $this->clearScore();
        $this->currentFrame = 1;
        $this->currentDelivery = 0;
        $this->gameEnded = false;

        return $this;
    }

    public function endGame()
    {
        $this->gameEnded = true;

        return $this;
    }

    public function resetPins()
    {
        $this->pinsStanding = self::MAX_PINS;

        return $this;
    }

    public function clearScore()
    {
        $this->bowlingScoreCard = [];

        return $this;
    }

    protected function calculateFinalScore(): int
    {
        $totalScore = 0;
        $frameNumber = 1;
        $deliveryNumber = 0;

        foreach ($this->bowlingScoreCard as $key => $pinsKnockedDown) {
            $deliveryNumber++;

            // Frames 1-9
            if ($frameNumber < self::TOTAL_FRAMES) {
                if ($deliveryNumber === 1) {
                    if ($pinsKnockedDown < self::MAX_PINS) {
                        $totalScore += $pinsKnockedDown;

                    // Strike: include next 2 deliveries
                    } elseif ($pinsKnockedDown === self::MAX_PINS) {
                        $totalScore += $pinsKnockedDown;
                        if (isset($this->bowlingScoreCard[$key+1])) {
                            $totalScore += $this->bowlingScoreCard[$key+1];
                        }
                        if (isset($this->bowlingScoreCard[$key+2])) {
                            $totalScore += $this->bowlingScoreCard[$key+2];
                        }

                        $frameNumber++;
                        $deliveryNumber = 0;
                    }
                
                } elseif ($deliveryNumber === 2) {
                    // Spare: include next delivery
                    if (($pinsKnockedDown + $this->bowlingScoreCard[$key-1]) === self::MAX_PINS) {
                        $totalScore += $pinsKnockedDown;
                        if (isset($this->bowlingScoreCard[$key+1])) {
                            $totalScore += $this->bowlingScoreCard[$key+1];
                        }

                    // Open Frame
                    } else {
                        $totalScore += $pinsKnockedDown;
                    }

                    $frameNumber++;
                    $deliveryNumber = 0;
                }

            // Frame 10
            } elseif ($frameNumber === self::TOTAL_FRAMES) {
                if ($deliveryNumber <= 3) {
                    $totalScore += $pinsKnockedDown;
                }
            }
        }

        return $totalScore;
    }
}
