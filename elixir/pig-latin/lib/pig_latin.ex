defmodule PigLatin do
  @doc """
  Given a `phrase`, translate it a word at a time to Pig Latin.

  Words beginning with consonants should have the consonant moved to the end of
  the word, followed by "ay".

  Words beginning with vowels (aeiou) should have "ay" added to the end of the
  word.

  Some groups of letters are treated like consonants, including "ch", "qu",
  "squ", "th", "thr", and "sch".

  Some groups are treated like vowels, including "yt" and "xr".
  """


  @con_clusters ["ch", "qu", "squ", "th", "thr", "sch"]
  @vow_morphs [?a, ?e, ?i, ?o, ?u, "yt", "xr"]
  @pig_morph "ay"

  @spec translate(phrase :: String.t()) :: String.t()
  def translate(phrase) do
    phrase
    |> String.split
    |> Enum.map(&parse/1)
    |> List.foldl("", fn x, acc -> acc <> " " <> x end)
    |> String.trim
  end

  def parse(<<"ch", rest::binary>>), do: rest <> "ch" <> @pig_morph
  def parse(<<"qu", rest::binary>>), do: rest <> "qu" <> @pig_morph
  def parse(<<"squ", rest::binary>>), do: rest <> "squ" <> @pig_morph
  def parse(<<"th", rest::binary>>), do: rest <> "th" <> @pig_morph
  def parse(<<"thr", rest::binary>>), do: rest <> "thr" <> @pig_morph
  def parse(<<"sch", rest::binary>>), do: rest <> "sch" <> @pig_morph
  def parse(<<vow::size(8), _::binary>> = str) when vow in @vow_morphs, do: str <> @pig_morph
  # def parse(str), do: str <> @pig_morph


  # def parse(<<two::size(16), rest::binary>>) when two in @con_clusters do
  #   rest <> two <> @pig_morph
  # end


  # def parse(<<two::size(16), rest::binary>> = str) when two in @vow_morphs do
  #   str <> @pig_morph
  # end

  # # def parse(<<"xr", rest::binary>> = str), do: str <> @pig_morph
  # def parse(<<"p", rest::binary>>), do: rest <> "p" <> @pig_morph
  # def parse(str), do: str <> @pig_morph
end
