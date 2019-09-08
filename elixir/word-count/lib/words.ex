defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    sentence |> clean_and_split |> make(%{})
  end

  defp clean_and_split(string) do
    string
    |> String.replace(~r/[^0-9A-Za-zäöüÄÖÜß-]+/, " ")
    |> String.downcase()
    |> String.split
  end

  defp make([head | tail], map) do
    with true <- Map.has_key?(map, head) do
      make(tail, %{map | head => map[head]+1})
    else
      _ ->
        make(tail, Map.put(map, head, 1))
    end
  end

  defp make([], map), do: map
end
