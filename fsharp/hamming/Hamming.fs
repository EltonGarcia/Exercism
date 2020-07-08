module Hamming

let distance (strand1: string) (strand2: string): int option = 
    if strand1.Length <> strand2.Length then None
    else Seq.zip strand1 strand2 
        |> Seq.sumBy (fun (a, b) -> if a <> b then 1 else 0)
        |> Some