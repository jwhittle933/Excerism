defmodule SpaceAge do
  @type planet ::
          :mercury
          | :venus
          | :earth
          | :mars
          | :jupiter
          | :saturn
          | :uranus
          | :neptune

  @doc """
  Return the number of years a person that has lived for 'seconds' seconds is
  aged on 'planet'.
  """
  @spec age_on(planet, pos_integer) :: float
  def age_on(planet, seconds) do
    apply(__MODULE__, planet, [seconds])
  end

  def mercury(seconds), do: seconds / toSeconds(0.2408467)

  def venus(seconds), do: seconds / toSeconds(0.61519726)

  def earth(seconds), do: seconds / 31557600

  def mars(seconds), do: seconds / toSeconds(1.8808158)

  def jupiter(seconds), do: seconds / toSeconds(11.862615)

  def saturn(seconds), do: seconds / toSeconds(29.447489)

  def uranus(seconds), do: seconds / toSeconds(84.016846)

  def neptune(seconds), do: seconds / toSeconds(164.79132)

  defp toSeconds(relEarthYears), do: relEarthYears * 365.25 * 24 * 60 * 60
end
