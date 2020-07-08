module Pangram

let isPangram (input: string): bool = 
    input.ToLowerInvariant()
    |> set
    |> Set.isSubset (Set ['a'..'z'])