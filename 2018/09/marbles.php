#!/usr/bin/env php
<?php

const INPUT = __DIR__ . "/input";

class Marbles {
  private $marbles;
  private $current;
  private $marble;

  public function __construct() {
    $this->marbles = new Ds\Deque([0]);
    $this->current = 0;
    $this->marble = 0;
  }

  public function play() {
    $score = 0;
    $this->marble++;

    if ($this->marble % 23 == 0) {
      $score += $this->marble;
      $this->marbles->rotate(-7);
      $score += $this->marbles->shift();
    } else {
      $this->marbles->rotate(2);
      $this->marbles->unshift($this->marble);
    }

    return $score;
  }
}

$finput = fopen(INPUT, 'r');

while ($line = fgets($finput)) {
  preg_match('/(?<players>[0-9]+) players; last marble is worth (?<limit>[0-9]+) points/', $line, $matches);

  $players = $matches["players"];
  $limit = $matches["limit"] * 100;

  $scores = array_fill(0, $players, 0);

  $marbles = new Marbles();

  $player = 0;

  for ($turn = 0; $turn < $limit; $turn++) {
    $score = $marbles->play();
    $scores[$player] += $score;

    $player++;
    $player %= $players;
  }

  $highscore = max($scores);

  printf("%d players; last marble is worth %d points: high score is %d\n", $players, $limit, $highscore);
}

fclose($finput);

?>
