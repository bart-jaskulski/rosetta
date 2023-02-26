<?php

declare(strict_types=1);

class IsbnChecker
{

  public function isValid(string $isbn): bool
  {
    $digitsArray = array_filter(
        str_split($isbn),
        static fn($i) => is_numeric($i)
    );
    return array_sum(
      array_map(
        static fn ($v, $k) => $k % 2 === 0 ? $v * 3 : $v,
        $digitsArray,
        array_keys($digitsArray)
      )
    ) % 10 === 0;
  }
}
