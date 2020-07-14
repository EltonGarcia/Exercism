module PhoneNumber
open System.Text.RegularExpressions
open System

let extractCountryCode (numbers: List<int>) = 
    if numbers.Length <> 11 then numbers
    else 
    match numbers with
    | 0::ls -> ls
    | 1::ls -> ls
    | ls -> ls

let validateCountryCode numbers =
    if numbers |> List.length <> 11 then Ok numbers
    else
    match numbers with
    | 1::_ -> Ok (extractCountryCode numbers)
    | _ -> Error("11 digits must start with 1")

let validateInputSize (numbers: List<int>) =
    if numbers.Length > 11 then Error "more than 11 digits"
    else if numbers.Length < 10 then Error "incorrect number of digits"
    else Ok numbers

let validateAreaCode numbers =
    match numbers with
    | 0::_ -> Error("area code cannot start with zero")
    | 1::_ -> Error("area code cannot start with one")
    | phone -> Ok phone

let validateExchangeCode numbers =
    match numbers with
    | _::_::_::0::_ -> Error("exchange code cannot start with zero")
    | _::_::_::1::_ -> Error("exchange code cannot start with one")
    | phone -> Ok phone

let checkLetters input =
    match Regex.Match(input, "[a-zA-Z]+").Success with
    | true -> Error "letters not permitted"
    | false -> Ok input

let checkPunctuations input =
    match Regex.Match(input, "[@:!]+").Success with
    | true -> Error "punctuations not permitted"
    | false -> Ok input

let extractNumbers input = 
    Regex.Matches(input, "[\d]+")
        |> Seq.filter (fun o -> o.Success)
        |> Seq.collect (fun o -> o.Value.ToCharArray())
        |> Seq.map (Char.GetNumericValue >> int)
        |> Seq.toList

let (>=>) switch1 switch2 = 
    switch1 >> (Result.bind switch2)

let clean input = 
    
    let valid = 
        checkLetters 
        >=> checkPunctuations
        >> Result.map extractNumbers
        >> Result.bind validateInputSize
        >=> validateCountryCode
        >=> validateAreaCode
        >=> validateExchangeCode
    
    match valid input with
    | Ok number -> number |> List.map (fun o -> o.ToString()) |> String.concat "" |> uint64 |> Ok
    | Error msg -> Error msg