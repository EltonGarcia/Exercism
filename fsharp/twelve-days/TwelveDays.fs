module TwelveDays

let map = [
    ("first", "a Partridge");
    ("second", "two Turtle Doves");
    ("third", "three French Hens");
    ("fourth", "four Calling Birds");
    ("fifth", "five Gold Rings");
    ("sixth", "six Geese-a-Laying");
    ("seventh", "seven Swans-a-Swimming");
    ("eighth", "eight Maids-a-Milking");
    ("ninth", "nine Ladies Dancing");
    ("tenth", "ten Lords-a-Leaping");
    ("eleventh", "eleven Pipers Piping");
    ("twelfth", "twelve Drummers Drumming");
]

let rec formatList gifts =
    match gifts with
    | [head] -> head
    | [head; tail] -> (sprintf "%s, and %s" head tail)
    | _ -> sprintf "%s, %s" gifts.Head (formatList gifts.Tail)

let sentence index day =
    let gifts =
        map 
        |> List.take index
        |> List.rev
        |> List.map snd
        |> formatList
    sprintf "On the %s day of Christmas my true love gave to me: %s in a Pear Tree." day gifts
    
let recite start stop = 
    map 
    |> List.indexed
    |> List.skip (start - 1)
    |> List.take (stop - start + 1)
    |> List.map (fun (i, (day, _)) -> sentence (i+1) day)
    