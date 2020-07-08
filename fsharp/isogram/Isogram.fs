module Isogram
open System

let isIsogram str = 
    str
    |> Seq.filter Char.IsLetter
    |> Seq.countBy Char.ToLowerInvariant
    |> Seq.forall (fun count -> snd count = 1)