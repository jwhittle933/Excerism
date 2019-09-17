defmodule RomanNumerals do
  @doc """
  Convert the number to a roman number.
  """
  @spec numeral(pos_integer) :: String.t()
  def numeral(number) do
    number |> resolve
  end

  def resolve(number, str \\ "")
  def resolve(number, str) when number >= 1000 do
    resolve(rem(number, 1000), str <> String.duplicate("M", div(number, 1000)))
  end

  def resolve(number, str) when number >= 900 do
    resolve(rem(number, 100), str <> "CM")
  end

  def resolve(number, str) when number >= 500 do
    resolve(rem(number, 100), str <> "D" <> String.duplicate("C", div(number - 500, 100)))
  end

  def resolve(number, str) when number >= 400 do
    resolve(rem(number, 100), str <> "CD")
  end

  def resolve(number, str) when number >= 100 do
    resolve(rem(number, 100), str <> String.duplicate("C", div(number, 100)))
  end

  def resolve(number, str) when number >= 90 do
    resolve(rem(number, 10), str <> "XC")
  end

  def resolve(number, str) when number >= 50 do
    resolve(rem(number, 10), str <> "L" <> String.duplicate("X", div(number - 50, 10)))
  end

  def resolve(number, str) when number >= 40 do
    resolve(rem(number, 10), str <> "XL")
  end

  def resolve(number, str) when number >= 10 do
    resolve(rem(number, 10), str <> String.duplicate("X", div(number, 10)))
  end

  def resolve(number, str) when number == 9 do
    resolve(rem(number, 1), str <> "IX")
  end

  def resolve(number, str) when number >= 5 do
    resolve(rem(number, 1), str <> "V" <> String.duplicate("I", div(number - 5, 1)))
  end

  def resolve(number, str) when number == 4 do
    resolve(rem(number, 1), str <> "IV")
  end

  def resolve(number, str) when number >= 1 do
    resolve(0, str <> String.duplicate("I", number))
  end

  def resolve(0, str), do: str
end
