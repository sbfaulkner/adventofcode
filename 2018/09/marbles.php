#!/usr/bin/env php
<?php

const INPUT = __DIR__ . "/input";

class Marbles {
  private $marbles = array(0);
  private $current = 0;
  private $marble = 0;

  public function Marbles() {
    $this->marbles = array(0);
    $this->current = 0;
    $this->marble = 0;
  }

  public function play() {
    $score = 0;
    $this->marble++;
    $count = count($this->marbles);

    if ($this->marble % 23 == 0) {
      $score += $this->marble;
      $this->current -= 7;
      $this->current = ($count + ($this->current % $count)) % $count;
      $score += $this->marbles[$this->current];
      array_splice($this->marbles, $this->current, 1);
    } else {
      $this->current = ($this->current + 1) % $count + 1;
      array_splice($this->marbles, $this->current, 0, array($this->marble));
    }

    return $score;
  }

  public function print() {
    $count = count($this->marbles);

    for ($m = 0; $m < $count; $m++) {
      if ($m == $this->current) {
        printf(" (%d)", $this->marbles[$m]);
      } else {
        printf(" %d", $this->marbles[$m]);
      }
    }

    print("\n");
  }
}

$finput = fopen(INPUT, 'r');

while ($line = fgets($finput)) {
  preg_match('/(?<players>[0-9]+) players; last marble is worth (?<limit>[0-9]+) points/', $line, $matches);

  $players = $matches["players"];
  $limit = $matches["limit"];

  $scores = array_fill(0, $players, 0);

  $marbles = new Marbles();

  if ($debug) {
    print("[-] ");
    $marbles->print();
  }

  $player = 0;

  for ($turn = 0; $turn < $limit; $turn++) {
    $score = $marbles->play();
    $scores[$player] += $score;

    if ($debug) {
      printf("[%d]", $player+1);
      $marbles->print();
    }

    $player++;
    $player %= $players;
  }

  $highscore = max($scores);

  printf("%d players; last marble is worth %d points: high score is %d\n", $players, $limit, $highscore);
}

fclose($finput);

?>
