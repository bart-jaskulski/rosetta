<?php

use PHPUnit\Framework\Attributes\DataProvider;
use PHPUnit\Framework\TestCase;

class IsbnCheckerTest extends TestCase
{

  #[DataProvider('isbnProvider')]
  public function testValidDigits(string $isbn, bool $expectedValid): void
  {
    $checker = new IsbnChecker();
    $this->assertEquals($checker->isValid($isbn), $expectedValid);
  }

  public static function isbnProvider(): iterable
  {
    yield ['978-0596528126', true];
    yield ['978-0596528120', false];
    yield ['978-1788399081', true];
    yield ['978-1788399083', false];
  }
}
